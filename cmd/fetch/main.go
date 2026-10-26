// Command fetch queries the HoYoverse APIs, merges the answers into
// data/versions.json and regenerates README.md and docs/versions/.
//
//	go run ./cmd/fetch               daily update
//	go run ./cmd/fetch -backfill     also probe Sophon builds of all older versions
//	go run ./cmd/fetch -render-only  only regenerate README.md and docs/versions/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/tiennm99/genshin-repository/internal/hoyo"
	"github.com/tiennm99/genshin-repository/internal/render"
	"github.com/tiennm99/genshin-repository/internal/store"
)

func main() {
	dataPath := flag.String("data", "data/versions.json", "version store")
	out := flag.String("out", ".", "repository root that receives README.md and docs/versions/")
	backfill := flag.Bool("backfill", false, "probe getBuild for older X.Y.0 tags and recorded versions that lack Sophon data")
	renderOnly := flag.Bool("render-only", false, "skip the API and only regenerate README.md and docs/versions/")
	date := flag.String("date", time.Now().UTC().Format(time.DateOnly), "observation date (UTC, YYYY-MM-DD)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	err := run(ctx, *dataPath, *out, *date, *backfill, *renderOnly)
	stop()
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, dataPath, out, date string, backfill, renderOnly bool) error {
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return fmt.Errorf("-date: %w", err)
	}
	s, err := store.Load(dataPath)
	if err != nil {
		return err
	}
	if !renderOnly {
		err := hoyo.Sync(ctx, hoyo.New(), s, hoyo.Options{Date: date, Backfill: backfill, Logf: log.Printf})
		if err != nil {
			return err
		}
		if err := s.Save(dataPath); err != nil {
			return err
		}
	}
	files, err := render.Render(s)
	if err != nil {
		return err
	}
	return render.Write(out, files)
}
