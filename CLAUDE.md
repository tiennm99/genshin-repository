# genshin-repository

Download links and Sophon build data for every Genshin Impact (global PC)
version. A Go fetcher updates `data/versions.json` from the HoYoverse API
daily and regenerates `README.md`.

## Commands

- `go run ./cmd/fetch`: query the API, merge into `data/versions.json`, regenerate `README.md`
- `go run ./cmd/fetch -backfill`: also probe Sophon builds for older `X.Y.0` tags from 5.6.0
- `go run ./cmd/fetch -render-only`: regenerate `README.md` without the network
- `go run ./cmd/seed`: one-off import of the upstream history at pinned commits (merges, never removes)
- Checks: `go vet ./...`, `golangci-lint run`, `go test ./...`, `CGO_ENABLED=0 go build ./...`

## Layout

- `internal/store`: data model, merge rules, canonical JSON
- `internal/hoyo`: API client (getGamePackages, getGameBranches, getBuild, POST getPatchBuild) and `Sync`
- `internal/upstream`: parser for JRSKelvin's Markdown lists and API snapshots; HEAD link check
- `internal/readme`: README renderer; static text in `intro.md` and `guide.md`
- `.github/workflows/update.yml`: daily fetch and bot commit; `ci.yml`: vet, lint, test, build, README sync

## Rules

- `README.md` and `data/` are generated. Edit the generator or its `.md` templates, then run `-render-only`.
- The store only grows: never delete a version, package or patch.
- Output must be deterministic: no run timestamps, stable sort order.
- The zip API is frozen at 5.5.0; only getGameBranches decides the live version and its dates.
- Stdlib only, pure Go (`CGO_ENABLED=0`).
- Code is Apache-2.0; `data/` and README content are CC BY-NC-SA 4.0 (`LICENSE-DATA`).
