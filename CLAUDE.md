# genshin-repository

Download links and Sophon build data for every Genshin Impact (global PC)
version. A Go fetcher updates `data/versions.json` from the HoYoverse API
daily and regenerates `README.md` and `docs/versions/`.

## Commands

- `go run ./cmd/fetch`: query the API, re-check every archive link, merge into `data/versions.json`, regenerate `README.md` and `docs/versions/`; `-check-links=false` skips the link check
- `go run ./cmd/fetch -backfill`: also probe Sophon builds for every older `X.Y.0` tag and recorded version
- `go run ./cmd/fetch -render-only`: regenerate `README.md` and `docs/versions/` without the network; `-out DIR` writes them under another root
- `go run ./cmd/seed`: one-off import of the upstream history at pinned commits (merges, never removes)
- Checks: `go vet ./...`, `golangci-lint run`, `go test ./...`, `CGO_ENABLED=0 go build ./...`

## Layout

- `internal/store`: data model, merge rules, canonical JSON
- `internal/hoyo`: API client (getGamePackages, getGameBranches, getBuild, POST getPatchBuild) and `Sync`
- `internal/upstream`: parser for JRSKelvin's Markdown lists and API snapshots
- `internal/linkcheck`: HEAD check of every archive link, run by `fetch` daily and by `seed`
- `internal/render`: README and `docs/versions/<N>.x.md` renderer; page shells in `*.md.tmpl`, tables in `tables.go`
- `docs/download-guide.md`, `docs/sophon.md`, `docs/data.md`: hand-written guides (no store data)
- `.github/workflows/update.yml`: daily fetch and bot commit; `ci.yml`: vet, lint, test, build, generated-files sync

## Rules

- `README.md`, `docs/versions/` and `data/` are generated. Edit the generator or its `.md.tmpl` shells, then run `-render-only`. The other `docs/*.md` are hand-written.
- Version headings are the bare version (`## 5.5.0`, anchor `#550`); status goes on the line below, so deep links stay stable.
- The store only grows: never delete a version, package or patch.
- Output must be deterministic: no run timestamps, stable sort order.
- The zip API is frozen at 5.5.0; only getGameBranches decides the live version and its dates.
- Stdlib only, pure Go (`CGO_ENABLED=0`).
- Code is Apache-2.0; `data/`, README and `docs/` content are CC BY-NC-SA 4.0 (`LICENSE-DATA`).
