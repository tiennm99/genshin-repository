# Genshin Impact Repository

Direct download links and Sophon build data for every Genshin Impact
version on the global PC client, from 1.3.0 to the live version. Use them
with a download manager instead of the launcher. A GitHub Actions job
checks the official HoYoverse API every day and updates this list.

Every link points to HoYoverse's own servers. This repository hosts no game
files. HoYoverse does not endorse it, and the two are not affiliated.

## Credits

This repository continues the work of:

- [JRSKelvin/GenshinRepository](https://github.com/JRSKelvin/GenshinRepository):
  one link list per version. The 1.3 to 5.5 links here come from it.
- [JRSKelvin/Genshin-Repository-List](https://github.com/JRSKelvin/Genshin-Repository-List):
  a daily GitHub Action that saves the HoYoverse API answer. Its snapshots
  from 2023 to 2026 supply the checksums, sizes and release dates.
- [rapdodge/GenshinRepository](https://github.com/rapdodge/GenshinRepository):
  the original repository, which JRSKelvin continued after it was archived.

The difference: this repository keeps both ideas in one place. It stores the
history in [data/versions.json](data/versions.json), and a Go program updates
it and this README from the API each day.

## Latest

| | Version |
| --- | --- |
| Live version | 7.1.0 |
| Pre-download | none |
| Newest version with zip archives | 5.5.0 |

## How to download 5.5.0 and earlier

HoYoverse published these versions as zip archives. Each version below has:

- **Full install:** the game archive, split into numbered parts
  (`.zip.001`, `.zip.002`, ...) for newer versions, plus one archive per
  voice-over language.
- **Update from X:** smaller archives that update an existing install of
  version X. They hold hdiff patch files that a launcher or patch tool
  applies to the old files, not a full game. The upstream lists call
  them pre-install files.

Steps for a full install:

1. Copy the links into a download manager (IDM, FDM, XDM, uGet,
   JDownloader, or `aria2c -x 16 -s 16 <url>`). Download every game part and
   the voice-over packs you want into one folder.
2. Check each file against the MD5 in the table: `certutil -hashfile <file> MD5`
   on Windows, `md5sum <file>` on Linux and macOS. Where the table says the
   MD5 was published for another size, the served file changed after the
   API listed it, and its MD5 is unknown.
3. Extract the game archive. For a split archive, open the `.zip.001` file
   with 7-Zip, which joins the parts. Extract the voice-over packs into the
   same folder.

Some old archives are no longer served; they are marked as dead links.

### Single files: ScatteredFiles

Releases that list a scattered-files URL (`res_list_url`) offer one more
route. Below that URL, `pkg_version` is a JSON-lines list of every game
file (`remoteName`, `md5`, `fileSize`). Each file can be downloaded on its own from
`<res_list_url>/<remoteName>`, which helps to repair a single file without
the full archive. HoYoverse may remove these URLs at any time.

## How to download 5.6 and later

HoYoverse stopped publishing full-game and patch zip archives after
**5.5.0**. Since 5.6, Genshin Impact (global, `hk4e_global`) ships only
through **Sophon**. In Sophon the game is split into small compressed
chunks, and a manifest lists them. Downloaders fetch the manifest and the
chunks, then rebuild the game files on your disk. There are no direct
download links to share for 5.6 and later. The launchers below install the
current version from the API. Each version entry below lists the Sophon
manifests, the chunk URL prefix, the sizes and the diff tags. They are for
reference and for tools or scripts that rebuild files from a Sophon
manifest. This repository does not claim that older builds can be installed
with these launchers.

### Official route: HoYoPlay (Windows)

1. Download HoYoPlay from <https://hoyoplay.hoyoverse.com/> (also linked
   from <https://genshin.hoyoverse.com/>).
2. Install it, select **Genshin Impact**, choose an install folder and click
   **Download**. HoYoPlay fetches the current version through Sophon.
3. Updates and pre-downloads (pre-install) appear in HoYoPlay when
   HoYoverse publishes them.
4. To keep an existing install, use **Locate game** on its folder instead of
   downloading again. HoYoPlay checks the files and fetches only what is
   missing or changed.

### Collapse Launcher (Windows)

Open-source launcher for HoYoverse games. It installs, updates,
pre-downloads and repairs Genshin through Sophon, using its own
`Hi3Helper.Sophon` library.

1. Download `CollapseLauncher-stable-Setup.exe` or the portable `.7z` from
   the [releases page](https://github.com/CollapseLauncher/Collapse/releases).
   It needs Windows 10 1809 or later.
2. In the region list, pick **Genshin Impact → Global**.
3. Click **Install** for a new install, or point it at an existing game
   folder. Updates and pre-downloads then go through Sophon.

### Twintail Launcher (Windows and Linux)

Open-source launcher with Genshin Impact support. Its download backend
(`fischl-rs`) has a Sophon mode for full installs, updates (diffs),
preloads and repairs.

1. Download the Windows installer (`*_x64-setup.exe`) or the Linux `.deb`
   (amd64 or arm64) from the
   [releases page](https://github.com/TwintailTeam/TwintailLauncher/releases).
   Other install methods, such as a Nix flake, are listed in its README.
2. Add **Genshin Impact**, choose an install path and start the download.

### An Anime Game Launcher (Linux)

Linux launcher that runs the game through Wine. It installs and updates
through Sophon.

1. Install it with the method for your distribution. Flatpak
   (`moe.launcher.an-anime-game-launcher` on Flathub) is the officially
   supported one. See the
   [installation wiki](https://github.com/an-anime-team/an-anime-game-launcher/wiki/Installation).
2. On first launch, follow the setup steps (Wine and DXVK components, game
   folder). Then click **Download** to install the game, or **Update** to
   update it.

### Sophon Downloader (Windows, files only)

A stand-alone GUI that downloads game files from Sophon manifests. It is not
a launcher. It covers `hk4e_global`, including full builds and patch builds.

1. Download `SophonDownloader.7z` from the
   [releases page](https://github.com/GesthosNetwork/Sophon-Downloader/releases)
   and extract it. It runs on Windows x64 and needs a recent .NET runtime;
   see its README.
2. Select Genshin Impact (global), the version or branch and the content
   (game and voice packs), then download to a folder.

### What the Sophon data below means

- **Manifest:** a compressed file that lists every game file and the chunks
  it is made of. Its URL is the manifest URL prefix plus the manifest ID.
- **Chunk URL prefix:** each chunk named in a manifest downloads from
  `<prefix>/<chunk name>`.
- **Download / Installed:** the compressed chunk total and the size on disk.
- **Sophon update ... (from X):** diff data that updates an install of
  version X (the diff tags). The Sophon API serves update data for the
  current version only, so older versions list the full build alone.
- The Sophon API also answers for several zip-era 4.x and 5.x versions.
  This list records Sophon builds from 5.6.0 on, where Sophon is the only
  way to download.

> The tools above are independent community projects. HoYoverse does not
> endorse them, and this repository has no ties to them. Third-party
> launchers may break when the game or API changes. Use them at your own
> risk.

## Versions

Newest first. Expand a major version to see its releases and links. Dates are the days a
version was seen in the API: up to 5.5.0 from the upstream daily snapshots, and from
7.1.0 on from this repository's daily check.

<details>
<summary><b>Version 7.x</b>: 7.0.0 to 7.1.0 (2 versions)</summary>

Jump to: [7.1.0](#710-live) · [7.0.0](#700)

### 7.1.0 (live)

Seen in the API as the live version since 2026-10-26 (UTC).

#### Sophon full build `sMXGW2ll3Fuu`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 119.50 GiB | 122.92 GiB | 2744 | 113610 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/manifest_671e1a92a6cf53ff_8d4dfb34d2ee2cf64aae45a9b1ecf58d) |
| Chinese voice-over | 13.13 GiB | 13.13 GiB | 177 | 12069 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/manifest_bff739d9bcacde9c_cacb6fd0b9768a5ef674666981aa00cc) |
| English voice-over | 17.05 GiB | 17.05 GiB | 177 | 15565 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/manifest_3e2d1468c5bd2260_1868f6fa795f58dc0439b7f83e44131d) |
| Japanese voice-over | 15.47 GiB | 15.48 GiB | 177 | 14154 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/manifest_1239c26765d1e8b1_447979cbd58a396ca944ab8c35b0eb64) |
| Korean voice-over | 13.78 GiB | 13.79 GiB | 177 | 12623 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/manifest_e76a11262e8cae1a_defe2edff1e45c10584cbfa50a0429b0) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/sMXGW2ll3Fuu`

#### Sophon update `t663grrK4Of4` (from 7.0.0, 6.7.0)

| Content | From 7.0.0 | From 6.7.0 | Manifest | Diff URL prefix |
| --- | --- | --- | --- | --- |
| Game | 9.31 GiB | 14.14 GiB | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/t663grrK4Of4/manifest_dc3c7027c0903da8_7ed824206b7559bbbb57ab2197c97b12) | `https://autopatchhk.yuanshen.com/client_app/sophon/diffs/cxhpq4g4rgg0/t663grrK4Of4/10016` |
| Chinese voice-over | 485.15 MiB | 976.96 MiB | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/t663grrK4Of4/manifest_b9e34d1fd9b1dbe5_d05a7eba567d59af9d44a05151ced4f0) | `https://autopatchhk.yuanshen.com/client_app/sophon/diffs/cxhpq4g4rgg0/t663grrK4Of4/10017` |
| English voice-over | 600.10 MiB | 1.16 GiB | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/t663grrK4Of4/manifest_a530130ca8a42f41_38bcdb95b27cf25b117568d4c8e7f870) | `https://autopatchhk.yuanshen.com/client_app/sophon/diffs/cxhpq4g4rgg0/t663grrK4Of4/10018` |
| Japanese voice-over | 555.28 MiB | 1.06 GiB | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/t663grrK4Of4/manifest_6d66db98af190469_85ee3050184980dd4c602e8204176a63) | `https://autopatchhk.yuanshen.com/client_app/sophon/diffs/cxhpq4g4rgg0/t663grrK4Of4/10020` |
| Korean voice-over | 513.35 MiB | 1019.71 MiB | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/t663grrK4Of4/manifest_c273254535796e52_b6170f9c52ca82ae3c43ef605268bda7) | `https://autopatchhk.yuanshen.com/client_app/sophon/diffs/cxhpq4g4rgg0/t663grrK4Of4/10019` |

Sizes are the download size of the diff data for each starting version.

### 7.0.0

#### Sophon full build `3XnPviMfoILi`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 112.98 GiB | 116.25 GiB | 2673 | 107500 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3XnPviMfoILi/manifest_2efe39d995f47170_29ad2bc6e42fe4af2086ae58c7e81b33) |
| Chinese voice-over | 12.65 GiB | 12.65 GiB | 175 | 11644 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3XnPviMfoILi/manifest_ff518d0444259c40_db9ab48c38e34403b8515e4a91553eb3) |
| English voice-over | 16.46 GiB | 16.46 GiB | 175 | 15053 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3XnPviMfoILi/manifest_e818d7fa91947bc8_c740684b3e4f74ba0f4863e90698c09d) |
| Japanese voice-over | 14.93 GiB | 14.94 GiB | 175 | 13660 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3XnPviMfoILi/manifest_b1991f905278e46b_9823c7c67957005cdb9b9538a4b36d4c) |
| Korean voice-over | 13.28 GiB | 13.29 GiB | 175 | 12158 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3XnPviMfoILi/manifest_195b6fbb7d7251d9_f83cb2ccfc3f4f4186b6a9d9282f6080) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/3XnPviMfoILi`

Only the full build is recorded: the Sophon API serves update data for the current version only.

</details>

<details>
<summary><b>Version 6.x</b>: 6.0.0 to 6.7.0 (8 versions)</summary>

Jump to: [6.7.0](#670) · [6.6.0](#660) · [6.5.0](#650) · [6.4.0](#640) · [6.3.0](#630) · [6.2.0](#620) · [6.1.0](#610) · [6.0.0](#600)

### 6.7.0

#### Sophon full build `E2zvp18oyzw6`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 108.22 GiB | 111.31 GiB | 2579 | 102928 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/E2zvp18oyzw6/manifest_ccce2f7066358067_b659db7d40b8ac7311addf76f7edf5b5) |
| Chinese voice-over | 12.17 GiB | 12.17 GiB | 173 | 11184 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/E2zvp18oyzw6/manifest_b9aeedd170da756b_f2dd191b0a0e9ec7a8fbcb99f4b4b2bc) |
| English voice-over | 15.89 GiB | 15.89 GiB | 173 | 14544 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/E2zvp18oyzw6/manifest_077d08137b5b7baf_4f27938035bf311b49f029448501fdc4) |
| Japanese voice-over | 14.41 GiB | 14.42 GiB | 173 | 13202 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/E2zvp18oyzw6/manifest_827f5588136bb41d_317b9fc7ddccbcf73a8d35720dbfd0a4) |
| Korean voice-over | 12.79 GiB | 12.79 GiB | 173 | 11706 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/E2zvp18oyzw6/manifest_4c0dc71303355859_2f6d9057677c5fdd06d018605812beb2) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/E2zvp18oyzw6`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.6.0

#### Sophon full build `L7CJe4FP1z71`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 104.57 GiB | 107.59 GiB | 2508 | 99486 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/L7CJe4FP1z71/manifest_bd262679a7e69d13_4d78eff35c7ddc2f37df853f06a1d58e) |
| Chinese voice-over | 11.89 GiB | 11.89 GiB | 172 | 10930 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/L7CJe4FP1z71/manifest_65e541fd82487e82_01149873f0fbe0c7fb536e4d4bda2633) |
| English voice-over | 15.54 GiB | 15.54 GiB | 172 | 14206 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/L7CJe4FP1z71/manifest_2a34a9aa3ca50950_5f8cfafb3250c0ff7643e6b1cb1aaa37) |
| Japanese voice-over | 14.06 GiB | 14.06 GiB | 172 | 12880 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/L7CJe4FP1z71/manifest_dcb4393fa2a81fc6_15a2fe08d0fd3c9fbbf953caff601ac2) |
| Korean voice-over | 12.48 GiB | 12.48 GiB | 172 | 11436 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/L7CJe4FP1z71/manifest_a6a5b050f9ede572_d3a4f3366045e665d9e7ad5906521ff7) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/L7CJe4FP1z71`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.5.0

#### Sophon full build `3clmr46VnNoZ`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 102.35 GiB | 105.31 GiB | 2468 | 97332 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3clmr46VnNoZ/manifest_308b3b2a4f03ebe4_6d705e93c11f945830cc00acb0128633) |
| Chinese voice-over | 11.40 GiB | 11.40 GiB | 170 | 10488 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3clmr46VnNoZ/manifest_462c9ebea79535eb_c43098e42e9964a9ade1873be2935aa8) |
| English voice-over | 14.97 GiB | 14.97 GiB | 170 | 13687 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3clmr46VnNoZ/manifest_a9e6e84b635a6b41_56a784c2845309c71b63eeec1eaebfe2) |
| Japanese voice-over | 13.50 GiB | 13.50 GiB | 170 | 12366 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3clmr46VnNoZ/manifest_fb412d2cab306ab9_08f4d50a49c92f57bb2a27d00dabb599) |
| Korean voice-over | 11.97 GiB | 11.98 GiB | 170 | 10964 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3clmr46VnNoZ/manifest_fee64a1a45f0a5ea_326ea9131c5fd7e5892de3f9bbe21eed) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/3clmr46VnNoZ`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.4.0

#### Sophon full build `jJ6hfpe1NGBu`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 99.65 GiB | 102.53 GiB | 2566 | 94863 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/jJ6hfpe1NGBu/manifest_1db0d950b8106744_1b75295180adc9458e6c61c1d14a90e6) |
| Chinese voice-over | 11.20 GiB | 11.20 GiB | 169 | 10316 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/jJ6hfpe1NGBu/manifest_2b64f503ce986cf8_54519c80d7d198c3484e3da42e19fa95) |
| English voice-over | 14.62 GiB | 14.63 GiB | 169 | 13383 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/jJ6hfpe1NGBu/manifest_1da14cf59d915ca5_e96b1e602c36ddeb8ecbe7a018f6810f) |
| Japanese voice-over | 13.27 GiB | 13.27 GiB | 169 | 12155 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/jJ6hfpe1NGBu/manifest_f8ec9c90094c2bb9_55e7d16dcf9bb91e5191b62ac9ef3969) |
| Korean voice-over | 11.76 GiB | 11.76 GiB | 169 | 10776 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/jJ6hfpe1NGBu/manifest_92554b4a34378a9c_3b096026e6f2842650da83c1f469d62e) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/jJ6hfpe1NGBu`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.3.0

#### Sophon full build `y7hlOmvUTl68`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 97.77 GiB | 100.61 GiB | 2532 | 93115 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/y7hlOmvUTl68/manifest_c2e2c8c6771c3be3_175784e700092390afc478e4a3b9809d) |
| Chinese voice-over | 10.89 GiB | 10.89 GiB | 168 | 10041 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/y7hlOmvUTl68/manifest_84f7762dd9d1f3c9_24127b75052f9132665c1e2d12b0fd42) |
| English voice-over | 14.22 GiB | 14.23 GiB | 168 | 13030 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/y7hlOmvUTl68/manifest_413eb26e4527e2f3_727bc36389863ddee65db2ab728e57e3) |
| Japanese voice-over | 12.94 GiB | 12.94 GiB | 168 | 11855 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/y7hlOmvUTl68/manifest_1c26ed214ccfe9cc_97019f966fb96a32936603b5fd834817) |
| Korean voice-over | 11.44 GiB | 11.45 GiB | 168 | 10465 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/y7hlOmvUTl68/manifest_e12615a9f6d2ebcf_11f125940e9e9f8dd625c5d2ec062f90) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/y7hlOmvUTl68`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.2.0

#### Sophon full build `b1yX8g30suHF`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 92.74 GiB | 95.46 GiB | 2458 | 88336 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/b1yX8g30suHF/manifest_ad377d2ea8aeeee6_bde7cb1239d3ef40862a8c0481a551a8) |
| Chinese voice-over | 11.11 GiB | 11.11 GiB | 166 | 10237 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/b1yX8g30suHF/manifest_74794db7102b8e56_b362350320c6f8af6a20b71655cadbfa) |
| English voice-over | 14.32 GiB | 14.33 GiB | 166 | 13070 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/b1yX8g30suHF/manifest_084d98f2b0b73417_f76f2698e719dbc07b5c3b689f0c5148) |
| Japanese voice-over | 13.20 GiB | 13.21 GiB | 166 | 12094 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/b1yX8g30suHF/manifest_e6602b44c4d9a827_01aedf0468eb7ff42026a26d195e6447) |
| Korean voice-over | 11.69 GiB | 11.69 GiB | 166 | 10693 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/b1yX8g30suHF/manifest_bdb81d5f5a7372a0_c7aab368be02b854542ee049e250ccdf) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/b1yX8g30suHF`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.1.0

#### Sophon full build `xCk1JZqwnbo3`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 90.81 GiB | 93.49 GiB | 2418 | 86567 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/xCk1JZqwnbo3/manifest_fd962d2133c651a3_0728111bf8628aec07636bc532032aa0) |
| Chinese voice-over | 15.34 GiB | 17.79 GiB | 164 | 16752 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/xCk1JZqwnbo3/manifest_a3f1a44279c8e6fd_6a793b6b444a51232f63ca4c56e7ed28) |
| English voice-over | 18.09 GiB | 20.43 GiB | 164 | 18461 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/xCk1JZqwnbo3/manifest_e28b12d0beb23ad8_991929bfee8bbb8db76862995a446ae5) |
| Japanese voice-over | 17.85 GiB | 23.06 GiB | 164 | 20508 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/xCk1JZqwnbo3/manifest_99c25d64b3ff45af_202931be1b2dd13b0728635c4a8d4cd5) |
| Korean voice-over | 14.94 GiB | 17.60 GiB | 164 | 15732 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/xCk1JZqwnbo3/manifest_2136f357ef60e145_897609d8f2e19629f6acfd46ae663ab4) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/xCk1JZqwnbo3`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 6.0.0

#### Sophon full build `R6VKQTMPUUBa`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 83.25 GiB | 85.41 GiB | 2298 | 79210 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/R6VKQTMPUUBa/manifest_e64ca10c417ae850_028ff30ff105d093df1df15a46fed16a) |
| Chinese voice-over | 14.84 GiB | 17.20 GiB | 162 | 16176 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/R6VKQTMPUUBa/manifest_a1cb4043bc36ff61_96cc45be5678146bc382ba6663163b4d) |
| English voice-over | 17.50 GiB | 19.71 GiB | 162 | 17821 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/R6VKQTMPUUBa/manifest_7e97c4876f493b2a_9ed81c934600a7c1d06298cc2e6a5a39) |
| Japanese voice-over | 17.29 GiB | 22.31 GiB | 162 | 19879 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/R6VKQTMPUUBa/manifest_0c6f0d6a63acad88_6bd4c453ac69aeee447b37ea1b764c10) |
| Korean voice-over | 14.47 GiB | 17.02 GiB | 162 | 15247 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/R6VKQTMPUUBa/manifest_da74d3a2bd8ec17b_8e0990bede9c5d24136ed9fc195a479b) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/R6VKQTMPUUBa`

Only the full build is recorded: the Sophon API serves update data for the current version only.

</details>

<details>
<summary><b>Version 5.x</b>: 5.0.0 to 5.8.0 (9 versions)</summary>

Jump to: [5.8.0](#580) · [5.7.0](#570) · [5.6.0](#560) · [5.5.0](#550) · [5.4.0](#540) · [5.3.0](#530) · [5.2.0](#520) · [5.1.0](#510) · [5.0.0](#500)

### 5.8.0

#### Sophon full build `d6lo39gA1LIA`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 77.96 GiB | 79.95 GiB | 2216 | 74178 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/d6lo39gA1LIA/manifest_5c3b9426a2f209d1_3934c9f0cc2b5ef1b468bad832be87bb) |
| Chinese voice-over | 14.15 GiB | 16.37 GiB | 160 | 15395 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/d6lo39gA1LIA/manifest_25ade311d90d6dc5_95ca3d1f6d10d9cacf9c0bcde55a8901) |
| English voice-over | 16.64 GiB | 18.69 GiB | 160 | 16887 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/d6lo39gA1LIA/manifest_920dc481d2415f49_69ed66dd8bd7b058fffb74f53e5283f4) |
| Japanese voice-over | 16.43 GiB | 21.21 GiB | 160 | 18894 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/d6lo39gA1LIA/manifest_c6b2b260b6d6dc9b_feb083d29890cff551f7fc3ba1ee6c5b) |
| Korean voice-over | 13.77 GiB | 16.17 GiB | 160 | 14475 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/d6lo39gA1LIA/manifest_4dd5c2ad064d15b1_8398182c59ee5cbe0f1a6a1b7b149986) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/d6lo39gA1LIA`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 5.7.0

