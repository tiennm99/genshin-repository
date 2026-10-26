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
