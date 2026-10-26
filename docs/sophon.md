# Sophon: 5.6 and later

[← README](../README.md) · [Versions](../README.md#versions) · [Download guide (zip era)](download-guide.md) · [Data format](data.md)

HoYoverse stopped publishing full-game and patch zip archives after
**5.5.0**. Since 5.6, Genshin Impact (global, `hk4e_global`) ships only
through **Sophon**. In Sophon the game is split into small compressed
chunks, and a manifest lists them. Downloaders fetch the manifest and the
chunks, then rebuild the game files on your disk. There are no direct
download links to share for 5.6 and later.

The launchers on this page install the current version from the API. The
version pages under [versions/](versions/) list each version's Sophon
manifests, chunk URL prefix, sizes and diff tags. They are for reference and
for tools or scripts that rebuild files from a Sophon manifest. This
repository does not claim that older builds can be installed with these
launchers.

## Official route: HoYoPlay (Windows)

1. Download HoYoPlay from <https://hoyoplay.hoyoverse.com/> (also linked
   from <https://genshin.hoyoverse.com/>).
2. Install it, select **Genshin Impact**, choose an install folder and click
   **Download**. HoYoPlay fetches the current version through Sophon.
3. Updates and pre-downloads (pre-install) appear in HoYoPlay when
   HoYoverse publishes them.
4. To keep an existing install, use **Locate game** on its folder instead of
   downloading again. HoYoPlay checks the files and fetches only what is
   missing or changed.

## Collapse Launcher (Windows)

Open-source launcher for HoYoverse games. It installs, updates,
pre-downloads and repairs Genshin through Sophon, using its own
`Hi3Helper.Sophon` library.

1. Download `CollapseLauncher-stable-Setup.exe` or the portable `.7z` from
   the [releases page](https://github.com/CollapseLauncher/Collapse/releases).
   It needs Windows 10 1809 or later.
2. In the region list, pick **Genshin Impact → Global**.
3. Click **Install** for a new install, or point it at an existing game
   folder. Updates and pre-downloads then go through Sophon.

## Twintail Launcher (Windows and Linux)

Open-source launcher with Genshin Impact support. Its download backend
(`fischl-rs`) has a Sophon mode for full installs, updates (diffs),
preloads and repairs.

1. Download the Windows installer (`*_x64-setup.exe`) or the Linux `.deb`
   (amd64 or arm64) from the
   [releases page](https://github.com/TwintailTeam/TwintailLauncher/releases).
   Other install methods, such as a Nix flake, are listed in its README.
2. Add **Genshin Impact**, choose an install path and start the download.

## An Anime Game Launcher (Linux)

Linux launcher that runs the game through Wine. It installs and updates
through Sophon.

1. Install it with the method for your distribution. Flatpak
   (`moe.launcher.an-anime-game-launcher` on Flathub) is the officially
   supported one. See the
   [installation wiki](https://github.com/an-anime-team/an-anime-game-launcher/wiki/Installation).
2. On first launch, follow the setup steps (Wine and DXVK components, game
   folder). Then click **Download** to install the game, or **Update** to
   update it.

## Sophon Downloader (Windows, files only)

A stand-alone GUI that downloads game files from Sophon manifests. It is not
a launcher. It covers `hk4e_global`, including full builds and patch builds.

1. Download `SophonDownloader.7z` from the
   [releases page](https://github.com/GesthosNetwork/Sophon-Downloader/releases)
   and extract it. It runs on Windows x64 and needs a recent .NET runtime;
   see its README.
2. Select Genshin Impact (global), the version or branch and the content
   (game and voice packs), then download to a folder.

## What the Sophon tables mean

- **Manifest:** a compressed file that lists every game file and the chunks
  it is made of. Its URL is the manifest URL prefix plus the manifest ID.
- **Chunk URL prefix:** each chunk named in a manifest downloads from
  `<prefix>/<chunk name>`.
- **Download / Installed:** the compressed chunk total and the size on disk.
- **Sophon update ... (from X):** diff data that updates an install of
  version X (the diff tags). The Sophon API serves update data for the
  current version only, so older versions list the full build alone.
- Versions that still shipped zip archives when Sophon appeared list both
  their archives and their Sophon full build. Versions with no recorded
  build show `—` in the Sophon column of the index.

> The tools on this page are independent community projects. HoYoverse does
> not endorse them, and this repository has no ties to them. Third-party
> launchers may break when the game or API changes. Use them at your own
> risk.