#### Sophon full build `3j90WyoOpfjt`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 75.51 GiB | 77.46 GiB | 2178 | 71843 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3j90WyoOpfjt/manifest_b39180b4ad2cabd3_5bd9ec245492bc76dc13c0204b43cce8) |
| Chinese voice-over | 13.80 GiB | 15.95 GiB | 158 | 14990 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3j90WyoOpfjt/manifest_ae64dd00dac4975a_f675baffb98d7ca30faa42348a228ad5) |
| English voice-over | 16.21 GiB | 18.17 GiB | 158 | 16420 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3j90WyoOpfjt/manifest_498715b02da32adc_914c1bb08c35d2b628d7dff2494f3153) |
| Japanese voice-over | 16.03 GiB | 20.70 GiB | 158 | 18448 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3j90WyoOpfjt/manifest_ac61814bcec0eda7_c09fa48e336ce695f0c649ce4e4c235e) |
| Korean voice-over | 13.43 GiB | 15.75 GiB | 158 | 14101 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/3j90WyoOpfjt/manifest_bb55ed3f9194105c_284609abfaa50f7ab6b235b941fb4cdf) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/3j90WyoOpfjt`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 5.6.0

#### Sophon full build `q3h361jUEuu0`

| Content | Download | Installed | Files | Chunks | Manifest |
| --- | --- | --- | --- | --- | --- |
| Game | 74.48 GiB | 76.41 GiB | 2153 | 70917 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/q3h361jUEuu0/manifest_6194a90dbfdae455_f2f91f7cf5f869009f4816d8489f66ca) |
| Chinese voice-over | 13.44 GiB | 15.52 GiB | 163 | 14608 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/q3h361jUEuu0/manifest_a051fcb00ffd2801_28580ffe5eec0403ece42760eedf7845) |
| English voice-over | 15.86 GiB | 17.68 GiB | 163 | 16017 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/q3h361jUEuu0/manifest_fa8dc928000f7107_7637d81dc91dcbef3ee7673d38a2366c) |
| Japanese voice-over | 15.62 GiB | 20.14 GiB | 163 | 17974 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/q3h361jUEuu0/manifest_c534564a3ba8c422_cdaf588e766eaef62b37270fb40dee96) |
| Korean voice-over | 13.07 GiB | 15.32 GiB | 163 | 13743 | [manifest](https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/q3h361jUEuu0/manifest_ca6057d20a260850_fce8310076eff284b2c3f8701e497f45) |

Chunk URL prefix: `https://autopatchhk.yuanshen.com/client_app/sophon/chunks/cxhpq4g4rgg0/q3h361jUEuu0`

Only the full build is recorded: the Sophon API serves update data for the current version only.

### 5.5.0

Seen in the API as the live version from 2025-03-26 to 2025-05-06, pre-download from 2025-03-24 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.5.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.001) | Game, part 1 of 8 | 10.00 GiB | `e65a66073f10325a384c4d8f01199e21` |
| [GenshinImpact_5.5.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.002) | Game, part 2 of 8 | 10.00 GiB | `e262578cc6791c927a77c2a784da84dd` |
| [GenshinImpact_5.5.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.003) | Game, part 3 of 8 | 10.00 GiB | `e9dd5ce69e4ee8025c3e3319da8fb95d` |
| [GenshinImpact_5.5.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.004) | Game, part 4 of 8 | 10.00 GiB | `c3431ba39643f99bab5af37f03c81139` |
| [GenshinImpact_5.5.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.005) | Game, part 5 of 8 | 10.00 GiB | `57e55567b5292f81e3da26636da3124f` |
| [GenshinImpact_5.5.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.006) | Game, part 6 of 8 | 10.00 GiB | `1cbc02e908b8a1ad7d421f23e9fe689a` |
| [GenshinImpact_5.5.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.007) | Game, part 7 of 8 | 10.00 GiB | `ba292fa02c2ce8887fbef86ae45f6100` |
| [GenshinImpact_5.5.0.zip.008](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/GenshinImpact_5.5.0.zip.008) | Game, part 8 of 8 | 4.61 GiB | `276e433ab28c35748c6deef72d28d3c2` |
| [Audio_Chinese_5.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/Audio_Chinese_5.5.0.zip) | Chinese voice-over | 15.02 GiB | `724eb66fac6c39199faecf8117b036a5` |
| [Audio_English(US)_5.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/Audio_English%28US%29_5.5.0.zip) | English voice-over | 17.15 GiB | `e14821e9236844f1fae27396f30f094a` |
| [Audio_Japanese_5.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/Audio_Japanese_5.5.0.zip) | Japanese voice-over | 19.50 GiB | `21117c63e8a55abe0f9a704528bcdb1c` |
| [Audio_Korean_5.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/Audio_Korean_5.5.0.zip) | Korean voice-over | 14.80 GiB | `c52cd0575dfe126a6604fe2b43650dae` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250314110016_HcIQuDGRmsbByeAE/ScatteredFiles` (file list in `pkg_version`).

#### Update from 5.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.4.0_5.5.0_hdiff_IlvHovyEdpXnwiCH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.4.0_5.5.0_hdiff_IlvHovyEdpXnwiCH.zip) | Game | 15.20 GiB | `351503a9da9d8c657c7c33e544dd48d8` |
| [audio_zh-cn_5.4.0_5.5.0_hdiff_DuzyGAJXDnTEDcDG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.4.0_5.5.0_hdiff_DuzyGAJXDnTEDcDG.zip) | Chinese voice-over | 281.48 MiB | `20f67f06f746f96cd6176eff9ae2491d` |
| [audio_en-us_5.4.0_5.5.0_hdiff_HveRbpmNrNejbGYL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.4.0_5.5.0_hdiff_HveRbpmNrNejbGYL.zip) | English voice-over | 332.17 MiB | `9a922acc8f68bbcab55c06847dcbe37f` |
| [audio_ja-jp_5.4.0_5.5.0_hdiff_RCsszbNmxZciMvkK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.4.0_5.5.0_hdiff_RCsszbNmxZciMvkK.zip) | Japanese voice-over | 326.91 MiB | `471020c4184331a9159b29ed3107804d` |
| [audio_ko-kr_5.4.0_5.5.0_hdiff_gKGHXRscunFdPzdH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.4.0_5.5.0_hdiff_gKGHXRscunFdPzdH.zip) | Korean voice-over | 274.61 MiB | `afca3b2d11fa6068b59821ec0dfa0948` |

#### Update from 5.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.3.0_5.5.0_hdiff_hLjncNZamXTtcKKo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.3.0_5.5.0_hdiff_hLjncNZamXTtcKKo.zip) | Game | 22.01 GiB | `9b673560692022ba2147c9557d555e64` |
| [audio_zh-cn_5.3.0_5.5.0_hdiff_IjUYCPouJjYnAVzM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.3.0_5.5.0_hdiff_IjUYCPouJjYnAVzM.zip) | Chinese voice-over | 705.90 MiB | `af73c0e343da127b27afb7cee8d5f828` |
| [audio_en-us_5.3.0_5.5.0_hdiff_hyPqxqduaCinaiJo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.3.0_5.5.0_hdiff_hyPqxqduaCinaiJo.zip) | English voice-over | 686.25 MiB | `c09ff2a7f45a6fdd58046ea7f16db4a3` |
| [audio_ja-jp_5.3.0_5.5.0_hdiff_beYoljTPLiGPgFOz.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.3.0_5.5.0_hdiff_beYoljTPLiGPgFOz.zip) | Japanese voice-over | 810.74 MiB | `28f2866151a4f333bfa27db6a415ca73` |
| [audio_ko-kr_5.3.0_5.5.0_hdiff_IuoBNsASjDzeqgNY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.3.0_5.5.0_hdiff_IuoBNsASjDzeqgNY.zip) | Korean voice-over | 682.64 MiB | `a62129453431421ddb4fd6d4ee3f9413` |

### 5.4.0

Seen in the API as the live version from 2025-02-12 to 2025-03-25, pre-download from 2025-02-10 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.4.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.001) | Game, part 1 of 8 | 10.00 GiB | `e25ba37b102aeea9c32f0f1d38ca1318` |
| [GenshinImpact_5.4.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.002) | Game, part 2 of 8 | 10.00 GiB | `454e205e080d0c91f99bd5968ff5baec` |
| [GenshinImpact_5.4.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.003) | Game, part 3 of 8 | 10.00 GiB | `ef73d348e464ce4e43b26563a017a5c1` |
| [GenshinImpact_5.4.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.004) | Game, part 4 of 8 | 10.00 GiB | `d76b1b2d12e5c643f09dc5a9f54e238b` |
| [GenshinImpact_5.4.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.005) | Game, part 5 of 8 | 10.00 GiB | `7b8da4761346066d84e8dd930a3a3733` |
| [GenshinImpact_5.4.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.006) | Game, part 6 of 8 | 10.00 GiB | `d1d8f6e36d421e0da6e1d7b103b1a36b` |
| [GenshinImpact_5.4.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.007) | Game, part 7 of 8 | 10.00 GiB | `894d8ee64fe7dc183c3f0bd468604e89` |
| [GenshinImpact_5.4.0.zip.008](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/GenshinImpact_5.4.0.zip.008) | Game, part 8 of 8 | 1.71 GiB | `0cba1f6ffdc0170d6ed53be045b1038b` |
| [Audio_Chinese_5.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/Audio_Chinese_5.4.0.zip) | Chinese voice-over | 14.70 GiB | `05707a6a95f0812bae18a2ef5b0b164e` |
| [Audio_English(US)_5.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/Audio_English%28US%29_5.4.0.zip) | English voice-over | 16.76 GiB | `11789d4274d0cbf9d2b59694223ecd0d` |
| [Audio_Japanese_5.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/Audio_Japanese_5.4.0.zip) | Japanese voice-over | 19.08 GiB | `7c989a875bc30a6f9e06e70ac61819f2` |
| [Audio_Korean_5.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/Audio_Korean_5.4.0.zip) | Korean voice-over | 14.48 GiB | `519f71c38b184ba2a6fe9994ba42dc79` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20250125201352_EiPmYLKVptWspsHf/ScatteredFiles` (file list in `pkg_version`).

#### Update from 5.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.3.0_5.4.0_hdiff_vlIbVzLpRIFHZcym.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.3.0_5.4.0_hdiff_vlIbVzLpRIFHZcym.zip) | Game | 16.75 GiB | `87b6570843712191de2c1fd2d0226ea8` |
| [audio_zh-cn_5.3.0_5.4.0_hdiff_juLXEdublvTHwOPL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.3.0_5.4.0_hdiff_juLXEdublvTHwOPL.zip) | Chinese voice-over | 425.92 MiB | `5d6a61771530949e7b4c9e037592b917` |
| [audio_en-us_5.3.0_5.4.0_hdiff_rUjNMXDApbBRrLtB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.3.0_5.4.0_hdiff_rUjNMXDApbBRrLtB.zip) | English voice-over | 354.91 MiB | `78b3d7d80d558ff37151e10e562ee8cb` |
| [audio_ja-jp_5.3.0_5.4.0_hdiff_GYvLQJtloHewyhAd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.3.0_5.4.0_hdiff_GYvLQJtloHewyhAd.zip) | Japanese voice-over | 485.99 MiB | `5cbfe6c148a280bdfd8bea4054667587` |
| [audio_ko-kr_5.3.0_5.4.0_hdiff_bHMaFGHIWezvVoGL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.3.0_5.4.0_hdiff_bHMaFGHIWezvVoGL.zip) | Korean voice-over | 409.76 MiB | `3c7bc585055e508af6ac15aa768b16f9` |

#### Update from 5.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.2.0_5.4.0_hdiff_nhbCdqtBokgjnsTo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.2.0_5.4.0_hdiff_nhbCdqtBokgjnsTo.zip) | Game | 23.07 GiB | `5212f460443e46bb679d7fa97286fc46` |
| [audio_zh-cn_5.2.0_5.4.0_hdiff_qvUdBSGRwBauLmne.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.2.0_5.4.0_hdiff_qvUdBSGRwBauLmne.zip) | Chinese voice-over | 1.06 GiB | `a897669ad4f0627bd3bf7fdf4c6e2c76` |
| [audio_en-us_5.2.0_5.4.0_hdiff_PzaSwHqqvaxDVWHd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.2.0_5.4.0_hdiff_PzaSwHqqvaxDVWHd.zip) | English voice-over | 1007.04 MiB | `abcb7389dcde8e19b086dbc19b597101` |
| [audio_ja-jp_5.2.0_5.4.0_hdiff_JaDVQGJbqLlQwrjN.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.2.0_5.4.0_hdiff_JaDVQGJbqLlQwrjN.zip) | Japanese voice-over | 1.26 GiB | `dff50933e6d363879ef5334fcf895f06` |
| [audio_ko-kr_5.2.0_5.4.0_hdiff_glXIlRCjmGCRUFmu.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.2.0_5.4.0_hdiff_glXIlRCjmGCRUFmu.zip) | Korean voice-over | 1.09 GiB | `830a4fa257297a01e9758989cec581dd` |

### 5.3.0

Seen in the API as the live version from 2025-01-01 to 2025-02-11, pre-download from 2024-12-30 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.3.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.001) | Game, part 1 of 8 | 10.00 GiB | `d7ea7d49334e03e590db3f047cd9ea88` |
| [GenshinImpact_5.3.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.002) | Game, part 2 of 8 | 10.00 GiB | `b4178034c1d09e889e43fd76b3fb4d3c` |
| [GenshinImpact_5.3.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.003) | Game, part 3 of 8 | 10.00 GiB | `43b70975fcb957abaaaf7d940969679a` |
| [GenshinImpact_5.3.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.004) | Game, part 4 of 8 | 10.00 GiB | `d734b1edeb1b2b0d47d4d4bab7af6778` |
| [GenshinImpact_5.3.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.005) | Game, part 5 of 8 | 10.00 GiB | `95abe987ff924c21f3e5085492448760` |
| [GenshinImpact_5.3.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.006) | Game, part 6 of 8 | 10.00 GiB | `492510ae74ae8ac696ee59b4e831d039` |
| [GenshinImpact_5.3.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.007) | Game, part 7 of 8 | 10.00 GiB | `0c68334b33ee878c5beac321339b9447` |
| [GenshinImpact_5.3.0.zip.008](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/GenshinImpact_5.3.0.zip.008) | Game, part 8 of 8 | 973.40 MiB | `18d44596a5f1467682f5e038c80bd92a` |
| [Audio_Chinese_5.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/Audio_Chinese_5.3.0.zip) | Chinese voice-over | 14.20 GiB | `2727087a20d630d35efe804ae683e72e` |
| [Audio_English(US)_5.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/Audio_English%28US%29_5.3.0.zip) | English voice-over | 16.24 GiB | `76f338d1925ff39cbf73f0418e9ae354` |
| [Audio_Japanese_5.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/Audio_Japanese_5.3.0.zip) | Japanese voice-over | 18.44 GiB | `f2b1c1f217dafbcdf27aeece987256b7` |
| [Audio_Korean_5.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/Audio_Korean_5.3.0.zip) | Korean voice-over | 14.00 GiB | `6356a494c7cce397bdbb1213aa6e7298` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241219110745_1vT3FzXdDTDFZFrL/ScatteredFiles` (file list in `pkg_version`).

#### Update from 5.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.2.0_5.3.0_hdiff_cixFUHQZyAJuOKIN.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.2.0_5.3.0_hdiff_cixFUHQZyAJuOKIN.zip) | Game | 17.17 GiB | `46b32c38dbb348a172df8352074da5a1` |
| [audio_zh-cn_5.2.0_5.3.0_hdiff_zFnGPPbohnwLAFxC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.2.0_5.3.0_hdiff_zFnGPPbohnwLAFxC.zip) | Chinese voice-over | 661.34 MiB | `3b4f264bd791b5f81eb165fe6d36676e` |
| [audio_en-us_5.2.0_5.3.0_hdiff_kFXfWOqklZrcycKa.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.2.0_5.3.0_hdiff_kFXfWOqklZrcycKa.zip) | English voice-over | 665.68 MiB | `c7189fa380b90a343cc56b7d8438eaba` |
| [audio_ja-jp_5.2.0_5.3.0_hdiff_mqXJGZjryItulRRx.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.2.0_5.3.0_hdiff_mqXJGZjryItulRRx.zip) | Japanese voice-over | 803.66 MiB | `68146dc2e2ea63b0cae452ea01b23136` |
| [audio_ko-kr_5.2.0_5.3.0_hdiff_oyddjWZZfxEwaTYM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.2.0_5.3.0_hdiff_oyddjWZZfxEwaTYM.zip) | Korean voice-over | 703.04 MiB | `fcee09acf85cdd2dd1d30bdea69c9065` |

#### Update from 5.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.1.0_5.3.0_hdiff_EqAXgBXKuyfZhuCC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.1.0_5.3.0_hdiff_EqAXgBXKuyfZhuCC.zip) | Game | 23.39 GiB | `f5dba04301e9c3c4985ef5f68b342f69` |
| [audio_zh-cn_5.1.0_5.3.0_hdiff_bwYLrQOzBulwMjsk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.1.0_5.3.0_hdiff_bwYLrQOzBulwMjsk.zip) | Chinese voice-over | 942.31 MiB | `574d39b7d624be84a4b718a7da2f7071` |
| [audio_en-us_5.1.0_5.3.0_hdiff_XntHHgeQMKmMPDPD.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.1.0_5.3.0_hdiff_XntHHgeQMKmMPDPD.zip) | English voice-over | 1.05 GiB | `eda197bbb6f35f27d04443dc20e63d4e` |
| [audio_ja-jp_5.1.0_5.3.0_hdiff_gsmDWqeCKRCyKocR.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.1.0_5.3.0_hdiff_gsmDWqeCKRCyKocR.zip) | Japanese voice-over | 1.09 GiB | `9d35e1a6dca5e89659e5a8e62ee72ec3` |
| [audio_ko-kr_5.1.0_5.3.0_hdiff_ZNxbHFxkCPDnTaZZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.1.0_5.3.0_hdiff_ZNxbHFxkCPDnTaZZ.zip) | Korean voice-over | 984.31 MiB | `8421256053f495379f1bef6247cb5cdc` |

### 5.2.0

Seen in the API as the live version from 2024-11-20 to 2024-12-31, pre-download from 2024-11-18 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.2.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `651a72b1adb3bb5be28dcffded9ee7ec` |
| [GenshinImpact_5.2.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `ecfcb0aa4b403608aa0089e8cbc0c618` |
| [GenshinImpact_5.2.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `5bac8cd01f87d13dd3e20cbf897731d9` |
| [GenshinImpact_5.2.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `84501620dded6000f758bfc4bde68d50` |
| [GenshinImpact_5.2.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `2d47ee68d07fd698c02001311998498f` |
| [GenshinImpact_5.2.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `8246d93ee1f93ed787bf87314220ed4e` |
| [GenshinImpact_5.2.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/GenshinImpact_5.2.0.zip.007) | Game, part 7 of 7 | 7.24 GiB | `fda38b6c9a5f746829abd8b561f51f54` |
| [Audio_Chinese_5.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/Audio_Chinese_5.2.0.zip) | Chinese voice-over | 13.42 GiB | `10dd6595c3687dec572afdc329c47eec` |
| [Audio_English(US)_5.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/Audio_English%28US%29_5.2.0.zip) | English voice-over | 15.46 GiB | `c67d7e98f823afce67fe662f76d27388` |
| [Audio_Japanese_5.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/Audio_Japanese_5.2.0.zip) | Japanese voice-over | 17.41 GiB | `285fba0cf723fc0ec033944caa1ada5c` |
| [Audio_Korean_5.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/Audio_Korean_5.2.0.zip) | Korean voice-over | 13.17 GiB | `071cfaab116ac1578fb6d755c34634e0` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20241108173711_97p3DLcQqGoVFVR0/ScatteredFiles` (file list in `pkg_version`).

#### Update from 5.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.1.0_5.2.0_hdiff_pZdCtRzXGSlIRTmx.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.1.0_5.2.0_hdiff_pZdCtRzXGSlIRTmx.zip) | Game | 16.23 GiB | `8fb18e51497e77d8f771cccb882a01e6` |
| [audio_zh-cn_5.1.0_5.2.0_hdiff_uLyHclofXUAJoSHK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.1.0_5.2.0_hdiff_uLyHclofXUAJoSHK.zip) | Chinese voice-over | 281.47 MiB | `7774ea4a38100d93b0bf893bb1a4fdca` |
| [audio_en-us_5.1.0_5.2.0_hdiff_YFHqlJfMKPdxNJAj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.1.0_5.2.0_hdiff_YFHqlJfMKPdxNJAj.zip) | English voice-over | 413.63 MiB | `d4a352c3bbc1d4bbc8471898ad2bfe1c` |
| [audio_ja-jp_5.1.0_5.2.0_hdiff_viAIYCwPGzRJVDqV.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.1.0_5.2.0_hdiff_viAIYCwPGzRJVDqV.zip) | Japanese voice-over | 311.53 MiB | `dea044b964457116a74656627906e104` |
| [audio_ko-kr_5.1.0_5.2.0_hdiff_kpURIBxLmNZpuqWi.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.1.0_5.2.0_hdiff_kpURIBxLmNZpuqWi.zip) | Korean voice-over | 281.78 MiB | `1ed412136f723515f3da76a8349e9ab8` |

#### Update from 5.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.0.0_5.2.0_hdiff_nsPbABFFJkqIWACo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.0.0_5.2.0_hdiff_nsPbABFFJkqIWACo.zip) | Game | 23.66 GiB | `f4ffe9970f71f6f3789c4d1d7fa01d8b` |
| [audio_zh-cn_5.0.0_5.2.0_hdiff_oHPVZfgwnennEoHX.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.0.0_5.2.0_hdiff_oHPVZfgwnennEoHX.zip) | Chinese voice-over | 831.34 MiB | `d4823b881a44b8c251d4bd53d4d77e90` |
| [audio_en-us_5.0.0_5.2.0_hdiff_FAoZwUpiPebLlqjz.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.0.0_5.2.0_hdiff_FAoZwUpiPebLlqjz.zip) | English voice-over | 852.78 MiB | `9c3d6b24a29884c8f5a78a24e955b2d4` |
| [audio_ja-jp_5.0.0_5.2.0_hdiff_oQYCpQTDDuJqFLUM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.0.0_5.2.0_hdiff_oQYCpQTDDuJqFLUM.zip) | Japanese voice-over | 917.86 MiB | `9e83a0cb36f3facdd555ec905df9fa5d` |
| [audio_ko-kr_5.0.0_5.2.0_hdiff_UzYUcmxtPnkLSYLe.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.0.0_5.2.0_hdiff_UzYUcmxtPnkLSYLe.zip) | Korean voice-over | 806.79 MiB | `4e6a4c613273c3e0ae630ed5243719b3` |

