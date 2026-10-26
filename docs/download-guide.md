# Download guide: zip archives (5.5.0 and earlier)

[← README](../README.md) · [Versions](../README.md#versions) · [Sophon (5.6 and later)](sophon.md) · [Data format](data.md)

HoYoverse published Genshin Impact 5.5.0 and every earlier version as zip
archives. Each version page under [versions/](versions/) lists them in two
kinds of table:

- **Full install:** the game archive, split into numbered parts
  (`.zip.001`, `.zip.002`, ...) for newer versions, plus one archive per
  voice-over language.
- **Update from X:** smaller archives that update an existing install of
  version X. They hold hdiff patch files that a launcher or patch tool
  applies to the old files, not a full game. The upstream lists call them
  pre-install files.

From 5.6 on there are no zip archives; see [Sophon](sophon.md).

## Full install, step by step

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

Some old archives are no longer served. The version tables mark them as
dead links, and the version index notes how many a version has.

## Single files: ScatteredFiles

Versions that list a scattered-files URL (`res_list_url`) offer one more
route. Under that URL, `pkg_version` is a JSON-lines list of every game file
(`remoteName`, `md5`, `fileSize`). Each file can be downloaded on its own
from `<res_list_url>/<remoteName>`, which helps to repair a single file
without the full archive. HoYoverse may remove these URLs at any time.
