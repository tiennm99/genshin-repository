// Package linkcheck records which archive links still answer.
package linkcheck

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/tiennm99/genshin-repository/internal/store"
)

// Check sends a HEAD request to every archive URL and records 200, 404
// and 410 answers. Other answers (the CDN sometimes answers 400 for archives
// over 50 GB) leave the previous status alone and count as unknown. For a
// live link, Content-Length replaces the recorded size: legacy snapshot sizes
// are a few MB off for some archives.
func Check(ctx context.Context, hc *http.Client, s *store.Store, workers int) (live, dead, unknown int) {
	var pkgs []*store.Package
	for _, v := range s.Versions {
		pkgs = append(pkgs, v.Full...)
		for _, p := range v.Patches {
			pkgs = append(pkgs, p.Packages...)
		}
	}
	jobs := make(chan *store.Package)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for p := range jobs {
				status, length := head(ctx, hc, p.URL)
				mu.Lock()
				switch status {
				case http.StatusOK:
					live++
					p.HTTPStatus = status
					p.SetServedSize(length)
				case http.StatusNotFound, http.StatusGone:
					dead++
					p.HTTPStatus = status
				default:
					unknown++
				}
				mu.Unlock()
			}
		})
	}
	for _, p := range pkgs {
		jobs <- p
	}
	close(jobs)
	wg.Wait()
	return live, dead, unknown
}

// retryDelay is the back-off step between HEAD attempts.
var retryDelay = time.Second

func head(ctx context.Context, hc *http.Client, u string) (status int, length int64) {
	for attempt := range 4 {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * retryDelay)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)
		if err != nil {
			return 0, 0
		}
		resp, err := hc.Do(req)
		if err != nil {
			continue
		}
		_ = resp.Body.Close()
		status, length = resp.StatusCode, resp.ContentLength
		// Only 200 and "gone" answers are final; retry anything else.
		if status == http.StatusOK || status == http.StatusNotFound || status == http.StatusGone {
			return status, length
		}
	}
	return status, length
}