### 5.1.0

Seen in the API as the live version from 2024-10-09 to 2024-11-19, pre-download from 2024-10-07 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.1.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `cbe2934260fda10ab8cdaca80ef69cb3` |
| [GenshinImpact_5.1.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `df706bb5e83dde4d3df7276a1210a8fb` |
| [GenshinImpact_5.1.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `dd7062a4abde83f02f184fe081eb2006` |
| [GenshinImpact_5.1.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `bd2ffa879db43c05f88f7f1235f50df8` |
| [GenshinImpact_5.1.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `e3beb938f521be5a74431716394baee3` |
| [GenshinImpact_5.1.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `f59a086e986cce9f9e98e2d3c44212c8` |
| [GenshinImpact_5.1.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/GenshinImpact_5.1.0.zip.007) | Game, part 7 of 7 | 5.94 GiB | `0e165aec640ca7025f9d52b0f53a4223` |
| [Audio_Chinese_5.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/Audio_Chinese_5.1.0.zip) | Chinese voice-over | 13.09 GiB | `5c2e8a5e8a79a2e04a43e3a9d6df1ade` |
| [Audio_English(US)_5.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/Audio_English%28US%29_5.1.0.zip) | English voice-over | 15.09 GiB | `2504bd3606681c20cbdb5f82e8ca361b` |
| [Audio_Japanese_5.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/Audio_Japanese_5.1.0.zip) | Japanese voice-over | 17.00 GiB | `b2a2d965ccc36b1c583381cf4b3962e9` |
| [Audio_Korean_5.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/Audio_Korean_5.1.0.zip) | Korean voice-over | 12.84 GiB | `89f4668220ceaba401244c5d506defeb` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240927184459_CVopfb3tD4Zi3As6/ScatteredFiles` (file list in `pkg_version`).

#### Update from 5.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_5.0.0_5.1.0_hdiff_YqcpfTRBIuOIwHrO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_5.0.0_5.1.0_hdiff_YqcpfTRBIuOIwHrO.zip) | Game | 18.00 GiB | `3d628fcde7aaf5b242e253188885ef8d` |
| [audio_zh-cn_5.0.0_5.1.0_hdiff_bktSOAPsrWSJHOII.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_5.0.0_5.1.0_hdiff_bktSOAPsrWSJHOII.zip) | Chinese voice-over | 550.23 MiB | `1e988fbd3a6854e885564e3cca56fb65` |
| [audio_en-us_5.0.0_5.1.0_hdiff_jRxUUsOrikPdkOPL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_5.0.0_5.1.0_hdiff_jRxUUsOrikPdkOPL.zip) | English voice-over | 522.03 MiB | `501590a0aa8a67e0101e827be5c16cef` |
| [audio_ja-jp_5.0.0_5.1.0_hdiff_cSIWEmsNoqluMIrF.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_5.0.0_5.1.0_hdiff_cSIWEmsNoqluMIrF.zip) | Japanese voice-over | 608.07 MiB | `de00c4d5d02dead7b8e80c68ab57db17` |
| [audio_ko-kr_5.0.0_5.1.0_hdiff_PumQPsFuEgvDnZyX.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_5.0.0_5.1.0_hdiff_PumQPsFuEgvDnZyX.zip) | Korean voice-over | 526.03 MiB | `0d77474fe49e46cdfc20298034501019` |

#### Update from 4.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.8.0_5.1.0_hdiff_GXYqlkIrnPhGxgBL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_4.8.0_5.1.0_hdiff_GXYqlkIrnPhGxgBL.zip) | Game | 24.95 GiB | `a69026b5c131b317850de7df162c1f0d` |
| [audio_zh-cn_4.8.0_5.1.0_hdiff_jTthnkzFifLoNpMG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_4.8.0_5.1.0_hdiff_jTthnkzFifLoNpMG.zip) | Chinese voice-over | 1.17 GiB | `3e874f75a561e606606a25a23a091d1a` |
| [audio_en-us_4.8.0_5.1.0_hdiff_FxnkfCmaqJMtphJu.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_4.8.0_5.1.0_hdiff_FxnkfCmaqJMtphJu.zip) | English voice-over | 1.25 GiB | `3e3ed6a14926ac7e87eab7d776fbe0be` |
| [audio_ja-jp_4.8.0_5.1.0_hdiff_SCCqvOQJFVpHQuqn.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_4.8.0_5.1.0_hdiff_SCCqvOQJFVpHQuqn.zip) | Japanese voice-over | 1.30 GiB | `23a0f5d02f259592c2dcfbe999aedeed` |
| [audio_ko-kr_4.8.0_5.1.0_hdiff_xyLFEAxJPVfQIEdZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_4.8.0_5.1.0_hdiff_xyLFEAxJPVfQIEdZ.zip) | Korean voice-over | 1.11 GiB | `ebe8365da2247a7bb26401fc1d9e63a0` |

### 5.0.0

Seen in the API as the live version from 2024-08-28 to 2024-10-08, pre-download from 2024-08-26 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_5.0.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `1ebf5dbcbe43bebcda7a57a8d789092e` |
| [GenshinImpact_5.0.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `57a67026c45d57c28e5b52e24e84cc04` |
| [GenshinImpact_5.0.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `5e66ff28eaf6ba89e49f153c0f077d34` |
| [GenshinImpact_5.0.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `39f014a760e27f77abed1989739c74c6` |
| [GenshinImpact_5.0.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `15f9405a199afba833f18fce288b9c7f` |
| [GenshinImpact_5.0.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `881432ceab27987b1297c9eefb39f192` |
| [GenshinImpact_5.0.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/GenshinImpact_5.0.0.zip.007) | Game, part 7 of 7 | 3.78 GiB | `951f91992b428385294baf9b6c764d49` |
| [Audio_Chinese_5.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/Audio_Chinese_5.0.0.zip) | Chinese voice-over | 14.78 GiB | `216b3e53f3c5c7e1290891696b2bbc66` |
| [Audio_English(US)_5.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/Audio_English%28US%29_5.0.0.zip) | English voice-over | 17.23 GiB | `ecd59f31ec48c50f9051fdad39603d67` |
| [Audio_Japanese_5.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/Audio_Japanese_5.0.0.zip) | Japanese voice-over | 19.38 GiB | `95efbd23e1bde2eb574f8090cc118109` |
| [Audio_Korean_5.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/Audio_Korean_5.0.0.zip) | Korean voice-over | 14.52 GiB | `8bf09bd07413189b69a5a0512df97335` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240816185649_LtymMnnIZVQfbLZ2/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.8.0_5.0.0_hdiff_wZvKsUhQtnBEutrh.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_4.8.0_5.0.0_hdiff_wZvKsUhQtnBEutrh.zip) | Game | 19.52 GiB | `ec0cc740e01c374cd53d48219c1a4aff` |
| [audio_zh-cn_4.8.0_5.0.0_hdiff_ZZiDHvYQeHGKuFeP.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_4.8.0_5.0.0_hdiff_ZZiDHvYQeHGKuFeP.zip) | Chinese voice-over | 643.26 MiB | `c1073373d6c7b3680217335dc346de50` |
| [audio_en-us_4.8.0_5.0.0_hdiff_vsfAECOkroqoZSqK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_4.8.0_5.0.0_hdiff_vsfAECOkroqoZSqK.zip) | English voice-over | 763.68 MiB | `b5cb77749a0e2fc0e85b6b1ee319a7e9` |
| [audio_ja-jp_4.8.0_5.0.0_hdiff_LumPhRraNOjGJMnG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_4.8.0_5.0.0_hdiff_LumPhRraNOjGJMnG.zip) | Japanese voice-over | 733.02 MiB | `934899593f7234bec937fca98c0b0bed` |
| [audio_ko-kr_4.8.0_5.0.0_hdiff_gTmBNUAGPpXxrRKC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_4.8.0_5.0.0_hdiff_gTmBNUAGPpXxrRKC.zip) | Korean voice-over | 617.06 MiB | `64c52d4065a5d5983b37faeb9796a3eb` |

#### Update from 4.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.7.0_5.0.0_hdiff_yIzQMgpcGyexxigH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_4.7.0_5.0.0_hdiff_yIzQMgpcGyexxigH.zip) | Game | 23.50 GiB | `1589950b2fb98962df67dbde5011f782` |
| [audio_zh-cn_4.7.0_5.0.0_hdiff_oiHvhPpyOEfmnCRC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_4.7.0_5.0.0_hdiff_oiHvhPpyOEfmnCRC.zip) | Chinese voice-over | 1000.82 MiB | `6d50b43f4b8ae5b3034a22bd51831ca8` |
| [audio_en-us_4.7.0_5.0.0_hdiff_peHdUQzpeMePNdRd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_4.7.0_5.0.0_hdiff_peHdUQzpeMePNdRd.zip) | English voice-over | 1.17 GiB | `f2ada108661d82729d2aa6d5acb585ea` |
| [audio_ja-jp_4.7.0_5.0.0_hdiff_IjbPGEoMTmdwazkV.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_4.7.0_5.0.0_hdiff_IjbPGEoMTmdwazkV.zip) | Japanese voice-over | 1.10 GiB | `1601befcfcd638cd7b4aed3091f28155` |
| [audio_ko-kr_4.7.0_5.0.0_hdiff_jtnGBZilUVfWRxlB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_4.7.0_5.0.0_hdiff_jtnGBZilUVfWRxlB.zip) | Korean voice-over | 939.50 MiB | `c71685deb67a733059d73072e1199b12` |

</details>

<details>
<summary><b>Version 4.x</b>: 4.0.0 to 4.8.0 (10 versions)</summary>

Jump to: [4.8.0](#480) · [4.7.0](#470) · [4.6.0](#460) · [4.5.0](#450) · [4.4.0](#440) · [4.3.0](#430) · [4.2.0](#420) · [4.1.0](#410) · [4.0.1](#401) · [4.0.0](#400)

### 4.8.0

Seen in the API as the live version from 2024-07-18 to 2024-08-27 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.8.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.001) | Game, part 1 of 6 | 10.00 GiB | `46bf7f53e8f78fc540d72a66bdfd1f9b` |
| [GenshinImpact_4.8.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.002) | Game, part 2 of 6 | 10.00 GiB | `87c735e51721b872fcaec990db078716` |
| [GenshinImpact_4.8.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.003) | Game, part 3 of 6 | 10.00 GiB | `2976acd8fd24367a0e3f4caa2eb4f5f6` |
| [GenshinImpact_4.8.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.004) | Game, part 4 of 6 | 10.00 GiB | `6510b9288599eef831a0a83f2dcf6f36` |
| [GenshinImpact_4.8.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.005) | Game, part 5 of 6 | 10.00 GiB | `2eb292ad6deabbe9fb4f17f584c9c744` |
| [GenshinImpact_4.8.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/GenshinImpact_4.8.0.zip.006) | Game, part 6 of 6 | 8.55 GiB | `378a0ad2b7a76c9402fc66d408b2cdfa` |
| [Audio_Chinese_4.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/Audio_Chinese_4.8.0.zip) | Chinese voice-over | 14.05 GiB | `d5ef6d872a3f3f080eb2bd5132670aaa` |
| [Audio_English(US)_4.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/Audio_English%28US%29_4.8.0.zip) | English voice-over | 16.41 GiB | `0ea70602f7eb578acfb9ef3604479118` |
| [Audio_Japanese_4.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/Audio_Japanese_4.8.0.zip) | Japanese voice-over | 18.41 GiB | `e5ee37135fa7ee787606bd33ef25c1da` |
| [Audio_Korean_4.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/Audio_Korean_4.8.0.zip) | Korean voice-over | 13.79 GiB | `9e4e588b658f39419aa0c6a4d7f0052b` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240705190845_sEVV7ebRBuMrOheE/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.7.0_4.8.0_hdiff_WCGypepmldFKYWmS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_4.7.0_4.8.0_hdiff_WCGypepmldFKYWmS.zip) | Game | 13.84 GiB | `f18a0e18b4df64ae2109c6d951cdc37c` |
| [audio_zh-cn_4.7.0_4.8.0_hdiff_XunTGFzRehXbdfBM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_4.7.0_4.8.0_hdiff_XunTGFzRehXbdfBM.zip) | Chinese voice-over | 357.67 MiB | `03082217795da95eb4a2b117a8a01ff7` |
| [audio_en-us_4.7.0_4.8.0_hdiff_iRPQrRliVLmISmPj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_4.7.0_4.8.0_hdiff_iRPQrRliVLmISmPj.zip) | English voice-over | 432.67 MiB | `e2dadcc79476359c1a6cee9c7141155f` |
| [audio_ja-jp_4.7.0_4.8.0_hdiff_TpngPXQdYNIqGmBB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_4.7.0_4.8.0_hdiff_TpngPXQdYNIqGmBB.zip) | Japanese voice-over | 398.03 MiB | `66ec311c85b05fda917526f36361c9ef` |
| [audio_ko-kr_4.7.0_4.8.0_hdiff_VAmyxLulAOgLBfEQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_4.7.0_4.8.0_hdiff_VAmyxLulAOgLBfEQ.zip) | Korean voice-over | 322.57 MiB | `47daefe2506360faef5a46c1f8661bdc` |

#### Update from 4.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.6.0_4.8.0_hdiff_glbymjcdFUJUkjaj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/game_4.6.0_4.8.0_hdiff_glbymjcdFUJUkjaj.zip) | Game | 27.23 GiB | `40a04f8f74124dfcfc66525da4d49aa3` |
| [audio_zh-cn_4.6.0_4.8.0_hdiff_TqEURkOKyCJnpLhD.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_zh-cn_4.6.0_4.8.0_hdiff_TqEURkOKyCJnpLhD.zip) | Chinese voice-over | 869.81 MiB | `dc0055e4282205fe082a217ef50f0620` |
| [audio_en-us_4.6.0_4.8.0_hdiff_AFapvvrqGdvNUglt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_en-us_4.6.0_4.8.0_hdiff_AFapvvrqGdvNUglt.zip) | English voice-over | 1.02 GiB | `3de090a8c277665af5a609fa9e308787` |
| [audio_ja-jp_4.6.0_4.8.0_hdiff_jcwUMAzUnBmCDrEt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ja-jp_4.6.0_4.8.0_hdiff_jcwUMAzUnBmCDrEt.zip) | Japanese voice-over | 977.84 MiB | `cb56ffa644747381c0d863ab18b79421` |
| [audio_ko-kr_4.6.0_4.8.0_hdiff_dDPSxPwdcwUcIMsl.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/audio_ko-kr_4.6.0_4.8.0_hdiff_dDPSxPwdcwUcIMsl.zip) | Korean voice-over | 781.53 MiB | `fe4e5fe9b1d17bab4739371ce251e469` |

### 4.7.0

Seen in the API as the live version from 2024-06-05 to 2024-07-17, pre-download from 2024-06-03 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.7.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.001) | Game, part 1 of 6 | 10.00 GiB | `0790ed842a1732fb9e5530a826828440` |
| [GenshinImpact_4.7.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.002) | Game, part 2 of 6 | 10.00 GiB | `6ac391b6a3a185bc8ab1e431f67ecd25` |
| [GenshinImpact_4.7.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.003) | Game, part 3 of 6 | 10.00 GiB | `36460a467de4901f517f8ed9be6b877c` |
| [GenshinImpact_4.7.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.004) | Game, part 4 of 6 | 10.00 GiB | `d1c0d81ab7aff5d5fb490cff20b9b87f` |
| [GenshinImpact_4.7.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.005) | Game, part 5 of 6 | 10.00 GiB | `fec57d6f7f78c04309f16dfc2207cd6f` |
| [GenshinImpact_4.7.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/GenshinImpact_4.7.0.zip.006) | Game, part 6 of 6 | 6.66 GiB | `84135fa7008156965514a6ec99c55c66` |
| [Audio_Chinese_4.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/Audio_Chinese_4.7.0.zip) | Chinese voice-over | 13.65 GiB | `15837d2859a5f1c997bd0a85c8f4d630` |
| [Audio_English(US)_4.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/Audio_English%28US%29_4.7.0.zip) | English voice-over | 15.94 GiB | `fe58d4413dd6d1159bb345746a69c3aa` |
| [Audio_Japanese_4.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/Audio_Japanese_4.7.0.zip) | Japanese voice-over | 17.90 GiB | `7fab66b921af85d1065724d196488beb` |
| [Audio_Korean_4.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/Audio_Korean_4.7.0.zip) | Korean voice-over | 13.40 GiB | `d4423f11472c816b842d95dfd05044dd` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240524181522_P7n5afVhY8WeoVZb/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.6.0_4.7.0_hdiff_PIM7zBX3lxGjbaNV.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.6.0_4.7.0_hdiff_PIM7zBX3lxGjbaNV.zip) | Game | 24.36 GiB | `5e6e400e6f8c9efc4b0f5b2609edeac1` |
| [zh-cn_4.6.0_4.7.0_hdiff_vGOido1Us5Djnwb8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.6.0_4.7.0_hdiff_vGOido1Us5Djnwb8.zip) | Chinese voice-over | 513.59 MiB | `583a7cb93f7d0651af20dc2ca16f8b24` |
| [en-us_4.6.0_4.7.0_hdiff_ya8F9GL1ojP4D7nR.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.6.0_4.7.0_hdiff_ya8F9GL1ojP4D7nR.zip) | English voice-over | 607.33 MiB | `cb9dcddc80cf3e336c280c05f67627d2` |
| [ja-jp_4.6.0_4.7.0_hdiff_NeDO7m6dnx3cLRpy.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.6.0_4.7.0_hdiff_NeDO7m6dnx3cLRpy.zip) | Japanese voice-over | 579.81 MiB | `8406125f361fc533e3b174e749c87b39` |
| [ko-kr_4.6.0_4.7.0_hdiff_3qryOQWP2YDnKmBA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.6.0_4.7.0_hdiff_3qryOQWP2YDnKmBA.zip) | Korean voice-over | 458.70 MiB | `a2c6d59146ec07f05f5a003ff93aa6e0` |

#### Update from 4.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.5.0_4.7.0_hdiff_hjfZ1LgT52ceYDOv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.5.0_4.7.0_hdiff_hjfZ1LgT52ceYDOv.zip) | Game | 29.80 GiB | `c137e40dfb506146d05adb294a5d6316` |
| [zh-cn_4.5.0_4.7.0_hdiff_TQp5Jwn0FAfIDsE4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.5.0_4.7.0_hdiff_TQp5Jwn0FAfIDsE4.zip) | Chinese voice-over | 820.08 MiB | `43949e7bfbe70e61877edb0b208ab156` |
| [en-us_4.5.0_4.7.0_hdiff_xEHpcAnqNb48Ff3S.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.5.0_4.7.0_hdiff_xEHpcAnqNb48Ff3S.zip) | English voice-over | 1006.46 MiB | `c5b4f308bf22c7825c98538d3f774f7a` |
| [ja-jp_4.5.0_4.7.0_hdiff_GCOBxKrAuTwDXqeQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.5.0_4.7.0_hdiff_GCOBxKrAuTwDXqeQ.zip) | Japanese voice-over | 918.34 MiB | `a6316c5733169a4f5b80a40cecc0653a` |
| [ko-kr_4.5.0_4.7.0_hdiff_TVIabdjZRsfUn1B9.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.5.0_4.7.0_hdiff_TVIabdjZRsfUn1B9.zip) | Korean voice-over | 732.59 MiB | `ac8e76c4744dfc9e8cd8cb74bfe3ee44` |

### 4.6.0

Seen in the API as the live version from 2024-04-24 to 2024-06-04, pre-download from 2024-04-22 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.6.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.001) | Game, part 1 of 6 | 10.00 GiB | `03ce6a5d4b7d8ba192cc6d9133a00bef` |
| [GenshinImpact_4.6.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.002) | Game, part 2 of 6 | 10.00 GiB | `a0320432c00779630df4a421ac7f3999` |
| [GenshinImpact_4.6.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.003) | Game, part 3 of 6 | 10.00 GiB | `335a750e6576205970a0bb49fbb2ce72` |
| [GenshinImpact_4.6.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.004) | Game, part 4 of 6 | 10.00 GiB | `2210fca193fc5261c5ae75841ce17002` |
| [GenshinImpact_4.6.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.005) | Game, part 5 of 6 | 10.00 GiB | `f96757dbce7764e65150449f4a476b3e` |
| [GenshinImpact_4.6.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/GenshinImpact_4.6.0.zip.006) | Game, part 6 of 6 | 5.28 GiB | `6b43dac3616027ce7a980c41092f72fa` |
| [Audio_Chinese_4.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/Audio_Chinese_4.6.0.zip) | Chinese voice-over | 13.07 GiB | `974bc5df3f7d96de01b68f1aab0371e5` |
| [Audio_English(US)_4.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/Audio_English%28US%29_4.6.0.zip) | English voice-over | 15.29 GiB | `f577ea4624981b36789e24913d0f3745` |
| [Audio_Japanese_4.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/Audio_Japanese_4.6.0.zip) | Japanese voice-over | 17.13 GiB | `f039da619e670259d7e57c4f5e84be9e` |
| [Audio_Korean_4.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/Audio_Korean_4.6.0.zip) | Korean voice-over | 12.85 GiB | `58c3fc8782f3d59b324dfd603fa93e6d` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240412191920_HThGUGlY3Gjy6AXP/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.5.0_4.6.0_hdiff_MyQqNZPJWOktACjI.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.5.0_4.6.0_hdiff_MyQqNZPJWOktACjI.zip) | Game | 28.42 GiB | `2914b6cd5decb1a1664d0ff2eb24f74e` |
| [zh-cn_4.5.0_4.6.0_hdiff_q8fZULgOsl1SoipB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.5.0_4.6.0_hdiff_q8fZULgOsl1SoipB.zip) | Chinese voice-over | 306.68 MiB | `bcc70907ad4b514270283a4cf57cf955` |
| [en-us_4.5.0_4.6.0_hdiff_aRk4uMi1H9OALJ7d.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.5.0_4.6.0_hdiff_aRk4uMi1H9OALJ7d.zip) | English voice-over | 399.28 MiB | `086a30a2cf1f340efd70a1e3051cd3dc` |
| [ja-jp_4.5.0_4.6.0_hdiff_dNhr1mUeITOnZ76R.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.5.0_4.6.0_hdiff_dNhr1mUeITOnZ76R.zip) | Japanese voice-over | 338.69 MiB | `ee4e47e80478a378cd78b0467b7aa1ba` |
| [ko-kr_4.5.0_4.6.0_hdiff_wrnM3Vumg5xSbKEa.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.5.0_4.6.0_hdiff_wrnM3Vumg5xSbKEa.zip) | Korean voice-over | 273.97 MiB | `6aec8a51b5bbed306251534da406e505` |

#### Update from 4.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.4.0_4.6.0_hdiff_rUuHtsVZCRbiNGMd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.4.0_4.6.0_hdiff_rUuHtsVZCRbiNGMd.zip) | Game | 28.76 GiB | `5f75430f99aaba0b20164defec9ed7f0` |
| [zh-cn_4.4.0_4.6.0_hdiff_WlFJDQ7o4kLxdvEY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.4.0_4.6.0_hdiff_WlFJDQ7o4kLxdvEY.zip) | Chinese voice-over | 610.15 MiB | `de07bdbf77a152521727137697fd3457` |
| [en-us_4.4.0_4.6.0_hdiff_yuFxY5cgThAqwWiU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.4.0_4.6.0_hdiff_yuFxY5cgThAqwWiU.zip) | English voice-over | 851.41 MiB | `f875d21c3e3f4234437c037ae2e324fd` |
| [ja-jp_4.4.0_4.6.0_hdiff_5VqBYzo3K6pfjhGd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.4.0_4.6.0_hdiff_5VqBYzo3K6pfjhGd.zip) | Japanese voice-over | 705.16 MiB | `df3c9c8a089557ff184a9db14062e49f` |
| [ko-kr_4.4.0_4.6.0_hdiff_M7chPZ8Dg2dIbtSf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.4.0_4.6.0_hdiff_M7chPZ8Dg2dIbtSf.zip) | Korean voice-over | 573.81 MiB | `65b3756cd867abb3845d02fc6e388a88` |

### 4.5.0

Seen in the API as the live version from 2024-03-13 to 2024-04-23, pre-download from 2024-03-11 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.5.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.001) | Game, part 1 of 8 | 10.00 GiB | `f90ff656a5c84c299a9a0359e67d6133` |
| [GenshinImpact_4.5.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.002) | Game, part 2 of 8 | 10.00 GiB | `8503c15c090ec71d699591d4782ac9ac` |
| [GenshinImpact_4.5.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.003) | Game, part 3 of 8 | 10.00 GiB | `f57a42ced76275035216984ef3edb1b0` |
| [GenshinImpact_4.5.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.004) | Game, part 4 of 8 | 10.00 GiB | `47b905fb1a6a3a3e8518002957ef543d` |
| [GenshinImpact_4.5.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.005) | Game, part 5 of 8 | 10.00 GiB | `9e2200dd6334f27ece3afd3008c83ee4` |
| [GenshinImpact_4.5.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.006) | Game, part 6 of 8 | 10.00 GiB | `1944ef5e7b24adb38157b5920d2662be` |
| [GenshinImpact_4.5.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.007) | Game, part 7 of 8 | 10.00 GiB | `815eeee0c40f125c8ac355551893e4cc` |
| [GenshinImpact_4.5.0.zip.008](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/GenshinImpact_4.5.0.zip.008) | Game, part 8 of 8 | 175.76 MiB | `95be2c4d704b1e2a8dad719344e66e63` |
| [Audio_Chinese_4.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/Audio_Chinese_4.5.0.zip) | Chinese voice-over | 12.73 GiB | `53c786603c74b1085cb445a306d75644` |
| [Audio_English(US)_4.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/Audio_English%28US%29_4.5.0.zip) | English voice-over | 14.86 GiB | `6dc1ebd7e246d04f91ee3a1b5d73f4bf` |
| [Audio_Japanese_4.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/Audio_Japanese_4.5.0.zip) | Japanese voice-over | 16.68 GiB | `54eda6310c08932887ac1b40386d3769` |
| [Audio_Korean_4.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/Audio_Korean_4.5.0.zip) | Korean voice-over | 12.51 GiB | `7ca3127b5f14e398147fa4e2db8f2670` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240301203033_RZSIny3hwJ5nq959/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.4.0_4.5.0_hdiff_nodDY9LQAgruZ6Ns.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.4.0_4.5.0_hdiff_nodDY9LQAgruZ6Ns.zip) | Game | 4.84 GiB | `cfe53f50848b744d72f72adaa2249fc0` |
| [zh-cn_4.4.0_4.5.0_hdiff_YmheEAnjvstCbGxJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.4.0_4.5.0_hdiff_YmheEAnjvstCbGxJ.zip) | Chinese voice-over | 303.73 MiB | `9dc166ad2425ed760cb2620f75c6073e` |
| [en-us_4.4.0_4.5.0_hdiff_D3JK1CQZHt46IcBv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.4.0_4.5.0_hdiff_D3JK1CQZHt46IcBv.zip) | English voice-over | 452.70 MiB | `12e5e72abd1894dbce8e8be8ef32c5b3` |
| [ja-jp_4.4.0_4.5.0_hdiff_D16LijOxqwgfCSWa.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.4.0_4.5.0_hdiff_D16LijOxqwgfCSWa.zip) | Japanese voice-over | 366.54 MiB | `74a1342b469072bf77d1ab1ded5f1569` |
| [ko-kr_4.4.0_4.5.0_hdiff_v7HhnQibzTMAuO5B.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.4.0_4.5.0_hdiff_v7HhnQibzTMAuO5B.zip) | Korean voice-over | 299.97 MiB | `16aa71f9aec2bd13afa7e45ad7274d18` |

#### Update from 4.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.3.0_4.5.0_hdiff_4HTYIGcwxNvzoMms.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.3.0_4.5.0_hdiff_4HTYIGcwxNvzoMms.zip) | Game | 9.31 GiB | `e805219cc6d3f78de5f06938b7cf2c25` |
| [zh-cn_4.3.0_4.5.0_hdiff_FBmv75fUwNpaX1er.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.3.0_4.5.0_hdiff_FBmv75fUwNpaX1er.zip) | Chinese voice-over | 671.70 MiB | `92442a3ac434a88e46775832ae7579d6` |
| [en-us_4.3.0_4.5.0_hdiff_6kNCAZ30EecPULtH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.3.0_4.5.0_hdiff_6kNCAZ30EecPULtH.zip) | English voice-over | 928.17 MiB | `171d4dd9ffdb30b4209e915e048a6fe3` |
| [ja-jp_4.3.0_4.5.0_hdiff_NS9Ju75gDp0ybfFO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.3.0_4.5.0_hdiff_NS9Ju75gDp0ybfFO.zip) | Japanese voice-over | 806.36 MiB | `1ae45dadd301042545deb615cb187cfa` |
| [ko-kr_4.3.0_4.5.0_hdiff_icrsofHZMBLgj6em.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.3.0_4.5.0_hdiff_icrsofHZMBLgj6em.zip) | Korean voice-over | 652.73 MiB | `d4d7c4127f2ff97b602803119387a397` |

### 4.4.0

Seen in the API as the live version from 2024-01-31 to 2024-03-12, pre-download from 2024-01-29 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.4.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `e88cc4c3a692cbe2bcc6bafdf1b6faa4` |
| [GenshinImpact_4.4.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `f8b33f2c761c6717bc773314008a6ead` |
| [GenshinImpact_4.4.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `d88d3be3b856121836f66c1c7e3a5f15` |
| [GenshinImpact_4.4.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `4a1e1d951b35af65f6865d7b438893bc` |
| [GenshinImpact_4.4.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `9c055ecac31b34d961758a6e97918698` |
| [GenshinImpact_4.4.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `9daae514dcdda06e777514c226addb60` |
| [GenshinImpact_4.4.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/GenshinImpact_4.4.0.zip.007) | Game, part 7 of 7 | 9.77 GiB | `d2a532642f23eef30905319053da71ca` |
| [Audio_Chinese_4.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/Audio_Chinese_4.4.0.zip) | Chinese voice-over | 12.39 GiB | `c5450004abd3740e3d5590d98a80e008` |
| [Audio_English(US)_4.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/Audio_English%28US%29_4.4.0.zip) | English voice-over | 14.38 GiB | `82596a31081a0eb3fce46bb0b1c2fa8c` |
| [Audio_Japanese_4.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/Audio_Japanese_4.4.0.zip) | Japanese voice-over | 16.21 GiB | `1e33566476d0d7065612232024b5cabb` |
| [Audio_Korean_4.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/Audio_Korean_4.4.0.zip) | Korean voice-over | 12.15 GiB | `0c40f7b6c49932006bc36291871dc6a5` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20240119183743_YjHC1oBl0Hsgxkub/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.3.0_4.4.0_hdiff_7lGqkpy9saiZYfXS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.3.0_4.4.0_hdiff_7lGqkpy9saiZYfXS.zip) | Game | 7.61 GiB | `7ec96297548f58e9fbe568c9cf5a24be` |
| [zh-cn_4.3.0_4.4.0_hdiff_iGxba8mpHN694IBY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.3.0_4.4.0_hdiff_iGxba8mpHN694IBY.zip) | Chinese voice-over | 368.08 MiB | `88071a288298732c35ca60254333fe59` |
| [en-us_4.3.0_4.4.0_hdiff_q8uIUt3RlZoyKWcA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.3.0_4.4.0_hdiff_q8uIUt3RlZoyKWcA.zip) | English voice-over | 476.03 MiB | `d0a08ce1bcc39916f46e5865f106aa8e` |
| [ja-jp_4.3.0_4.4.0_hdiff_rnLX9isdJVu81afS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.3.0_4.4.0_hdiff_rnLX9isdJVu81afS.zip) | Japanese voice-over | 439.98 MiB | `ea2a5429ae24f49984bd4f4ee2f0a80d` |
| [ko-kr_4.3.0_4.4.0_hdiff_VzPIgE8jrkbZc1T4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.3.0_4.4.0_hdiff_VzPIgE8jrkbZc1T4.zip) | Korean voice-over | 353.10 MiB | `5d1d91973e71236d13a9638fe5e1ac7a` |

#### Update from 4.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.2.0_4.4.0_hdiff_1aB6HkjiezJR30fI.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.2.0_4.4.0_hdiff_1aB6HkjiezJR30fI.zip) | Game | 15.50 GiB | `2f15d128c46ff89749df10b59cc1fac0` |
| [zh-cn_4.2.0_4.4.0_hdiff_VjqxXHSlyGUbOdWM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.2.0_4.4.0_hdiff_VjqxXHSlyGUbOdWM.zip) | Chinese voice-over | 687.29 MiB | `e1e9c06038d91e6fc2cbe10d8b51f718` |
| [en-us_4.2.0_4.4.0_hdiff_R5vqmbc1UFEApWDg.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.2.0_4.4.0_hdiff_R5vqmbc1UFEApWDg.zip) | English voice-over | 900.30 MiB | `8c6597a37fc5282bbe4f9cf897460397` |
| [ja-jp_4.2.0_4.4.0_hdiff_pXCWv90ckibjMDZS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.2.0_4.4.0_hdiff_pXCWv90ckibjMDZS.zip) | Japanese voice-over | 833.14 MiB | `c46f853d18f676b522c4afe2965222d8` |
| [ko-kr_4.2.0_4.4.0_hdiff_CADJ8w3uMhLEPlGO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.2.0_4.4.0_hdiff_CADJ8w3uMhLEPlGO.zip) | Korean voice-over | 664.29 MiB | `54022b9fb8fdf6eb16b65a52106dec12` |

### 4.3.0

Seen in the API as the live version from 2023-12-20 to 2024-01-30, pre-download from 2023-12-18 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.3.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `a878723bc55799a9f44d419ee92e65db` |
| [GenshinImpact_4.3.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `3f42843fd6544af39664f059a4bb8c86` |
| [GenshinImpact_4.3.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `a6972a1053dba8697a3f18a4b49a5d2b` |
| [GenshinImpact_4.3.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `a29ed111ae6a6315ea1eae0f99e4a106` |
| [GenshinImpact_4.3.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `f6d4d05de27af51bfb7ce086538045ed` |
| [GenshinImpact_4.3.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `f484c07cb27fea8a0f2c48dd95a50b71` |
| [GenshinImpact_4.3.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/GenshinImpact_4.3.0.zip.007) | Game, part 7 of 7 | 6.08 GiB | `7a7713bf4822c2fed89332e5aa790290` |
| [Audio_Chinese_4.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/Audio_Chinese_4.3.0.zip) | Chinese voice-over | 11.98 GiB | `c1b32bf61b70f41d939c1afaf4bb6c06` |
| [Audio_English(US)_4.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/Audio_English%28US%29_4.3.0.zip) | English voice-over | 13.86 GiB | `fe3094550f7616595483b99d3e34e870` |
| [Audio_Japanese_4.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/Audio_Japanese_4.3.0.zip) | Japanese voice-over | 15.65 GiB | `c725c3cb235713ffe95ee0e7b0849685` |
| [Audio_Korean_4.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/Audio_Korean_4.3.0.zip) | Korean voice-over | 11.74 GiB | `8b7b91492a81e1b91930e5f960c1bf69` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231208190631_76nrKFC3l3HQt1yg/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.2.0_4.3.0_hdiff_YQPfi2FIskEr085l.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.2.0_4.3.0_hdiff_YQPfi2FIskEr085l.zip) | Game | 11.20 GiB | `15e9358cded433f9adaf51026819b3ee` |
| [zh-cn_4.2.0_4.3.0_hdiff_Ztz5oPHxBgrT83V2.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.2.0_4.3.0_hdiff_Ztz5oPHxBgrT83V2.zip) | Chinese voice-over | 319.50 MiB | `69ab245728f2d962486df0d911ddb9bd` |
| [en-us_4.2.0_4.3.0_hdiff_EnFvWP8B9SQTaAgO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.2.0_4.3.0_hdiff_EnFvWP8B9SQTaAgO.zip) | English voice-over | 424.33 MiB | `ad292027bc24e6f6737eda5e741607d1` |
| [ja-jp_4.2.0_4.3.0_hdiff_5hwF9mPYkgCy7uca.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.2.0_4.3.0_hdiff_5hwF9mPYkgCy7uca.zip) | Japanese voice-over | 393.26 MiB | `10507f960efb3bd21e7a104382082e64` |
| [ko-kr_4.2.0_4.3.0_hdiff_dsziMCt2DKNYe3IR.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.2.0_4.3.0_hdiff_dsziMCt2DKNYe3IR.zip) | Korean voice-over | 311.52 MiB | `301a330da0c23b4634f9cc8a505d169d` |

#### Update from 4.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.1.0_4.3.0_hdiff_aHVFC5jtk4JnmSQy.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.1.0_4.3.0_hdiff_aHVFC5jtk4JnmSQy.zip) | Game | 18.06 GiB | `cdd90e29dcee2e3b8691c0e24dcafe41` |
| [zh-cn_4.1.0_4.3.0_hdiff_bjeqMJf6EwOFsS3Q.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.1.0_4.3.0_hdiff_bjeqMJf6EwOFsS3Q.zip) | Chinese voice-over | 767.97 MiB | `064a69dacb5bd172ac6f1e7220fd98d1` |
| [en-us_4.1.0_4.3.0_hdiff_ViRbAxKhFm5ZM76s.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.1.0_4.3.0_hdiff_ViRbAxKhFm5ZM76s.zip) | English voice-over | 1010.15 MiB | `fa69771fce60439157493866863e7742` |
| [ja-jp_4.1.0_4.3.0_hdiff_gbIDuPQWRrc12NeO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.1.0_4.3.0_hdiff_gbIDuPQWRrc12NeO.zip) | Japanese voice-over | 924.06 MiB | `ff2a2351d5351b435d3120b83ea9ddfe` |
| [ko-kr_4.1.0_4.3.0_hdiff_qgPnM26ijOAyFsoJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.1.0_4.3.0_hdiff_qgPnM26ijOAyFsoJ.zip) | Korean voice-over | 777.89 MiB | `65f596ff107f540f9a386c96998a7094` |

### 4.2.0

Seen in the API as the live version from 2023-11-08 to 2023-12-19, pre-download from 2023-11-06 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip) (dead link, HTTP 404) | Game, single archive | unknown |  |
| [GenshinImpact_4.2.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `885b3c28d97d7d2186518182c4ab2d6c` |
| [GenshinImpact_4.2.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `b244ac40b994c6d8b1c094251b5b4140` |
| [GenshinImpact_4.2.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `b7ef7305eab3a7fb74772a576a034b4c` |
| [GenshinImpact_4.2.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `20e9d2b4fb096a98b7912122257b26d5` |
| [GenshinImpact_4.2.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `71be9a0ab6eaeb1748acc84820e44c32` |
| [GenshinImpact_4.2.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `84b1debd0bee45309157400096f64620` |
| [GenshinImpact_4.2.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/GenshinImpact_4.2.0.zip.007) | Game, part 7 of 7 | 4.92 GiB | `4d43a29fdef65a48f0bc03d7b2ea30a7` |
| [Audio_Chinese_4.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/Audio_Chinese_4.2.0.zip) | Chinese voice-over | 11.62 GiB | `87ef994d2f0227980b1bf22e2096cbd6` |
| [Audio_English(US)_4.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/Audio_English%28US%29_4.2.0.zip) | English voice-over | 13.40 GiB | `ee73904bf5ea10ad321b60cabfaaf5de` |
| [Audio_Japanese_4.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/Audio_Japanese_4.2.0.zip) | Japanese voice-over | 15.15 GiB | `1ac60012b267a710d06dc053e1687ebb` |
| [Audio_Korean_4.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/Audio_Korean_4.2.0.zip) | Korean voice-over | 11.38 GiB | `053ef9f2daf97e61cb0505c409fba5ca` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20231030132335_iOEfPMcbrXpiA8Ca/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.1.0_4.2.0_hdiff_pORAhLFjCN9dvuW7.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.1.0_4.2.0_hdiff_pORAhLFjCN9dvuW7.zip) | Game | 13.19 GiB | `d69cce82723e7471e79aab927d0b0fa7` |
| [zh-cn_4.1.0_4.2.0_hdiff_gmLkOPoXpj9xCeKJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.1.0_4.2.0_hdiff_gmLkOPoXpj9xCeKJ.zip) | Chinese voice-over | 448.74 MiB | `fc2cdaac7eec75578fb62dfcef0fe7be` |
| [en-us_4.1.0_4.2.0_hdiff_DUdAszbXT2OeCm5h.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.1.0_4.2.0_hdiff_DUdAszbXT2OeCm5h.zip) | English voice-over | 586.78 MiB | `62349d013e4d4d3173d5ef1fdf1f10b7` |
| [ja-jp_4.1.0_4.2.0_hdiff_JFoP2gn0tx7ks1Hd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.1.0_4.2.0_hdiff_JFoP2gn0tx7ks1Hd.zip) | Japanese voice-over | 531.08 MiB | `e403089a19c2e87a14b0271bf50c55d3` |
| [ko-kr_4.1.0_4.2.0_hdiff_PoiVS3Y0REXAgfeJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.1.0_4.2.0_hdiff_PoiVS3Y0REXAgfeJ.zip) | Korean voice-over | 466.59 MiB | `66ab9e4930b87dd1b551dcc6c3f45a72` |

#### Update from 4.0.1

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.0.1_4.2.0_hdiff_fkU1mJghZD5ScQvI.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.0.1_4.2.0_hdiff_fkU1mJghZD5ScQvI.zip) | Game | 19.01 GiB | `bcc371d3bb14b096de9cd3c424b99ec1` |
| [zh-cn_4.0.1_4.2.0_hdiff_cRPEbDns7e9wXh02.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.0.1_4.2.0_hdiff_cRPEbDns7e9wXh02.zip) | Chinese voice-over | 983.67 MiB | `4452d6d82e85fbe1d72ae9b86203de01` |
| [en-us_4.0.1_4.2.0_hdiff_fTHol786k5ctKZ30.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.0.1_4.2.0_hdiff_fTHol786k5ctKZ30.zip) | English voice-over | 1.24 GiB | `0edc24a7ce73548f9a8309bb5517d2ae` |
| [ja-jp_4.0.1_4.2.0_hdiff_9VMCIDbRsGuFx8UK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.0.1_4.2.0_hdiff_9VMCIDbRsGuFx8UK.zip) | Japanese voice-over | 1.16 GiB | `c78ccdd9a5fe057c14c5b9a6c72ff910` |
| [ko-kr_4.0.1_4.2.0_hdiff_eVUcYlKhFQPHnZwu.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.0.1_4.2.0_hdiff_eVUcYlKhFQPHnZwu.zip) | Korean voice-over | 1012.02 MiB | `b400b146f4aa6bbb5c82ca60cbeeed4b` |

#### Update from 4.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.0.0_4.2.0_hdiff_NGAnaC4bD0z6Excs.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.0.0_4.2.0_hdiff_NGAnaC4bD0z6Excs.zip) | Game | 19.01 GiB | `3c27d0f7f537b017979ee64a5a462ef0` |
| [zh-cn_4.0.0_4.2.0_hdiff_ZpFLgWbDaOB5mfl9.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.0.0_4.2.0_hdiff_ZpFLgWbDaOB5mfl9.zip) | Chinese voice-over | 983.67 MiB | `b851f404343cb8fdc78671ab968974a2` |
| [en-us_4.0.0_4.2.0_hdiff_tFGfxP61Nwo30zhE.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.0.0_4.2.0_hdiff_tFGfxP61Nwo30zhE.zip) | English voice-over | 1.24 GiB | `edc1418021d195b8bde3a2780362435d` |
| [ja-jp_4.0.0_4.2.0_hdiff_p7aB1CtbmGrwPUXy.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.0.0_4.2.0_hdiff_p7aB1CtbmGrwPUXy.zip) | Japanese voice-over | 1.16 GiB | `35f2f08a2987cdaaabea65448cbbac07` |
| [ko-kr_4.0.0_4.2.0_hdiff_kKGcq40lE9WOJpBy.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.0.0_4.2.0_hdiff_kKGcq40lE9WOJpBy.zip) | Korean voice-over | 1012.02 MiB | `fc64082d6c4e64e3619ff9fbe9aa4789` |

### 4.1.0

Seen in the API as the live version from 2023-09-27 to 2023-11-07, pre-download from 2023-09-25 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip) (dead link, HTTP 404) | Game, single archive | unknown |  |
| [GenshinImpact_4.1.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.001) | Game, part 1 of 7 | 10.00 GiB | `6e16cd1032ac0c8a43cc84af09b745d8` |
| [GenshinImpact_4.1.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.002) | Game, part 2 of 7 | 10.00 GiB | `6847fa749dde8bb8431ab8fb9d8ba711` |
| [GenshinImpact_4.1.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.003) | Game, part 3 of 7 | 10.00 GiB | `6d1413d8c713327d5839387c00239a34` |
| [GenshinImpact_4.1.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.004) | Game, part 4 of 7 | 10.00 GiB | `d3116dd5e824134af11727d7d1805921` |
| [GenshinImpact_4.1.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.005) | Game, part 5 of 7 | 10.00 GiB | `9b39eb1d88066591098d86e4bcd0deeb` |
| [GenshinImpact_4.1.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.006) | Game, part 6 of 7 | 10.00 GiB | `f4bddee45838cffd8e558e37762cac30` |
| [GenshinImpact_4.1.0.zip.007](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/GenshinImpact_4.1.0.zip.007) | Game, part 7 of 7 | 922.33 MiB | `af703fb9f1760ab690a8c6eccfb0b039` |
| [Audio_Chinese_4.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/Audio_Chinese_4.1.0.zip) | Chinese voice-over | 11.07 GiB | `2255c0b293c399612d0c90ca82d1538d` |
| [Audio_English(US)_4.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/Audio_English%28US%29_4.1.0.zip) | English voice-over | 12.77 GiB | `183ab4443c35903c3d6a66c8ca3b3f41` |
| [Audio_Japanese_4.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/Audio_Japanese_4.1.0.zip) | Japanese voice-over | 14.44 GiB | `c47e36a3784391cc65ff9fdecd7925e8` |
| [Audio_Korean_4.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/Audio_Korean_4.1.0.zip) | Korean voice-over | 10.84 GiB | `e1aa1dbe48a6e1bbb69b30a1da7af00f` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230916101725_v2XCKuuvzCVh3BdI/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.0.1

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.0.1_4.1.0_hdiff_QSwRBvbj1gaAs7zG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.0.1_4.1.0_hdiff_QSwRBvbj1gaAs7zG.zip) | Game | 13.58 GiB | `f8fe75d3400a8918236c42b242218228` |
| [zh-cn_4.0.1_4.1.0_hdiff_jaS15ubBrgoGXwcH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.0.1_4.1.0_hdiff_jaS15ubBrgoGXwcH.zip) | Chinese voice-over | 523.93 MiB | `02d6201951bed27c418a11ececced203` |
| [en-us_4.0.1_4.1.0_hdiff_xfopIGCW5BzYXdrv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.0.1_4.1.0_hdiff_xfopIGCW5BzYXdrv.zip) | English voice-over | 678.62 MiB | `da0fe1a2099f73501700e89a374fe5ec` |
| [ja-jp_4.0.1_4.1.0_hdiff_otzl30WOYGHJnqec.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.0.1_4.1.0_hdiff_otzl30WOYGHJnqec.zip) | Japanese voice-over | 649.36 MiB | `c245337aad41305b51b4db2731a0b8c6` |
| [ko-kr_4.0.1_4.1.0_hdiff_wPMhbF6pWidXovEe.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.0.1_4.1.0_hdiff_wPMhbF6pWidXovEe.zip) | Korean voice-over | 532.69 MiB | `45b22080ba7ef81265a16e96c593738b` |

#### Update from 4.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.0.0_4.1.0_hdiff_NHu1vzpoal6X54fQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.0.0_4.1.0_hdiff_NHu1vzpoal6X54fQ.zip) | Game | 13.58 GiB | `3555cbc552417b679cce4f0b7c277976` |
| [zh-cn_4.0.0_4.1.0_hdiff_HluXU6oGnSwvFxWZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.0.0_4.1.0_hdiff_HluXU6oGnSwvFxWZ.zip) | Chinese voice-over | 523.93 MiB | `c89c3628db171aa27471ab2eca626ee3` |
| [en-us_4.0.0_4.1.0_hdiff_rXQdZ8kuNhMbEUFJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.0.0_4.1.0_hdiff_rXQdZ8kuNhMbEUFJ.zip) | English voice-over | 678.62 MiB | `cb38f6e8bdd983b99d6c81e9aaa779ff` |
| [ja-jp_4.0.0_4.1.0_hdiff_2SE4Oagr0ZNxF9lY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.0.0_4.1.0_hdiff_2SE4Oagr0ZNxF9lY.zip) | Japanese voice-over | 649.36 MiB | `52e1d6e75bfdb10dbdab83ae45d78865` |
| [ko-kr_4.0.0_4.1.0_hdiff_Kg5FVU1AJ4hEMzC8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.0.0_4.1.0_hdiff_Kg5FVU1AJ4hEMzC8.zip) | Korean voice-over | 532.69 MiB | `d983bbe80f51839e221f6f58507fb7e6` |

#### Update from 3.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.8.0_4.1.0_hdiff_oC7ajt1xOVZnp54s.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.8.0_4.1.0_hdiff_oC7ajt1xOVZnp54s.zip) | Game | 37.48 GiB | `b2b8c4bc82b4a40af1a0efe0b55e5323` |
| [zh-cn_3.8.0_4.1.0_hdiff_X1jCLiI8Hpbr4Vxt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.8.0_4.1.0_hdiff_X1jCLiI8Hpbr4Vxt.zip) | Chinese voice-over | 1.12 GiB | `95bb42386bba79c9a871a32d5e6e5aaf` |
| [en-us_3.8.0_4.1.0_hdiff_McWPihmf2D7K95xX.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.8.0_4.1.0_hdiff_McWPihmf2D7K95xX.zip) | English voice-over | 1.42 GiB | `ed2e3be1f4e0cd76908205d30ddd4353` |
| [ja-jp_3.8.0_4.1.0_hdiff_arBPjs7FmUKtDpIo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.8.0_4.1.0_hdiff_arBPjs7FmUKtDpIo.zip) | Japanese voice-over | 1.33 GiB | `5dc89136d47dd1aec54b595153000d92` |
| [ko-kr_3.8.0_4.1.0_hdiff_Mn70mAKotTFVwepf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.8.0_4.1.0_hdiff_Mn70mAKotTFVwepf.zip) | Korean voice-over | 1.13 GiB | `1b4461783ffe1f10bc323309673a8ed0` |

### 4.0.1

Seen in the API as the live version from 2023-08-23 to 2023-09-26 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.0.1.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.001) | Game, part 1 of 6 | 10.00 GiB | `75e48096a77fd2a728d3eb261135c907` |
| [GenshinImpact_4.0.1.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.002) | Game, part 2 of 6 | 10.00 GiB | `6c6d52cd6b62deacc8a2ab1e056a6122` |
| [GenshinImpact_4.0.1.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.003) | Game, part 3 of 6 | 10.00 GiB | `8b75c40aee270f227710857eacdb654c` |
| [GenshinImpact_4.0.1.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.004) | Game, part 4 of 6 | 10.00 GiB | `7bf5d629f85f20c7784045266ac4a9e1` |
| [GenshinImpact_4.0.1.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.005) | Game, part 5 of 6 | 10.00 GiB | `26f042472fac3308f7fddd62d61155f5` |
| [GenshinImpact_4.0.1.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/GenshinImpact_4.0.1.zip.006) | Game, part 6 of 6 | 4.99 GiB | `9982d6bdb0fc65fa724a9b04e5605d55` |
| [Audio_Chinese_4.0.1.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/Audio_Chinese_4.0.1.zip) | Chinese voice-over | 10.45 GiB | `9ad5e51ecc20431c732d8dae3f8621a9` |
| [Audio_English(US)_4.0.1.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/Audio_English%28US%29_4.0.1.zip) | English voice-over | 12.04 GiB | `0928ddbf27eb0886588dfd10f38bdaab` |
| [Audio_Japanese_4.0.1.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/Audio_Japanese_4.0.1.zip) | Japanese voice-over | 13.59 GiB | `e79f9d0686fc1ee8ff01b5bc3dba0536` |
| [Audio_Korean_4.0.1.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/Audio_Korean_4.0.1.zip) | Korean voice-over | 10.23 GiB | `e35f1442360ab336d5c6d971440e18e7` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230821151229_NAlBxGiyKlVXWZQJ/ScatteredFiles` (file list in `pkg_version`).

#### Update from 4.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_4.0.0_4.0.1_hdiff_JNhULRDxB2ntig48.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_4.0.0_4.0.1_hdiff_JNhULRDxB2ntig48.zip) | Game | 395.14 MiB | `d9eddebc5a43900dda40b1aa993a35e0` |
| [zh-cn_4.0.0_4.0.1_hdiff_yx71braIALoB502f.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_4.0.0_4.0.1_hdiff_yx71braIALoB502f.zip) | Chinese voice-over | 0.02 KiB | `76cdb2bad9582d23c1f6f4d868218d6c` |
| [en-us_4.0.0_4.0.1_hdiff_lHMOr4c5a1mtS836.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_4.0.0_4.0.1_hdiff_lHMOr4c5a1mtS836.zip) | English voice-over | 0.02 KiB | `76cdb2bad9582d23c1f6f4d868218d6c` |
| [ja-jp_4.0.0_4.0.1_hdiff_0ja7pUAEQLc6KRvi.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_4.0.0_4.0.1_hdiff_0ja7pUAEQLc6KRvi.zip) | Japanese voice-over | 0.02 KiB | `76cdb2bad9582d23c1f6f4d868218d6c` |
| [ko-kr_4.0.0_4.0.1_hdiff_r1PQymX6DeFWcdv5.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_4.0.0_4.0.1_hdiff_r1PQymX6DeFWcdv5.zip) | Korean voice-over | 0.02 KiB | `76cdb2bad9582d23c1f6f4d868218d6c` |

#### Update from 3.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.8.0_4.0.1_hdiff_h69nuw1XgeJ2qA3i.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.8.0_4.0.1_hdiff_h69nuw1XgeJ2qA3i.zip) | Game | 31.46 GiB | `2c43482670866e5950fea8052790a16c` |
| [zh-cn_3.8.0_4.0.1_hdiff_zwvWKSEmiZ3CfaVT.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.8.0_4.0.1_hdiff_zwvWKSEmiZ3CfaVT.zip) | Chinese voice-over | 627.52 MiB | `76e3be1be2c805f94d4ffd7d394fd743` |
| [en-us_3.8.0_4.0.1_hdiff_UEbY1IBCtfZz5HRs.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.8.0_4.0.1_hdiff_UEbY1IBCtfZz5HRs.zip) | English voice-over | 778.28 MiB | `dc86b267ec0ddbd088344f010cf0d3ad` |
| [ja-jp_3.8.0_4.0.1_hdiff_zpgFW8dkZ41baHns.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.8.0_4.0.1_hdiff_zpgFW8dkZ41baHns.zip) | Japanese voice-over | 713.14 MiB | `9a1a601cdbdd9e21870a8743aca00eca` |
| [ko-kr_3.8.0_4.0.1_hdiff_i32BMSKVAygECRaZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.8.0_4.0.1_hdiff_i32BMSKVAygECRaZ.zip) | Korean voice-over | 619.41 MiB | `3755a167ecdfa63b33a08a4117619ba8` |

### 4.0.0

Seen in the API as the live version from 2023-08-16 to 2023-08-22, pre-download from 2023-08-14 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_4.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip) | Game, single archive | 56.16 GiB |  |
| [GenshinImpact_4.0.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.001) | Game, part 1 of 6 | 10.00 GiB | `341db4fb0e88f8d72d4e52a3f67954ea` |
| [GenshinImpact_4.0.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.002) | Game, part 2 of 6 | 10.00 GiB | `f399e762b2798671d6d5c79f61643645` |
| [GenshinImpact_4.0.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.003) | Game, part 3 of 6 | 10.00 GiB | `f77427b26c61097daf3630806953cca5` |
| [GenshinImpact_4.0.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.004) | Game, part 4 of 6 | 10.00 GiB | `0efc329959216ccc624c7d4d8b7771bd` |
| [GenshinImpact_4.0.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.005) | Game, part 5 of 6 | 10.00 GiB | `a39dcc95b8e75ee3281a87d9ae203f96` |
| [GenshinImpact_4.0.0.zip.006](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/GenshinImpact_4.0.0.zip.006) | Game, part 6 of 6 | 4.99 GiB | `71ee85d6d7f73a533538e15ac2cacf74` |
| [Audio_Chinese_4.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/Audio_Chinese_4.0.0.zip) | Chinese voice-over | 10.45 GiB | `01769babfc856d24c31fe9e928db0cf7` |
| [Audio_English(US)_4.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/Audio_English%28US%29_4.0.0.zip) | English voice-over | 12.04 GiB | `1154991a6f1c7cb4b1a95f9cc0e38e28` |
| [Audio_Japanese_4.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/Audio_Japanese_4.0.0.zip) | Japanese voice-over | 13.59 GiB | `8eda06aa393d01cbb6d0d0da1835724e` |
| [Audio_Korean_4.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/Audio_Korean_4.0.0.zip) | Korean voice-over | 10.23 GiB | `a41b20f18637e1293c89990df6a12434` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230804185804_eTmE8EZjJZdAJapq/ScatteredFiles` (file list in `pkg_version`).

#### Update from 3.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.8.0_4.0.0_hdiff_Brgs45clx6Teu1wO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.8.0_4.0.0_hdiff_Brgs45clx6Teu1wO.zip) | Game | 31.46 GiB | `d2cc88a7e6f3b9081bfd5d1f155bc6fb` |
| [zh-cn_3.8.0_4.0.0_hdiff_mEnWBIaviZwz9rqH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.8.0_4.0.0_hdiff_mEnWBIaviZwz9rqH.zip) | Chinese voice-over | 627.52 MiB | `4b8ff5c7c9192275afcf3c894a0cd748` |
| [en-us_3.8.0_4.0.0_hdiff_M9dPNbJx1hiWoZuY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.8.0_4.0.0_hdiff_M9dPNbJx1hiWoZuY.zip) | English voice-over | 778.28 MiB | `a3247294906db6bad72f166c4d8e256b` |
| [ja-jp_3.8.0_4.0.0_hdiff_lN024VYf9jhzPbBo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.8.0_4.0.0_hdiff_lN024VYf9jhzPbBo.zip) | Japanese voice-over | 713.14 MiB | `860d203710e0e9fd54bad6d64ad7bed4` |
| [ko-kr_3.8.0_4.0.0_hdiff_GR7WVCPuATEyjxK2.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.8.0_4.0.0_hdiff_GR7WVCPuATEyjxK2.zip) | Korean voice-over | 619.41 MiB | `a363160326f45589ba71844a82ac580a` |

#### Update from 3.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.7.0_4.0.0_hdiff_ny78JZtwupS1F6PL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.7.0_4.0.0_hdiff_ny78JZtwupS1F6PL.zip) | Game | 32.25 GiB | `0a39917c0ea950f67b98e422867b0ffc` |
| [zh-cn_3.7.0_4.0.0_hdiff_6bH0wmErf3kdAinZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.7.0_4.0.0_hdiff_6bH0wmErf3kdAinZ.zip) | Chinese voice-over | 933.94 MiB | `5b48678a0364c243c5e4ad98e43eced8` |
| [en-us_3.7.0_4.0.0_hdiff_5XpQ2rw03n8e9Yom.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.7.0_4.0.0_hdiff_5XpQ2rw03n8e9Yom.zip) | English voice-over | 1.16 GiB | `18fc6b5e94ad661248365e5411d22cd8` |
| [ja-jp_3.7.0_4.0.0_hdiff_OSM7inRWL8vjg0UH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.7.0_4.0.0_hdiff_OSM7inRWL8vjg0UH.zip) | Japanese voice-over | 1.08 GiB | `9957b3f9ea4f1a2c4f6f92d0a40f6418` |
| [ko-kr_3.7.0_4.0.0_hdiff_mFR6qzxTSlaYjdOh.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.7.0_4.0.0_hdiff_mFR6qzxTSlaYjdOh.zip) | Korean voice-over | 927.98 MiB | `8d7fe2c5097a7565d8b95760c1255f75` |

</details>

<details>
<summary><b>Version 3.x</b>: 3.0.0 to 3.8.0 (9 versions)</summary>

Jump to: [3.8.0](#380) · [3.7.0](#370) · [3.6.0](#360) · [3.5.0](#350) · [3.4.0](#340) · [3.3.0](#330) · [3.2.0](#320) · [3.1.0](#310) · [3.0.0](#300)

### 3.8.0

Seen in the API as the live version from 2023-07-05 to 2023-08-15, pre-download from 2023-07-03 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip) | Game, single archive | 49.36 GiB |  |
| [GenshinImpact_3.8.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip.001) | Game, part 1 of 5 | 10.00 GiB | `f7ea3a2c2ca05073ae809f1164f1a50f` |
| [GenshinImpact_3.8.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip.002) | Game, part 2 of 5 | 10.00 GiB | `3ccb011b33ab1fc4a8e47d16a2868d84` |
| [GenshinImpact_3.8.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip.003) | Game, part 3 of 5 | 10.00 GiB | `4c89372799357cd9b3bb6a2a62d017cc` |
| [GenshinImpact_3.8.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip.004) | Game, part 4 of 5 | 10.00 GiB | `a41db2303e7e9579091a2269b2052d6b` |
| [GenshinImpact_3.8.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/GenshinImpact_3.8.0.zip.005) | Game, part 5 of 5 | 8.28 GiB | `1916328ca260b0a51e8de10b80fd4fa1` |
| [Audio_Chinese_3.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/Audio_Chinese_3.8.0.zip) | Chinese voice-over | 9.72 GiB | `cab24930c10ca0af4a48b8f806d0bd33` |
| [Audio_English(US)_3.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/Audio_English%28US%29_3.8.0.zip) | English voice-over | 11.21 GiB | `43adbc72d7c8967a852c8567721330ac` |
| [Audio_Japanese_3.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/Audio_Japanese_3.8.0.zip) | Japanese voice-over | 12.64 GiB | `c02d958dac54852ae52af813b2ee0814` |
| [Audio_Korean_3.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/Audio_Korean_3.8.0.zip) | Korean voice-over | 9.52 GiB | `82a2f403d24fc77037c1ad93f7c0ad85` |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230625120419_BEKxsz5x6A7uZvGW/ScatteredFiles` (file list in `pkg_version`).

#### Update from 3.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.7.0_3.8.0_hdiff_wCu26d7XVB51UQbi.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.7.0_3.8.0_hdiff_wCu26d7XVB51UQbi.zip) | Game | 10.27 GiB | `ababc2ba9eb4244d7006c1ecdc69856f` |
| [zh-cn_3.7.0_3.8.0_hdiff_iouZLFje3a1nfxD5.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.7.0_3.8.0_hdiff_iouZLFje3a1nfxD5.zip) | Chinese voice-over | 306.46 MiB | `f2f9ca52cb7253c2ccb73568c610e956` |
| [en-us_3.7.0_3.8.0_hdiff_g4YF0QKZoAx1SBci.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.7.0_3.8.0_hdiff_g4YF0QKZoAx1SBci.zip) | English voice-over | 404.85 MiB | `8ab2acbd3e48ce6c1f3f9a1b6503dae3` |
| [ja-jp_3.7.0_3.8.0_hdiff_femTrvp6gAP05RM1.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.7.0_3.8.0_hdiff_femTrvp6gAP05RM1.zip) | Japanese voice-over | 396.80 MiB | `2afe15426ccf0360fe16e2d7377e65fc` |
| [ko-kr_3.7.0_3.8.0_hdiff_0mETUysXGcKuRja2.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.7.0_3.8.0_hdiff_0mETUysXGcKuRja2.zip) | Korean voice-over | 308.65 MiB | `4c76d498aa51fd12fc7bc9b1d7b62383` |

#### Update from 3.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.6.0_3.8.0_hdiff_y846MR7XilnIdUoK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.6.0_3.8.0_hdiff_y846MR7XilnIdUoK.zip) | Game | 11.66 GiB | `2b75516c57b8d908f3be4003a6d00bb3` |
| [zh-cn_3.6.0_3.8.0_hdiff_s67RMzLSKN09IljC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.6.0_3.8.0_hdiff_s67RMzLSKN09IljC.zip) | Chinese voice-over | 676.75 MiB | `7956844519fbbccbe735e17a575b227e` |
| [en-us_3.6.0_3.8.0_hdiff_qTPOvM2gpjHyakWZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.6.0_3.8.0_hdiff_qTPOvM2gpjHyakWZ.zip) | English voice-over | 886.55 MiB | `0814dc67b60b812d21ea02ba1d7787e2` |
| [ja-jp_3.6.0_3.8.0_hdiff_CqTocjlFw2s4POrB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.6.0_3.8.0_hdiff_CqTocjlFw2s4POrB.zip) | Japanese voice-over | 876.03 MiB | `1330acbbb895fd3ac2d5cf37b2d50722` |
| [ko-kr_3.6.0_3.8.0_hdiff_6fNz8wbocdnYOD1Q.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.6.0_3.8.0_hdiff_6fNz8wbocdnYOD1Q.zip) | Korean voice-over | 660.93 MiB | `73cfd478548eaee7a07fde9c349d3076` |

### 3.7.0

Seen in the API as the live version from 2023-05-24 to 2023-07-04, pre-download from 2023-05-22 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip) | Game, single archive | 48.60 GiB |  |
| [GenshinImpact_3.7.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip.001) | Game, part 1 of 5 | 10.00 GiB | `bd193afb5cf2d16e1e048bbdf07dde66` |
| [GenshinImpact_3.7.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip.002) | Game, part 2 of 5 | 10.00 GiB | `63353ee59d60619313fc9032fa87d697` |
| [GenshinImpact_3.7.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip.003) | Game, part 3 of 5 | 10.00 GiB | `989d7df8b11fe781322eae27260c86f3` |
| [GenshinImpact_3.7.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip.004) | Game, part 4 of 5 | 10.00 GiB | `91fc90cb6f09712619f9943a5b1603cb` |
| [GenshinImpact_3.7.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/GenshinImpact_3.7.0.zip.005) | Game, part 5 of 5 | 7.54 GiB | `7350eb36815994a0fe93ea94544198bf` |
| [Audio_Chinese_3.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/Audio_Chinese_3.7.0.zip) | Chinese voice-over | 9.38 GiB | `2b0fd39275a2974fc85ee1e9735b42b7` (published for 10071698310 bytes; the served file is 10075896724 bytes, so it may not match) |
| [Audio_English(US)_3.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/Audio_English%28US%29_3.7.0.zip) | English voice-over | 10.79 GiB | `43b360b3f5c71feb951a204f81a1d232` (published for 11585674281 bytes; the served file is 11583388687 bytes, so it may not match) |
| [Audio_Japanese_3.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/Audio_Japanese_3.7.0.zip) | Japanese voice-over | 12.15 GiB | `516ad9333520473c6ff3ac7fe6d1a5e9` (published for 13045963162 bytes; the served file is 13044193326 bytes, so it may not match) |
| [Audio_Korean_3.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/Audio_Korean_3.7.0.zip) | Korean voice-over | 9.18 GiB | `07251fd076e60111716c4ca970965a87` (published for 9856949945 bytes; the served file is 9858003494 bytes, so it may not match) |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230513200104_2odHBzbUAP5IOIvE/ScatteredFiles` (file list in `pkg_version`).

#### Update from 3.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.6.0_3.7.0_hdiff_SIQuq5CfVPEXnBNy.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.6.0_3.7.0_hdiff_SIQuq5CfVPEXnBNy.zip) | Game | 8.35 GiB | `f9c4f21fa785fbc48845c44b13472b9e` |
| [zh-cn_3.6.0_3.7.0_hdiff_a8YmD1HvK69sp7cS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.6.0_3.7.0_hdiff_a8YmD1HvK69sp7cS.zip) | Chinese voice-over | 371.13 MiB | `27247ad4e102f7dfc927c265e9191271` |
| [en-us_3.6.0_3.7.0_hdiff_K8ei6yaM5B1Edj9g.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.6.0_3.7.0_hdiff_K8ei6yaM5B1Edj9g.zip) | English voice-over | 482.55 MiB | `e6fe450c4e6d7cd04051e1236cda4fd0` |
| [ja-jp_3.6.0_3.7.0_hdiff_x3n1GpVAtsylQwv7.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.6.0_3.7.0_hdiff_x3n1GpVAtsylQwv7.zip) | Japanese voice-over | 480.29 MiB | `fe3a4c481f33f3dae646d164b4491685` |
| [ko-kr_3.6.0_3.7.0_hdiff_cT17sLmPlZfKkQuJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.6.0_3.7.0_hdiff_cT17sLmPlZfKkQuJ.zip) | Korean voice-over | 353.05 MiB | `19bd9c8b418adce59433bab4bf3d402a` |

#### Update from 3.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.5.0_3.7.0_hdiff_fy67SAIbUXdo1kMp.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.5.0_3.7.0_hdiff_fy67SAIbUXdo1kMp.zip) | Game | 12.02 GiB | `bcec39f882910cddde95d05272dc7f69` |
| [zh-cn_3.5.0_3.7.0_hdiff_IfGdXle2ZrkW96Fv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.5.0_3.7.0_hdiff_IfGdXle2ZrkW96Fv.zip) | Chinese voice-over | 836.37 MiB | `4326eb9e996ffdf0520a918e014dbb37` |
| [en-us_3.5.0_3.7.0_hdiff_dLDNvaoEmZBXePk5.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.5.0_3.7.0_hdiff_dLDNvaoEmZBXePk5.zip) | English voice-over | 1.19 GiB | `325bb01db9737ea2be14f4d2d2f180c4` |
| [ja-jp_3.5.0_3.7.0_hdiff_fqPtTy5jHG8412og.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.5.0_3.7.0_hdiff_fqPtTy5jHG8412og.zip) | Japanese voice-over | 1.00 GiB | `4645cad784d4776c2e7c48a08867a466` |
| [ko-kr_3.5.0_3.7.0_hdiff_5ZaDdoGt2q7yCJwP.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.5.0_3.7.0_hdiff_5ZaDdoGt2q7yCJwP.zip) | Korean voice-over | 835.58 MiB | `0f2b640755f87bb04b49190c677a36d1` |

### 3.6.0

Seen in the API as the live version from 2023-04-12 to 2023-05-23, pre-download from 2023-04-10 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip) | Game, single archive | 47.62 GiB | `2686795437c83749e975333216d99edf` (published for 51131585659 bytes; the served file is 51134869445 bytes, so it may not match) |
| [GenshinImpact_3.6.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip.001) | Game, part 1 of 5 | 10.00 GiB | `dfb4a52b7c1df52b575399c88387a834` |
| [GenshinImpact_3.6.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip.002) | Game, part 2 of 5 | 10.00 GiB | `732de764330a7e91caf797ea42d033eb` |
| [GenshinImpact_3.6.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip.003) | Game, part 3 of 5 | 10.00 GiB | `996ba698498e61f5c17e3930a34611a1` |
| [GenshinImpact_3.6.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip.004) | Game, part 4 of 5 | 10.00 GiB | `6af5899787aa820f55b03a2778a99d3e` |
| [GenshinImpact_3.6.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/GenshinImpact_3.6.0.zip.005) | Game, part 5 of 5 | 6.56 GiB | `4270ab2beb4fbeec4b97018c3c4dc66a` |
| [Audio_Chinese_3.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/Audio_Chinese_3.6.0.zip) | Chinese voice-over | 8.97 GiB | `69a082ebb35929d1392318c2c4ec045c` (published for 9631464162 bytes; the served file is 9626507183 bytes, so it may not match) |
| [Audio_English(US)_3.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/Audio_English%28US%29_3.6.0.zip) | English voice-over | 10.28 GiB | `91efd98ed1eb3b39ab93588789d402da` (published for 11038065951 bytes; the served file is 11041923626 bytes, so it may not match) |
| [Audio_Japanese_3.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/Audio_Japanese_3.6.0.zip) | Japanese voice-over | 11.56 GiB | `565cb3fddbe213c49308749b44e9a562` (published for 12412455486 bytes; the served file is 12412394814 bytes, so it may not match) |
| [Audio_Korean_3.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/Audio_Korean_3.6.0.zip) | Korean voice-over | 8.79 GiB | `7fc1859050f138fb13fe4eb6e2fb8573` (published for 9438190633 bytes; the served file is 9434740241 bytes, so it may not match) |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230331200258_xR748JQdRJ19pQH0/ScatteredFiles` (file list in `pkg_version`).

#### Update from 3.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.5.0_3.6.0_hdiff_70c4ojbL1mfTwY3Q.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.5.0_3.6.0_hdiff_70c4ojbL1mfTwY3Q.zip) | Game | 9.88 GiB | `ce07abfa216757b35da43401868dc65e` |
| [zh-cn_3.5.0_3.6.0_hdiff_hJrb6Zzsq0m7GQHK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.5.0_3.6.0_hdiff_hJrb6Zzsq0m7GQHK.zip) | Chinese voice-over | 465.30 MiB | `fa4420603967aae67451050b5a61c806` |
| [en-us_3.5.0_3.6.0_hdiff_xS3HAI51ksDQOl8c.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.5.0_3.6.0_hdiff_xS3HAI51ksDQOl8c.zip) | English voice-over | 740.14 MiB | `500fa465149a7dcc6a7a72b859b82302` |
| [ja-jp_3.5.0_3.6.0_hdiff_6ewT53aBmuNbdPJE.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.5.0_3.6.0_hdiff_6ewT53aBmuNbdPJE.zip) | Japanese voice-over | 545.94 MiB | `432385e20612038143e664ae6717e16d` |
| [ko-kr_3.5.0_3.6.0_hdiff_GDlPHvUgyXQwje9C.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.5.0_3.6.0_hdiff_GDlPHvUgyXQwje9C.zip) | Korean voice-over | 482.66 MiB | `30e21b012374f55fd5d123a2840a2e04` |

#### Update from 3.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.4.0_3.6.0_hdiff_t30d4K7DUnwNcmja.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.4.0_3.6.0_hdiff_t30d4K7DUnwNcmja.zip) | Game | 11.18 GiB | `69609f7a420b46ebe8ddd256625bed63` |
| [zh-cn_3.4.0_3.6.0_hdiff_wsQFNvrIdV4qLKCG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.4.0_3.6.0_hdiff_wsQFNvrIdV4qLKCG.zip) | Chinese voice-over | 919.58 MiB | `62780d82af053e46483b2f7b3f3148c0` |
| [en-us_3.4.0_3.6.0_hdiff_gmhB8dVWQJRKlyLj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.4.0_3.6.0_hdiff_gmhB8dVWQJRKlyLj.zip) | English voice-over | 1.27 GiB | `146cce1f94ff0499e582666c066f59ad` |
| [ja-jp_3.4.0_3.6.0_hdiff_akC64FrYKsS3ow5M.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.4.0_3.6.0_hdiff_akC64FrYKsS3ow5M.zip) | Japanese voice-over | 1.06 GiB | `ac2bd9bcdae4185766f3a27ac5e206d9` |
| [ko-kr_3.4.0_3.6.0_hdiff_bf2wae6FH7npkOgt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.4.0_3.6.0_hdiff_bf2wae6FH7npkOgt.zip) | Korean voice-over | 917.22 MiB | `9c5a77c7cb77da697764b5e57ee5d69c` |

### 3.5.0

Seen in the API as the live version from 2023-03-17 to 2023-04-11 (UTC).

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip) | Game, single archive | 44.66 GiB | `1e6334da5268caf47e17d2fa12e9bef1` (published for 47953309860 bytes; the served file is 47955087332 bytes, so it may not match) |
| [GenshinImpact_3.5.0.zip.001](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip.001) | Game, part 1 of 5 | 10.00 GiB | `4551716f735af6a5cf8b8565250d1eca` |
| [GenshinImpact_3.5.0.zip.002](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip.002) | Game, part 2 of 5 | 10.00 GiB | `869529539a9cedde466b68ad5c7de4d1` |
| [GenshinImpact_3.5.0.zip.003](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip.003) | Game, part 3 of 5 | 10.00 GiB | `c2704558567c6c02c0c676acf23d8efa` |
| [GenshinImpact_3.5.0.zip.004](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip.004) | Game, part 4 of 5 | 10.00 GiB | `2bf1060bf1b0d743562b3330eb6e2f7d` |
| [GenshinImpact_3.5.0.zip.005](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/GenshinImpact_3.5.0.zip.005) | Game, part 5 of 5 | 3.61 GiB | `0f5059a7538b35e959f7a4c0c4acccf4` |
| [Audio_Chinese_3.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/Audio_Chinese_3.5.0.zip) | Chinese voice-over | 8.44 GiB | `fa18f4adccdbffb7d19ff4f077250571` (published for 9062380995 bytes; the served file is 9062211706 bytes, so it may not match) |
| [Audio_English(US)_3.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/Audio_English%28US%29_3.5.0.zip) | English voice-over | 9.64 GiB | `b04087a83c0a9d588bd4dff665a2aba2` (published for 10350871184 bytes; the served file is 10352216788 bytes, so it may not match) |
| [Audio_Japanese_3.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/Audio_Japanese_3.5.0.zip) | Japanese voice-over | 10.84 GiB | `d8b640026a782cba817e59a196a6a747` (published for 11639361373 bytes; the served file is 11641998846 bytes, so it may not match) |
| [Audio_Korean_3.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/Audio_Korean_3.5.0.zip) | Korean voice-over | 8.25 GiB | `b947fc7b6dac078cc863223827e5c5f0` (published for 8858370048 bytes; the served file is 8862007457 bytes, so it may not match) |

Scattered files: `https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230220120928_iIYV9krGiWL06eeB/ScatteredFiles` (file list in `pkg_version`).

#### Update from 3.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.4.0_3.5.0_hdiff_rXYAZvelstu54E2c.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.4.0_3.5.0_hdiff_rXYAZvelstu54E2c.zip) | Game | 6.02 GiB | `50d65cef9f0ec192725876b7eb2f8650` |
| [zh-cn_3.4.0_3.5.0_hdiff_9T2Ac7RCKs0651G8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.4.0_3.5.0_hdiff_9T2Ac7RCKs0651G8.zip) | Chinese voice-over | 454.30 MiB | `e9cd6f32e5541adf1e77a007a681eb24` |
| [en-us_3.4.0_3.5.0_hdiff_bCHgs9NcSPI54hK3.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.4.0_3.5.0_hdiff_bCHgs9NcSPI54hK3.zip) | English voice-over | 571.50 MiB | `9490781f902796999b7d290cdc6c6038` |
| [ja-jp_3.4.0_3.5.0_hdiff_7GSo18Z9C2rOiDjp.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.4.0_3.5.0_hdiff_7GSo18Z9C2rOiDjp.zip) | Japanese voice-over | 543.32 MiB | `7452317bf85bda6764ae6006d75950f2` |
| [ko-kr_3.4.0_3.5.0_hdiff_CvwZI7isgHbkAanc.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.4.0_3.5.0_hdiff_CvwZI7isgHbkAanc.zip) | Korean voice-over | 434.66 MiB | `0b3129ddbf83f8735231966be080723a` |

#### Update from 3.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.3.0_3.5.0_hdiff_jFk1USEdyHepP3uG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.3.0_3.5.0_hdiff_jFk1USEdyHepP3uG.zip) | Game | 7.73 GiB | `da20792bcfcf890583f66efbc5a6b5e4` |
| [zh-cn_3.3.0_3.5.0_hdiff_DuVQLAq1nkay7Yo4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.3.0_3.5.0_hdiff_DuVQLAq1nkay7Yo4.zip) | Chinese voice-over | 828.53 MiB | `296710d5deaf87617932f8a609b1ea99` |
| [en-us_3.3.0_3.5.0_hdiff_ep1y7wDFiLnjb5rW.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.3.0_3.5.0_hdiff_ep1y7wDFiLnjb5rW.zip) | English voice-over | 1.00 GiB | `cae41a4a77af6fc62fa0884cb3df3baf` |
| [ja-jp_3.3.0_3.5.0_hdiff_lcO03yi4gv5jhmwG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.3.0_3.5.0_hdiff_lcO03yi4gv5jhmwG.zip) | Japanese voice-over | 994.23 MiB | `78b460abdb60a27290384af109083327` |
| [ko-kr_3.3.0_3.5.0_hdiff_YVM8xuZhWgNnjHAI.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.3.0_3.5.0_hdiff_YVM8xuZhWgNnjHAI.zip) | Korean voice-over | 788.16 MiB | `7c039cceabd052e9a29e1c69a5f7cb53` |

### 3.4.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230109135018_10QhExKHwAoa4ecr/GenshinImpact_3.4.0.zip) (dead link, HTTP 404) | Game | unknown |  |
| [Audio_Chinese_3.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230109135018_10QhExKHwAoa4ecr/Audio_Chinese_3.4.0.zip) (dead link, HTTP 404) | Chinese voice-over | unknown |  |
| [Audio_English(US)_3.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230109135018_10QhExKHwAoa4ecr/Audio_English%28US%29_3.4.0.zip) (dead link, HTTP 404) | English voice-over | unknown |  |
| [Audio_Japanese_3.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230109135018_10QhExKHwAoa4ecr/Audio_Japanese_3.4.0.zip) (dead link, HTTP 404) | Japanese voice-over | unknown |  |
| [Audio_Korean_3.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20230109135018_10QhExKHwAoa4ecr/Audio_Korean_3.4.0.zip) (dead link, HTTP 404) | Korean voice-over | unknown |  |

#### Update from 3.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.3.0_3.4.0_hdiff_IxKGMXhCLzZAJqUm.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.3.0_3.4.0_hdiff_IxKGMXhCLzZAJqUm.zip) (dead link, HTTP 404) | Game | unknown |  |
| [zh-cn_3.3.0_3.4.0_hdiff_CwIhpT8VfBUqlSEg.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.3.0_3.4.0_hdiff_CwIhpT8VfBUqlSEg.zip) | Chinese voice-over | 374.34 MiB |  |
| [en-us_3.3.0_3.4.0_hdiff_kXarZuSHnoEYzWiq.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.3.0_3.4.0_hdiff_kXarZuSHnoEYzWiq.zip) | English voice-over | 455.40 MiB |  |
| [ja-jp_3.3.0_3.4.0_hdiff_4RnJjQLArch3yGOt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.3.0_3.4.0_hdiff_4RnJjQLArch3yGOt.zip) | Japanese voice-over | 451.21 MiB |  |
| [ko-kr_3.3.0_3.4.0_hdiff_j3fribpBSCJaMUnE.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.3.0_3.4.0_hdiff_j3fribpBSCJaMUnE.zip) | Korean voice-over | 353.58 MiB |  |

#### Update from 3.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.2.0_3.4.0_hdiff_2Tv5e1BCLFnW0dPG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.2.0_3.4.0_hdiff_2Tv5e1BCLFnW0dPG.zip) (dead link, HTTP 404) | Game | unknown |  |
| [zh-cn_3.2.0_3.4.0_hdiff_sJLa0SElyqR5dFBo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.2.0_3.4.0_hdiff_sJLa0SElyqR5dFBo.zip) | Chinese voice-over | 774.57 MiB |  |
| [en-us_3.2.0_3.4.0_hdiff_Onc2ZSWqiHMLh80p.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.2.0_3.4.0_hdiff_Onc2ZSWqiHMLh80p.zip) | English voice-over | 975.69 MiB |  |
| [ja-jp_3.2.0_3.4.0_hdiff_QUa8erWEPHFwgGvc.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.2.0_3.4.0_hdiff_QUa8erWEPHFwgGvc.zip) | Japanese voice-over | 921.14 MiB |  |
| [ko-kr_3.2.0_3.4.0_hdiff_u1pfvG0C8lakLRKT.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.2.0_3.4.0_hdiff_u1pfvG0C8lakLRKT.zip) | Korean voice-over | 741.34 MiB |  |

### 3.3.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221128113321_cOH1qEM6iRagjJC6/GenshinImpact_3.3.0.zip) | Game | 42.50 GiB |  |
| [Audio_Chinese_3.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221128113321_cOH1qEM6iRagjJC6/Audio_Chinese_3.3.0.zip) | Chinese voice-over | 7.49 GiB |  |
| [Audio_English(US)_3.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221128113321_cOH1qEM6iRagjJC6/Audio_English%28US%29_3.3.0.zip) | English voice-over | 8.55 GiB |  |
| [Audio_Japanese_3.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221128113321_cOH1qEM6iRagjJC6/Audio_Japanese_3.3.0.zip) | Japanese voice-over | 9.55 GiB |  |
| [Audio_Korean_3.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221128113321_cOH1qEM6iRagjJC6/Audio_Korean_3.3.0.zip) | Korean voice-over | 7.35 GiB |  |

#### Update from 3.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.2.0_3.3.0_hdiff_fiuTl2cLqp198WVG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.2.0_3.3.0_hdiff_fiuTl2cLqp198WVG.zip) | Game | 5.92 GiB |  |
| [zh-cn_3.2.0_3.3.0_hdiff_Y7EPh3tBmMbfwWac.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.2.0_3.3.0_hdiff_Y7EPh3tBmMbfwWac.zip) | Chinese voice-over | 400.28 MiB |  |
| [en-us_3.2.0_3.3.0_hdiff_Mtg67G54pwYTnuEA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.2.0_3.3.0_hdiff_Mtg67G54pwYTnuEA.zip) | English voice-over | 520.35 MiB |  |
| [ja-jp_3.2.0_3.3.0_hdiff_DE9YUdTvyOecHkih.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.2.0_3.3.0_hdiff_DE9YUdTvyOecHkih.zip) | Japanese voice-over | 470.71 MiB |  |
| [ko-kr_3.2.0_3.3.0_hdiff_jfd7VWceoaE0A4uB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.2.0_3.3.0_hdiff_jfd7VWceoaE0A4uB.zip) | Korean voice-over | 387.82 MiB |  |

#### Update from 3.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.1.0_3.3.0_hdiff_IGL1F0SYrChkDmvf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.1.0_3.3.0_hdiff_IGL1F0SYrChkDmvf.zip) | Game | 8.83 GiB |  |
| [zh-cn_3.1.0_3.3.0_hdiff_LSgFINh859GtE0fp.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.1.0_3.3.0_hdiff_LSgFINh859GtE0fp.zip) | Chinese voice-over | 936.65 MiB |  |
| [en-us_3.1.0_3.3.0_hdiff_IAkYRNMdX7OrEaHS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.1.0_3.3.0_hdiff_IAkYRNMdX7OrEaHS.zip) | English voice-over | 1.15 GiB |  |
| [ja-jp_3.1.0_3.3.0_hdiff_nYG9Lxsc5EUWXyrT.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.1.0_3.3.0_hdiff_nYG9Lxsc5EUWXyrT.zip) | Japanese voice-over | 1.10 GiB |  |
| [ko-kr_3.1.0_3.3.0_hdiff_zkv4RZ7Y5jfPGauO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.1.0_3.3.0_hdiff_zkv4RZ7Y5jfPGauO.zip) | Korean voice-over | 909.00 MiB |  |

### 3.2.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221024103618_h2e3o3zijYKEqHnQ/GenshinImpact_3.2.0.zip) (dead link, HTTP 404) | Game | unknown |  |
| [Audio_Chinese_3.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221024103618_h2e3o3zijYKEqHnQ/Audio_Chinese_3.2.0.zip) (dead link, HTTP 404) | Chinese voice-over | unknown |  |
| [Audio_English(US)_3.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221024103618_h2e3o3zijYKEqHnQ/Audio_English%28US%29_3.2.0.zip) (dead link, HTTP 404) | English voice-over | unknown |  |
| [Audio_Japanese_3.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221024103618_h2e3o3zijYKEqHnQ/Audio_Japanese_3.2.0.zip) (dead link, HTTP 404) | Japanese voice-over | unknown |  |
| [Audio_Korean_3.2.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20221024103618_h2e3o3zijYKEqHnQ/Audio_Korean_3.2.0.zip) (dead link, HTTP 404) | Korean voice-over | unknown |  |

#### Update from 3.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.1.0_3.2.0_hdiff_3B7InFiY1SDdN8Lv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.1.0_3.2.0_hdiff_3B7InFiY1SDdN8Lv.zip) (dead link, HTTP 404) | Game | unknown |  |
| [zh-cn_3.1.0_3.2.0_hdiff_R8oNZkyzjwxYFGCQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.1.0_3.2.0_hdiff_R8oNZkyzjwxYFGCQ.zip) | Chinese voice-over | 536.72 MiB |  |
| [en-us_3.1.0_3.2.0_hdiff_ezXCZESyB4JuGs5g.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.1.0_3.2.0_hdiff_ezXCZESyB4JuGs5g.zip) | English voice-over | 661.94 MiB |  |
| [ja-jp_3.1.0_3.2.0_hdiff_mHfx98gkNpB36szQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.1.0_3.2.0_hdiff_mHfx98gkNpB36szQ.zip) | Japanese voice-over | 652.15 MiB |  |
| [ko-kr_3.1.0_3.2.0_hdiff_dZwqgxaJQusz5e7E.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.1.0_3.2.0_hdiff_dZwqgxaJQusz5e7E.zip) | Korean voice-over | 521.26 MiB |  |

#### Update from 3.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.0.0_3.2.0_hdiff_0XM45IeyvRhgYCLj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.0.0_3.2.0_hdiff_0XM45IeyvRhgYCLj.zip) (dead link, HTTP 404) | Game | unknown |  |
| [zh-cn_3.0.0_3.2.0_hdiff_1n6f7JSmpBiHOvz8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.0.0_3.2.0_hdiff_1n6f7JSmpBiHOvz8.zip) | Chinese voice-over | 951.34 MiB |  |
| [en-us_3.0.0_3.2.0_hdiff_ZkinY2VSEILsuAwq.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.0.0_3.2.0_hdiff_ZkinY2VSEILsuAwq.zip) | English voice-over | 1.15 GiB |  |
| [ja-jp_3.0.0_3.2.0_hdiff_vchfdp38qJjbgTPz.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.0.0_3.2.0_hdiff_vchfdp38qJjbgTPz.zip) | Japanese voice-over | 1.11 GiB |  |
| [ko-kr_3.0.0_3.2.0_hdiff_WuHL2AB9UVvayFsl.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.0.0_3.2.0_hdiff_WuHL2AB9UVvayFsl.zip) | Korean voice-over | 946.84 MiB |  |

### 3.1.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220917165430_NyMmj1Ta9KlZKgCZ/GenshinImpact_3.1.0.zip) | Game | 37.89 GiB |  |
| [Audio_Chinese_3.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220917165430_NyMmj1Ta9KlZKgCZ/Audio_Chinese_3.1.0.zip) | Chinese voice-over | 8.19 GiB |  |
| [Audio_English(US)_3.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220917165430_NyMmj1Ta9KlZKgCZ/Audio_English%28US%29_3.1.0.zip) | English voice-over | 9.46 GiB |  |
| [Audio_Japanese_3.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220917165430_NyMmj1Ta9KlZKgCZ/Audio_Japanese_3.1.0.zip) | Japanese voice-over | 10.45 GiB |  |
| [Audio_Korean_3.1.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220917165430_NyMmj1Ta9KlZKgCZ/Audio_Korean_3.1.0.zip) | Korean voice-over | 8.08 GiB |  |

#### Update from 3.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_3.0.0_3.1.0_hdiff_JajVqm1LS2G5Wd3g.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_3.0.0_3.1.0_hdiff_JajVqm1LS2G5Wd3g.zip) | Game | 6.61 GiB |  |
| [zh-cn_3.0.0_3.1.0_hdiff_MZkPImOajGWFoehp.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_3.0.0_3.1.0_hdiff_MZkPImOajGWFoehp.zip) | Chinese voice-over | 540.32 MiB |  |
| [en-us_3.0.0_3.1.0_hdiff_J254a8oHvQTGMwu3.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_3.0.0_3.1.0_hdiff_J254a8oHvQTGMwu3.zip) | English voice-over | 702.70 MiB |  |
| [ja-jp_3.0.0_3.1.0_hdiff_3NZcVbSxzvXgMaq9.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_3.0.0_3.1.0_hdiff_3NZcVbSxzvXgMaq9.zip) | Japanese voice-over | 648.68 MiB |  |
| [ko-kr_3.0.0_3.1.0_hdiff_AhytWDJvIYoPU9nZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_3.0.0_3.1.0_hdiff_AhytWDJvIYoPU9nZ.zip) | Korean voice-over | 562.00 MiB |  |

#### Update from 2.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.8.0_3.1.0_hdiff_vQp41GPwdscT3UxV.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.8.0_3.1.0_hdiff_vQp41GPwdscT3UxV.zip) | Game | 12.12 GiB |  |
| [zh-cn_2.8.0_3.1.0_hdiff_xfwy75Z1dKCN6sEu.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.8.0_3.1.0_hdiff_xfwy75Z1dKCN6sEu.zip) | Chinese voice-over | 1.09 GiB |  |
| [en-us_2.8.0_3.1.0_hdiff_Potvg1euXns4a5WK.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.8.0_3.1.0_hdiff_Potvg1euXns4a5WK.zip) | English voice-over | 1.36 GiB |  |
| [ja-jp_2.8.0_3.1.0_hdiff_SFedYgkphMqGvtWZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.8.0_3.1.0_hdiff_SFedYgkphMqGvtWZ.zip) | Japanese voice-over | 1.28 GiB |  |
| [ko-kr_2.8.0_3.1.0_hdiff_FpEC0r5WxkMgD1cN.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.8.0_3.1.0_hdiff_FpEC0r5WxkMgD1cN.zip) | Korean voice-over | 1.11 GiB |  |

### 3.0.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_3.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220815143807_dyIghvy1b5fjfzHU/GenshinImpact_3.0.0.zip) | Game | 34.46 GiB |  |
| [Audio_Chinese_3.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220815143807_dyIghvy1b5fjfzHU/Audio_Chinese_3.0.0.zip) | Chinese voice-over | 7.56 GiB |  |
| [Audio_English(US)_3.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220815143807_dyIghvy1b5fjfzHU/Audio_English%28US%29_3.0.0.zip) | English voice-over | 8.72 GiB |  |
| [Audio_Japanese_3.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220815143807_dyIghvy1b5fjfzHU/Audio_Japanese_3.0.0.zip) | Japanese voice-over | 9.61 GiB |  |
| [Audio_Korean_3.0.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220815143807_dyIghvy1b5fjfzHU/Audio_Korean_3.0.0.zip) | Korean voice-over | 7.43 GiB |  |

#### Update from 2.8.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.8.0_3.0.0_hdiff_5eW7vzqQBEwmrVCi.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.8.0_3.0.0_hdiff_5eW7vzqQBEwmrVCi.zip) | Game | 8.35 GiB |  |
| [zh-cn_2.8.0_3.0.0_hdiff_lUMQXapY5jnDKokJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.8.0_3.0.0_hdiff_lUMQXapY5jnDKokJ.zip) | Chinese voice-over | 575.84 MiB |  |
| [en-us_2.8.0_3.0.0_hdiff_6yWlQX5NphSDHovf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.8.0_3.0.0_hdiff_6yWlQX5NphSDHovf.zip) | English voice-over | 691.60 MiB |  |
| [ja-jp_2.8.0_3.0.0_hdiff_FiYR9adhksO62VcZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.8.0_3.0.0_hdiff_FiYR9adhksO62VcZ.zip) | Japanese voice-over | 659.59 MiB |  |
| [ko-kr_2.8.0_3.0.0_hdiff_rckI0qyC8V3xBaep.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.8.0_3.0.0_hdiff_rckI0qyC8V3xBaep.zip) | Korean voice-over | 577.34 MiB |  |

#### Update from 2.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.7.0_3.0.0_hdiff_b1u9GtgJhTUPryIL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.7.0_3.0.0_hdiff_b1u9GtgJhTUPryIL.zip) | Game | 9.94 GiB |  |
| [zh-cn_2.7.0_3.0.0_hdiff_0XJTgBrkQN1lsFUW.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.7.0_3.0.0_hdiff_0XJTgBrkQN1lsFUW.zip) | Chinese voice-over | 1.31 GiB |  |
| [en-us_2.7.0_3.0.0_hdiff_DIWRaXKtUSpPMQ0z.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.7.0_3.0.0_hdiff_DIWRaXKtUSpPMQ0z.zip) | English voice-over | 1.59 GiB |  |
| [ja-jp_2.7.0_3.0.0_hdiff_6jJPH9SyWtqAC01n.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.7.0_3.0.0_hdiff_6jJPH9SyWtqAC01n.zip) | Japanese voice-over | 1.47 GiB |  |
| [ko-kr_2.7.0_3.0.0_hdiff_kcoyYRQpLqT61zth.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.7.0_3.0.0_hdiff_kcoyYRQpLqT61zth.zip) | Korean voice-over | 1.31 GiB |  |

</details>

<details>
<summary><b>Version 2.x</b>: 2.0.0 to 2.8.0 (9 versions)</summary>

Jump to: [2.8.0](#280) · [2.7.0](#270) · [2.6.0](#260) · [2.5.0](#250) · [2.4.0](#240) · [2.3.0](#230) · [2.2.0](#220) · [2.1.0](#210) · [2.0.0](#200)

### 2.8.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220625012443_r2qBGYstD0u3Q9xK/GenshinImpact_2.8.0.zip) | Game | 29.62 GiB |  |
| [Audio_Chinese_2.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220625012443_r2qBGYstD0u3Q9xK/Audio_Chinese_2.8.0.zip) | Chinese voice-over | 6.93 GiB |  |
| [Audio_English(US)_2.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220625012443_r2qBGYstD0u3Q9xK/Audio_English%28US%29_2.8.0.zip) | English voice-over | 7.99 GiB |  |
| [Audio_Japanese_2.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220625012443_r2qBGYstD0u3Q9xK/Audio_Japanese_2.8.0.zip) | Japanese voice-over | 8.72 GiB |  |
| [Audio_Korean_2.8.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220625012443_r2qBGYstD0u3Q9xK/Audio_Korean_2.8.0.zip) | Korean voice-over | 6.78 GiB |  |

#### Update from 2.7.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.7.0_2.8.0_hdiff_JvqEs0RiP9OrAbY7.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.7.0_2.8.0_hdiff_JvqEs0RiP9OrAbY7.zip) | Game | 4.31 GiB |  |
| [zh-cn_2.7.0_2.8.0_hdiff_r7VPXQ36dRabwK2s.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.7.0_2.8.0_hdiff_r7VPXQ36dRabwK2s.zip) | Chinese voice-over | 769.79 MiB |  |
| [en-us_2.7.0_2.8.0_hdiff_UXsOvlyKdAijwHhI.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.7.0_2.8.0_hdiff_UXsOvlyKdAijwHhI.zip) | English voice-over | 937.86 MiB |  |
| [ja-jp_2.7.0_2.8.0_hdiff_ozU04InEe5NDf9RJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.7.0_2.8.0_hdiff_ozU04InEe5NDf9RJ.zip) | Japanese voice-over | 843.45 MiB |  |
| [ko-kr_2.7.0_2.8.0_hdiff_YB9GdcHPnj4KisZo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.7.0_2.8.0_hdiff_YB9GdcHPnj4KisZo.zip) | Korean voice-over | 762.81 MiB |  |

#### Update from 2.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.6.0_2.8.0_hdiff_CmcT2Qur7R3tM0hU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.6.0_2.8.0_hdiff_CmcT2Qur7R3tM0hU.zip) | Game | 6.52 GiB |  |
| [zh-cn_2.6.0_2.8.0_hdiff_s8EPzuG5wiTS4q3o.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.6.0_2.8.0_hdiff_s8EPzuG5wiTS4q3o.zip) | Chinese voice-over | 5.59 GiB |  |
| [en-us_2.6.0_2.8.0_hdiff_PLEAWdHpRrBMkYIc.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.6.0_2.8.0_hdiff_PLEAWdHpRrBMkYIc.zip) | English voice-over | 6.80 GiB |  |
| [ja-jp_2.6.0_2.8.0_hdiff_pWurcwZs9FYNgxE8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.6.0_2.8.0_hdiff_pWurcwZs9FYNgxE8.zip) | Japanese voice-over | 6.56 GiB |  |
| [ko-kr_2.6.0_2.8.0_hdiff_nYO46PmHztpgDBJZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.6.0_2.8.0_hdiff_nYO46PmHztpgDBJZ.zip) | Korean voice-over | 5.51 GiB |  |

### 2.7.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220429112415_dDweiEHDnBI6cKmM/GenshinImpact_2.7.0.zip) | Game | 28.47 GiB |  |
| [Audio_Chinese_2.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220429112415_dDweiEHDnBI6cKmM/Audio_Chinese_2.7.0.zip) | Chinese voice-over | 6.29 GiB |  |
| [Audio_English(US)_2.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220429112415_dDweiEHDnBI6cKmM/Audio_English%28US%29_2.7.0.zip) | English voice-over | 7.23 GiB |  |
| [Audio_Japanese_2.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220429112415_dDweiEHDnBI6cKmM/Audio_Japanese_2.7.0.zip) | Japanese voice-over | 7.91 GiB |  |
| [Audio_Korean_2.7.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220429112415_dDweiEHDnBI6cKmM/Audio_Korean_2.7.0.zip) | Korean voice-over | 6.14 GiB |  |

#### Update from 2.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.6.0_2.7.0_hdiff_zfS2oacTJG7DlBER.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.6.0_2.7.0_hdiff_zfS2oacTJG7DlBER.zip) | Game | 4.62 GiB |  |
| [zh-cn_2.6.0_2.7.0_hdiff_QWVgNjwzprHPRKxL.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.6.0_2.7.0_hdiff_QWVgNjwzprHPRKxL.zip) | Chinese voice-over | 5.07 GiB |  |
| [en-us_2.6.0_2.7.0_hdiff_QscinXF9ASaB2POU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.6.0_2.7.0_hdiff_QscinXF9ASaB2POU.zip) | English voice-over | 6.15 GiB |  |
| [ja-jp_2.6.0_2.7.0_hdiff_5yDzjBLv08wFRsmx.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.6.0_2.7.0_hdiff_5yDzjBLv08wFRsmx.zip) | Japanese voice-over | 5.99 GiB |  |
| [ko-kr_2.6.0_2.7.0_hdiff_GnYMmTClWghVKSQk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.6.0_2.7.0_hdiff_GnYMmTClWghVKSQk.zip) | Korean voice-over | 4.99 GiB |  |

#### Update from 2.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.5.0_2.7.0_hdiff_HteQ6VmOcgy1PYxw.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.5.0_2.7.0_hdiff_HteQ6VmOcgy1PYxw.zip) | Game | 7.27 GiB |  |
| [zh-cn_2.5.0_2.7.0_hdiff_26vALtSnHzYTIopE.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.5.0_2.7.0_hdiff_26vALtSnHzYTIopE.zip) | Chinese voice-over | 5.07 GiB |  |
| [en-us_2.5.0_2.7.0_hdiff_tTJgcA4QIPOHBpkS.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.5.0_2.7.0_hdiff_tTJgcA4QIPOHBpkS.zip) | English voice-over | 6.15 GiB |  |
| [ja-jp_2.5.0_2.7.0_hdiff_CArIRcsvPhTtWqow.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.5.0_2.7.0_hdiff_CArIRcsvPhTtWqow.zip) | Japanese voice-over | 5.99 GiB |  |
| [ko-kr_2.5.0_2.7.0_hdiff_wSYkWPafc6OUxILt.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.5.0_2.7.0_hdiff_wSYkWPafc6OUxILt.zip) | Korean voice-over | 4.99 GiB |  |

### 2.6.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220318211020_aWEQCaw5ZERt54rm/GenshinImpact_2.6.0.zip) | Game | 26.77 GiB |  |
| [Audio_Chinese_2.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220318211020_aWEQCaw5ZERt54rm/Audio_Chinese_2.6.0.zip) | Chinese voice-over | 6.07 GiB |  |
| [Audio_English(US)_2.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220318211020_aWEQCaw5ZERt54rm/Audio_English%28US%29_2.6.0.zip) | English voice-over | 6.87 GiB |  |
| [Audio_Japanese_2.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220318211020_aWEQCaw5ZERt54rm/Audio_Japanese_2.6.0.zip) | Japanese voice-over | 7.47 GiB |  |
| [Audio_Korean_2.6.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220318211020_aWEQCaw5ZERt54rm/Audio_Korean_2.6.0.zip) | Korean voice-over | 5.84 GiB |  |

#### Update from 2.5.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.5.0_2.6.0_hdiff_qr9GvX8ALUe6Puk3.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.5.0_2.6.0_hdiff_qr9GvX8ALUe6Puk3.zip) | Game | 4.58 GiB |  |
| [zh-cn_2.5.0_2.6.0_hdiff_TMDQ3PkIdt4ABcXZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.5.0_2.6.0_hdiff_TMDQ3PkIdt4ABcXZ.zip) | Chinese voice-over | 360.50 MiB |  |
| [en-us_2.5.0_2.6.0_hdiff_cTrqX6uxY1IH7zVs.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.5.0_2.6.0_hdiff_cTrqX6uxY1IH7zVs.zip) | English voice-over | 478.30 MiB |  |
| [ja-jp_2.5.0_2.6.0_hdiff_gRwTfd4VsJXLv9uY.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.5.0_2.6.0_hdiff_gRwTfd4VsJXLv9uY.zip) | Japanese voice-over | 452.34 MiB |  |
| [ko-kr_2.5.0_2.6.0_hdiff_20EY6qkb8IetwBK9.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.5.0_2.6.0_hdiff_20EY6qkb8IetwBK9.zip) | Korean voice-over | 375.99 MiB |  |

#### Update from 2.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.4.0_2.6.0_hdiff_IzpvafMJqow6lB2U.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.4.0_2.6.0_hdiff_IzpvafMJqow6lB2U.zip) | Game | 9.62 GiB |  |
| [zh-cn_2.4.0_2.6.0_hdiff_g8JQpXKHxcwSYTGn.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.4.0_2.6.0_hdiff_g8JQpXKHxcwSYTGn.zip) | Chinese voice-over | 641.83 MiB |  |
| [en-us_2.4.0_2.6.0_hdiff_1qv4AlX5YUhtgRbf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.4.0_2.6.0_hdiff_1qv4AlX5YUhtgRbf.zip) | English voice-over | 817.35 MiB |  |
| [ja-jp_2.4.0_2.6.0_hdiff_P5fa2yuwXFYKEnHN.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.4.0_2.6.0_hdiff_P5fa2yuwXFYKEnHN.zip) | Japanese voice-over | 798.25 MiB |  |
| [ko-kr_2.4.0_2.6.0_hdiff_sYVhI9LJd3yTe5EU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.4.0_2.6.0_hdiff_sYVhI9LJd3yTe5EU.zip) | Korean voice-over | 639.09 MiB |  |

### 2.5.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220125104720_x2gRaOdngikczohR/GenshinImpact_2.5.0.zip) | Game | 24.55 GiB |  |
| [Audio_Chinese_2.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220125104720_x2gRaOdngikczohR/Audio_Chinese_2.5.0.zip) | Chinese voice-over | 5.65 GiB |  |
| [Audio_English(US)_2.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220125104720_x2gRaOdngikczohR/Audio_English%28US%29_2.5.0.zip) | English voice-over | 6.37 GiB |  |
| [Audio_Japanese_2.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220125104720_x2gRaOdngikczohR/Audio_Japanese_2.5.0.zip) | Japanese voice-over | 6.89 GiB |  |
| [Audio_Korean_2.5.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20220125104720_x2gRaOdngikczohR/Audio_Korean_2.5.0.zip) | Korean voice-over | 5.42 GiB |  |

#### Update from 2.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.4.0_2.5.0_hdiff_G7AskHSpFPiXwRyU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.4.0_2.5.0_hdiff_G7AskHSpFPiXwRyU.zip) | Game | 6.92 GiB |  |
| [zh-cn_2.4.0_2.5.0_hdiff_Tx1P3el0KJ68Vu4Z.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.4.0_2.5.0_hdiff_Tx1P3el0KJ68Vu4Z.zip) | Chinese voice-over | 281.34 MiB |  |
| [en-us_2.4.0_2.5.0_hdiff_FGN5bsVa0BZLKqEA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.4.0_2.5.0_hdiff_FGN5bsVa0BZLKqEA.zip) | English voice-over | 339.06 MiB |  |
| [ja-jp_2.4.0_2.5.0_hdiff_t0MU7XpPc3ofbYwD.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.4.0_2.5.0_hdiff_t0MU7XpPc3ofbYwD.zip) | Japanese voice-over | 345.92 MiB |  |
| [ko-kr_2.4.0_2.5.0_hdiff_0BLZeojnu59RDlx4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.4.0_2.5.0_hdiff_0BLZeojnu59RDlx4.zip) | Korean voice-over | 263.11 MiB |  |

#### Update from 2.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.3.0_2.5.0_hdiff_DFWAwSmaTj7n90Ru.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.3.0_2.5.0_hdiff_DFWAwSmaTj7n90Ru.zip) | Game | 8.35 GiB |  |
| [zh-cn_2.3.0_2.5.0_hdiff_6Y1XIaq27jKkMGsz.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.3.0_2.5.0_hdiff_6Y1XIaq27jKkMGsz.zip) | Chinese voice-over | 670.83 MiB |  |
| [en-us_2.3.0_2.5.0_hdiff_5Mx1uveJjTQz9w3r.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.3.0_2.5.0_hdiff_5Mx1uveJjTQz9w3r.zip) | English voice-over | 879.87 MiB |  |
| [ja-jp_2.3.0_2.5.0_hdiff_BA6UcSsKf2VRpLFH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.3.0_2.5.0_hdiff_BA6UcSsKf2VRpLFH.zip) | Japanese voice-over | 865.23 MiB |  |
| [ko-kr_2.3.0_2.5.0_hdiff_7iPhaQjMp3fsRGdX.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.3.0_2.5.0_hdiff_7iPhaQjMp3fsRGdX.zip) | Korean voice-over | 654.21 MiB |  |

### 2.4.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211225051318_JHACtHpvJ2yRaZH0/GenshinImpact_2.4.0.zip) | Game | 25.80 GiB |  |
| [Audio_Chinese_2.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211225051318_JHACtHpvJ2yRaZH0/Audio_Chinese_2.4.0.zip) | Chinese voice-over | 5.33 GiB |  |
| [Audio_English(US)_2.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211225051318_JHACtHpvJ2yRaZH0/Audio_English%28US%29_2.4.0.zip) | English voice-over | 6.01 GiB |  |
| [Audio_Japanese_2.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211225051318_JHACtHpvJ2yRaZH0/Audio_Japanese_2.4.0.zip) | Japanese voice-over | 6.48 GiB |  |
| [Audio_Korean_2.4.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211225051318_JHACtHpvJ2yRaZH0/Audio_Korean_2.4.0.zip) | Korean voice-over | 5.13 GiB |  |

#### Update from 2.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.3.0_2.4.0_diff_jm6dxRrsaWQ0wT2X.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.3.0_2.4.0_diff_jm6dxRrsaWQ0wT2X.zip) | Game | 6.15 GiB |  |
| [zh-cn_2.3.0_2.4.0_diff_CMYcoWsF1EuhiLpx.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.3.0_2.4.0_diff_CMYcoWsF1EuhiLpx.zip) | Chinese voice-over | 389.50 MiB |  |
| [en-us_2.3.0_2.4.0_diff_nB9mKiPvfW8MhRZd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.3.0_2.4.0_diff_nB9mKiPvfW8MhRZd.zip) | English voice-over | 540.82 MiB |  |
| [ja-jp_2.3.0_2.4.0_diff_VnP5l3toBOKUapvi.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.3.0_2.4.0_diff_VnP5l3toBOKUapvi.zip) | Japanese voice-over | 519.32 MiB |  |
| [ko-kr_2.3.0_2.4.0_diff_nykbuZGrtEsg5eTz.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.3.0_2.4.0_diff_nykbuZGrtEsg5eTz.zip) | Korean voice-over | 391.30 MiB |  |

#### Update from 2.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.2.0_2.4.0_diff_9kTlUMvuWDeNOKap.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.2.0_2.4.0_diff_9kTlUMvuWDeNOKap.zip) | Game | 8.55 GiB |  |
| [zh-cn_2.2.0_2.4.0_diff_Ui7XBC3FITMPL0hl.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.2.0_2.4.0_diff_Ui7XBC3FITMPL0hl.zip) | Chinese voice-over | 753.73 MiB |  |
| [en-us_2.2.0_2.4.0_diff_XD7bFBn31MpLcVws.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.2.0_2.4.0_diff_XD7bFBn31MpLcVws.zip) | English voice-over | 1008.63 MiB |  |
| [ja-jp_2.2.0_2.4.0_diff_euy5Exgt3bZdmfUr.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.2.0_2.4.0_diff_euy5Exgt3bZdmfUr.zip) | Japanese voice-over | 1000.60 MiB |  |
| [ko-kr_2.2.0_2.4.0_diff_Eq0ekGyVaBIWFdM5.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.2.0_2.4.0_diff_Eq0ekGyVaBIWFdM5.zip) | Korean voice-over | 767.78 MiB |  |

### 2.3.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211117173404_G0gLRnxvOd4PvSu9/GenshinImpact_2.3.0.zip) | Game | 23.08 GiB |  |
| [Audio_Chinese_2.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211117173404_G0gLRnxvOd4PvSu9/Audio_Chinese_2.3.0.zip) | Chinese voice-over | 4.91 GiB |  |
| [Audio_English(US)_2.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211117173404_G0gLRnxvOd4PvSu9/Audio_English%28US%29_2.3.0.zip) | English voice-over | 5.44 GiB |  |
| [Audio_Japanese_2.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211117173404_G0gLRnxvOd4PvSu9/Audio_Japanese_2.3.0.zip) | Japanese voice-over | 5.88 GiB |  |
| [Audio_Korean_2.3.0.zip](https://autopatchhk.yuanshen.com/client_app/download/pc_zip/20211117173404_G0gLRnxvOd4PvSu9/Audio_Korean_2.3.0.zip) | Korean voice-over | 4.70 GiB |  |

#### Update from 2.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.2.0_2.3.0_diff_aLjNHzRTOQlyb2x8.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.2.0_2.3.0_diff_aLjNHzRTOQlyb2x8.zip) | Game | 4.51 GiB |  |
| [zh-cn_2.2.0_2.3.0_diff_24AhO9aFfJEXHqp5.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.2.0_2.3.0_diff_24AhO9aFfJEXHqp5.zip) | Chinese voice-over | 364.23 MiB |  |
| [en-us_2.2.0_2.3.0_diff_9vQCTEkUrs7qoNxg.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.2.0_2.3.0_diff_9vQCTEkUrs7qoNxg.zip) | English voice-over | 467.81 MiB |  |
| [ja-jp_2.2.0_2.3.0_diff_iVDQPpkNu4etM5Xo.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.2.0_2.3.0_diff_iVDQPpkNu4etM5Xo.zip) | Japanese voice-over | 481.30 MiB |  |
| [ko-kr_2.2.0_2.3.0_diff_r6zX1n48Hh5LUBK7.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.2.0_2.3.0_diff_r6zX1n48Hh5LUBK7.zip) | Korean voice-over | 376.49 MiB |  |

#### Update from 2.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.1.0_2.3.0_diff_IFqrKCzDYyUt697g.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.1.0_2.3.0_diff_IFqrKCzDYyUt697g.zip) | Game | 6.64 GiB |  |
| [zh-cn_2.1.0_2.3.0_diff_j4VLGKErh0w7WeFC.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.1.0_2.3.0_diff_j4VLGKErh0w7WeFC.zip) | Chinese voice-over | 620.90 MiB |  |
| [en-us_2.1.0_2.3.0_diff_9WOeMLlpzcFh7owj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.1.0_2.3.0_diff_9WOeMLlpzcFh7owj.zip) | English voice-over | 790.82 MiB |  |
| [ja-jp_2.1.0_2.3.0_diff_8nXfsMArgbPOU30c.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.1.0_2.3.0_diff_8nXfsMArgbPOU30c.zip) | Japanese voice-over | 812.59 MiB |  |
| [ko-kr_2.1.0_2.3.0_diff_MEpgFPsUZABmGt9b.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.1.0_2.3.0_diff_MEpgFPsUZABmGt9b.zip) | Korean voice-over | 628.15 MiB |  |

### 2.2.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.2.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20211013_407d783aa6b98191/GenshinImpact_2.2.0.zip) (dead link, HTTP 404) | Game | unknown |  |
| [Audio_Chinese_2.2.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20211013_407d783aa6b98191/Audio_Chinese_2.2.0.zip) | Chinese voice-over | 4.51 GiB |  |
| [Audio_English(US)_2.2.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20211013_407d783aa6b98191/Audio_English%28US%29_2.2.0.zip) | English voice-over | 4.94 GiB |  |
| [Audio_Japanese_2.2.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20211013_407d783aa6b98191/Audio_Japanese_2.2.0.zip) | Japanese voice-over | 5.33 GiB |  |
| [Audio_Korean_2.2.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20211013_407d783aa6b98191/Audio_Korean_2.2.0.zip) | Korean voice-over | 4.28 GiB |  |

#### Update from 2.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.1.0_2.2.0_diff_4v9prHYT3QOD628M.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.1.0_2.2.0_diff_4v9prHYT3QOD628M.zip) | Game | 5.16 GiB |  |
| [zh-cn_2.1.0_2.2.0_diff_OQanfLW1ClY5hswv.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.1.0_2.2.0_diff_OQanfLW1ClY5hswv.zip) | Chinese voice-over | 256.67 MiB |  |
| [en-us_2.1.0_2.2.0_diff_e8pIgj039NSGh7i4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.1.0_2.2.0_diff_e8pIgj039NSGh7i4.zip) | English voice-over | 323.01 MiB |  |
| [ja-jp_2.1.0_2.2.0_diff_UC6QMjzTyguFKomW.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.1.0_2.2.0_diff_UC6QMjzTyguFKomW.zip) | Japanese voice-over | 331.30 MiB |  |
| [ko-kr_2.1.0_2.2.0_diff_rynlNTw32CAubHjk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.1.0_2.2.0_diff_rynlNTw32CAubHjk.zip) | Korean voice-over | 251.66 MiB |  |

#### Update from 2.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.0.0_2.2.0_diff_4fBjlXcqzgVsQ5ZG.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.0.0_2.2.0_diff_4fBjlXcqzgVsQ5ZG.zip) | Game | 9.38 GiB |  |
| [zh-cn_2.0.0_2.2.0_diff_E7gjdsbeBx3InRhH.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.0.0_2.2.0_diff_E7gjdsbeBx3InRhH.zip) | Chinese voice-over | 804.16 MiB |  |
| [en-us_2.0.0_2.2.0_diff_7lxrW0qg8M9mRCwj.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.0.0_2.2.0_diff_7lxrW0qg8M9mRCwj.zip) | English voice-over | 1.25 GiB |  |
| [ja-jp_2.0.0_2.2.0_diff_e4x8DkLvUHPuiZEJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.0.0_2.2.0_diff_e4x8DkLvUHPuiZEJ.zip) | Japanese voice-over | 869.03 MiB |  |
| [ko-kr_2.0.0_2.2.0_diff_hrUCmFj4qz83XSwJ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.0.0_2.2.0_diff_hrUCmFj4qz83XSwJ.zip) | Korean voice-over | 980.25 MiB |  |

### 2.1.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_2.1.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210901_db535e31166295ac/GenshinImpact_2.1.0.zip) | Game | 20.22 GiB |  |
| [Audio_Chinese_2.1.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210901_db535e31166295ac/Audio_Chinese_2.1.0.zip) | Chinese voice-over | 4.22 GiB |  |
| [Audio_English(US)_2.1.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210901_db535e31166295ac/Audio_English%28US%29_2.1.0.zip) | English voice-over | 4.60 GiB |  |
| [Audio_Japanese_2.1.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210901_db535e31166295ac/Audio_Japanese_2.1.0.zip) | Japanese voice-over | 4.94 GiB |  |
| [Audio_Korean_2.1.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210901_db535e31166295ac/Audio_Korean_2.1.0.zip) | Korean voice-over | 4.00 GiB |  |

#### Update from 2.0.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_2.0.0_2.1.0_diff_kYU5QpgOiPB1woI4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_2.0.0_2.1.0_diff_kYU5QpgOiPB1woI4.zip) | Game | 6.62 GiB |  |
| [zh-cn_2.0.0_2.1.0_diff_UCr51Zxdu9aAjGYg.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_2.0.0_2.1.0_diff_UCr51Zxdu9aAjGYg.zip) | Chinese voice-over | 547.50 MiB |  |
| [en-us_2.0.0_2.1.0_diff_n8EeG923ZU1BPcgq.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_2.0.0_2.1.0_diff_n8EeG923ZU1BPcgq.zip) | English voice-over | 957.89 MiB |  |
| [ja-jp_2.0.0_2.1.0_diff_Ub6DpehQk19xKEVa.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_2.0.0_2.1.0_diff_Ub6DpehQk19xKEVa.zip) | Japanese voice-over | 537.73 MiB |  |
| [ko-kr_2.0.0_2.1.0_diff_SYrI5qspl1z4tBRn.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_2.0.0_2.1.0_diff_SYrI5qspl1z4tBRn.zip) | Korean voice-over | 728.59 MiB |  |

#### Update from 1.6.1

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_1.6.1_2.1.0_diff_lEOthRaCuMIB4rTN.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_1.6.1_2.1.0_diff_lEOthRaCuMIB4rTN.zip) | Game | 10.33 GiB |  |
| [zh-cn_1.6.1_2.1.0_diff_vh3oYQe1kctqNIbO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_1.6.1_2.1.0_diff_vh3oYQe1kctqNIbO.zip) | Chinese voice-over | 900.58 MiB |  |
| [en-us_1.6.1_2.1.0_diff_VLagihjWn3rFAm8X.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_1.6.1_2.1.0_diff_VLagihjWn3rFAm8X.zip) | English voice-over | 1.08 GiB |  |
| [ja-jp_1.6.1_2.1.0_diff_uFVREQ8SgUcOlNoP.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_1.6.1_2.1.0_diff_uFVREQ8SgUcOlNoP.zip) | Japanese voice-over | 1002.89 MiB |  |
| [ko-kr_1.6.1_2.1.0_diff_XyfOJVYxLSlR21zA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_1.6.1_2.1.0_diff_XyfOJVYxLSlR21zA.zip) | Korean voice-over | 876.36 MiB |  |

### 2.0.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210707145437_906c4d25PUR1BSt7.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210707145437_906c4d25PUR1BSt7.zip) | Launcher | 293.95 MiB |  |
| [GenshinImpact_2.0.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210721_3a3ca8dfe8b26ea2/GenshinImpact_2.0.0.zip) | Game | 17.44 GiB |  |
| [Audio_Chinese_2.0.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210721_3a3ca8dfe8b26ea2/Audio_Chinese_2.0.0.zip) | Chinese voice-over | 3.71 GiB |  |
| [Audio_English(US)_2.0.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210721_3a3ca8dfe8b26ea2/Audio_English%28US%29_2.0.0.zip) | English voice-over | 4.01 GiB |  |
| [Audio_Japanese_2.0.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210721_3a3ca8dfe8b26ea2/Audio_Japanese_2.0.0.zip) | Japanese voice-over | 4.29 GiB |  |
| [Audio_Korean_2.0.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210721_3a3ca8dfe8b26ea2/Audio_Korean_2.0.0.zip) | Korean voice-over | 3.50 GiB |  |

#### Update from 1.6.1

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_1.6.1_2.0.0_diff_GLrkVi02eEIJHmoZ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_1.6.1_2.0.0_diff_GLrkVi02eEIJHmoZ.zip) | Game | 5.98 GiB |  |
| [zh-cn_1.6.1_2.0.0_diff_kxn7JGptOMsUWeXf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_1.6.1_2.0.0_diff_kxn7JGptOMsUWeXf.zip) | Chinese voice-over | 436.61 MiB |  |
| [en-us_1.6.1_2.0.0_diff_pa7XSfuvwheYHlR4.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_1.6.1_2.0.0_diff_pa7XSfuvwheYHlR4.zip) | English voice-over | 547.86 MiB |  |
| [ja-jp_1.6.1_2.0.0_diff_Rr1nlK3O0Wjgo4M2.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_1.6.1_2.0.0_diff_Rr1nlK3O0Wjgo4M2.zip) | Japanese voice-over | 489.74 MiB |  |
| [ko-kr_1.6.1_2.0.0_diff_YHsxbir0lqzcvNIk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_1.6.1_2.0.0_diff_YHsxbir0lqzcvNIk.zip) | Korean voice-over | 425.13 MiB |  |

#### Update from 1.6.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_1.6.0_2.0.0_diff_NmOQGgYFcB60KLzk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_1.6.0_2.0.0_diff_NmOQGgYFcB60KLzk.zip) | Game | 5.98 GiB |  |
| [zh-cn_1.6.0_2.0.0_diff_w4oOYzASepP9Rlc0.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_1.6.0_2.0.0_diff_w4oOYzASepP9Rlc0.zip) | Chinese voice-over | 436.61 MiB |  |
| [en-us_1.6.0_2.0.0_diff_F9KsGgq5duEzpVi2.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_1.6.0_2.0.0_diff_F9KsGgq5duEzpVi2.zip) | English voice-over | 547.86 MiB |  |
| [ja-jp_1.6.0_2.0.0_diff_bGKRLxOIFYU740nT.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_1.6.0_2.0.0_diff_bGKRLxOIFYU740nT.zip) | Japanese voice-over | 489.74 MiB |  |
| [ko-kr_1.6.0_2.0.0_diff_VoMeKwuyAZ1jRItk.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_1.6.0_2.0.0_diff_VoMeKwuyAZ1jRItk.zip) | Korean voice-over | 425.13 MiB |  |

</details>

<details>
<summary><b>Version 1.x</b>: 1.3.0 to 1.6.1 (6 versions)</summary>

Jump to: [1.6.1](#161) · [1.6.0](#160) · [1.5.0](#150) · [1.4.0](#140) · [1.3.2](#132) · [1.3.0](#130)

### 1.6.1

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210525213303_d009a819aTBew3U6.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210525213303_d009a819aTBew3U6.zip) | Launcher | 293.88 MiB |  |
| [GenshinImpact_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/GenshinImpact_1.6.0.zip) | Game | 13.93 GiB |  |
| [Audio_Chinese_1.6.1.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_eea40505c6337fd1/Audio_Chinese_1.6.1.zip) | Chinese voice-over | 3.23 GiB |  |
| [Audio_English(US)_1.6.1.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_eea40505c6337fd1/Audio_English%28US%29_1.6.1.zip) | English voice-over | 3.45 GiB |  |
| [Audio_Japanese_1.6.1.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_eea40505c6337fd1/Audio_Japanese_1.6.1.zip) | Japanese voice-over | 3.67 GiB |  |
| [Audio_Korean_1.6.1.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_eea40505c6337fd1/Audio_Korean_1.6.1.zip) | Korean voice-over | 3.02 GiB |  |

The game archive `GenshinImpact_1.6.0.zip` is named for an earlier release: upstream lists it as the 1.6.1 game data.

### 1.6.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210525213303_d009a819aTBew3U6.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210525213303_d009a819aTBew3U6.zip) | Launcher | 293.88 MiB |  |
| [GenshinImpact_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/GenshinImpact_1.6.0.zip) | Game | 13.93 GiB |  |
| [Audio_Chinese_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/Audio_Chinese_1.6.0.zip) | Chinese voice-over | 3.23 GiB |  |
| [Audio_English(US)_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/Audio_English%28US%29_1.6.0.zip) | English voice-over | 3.45 GiB |  |
| [Audio_Japanese_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/Audio_Japanese_1.6.0.zip) | Japanese voice-over | 3.67 GiB |  |
| [Audio_Korean_1.6.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210609_15f555799e5d6233/Audio_Korean_1.6.0.zip) | Korean voice-over | 3.02 GiB |  |

#### Update from 1.5.1

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_1.5.1_1.6.0_diff_jVasZSurGCqb1PYx.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_1.5.1_1.6.0_diff_jVasZSurGCqb1PYx.zip) | Game | 3.28 GiB |  |
| [zh-cn_1.5.1_1.6.0_diff_jXGDLg1FMVdYZrRQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_1.5.1_1.6.0_diff_jXGDLg1FMVdYZrRQ.zip) | Chinese voice-over | 307.76 MiB |  |
| [en-us_1.5.1_1.6.0_diff_LvqpG1Srcm72t6YF.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_1.5.1_1.6.0_diff_LvqpG1Srcm72t6YF.zip) | English voice-over | 369.50 MiB |  |
| [ja-jp_1.5.1_1.6.0_diff_mdLSj8pM1ANnHZTB.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_1.5.1_1.6.0_diff_mdLSj8pM1ANnHZTB.zip) | Japanese voice-over | 339.28 MiB |  |
| [ko-kr_1.5.1_1.6.0_diff_HJxsoPmCvqB2liEA.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_1.5.1_1.6.0_diff_HJxsoPmCvqB2liEA.zip) | Korean voice-over | 307.45 MiB |  |

### 1.5.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210413104427.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210413104427.zip) | Launcher | 293.71 MiB |  |
| [GenshinImpact_1.5.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210428_de80a243f7474c39/GenshinImpact_1.5.0.zip) | Game | 12.13 GiB |  |

#### Update from 1.4.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [game_1.4.0_1.5.0_diff_H8LGSgKU.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/game_1.4.0_1.5.0_diff_H8LGSgKU.zip) | Game | 5.42 GiB |  |
| [zh-cn_1.4.0_1.5.0_diff_hvg2F3QM.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/zh-cn_1.4.0_1.5.0_diff_hvg2F3QM.zip) | Chinese voice-over | 382.45 MiB |  |
| [en-us_1.4.0_1.5.0_diff_KdXT5Pi0.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/en-us_1.4.0_1.5.0_diff_KdXT5Pi0.zip) | English voice-over | 461.88 MiB |  |
| [ja-jp_1.4.0_1.5.0_diff_hfBOW6ym.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ja-jp_1.4.0_1.5.0_diff_hfBOW6ym.zip) | Japanese voice-over | 389.97 MiB |  |
| [ko-kr_1.4.0_1.5.0_diff_NcX8UP1L.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/ko-kr_1.4.0_1.5.0_diff_NcX8UP1L.zip) | Korean voice-over | 381.14 MiB |  |

### 1.4.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210304190611.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210304190611.zip) | Launcher | 228.69 MiB |  |
| [GenshinImpact_1.4.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210317_67c8f1002bb26672/GenshinImpact_1.4.0.zip) | Game | 20.68 GiB |  |

#### Update from 1.3.2

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.3.2_1.4.0_diff_1P539V4J.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.3.2_1.4.0_diff_1P539V4J.zip) | Game | 4.30 GiB |  |

