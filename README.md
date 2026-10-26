# Genshin Impact Repository

[![Update](https://github.com/tiennm99/genshin-repository/actions/workflows/update.yml/badge.svg)](https://github.com/tiennm99/genshin-repository/actions/workflows/update.yml)
![Live 7.1.0](https://img.shields.io/badge/live-7.1.0-blue)
![53 versions](https://img.shields.io/badge/versions-53-informational)
![Last checked 2026-10-26](https://img.shields.io/badge/last%20checked-2026--10--26-success)
[![Data license CC BY-NC-SA 4.0](https://img.shields.io/badge/data-CC%20BY--NC--SA%204.0-lightgrey)](LICENSE-DATA)

Direct download links and Sophon build data for every Genshin Impact version on the
global PC client, from 1.3.0 to 7.1.0. A GitHub Actions job checks the official
HoYoverse API every day and keeps this list and [data/versions.json](data/versions.json)
current. Every link points to HoYoverse's own servers; this repository hosts no game
files, and HoYoverse does not endorse it.

## Latest

| Channel | Version | Live in API (UTC) | Links |
| --- | --- | --- | --- |
| Live | **7.1.0** | since 2026-10-26 | [Sophon build and update](docs/versions/7.x.md#710) |
| Pre-download | none | — | — |
| Newest zip archives | **5.5.0** | 2025-03-26 → 2025-05-06 | [Download links](docs/versions/5.x.md#550) |

"Live in API" is when daily checks of the HoYoverse API saw the version as live. It is
not the release date. The checks cover 2023-03-17 to 2025-05-06, and since 2026-10-26;
`—` means no check saw the version. "Game zip" and "Sophon" are the download size of the
game data alone (voice packs are extra); `—` there means no archive or an unknown size.

Every version, by major: [7.x](docs/versions/7.x.md) · [6.x](docs/versions/6.x.md) · [5.x](docs/versions/5.x.md) · [4.x](docs/versions/4.x.md) · [3.x](docs/versions/3.x.md) · [2.x](docs/versions/2.x.md) · [1.x](docs/versions/1.x.md)

## How to download

- **5.5.0 and earlier:** zip archives. Fetch them with a download manager, check
  each against its MD5 and extract with 7-Zip. See the [download guide](docs/download-guide.md).
- **After 5.5.0:** HoYoverse ships the game only as Sophon chunks, so there are no
  direct links. Install it with HoYoPlay or a community launcher. See [Sophon](docs/sophon.md).

## Versions

Newest first. Expand a major version for its index; each version links to its full
tables (file links, MD5s, Sophon manifests) on that major's page.

Columns are explained under [Latest](#latest).

<details>
<summary><b>Version 7.x</b>: 7.0.0 to 7.1.0 (2 versions) · <a href="docs/versions/7.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [7.1.0](docs/versions/7.x.md#710) | since 2026-10-26 | — | 119.50 GiB | live |
| [7.0.0](docs/versions/7.x.md#700) | — | — | 112.98 GiB |  |

</details>

<details>
<summary><b>Version 6.x</b>: 6.0.0 to 6.7.0 (8 versions) · <a href="docs/versions/6.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [6.7.0](docs/versions/6.x.md#670) | — | — | 108.22 GiB |  |
| [6.6.0](docs/versions/6.x.md#660) | — | — | 104.57 GiB |  |
| [6.5.0](docs/versions/6.x.md#650) | — | — | 102.35 GiB |  |
| [6.4.0](docs/versions/6.x.md#640) | — | — | 99.65 GiB |  |
| [6.3.0](docs/versions/6.x.md#630) | — | — | 97.77 GiB |  |
| [6.2.0](docs/versions/6.x.md#620) | — | — | 92.74 GiB |  |
| [6.1.0](docs/versions/6.x.md#610) | — | — | 90.81 GiB |  |
| [6.0.0](docs/versions/6.x.md#600) | — | — | 83.25 GiB |  |

</details>

<details>
<summary><b>Version 5.x</b>: 5.0.0 to 5.8.0 (9 versions) · <a href="docs/versions/5.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [5.8.0](docs/versions/5.x.md#580) | — | — | 77.96 GiB |  |
| [5.7.0](docs/versions/5.x.md#570) | — | — | 75.51 GiB |  |
| [5.6.0](docs/versions/5.x.md#560) | — | — | 74.48 GiB |  |
| [5.5.0](docs/versions/5.x.md#550) | 2025-03-26 → 2025-05-06 | 74.61 GiB | 74.85 GiB |  |
| [5.4.0](docs/versions/5.x.md#540) | 2025-02-12 → 2025-03-25 | 71.71 GiB | 71.93 GiB |  |
| [5.3.0](docs/versions/5.x.md#530) | 2025-01-01 → 2025-02-11 | 70.95 GiB | 71.18 GiB |  |
| [5.2.0](docs/versions/5.x.md#520) | 2024-11-20 → 2024-12-31 | 67.24 GiB | 67.45 GiB |  |
| [5.1.0](docs/versions/5.x.md#510) | 2024-10-09 → 2024-11-19 | 65.94 GiB | 66.15 GiB |  |
| [5.0.0](docs/versions/5.x.md#500) | 2024-08-28 → 2024-10-08 | 63.78 GiB | 63.98 GiB |  |

</details>

<details>
<summary><b>Version 4.x</b>: 4.0.0 to 4.8.0 (10 versions) · <a href="docs/versions/4.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [4.8.0](docs/versions/4.x.md#480) | 2024-07-18 → 2024-08-27 | 58.55 GiB | 58.73 GiB |  |
| [4.7.0](docs/versions/4.x.md#470) | 2024-06-05 → 2024-07-17 | 56.66 GiB | 56.83 GiB |  |
| [4.6.0](docs/versions/4.x.md#460) | 2024-04-24 → 2024-06-04 | 55.28 GiB | 55.46 GiB |  |
| [4.5.0](docs/versions/4.x.md#450) | 2024-03-13 → 2024-04-23 | 70.17 GiB | 70.32 GiB |  |
| [4.4.0](docs/versions/4.x.md#440) | 2024-01-31 → 2024-03-12 | 69.77 GiB | 69.92 GiB |  |
| [4.3.0](docs/versions/4.x.md#430) | 2023-12-20 → 2024-01-30 | 66.08 GiB | 66.22 GiB |  |
| [4.2.0](docs/versions/4.x.md#420) | 2023-11-08 → 2023-12-19 | 64.92 GiB | 65.06 GiB | 1 dead link |
| [4.1.0](docs/versions/4.x.md#410) | 2023-09-27 → 2023-11-07 | 60.90 GiB | — | 1 dead link |
| [4.0.1](docs/versions/4.x.md#401) | 2023-08-23 → 2023-09-26 | 54.99 GiB | — |  |
| [4.0.0](docs/versions/4.x.md#400) | 2023-08-16 → 2023-08-22 | 54.99 GiB | — |  |

</details>

<details>
<summary><b>Version 3.x</b>: 3.0.0 to 3.8.0 (9 versions) · <a href="docs/versions/3.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [3.8.0](docs/versions/3.x.md#380) | 2023-07-05 → 2023-08-15 | 48.28 GiB | — |  |
| [3.7.0](docs/versions/3.x.md#370) | 2023-05-24 → 2023-07-04 | 47.54 GiB | — | MD5 may not match |
| [3.6.0](docs/versions/3.x.md#360) | 2023-04-12 → 2023-05-23 | 46.56 GiB | — | MD5 may not match |
| [3.5.0](docs/versions/3.x.md#350) | 2023-03-17 → 2023-04-11 | 43.61 GiB | — | MD5 may not match |
| [3.4.0](docs/versions/3.x.md#340) | — | — | — | full install dead, 2 dead update links |
| [3.3.0](docs/versions/3.x.md#330) | — | 42.50 GiB | — |  |
| [3.2.0](docs/versions/3.x.md#320) | — | — | — | full install dead, 2 dead update links |
| [3.1.0](docs/versions/3.x.md#310) | — | 37.89 GiB | — |  |
| [3.0.0](docs/versions/3.x.md#300) | — | 34.46 GiB | — |  |

</details>

<details>
<summary><b>Version 2.x</b>: 2.0.0 to 2.8.0 (9 versions) · <a href="docs/versions/2.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [2.8.0](docs/versions/2.x.md#280) | — | 29.62 GiB | — |  |
| [2.7.0](docs/versions/2.x.md#270) | — | 28.47 GiB | — |  |
| [2.6.0](docs/versions/2.x.md#260) | — | 26.77 GiB | — |  |
| [2.5.0](docs/versions/2.x.md#250) | — | 24.55 GiB | — |  |
| [2.4.0](docs/versions/2.x.md#240) | — | 25.80 GiB | — |  |
| [2.3.0](docs/versions/2.x.md#230) | — | 23.08 GiB | — |  |
| [2.2.0](docs/versions/2.x.md#220) | — | — | — | 1 dead link |
| [2.1.0](docs/versions/2.x.md#210) | — | 20.22 GiB | — |  |
| [2.0.0](docs/versions/2.x.md#200) | — | 17.44 GiB | — |  |

</details>

<details>
<summary><b>Version 1.x</b>: 1.3.0 to 1.6.1 (6 versions) · <a href="docs/versions/1.x.md">full page</a></summary>

| Version | Live in API (UTC) | Game zip | Sophon | Notes |
| --- | --- | --- | --- | --- |
| [1.6.1](docs/versions/1.x.md#161) | — | 13.93 GiB | — |  |
| [1.6.0](docs/versions/1.x.md#160) | — | 13.93 GiB | — |  |
| [1.5.0](docs/versions/1.x.md#150) | — | 12.13 GiB | — |  |
| [1.4.0](docs/versions/1.x.md#140) | — | 20.68 GiB | — |  |
| [1.3.2](docs/versions/1.x.md#132) | — | 18.12 GiB | — |  |
| [1.3.0](docs/versions/1.x.md#130) | — | 18.12 GiB | — |  |

</details>

## Data

[data/versions.json](data/versions.json) is the single source for this README and the
pages under [docs/versions/](docs/versions/): every version, archive, checksum, size, link
status and Sophon manifest, merged from the upstream history and the live API. It only
grows, so links that leave the API stay listed. Format and field meanings:
[docs/data.md](docs/data.md).

## Credits

This repository continues the work of:

- [JRSKelvin/GenshinRepository](https://github.com/JRSKelvin/GenshinRepository):
  one link list per version. The 1.3 to 5.5 links here come from it.
- [JRSKelvin/Genshin-Repository-List](https://github.com/JRSKelvin/Genshin-Repository-List):
  a daily GitHub Action that saves the HoYoverse API answer. Its snapshots supply the
  checksums, the sizes and the 2023 to 2025 dates.
- [rapdodge/GenshinRepository](https://github.com/rapdodge/GenshinRepository):
  the original repository, which JRSKelvin continued after it was archived.

This repository keeps both ideas in one place: the history in one data file, and a Go
program that updates it and these pages from the API each day.

## License

Code (the Go fetcher and generator): [Apache License 2.0](LICENSE). The data in
[data/](data/) and the content of this README and [docs/](docs/):
[CC BY-NC-SA 4.0](LICENSE-DATA), the license of the upstream lists they build on.

---

Last checked: 2026-10-26 (UTC). Generated from [data/versions.json](data/versions.json) by
`go run ./cmd/fetch`; do not edit by hand.
