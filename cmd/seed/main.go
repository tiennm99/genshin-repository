// Command seed imports the upstream history (JRSKelvin/GenshinRepository
// and JRSKelvin/Genshin-Repository-List, at pinned commits) into
// data/versions.json and regenerates README.md and docs/versions/. It is a
// one-off: the daily job never re-reads upstream. Running it again merges
// the same data.
//
//	go run ./cmd/seed
//	go run ./cmd/seed -check-links=false
package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tiennm99/genshin-repository/internal/linkcheck"
	"github.com/tiennm99/genshin-repository/internal/render"
	"github.com/tiennm99/genshin-repository/internal/store"
	"github.com/tiennm99/genshin-repository/internal/upstream"
)

func main() {
	dataPath := flag.String("data", "data/versions.json", "version store")
	out := flag.String("out", ".", "repository root that receives README.md and docs/versions/")
	checkLinks := flag.Bool("check-links", true, "HEAD every archive URL to record dead links and real sizes")
	workers := flag.Int("workers", 16, "parallel link checks")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	err := run(ctx, *dataPath, *out, *checkLinks, *workers)
	stop()
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, dataPath, out string, checkLinks bool, workers int) error {
	s, err := store.Load(dataPath)
	if err != nil {
		return err
	}
	hc := &http.Client{Timeout: 2 * time.Minute}

	md, err := upstream.Download(ctx, hc, upstream.TarballURL(upstream.RepoMarkdown, upstream.CommitMD))
	if err != nil {
		return err
	}
	log.Printf("%s: %d Version files", upstream.RepoMarkdown, upstream.ImportMarkdownRepo(s, md))

	snaps, err := upstream.Download(ctx, hc, upstream.TarballURL(upstream.RepoSnapshots, upstream.CommitSnaps))
	if err != nil {
		return err
	}
	counts, err := upstream.ImportSnapshotRepo(s, snaps)
	if err != nil {
		return err
	}
	log.Printf("%s: snapshots per format %v", upstream.RepoSnapshots, counts)

	if checkLinks {
		live, dead, unknown := linkcheck.Check(ctx, &http.Client{Timeout: 30 * time.Second}, s, workers)
		log.Printf("link check: %d live, %d dead, %d unknown (status kept)", live, dead, unknown)
	}
	if err := s.Save(dataPath); err != nil {
		return err
	}
	files, err := render.Render(s)
	if err != nil {
		return err
	}
	return render.Write(out, files)
}
