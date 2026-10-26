# Data format: `data/versions.json`

[← README](../README.md) · [Versions](../README.md#versions) · [Download guide](download-guide.md) · [Sophon](sophon.md)

[`data/versions.json`](../data/versions.json) is the single source of this
repository. The README and every page under [versions/](versions/) are
generated from it by `go run ./cmd/fetch`, and a daily GitHub Actions job
merges the official HoYoverse API answer into it.

## Rules

- **It only grows.** A version, archive or patch that once appeared is never
  removed, so links that drop out of the API stay listed. Later
  observations fill in or correct fields (MD5, size, link status).
- **It is canonical.** Versions sort ascending; archives sort launcher,
  game parts, then voice-over packs; patches sort newest `from` first. The
  JSON uses two-space indents and ends with a newline, so unchanged data
  gives a byte-identical file and every diff is a real change.
- **Dates are UTC days** (`YYYY-MM-DD`) on which a daily API check saw the
  version. They are not release dates (see [Dates](#dates)).

## Top level

```json
{ "versions": [ { "version": "1.3.0", ... }, ... ] }
```

## Version

| Field | Type | Meaning |
| --- | --- | --- |
| `version` | string | Three-part version, such as `5.5.0` |
| `sources` | string[] | Where the record came from (see [Sources](#sources)) |
| `main_first_seen`, `main_last_seen` | date | First and last day a check saw it as the live version |
| `pre_download_first_seen`, `pre_download_last_seen` | date | First and last day a check saw it offered as a pre-download |
| `res_list_url` | string | Scattered-files base URL; `pkg_version` under it lists every game file |
| `full` | Package[] | Full-install archives |
| `patches` | Patch[] | Update archives, one entry per starting version |
| `sophon` | SophonBuild | Full Sophon build |
| `sophon_patch` | SophonPatch | Sophon update build; recorded only while the version is current |

Optional fields are left out when empty.

## Package

| Field | Type | Meaning |
| --- | --- | --- |
| `kind` | string | `game`, `audio` (voice-over pack) or `launcher` |
| `language` | string | Audio only: `zh-cn`, `en-us`, `ja-jp` or `ko-kr` |
| `part` | int | 1-based part of a split game archive; absent for a single archive |
| `url` | string | Download URL on HoYoverse's servers |
| `md5` | string | Lower-case MD5 as published by the API or upstream list |
| `size` | int | Download size in bytes (the served `Content-Length` when known) |
| `decompressed_size` | int | Unpacked size in bytes |
| `http_status` | int | Answer of the last link check (run daily); absent if never checked; anything but 200 is a dead link. Only 200, 404 and 410 are recorded: other answers, such as the occasional 400 the CDN gives for large archives, keep the previous status |
| `md5_size` | int | Size the MD5 was published for, set only when it differs from `size`: the MD5 then likely describes another upload |

## Patch

| Field | Type | Meaning |
| --- | --- | --- |
| `from` | string | Version the update applies to |
| `packages` | Package[] | The update archives (hdiff data) |

## SophonBuild and SophonPatch

| Field | Type | Meaning |
| --- | --- | --- |
| `build_id` | string | Sophon build ID |
| `patch_id` | string | Patch only: Sophon patch ID |
| `diff_tags` | string[] | Patch only: versions the update applies to, newest first |
| `manifests` | SophonManifest[] | One per category: game data, then each voice-over language |

## SophonManifest

| Field | Type | Meaning |
| --- | --- | --- |
| `category_id` | string | Sophon category ID |
| `matching_field` | string | `game` or a voice-over language code |
| `manifest_id`, `manifest_checksum` | string | Manifest file name and its checksum |
| `manifest_url_prefix`, `manifest_url_suffix` | string | Manifest URL = prefix + `/` + `manifest_id` + suffix |
| `data_url_prefix`, `data_url_suffix` | string | Chunk (build) or diff (patch) download prefix; a chunk is at prefix + `/` + name + suffix |
| `stats` | Stats | Build only: totals for this category |
| `stats_by_from` | map of Stats | Patch only: totals per starting version |

Stats hold `compressed_size` (download bytes), `uncompressed_size` (bytes on
disk), `file_count` and `chunk_count`.

## Sources

| Value | Origin |
| --- | --- |
| `upstream-md` | [JRSKelvin/GenshinRepository](https://github.com/JRSKelvin/GenshinRepository) per-version Markdown lists |
| `upstream-legacy` | [JRSKelvin/Genshin-Repository-List](https://github.com/JRSKelvin/Genshin-Repository-List) snapshots in the old `sdk-os-static` format |
| `upstream-hyp` | The same repository's snapshots of `getGamePackages` |
| `api-game-packages` | Live `getGamePackages` (zip archives; frozen at 5.5.0) |
| `api-sophon` | Live `getGameBranches`, `getBuild` and `getPatchBuild` |

## Dates

The `*_seen` dates record what a daily check observed, so they bound the
release date but are not it. The upstream snapshots start on 2023-03-17
(3.5.0 shows that day, though it shipped earlier) and their live-version data
stops at 5.5.0, when HoYoverse froze the zip API. This repository's own
daily check started on 2026-10-26. Versions in between (5.6.0 to 7.0.0) were
recorded through a Sophon backfill and carry no dates.

## Using the data

```sh
# Every full-install URL of 5.5.0
jq -r '.versions[] | select(.version == "5.5.0") | .full[].url' data/versions.json

# The live version (latest main_last_seen)
jq -r '.versions | max_by(.main_last_seen // "") | .version' data/versions.json

# Dead links
jq -r '.versions[] | .version as $v | (.full // [])[], ((.patches // [])[].packages[])
  | select(.http_status != null and .http_status != 200) | "\($v) \(.url)"' data/versions.json
```

From Go, `github.com/tiennm99/genshin-repository/internal/store` loads the
file (`store.Load`) inside this module.

The data is licensed under [CC BY-NC-SA 4.0](../LICENSE-DATA).