#### Update from 1.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.3.0_1.4.0_diff_hLOEnPIc.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.3.0_1.4.0_diff_hLOEnPIc.zip) | Game | 4.30 GiB |  |

#### Update from 1.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.2.0_1.4.0_diff_kO7vMycf.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.2.0_1.4.0_diff_kO7vMycf.zip) | Game | 7.16 GiB |  |

### 1.3.2

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [GenshinImpact_1.3.2.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210210_b0516858014fc457/GenshinImpact_1.3.2.zip) | Game | 18.12 GiB |  |

#### Update from 1.3.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.3.0_1.3.2_diff_ryqTKPYO.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.3.0_1.3.2_diff_ryqTKPYO.zip) | Game | 159.21 MiB |  |

#### Update from 1.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.2.0_1.3.2_diff_sHu9eOFd.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.2.0_1.3.2_diff_sHu9eOFd.zip) | Game | 4.19 GiB |  |

### 1.3.0

#### Full install

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [update_20210122154904.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/update_20210122154904.zip) | Launcher | 228.52 MiB |  |
| [GenshinImpact_1.3.0.zip](https://autopatchhk.yuanshen.com/client_app/pc_mihoyo/20210203_f97a39582b45649f/GenshinImpact_1.3.0.zip) | Game | 18.12 GiB |  |

#### Update from 1.2.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.2.0_1.3.0_diff_sBM8DJZc.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.2.0_1.3.0_diff_sBM8DJZc.zip) | Game | 4.19 GiB |  |

#### Update from 1.1.0

| File | Content | Size | MD5 |
| --- | --- | --- | --- |
| [1.1.0_1.3.0_diff_8cyweCdQ.zip](https://autopatchhk.yuanshen.com/client_app/update/hk4e_global/10/1.1.0_1.3.0_diff_8cyweCdQ.zip) | Game | 6.32 GiB |  |

</details>

## License

The code (the Go fetcher and generator) is licensed under the [Apache License 2.0](LICENSE).
The data in [data/](data/) and the content of this README are licensed under
[CC BY-NC-SA 4.0](LICENSE-DATA), the license of the upstream lists they build on.

---

Last checked: 2026-10-26 (UTC). This file is generated from
[data/versions.json](data/versions.json) by `go run ./cmd/fetch`; do not edit it by hand.
