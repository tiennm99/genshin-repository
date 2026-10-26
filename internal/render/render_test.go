package render

import (
	"maps"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/tiennm99/genshin-repository/internal/store"
)

func sample() *store.Store {
	s := &store.Store{}
	old := s.Get("4.8.0")
	old.SeenMain("2024-07-18")
	old.SeenMain("2024-08-27")
	old.AddPackage("", &store.Package{Kind: store.KindGame, Part: 1, URL: "https://example.invalid/GenshinImpact_4.8.0.zip.001", MD5: "aaa", Size: 10 << 30})
	old.AddPackage("", &store.Package{Kind: store.KindGame, Part: 2, URL: "https://example.invalid/GenshinImpact_4.8.0.zip.002", Size: 5 << 30, HTTPStatus: 404})
	old.AddPackage("", &store.Package{Kind: store.KindAudio, Language: "en-us", URL: "https://example.invalid/Audio_English(US)_4.8.0.zip", Size: 3 << 20})
	old.AddPackage("4.7.0", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/game_4.7.0_4.8.0_hdiff_x.zip"})

	s.Get("5.0.0").SeenMain("2024-08-28")
	s.Find("5.0.0").AddPackage("", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/GenshinImpact_5.0.0.zip"})

	hotfix := s.Get("5.0.1")
	hotfix.AddPackage("", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/GenshinImpact_5.0.0.zip"})
	hotfix.AddPackage("", &store.Package{Kind: store.KindAudio, Language: "ko-kr", URL: "https://example.invalid/Audio_Korean_5.0.1.zip", MD5: "bbb", Size: 100, HTTPStatus: 200})
	hotfix.Full[1].SetServedSize(90)

	live := s.Get("7.1.0")
	live.SeenMain("2026-10-26")
	man := func(field, cat string) *store.SophonManifest {
		return &store.SophonManifest{CategoryID: cat, MatchingField: field, ManifestID: "manifest_" + field,
			ManifestURLPrefix: "https://example.invalid/manifests/B", DataURLPrefix: "https://example.invalid/chunks/B",
			Stats: &store.Stats{CompressedSize: 2 << 30, UncompressedSize: 3 << 30, FileCount: 7, ChunkCount: 9}}
	}
	live.SetSophon(&store.SophonBuild{BuildID: "B", Manifests: []*store.SophonManifest{man("ja-jp", "10020"), man("game", "10016")}})
	pm := man("game", "10016")
	pm.Stats, pm.DataURLPrefix = nil, "https://example.invalid/diffs/P/10016"
	pm.StatsByFrom = map[string]*store.Stats{"7.0.0": {CompressedSize: 1 << 30}, "6.7.0": {CompressedSize: 2 << 30}}
	live.SetSophonPatch(&store.SophonPatch{PatchID: "P", BuildID: "B", DiffTags: []string{"6.7.0", "7.0.0"}, Manifests: []*store.SophonManifest{pm}})

	pre := s.Get("7.2.0")
	pre.SeenPreDownload("2026-10-26")
	pre.SetSophon(&store.SophonBuild{BuildID: "C", Manifests: []*store.SophonManifest{man("game", "10016")}})

	backfilled := s.Get("6.0.0")
	backfilled.SetSophon(&store.SophonBuild{BuildID: "D", Manifests: []*store.SophonManifest{man("game", "10016")}})
	return s
}

func render(t *testing.T, s *store.Store) Files {
	t.Helper()
	f, err := Render(s)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

// realData renders the committed data, or skips when it is absent.
func realData(t *testing.T) Files {
	t.Helper()
	s, err := store.Load("../../data/versions.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Versions) == 0 {
		t.Skip("no data/versions.json")
	}
	return render(t, s)
}

func TestRenderFileSet(t *testing.T) {
	got := slices.Sorted(maps.Keys(render(t, sample())))
	want := []string{"README.md", "docs/versions/4.x.md", "docs/versions/5.x.md", "docs/versions/6.x.md", "docs/versions/7.x.md"}
	if !slices.Equal(got, want) {
		t.Errorf("files = %q, want %q", got, want)
	}
}

func TestReadmeOrderAndSize(t *testing.T) {
	out := render(t, sample())["README.md"]
	last := -1
	for _, h := range []string{"## Latest", "## How to download", "## Versions", "## Data", "## Credits", "## License"} {
		i := strings.Index(out, "\n"+h+"\n")
		if i < 0 || i < last {
			t.Fatalf("section %q missing or out of order", h)
		}
		last = i
	}
	if strings.Index(out, "## Credits") < strings.LastIndex(out, "</details>") {
		t.Error("credits come before the version index")
	}
	if strings.Contains(out, "<details open") {
		t.Error("a major is expanded by default")
	}
	if strings.Count(out, "<details>") != 4 || strings.Count(out, "\n</details>\n") != 4 {
		t.Error("details blocks not balanced")
	}
	if strings.Contains(out, "|\n</details>") {
		t.Error("table runs into </details>")
	}
	for _, leak := range []string{"| File | Content |", "[manifest]("} {
		if strings.Contains(out, leak) {
			t.Errorf("README holds detail table %q", leak)
		}
	}
	summaries := regexp.MustCompile(`<summary><b>Version (\d)\.x</b>: ([^·]+) ·`).FindAllStringSubmatch(out, -1)
	var got []string
	for _, m := range summaries {
		got = append(got, m[1]+" "+m[2])
	}
	want := []string{"7 7.1.0 to 7.2.0 (2 versions)", "6 6.0.0 (1 version)", "5 5.0.0 to 5.0.1 (2 versions)", "4 4.8.0 (1 version)"}
	if !slices.Equal(got, want) {
		t.Errorf("summaries = %q, want %q", got, want)
	}
	if !strings.HasSuffix(out, "do not edit by hand.\n") || strings.HasSuffix(out, "\n\n") {
		t.Error("README does not end with exactly one newline after the footer")
	}

	// The daily job runs these tests before fetching, so the real-data budget
	// leaves room for decades of new versions (about 150 bytes each) and
	// only catches detail tables leaking back into the README.
	if f := realData(t); len(f["README.md"]) >= 60_000 {
		t.Errorf("README is %d bytes, want < 60000", len(f["README.md"]))
	}
}

func TestReadmeLatest(t *testing.T) {
	out := render(t, sample())["README.md"]
	for _, want := range []string{
		"| Channel | Version | Live in API (UTC) | Links |",
		"| Live | **7.1.0** | since 2026-10-26 | [Sophon build and update](docs/versions/7.x.md#710) |",
		"| Pre-download | **7.2.0** | pre-download since 2026-10-26 | [Sophon build](docs/versions/7.x.md#720) |",
		"| Newest zip archives | **5.0.1** | — | [Download links](docs/versions/5.x.md#501) |",
		"Every version, by major: [7.x](docs/versions/7.x.md) · [6.x](docs/versions/6.x.md) · [5.x](docs/versions/5.x.md) · [4.x](docs/versions/4.x.md)\n",
		"- **5.0.1 and earlier:**",
		"- **After 5.0.1:**",
		"The checks cover 2024-07-18 to 2024-08-28, and since 2026-10-26;",
		"or an unknown size.\n\nEvery version, by major:",
		"![Live 7.1.0](https://img.shields.io/badge/live-7.1.0-blue)",
		"JRSKelvin/Genshin-Repository-List",
		"rapdodge/GenshinRepository",
		"Last checked: 2026-10-26 (UTC)",
		"CC BY-NC-SA 4.0",
		"Apache License 2.0",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("README lacks %q", want)
		}
	}
}

func TestIndexRows(t *testing.T) {
	f := render(t, sample())
	out := f["README.md"]
	for _, want := range []string{
		"| [7.1.0](docs/versions/7.x.md#710) | since 2026-10-26 | — | 2.00 GiB | live |",
		"| [7.2.0](docs/versions/7.x.md#720) | pre-download since 2026-10-26 | — | 2.00 GiB | pre-download |",
		"| [4.8.0](docs/versions/4.x.md#480) | 2024-07-18 → 2024-08-27 | 15.00 GiB | — | 1 dead link |",
		"| [5.0.1](docs/versions/5.x.md#501) | — | — | — | MD5 may not match |",
		"| [6.0.0](docs/versions/6.x.md#600) | — | — | 2.00 GiB |  |",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("README lacks row %q", want)
		}
	}
	if !strings.Contains(f["docs/versions/4.x.md"], "| [4.8.0](#480) | 2024-07-18 → 2024-08-27 |") {
		t.Error("major page index does not link in-page")
	}

	s := &store.Store{}
	v := s.Get("3.2.0")
	v.AddPackage("", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/a.zip", HTTPStatus: 404})
	v.AddPackage("3.1.0", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/b.zip", HTTPStatus: 403})
	if out := render(t, s)["README.md"]; !strings.Contains(out, "| — | — | — | all links dead |") {
		t.Error("all-dead version is not flagged")
	}
	v.AddPackage("3.1.0", &store.Package{Kind: store.KindAudio, URL: "https://example.invalid/c.zip", HTTPStatus: 200})
	v.AddPackage("3.0.0", &store.Package{Kind: store.KindGame, URL: "https://example.invalid/d.zip", HTTPStatus: 404})
	if out := render(t, s)["README.md"]; !strings.Contains(out, "| full install dead, 2 dead update links |") {
		t.Error("dead full install is not flagged")
	}
}

func TestGameZip(t *testing.T) {
	game := func(part int, n int64) *store.Package {
		return &store.Package{Kind: store.KindGame, Part: part, Size: n}
	}
	audio := &store.Package{Kind: store.KindAudio, Size: 100}
	for _, c := range []struct {
		full []*store.Package
		want int64
	}{
		{[]*store.Package{game(0, 50), game(1, 30), game(2, 20), audio}, 50}, // parts win over the joined archive
		{[]*store.Package{game(0, 40), audio}, 40},
		{[]*store.Package{game(1, 30), game(2, 0)}, 0},
		{[]*store.Package{audio}, 0},
	} {
		if got := gameZip(c.full); got != c.want {
			t.Errorf("gameZip = %d, want %d", got, c.want)
		}
	}
}

// heading anchors of a Markdown file, with GitHub's -1, -2 suffixes for
// repeated headings.
func anchors(md string) map[string]bool {
	seen := map[string]int{}
	out := map[string]bool{}
	for _, m := range regexp.MustCompile(`(?m)^#{1,6} (.+)$`).FindAllStringSubmatch(md, -1) {
		a := anchor(m[1])
		if n := seen[a]; n > 0 {
			out[a+"-"+strconv.Itoa(n)] = true
		} else {
			out[a] = true
		}
		seen[a]++
	}
	return out
}

func checkLinks(t *testing.T, f Files) {
	t.Helper()
	handWritten := map[string]bool{
		"docs/download-guide.md": true, "docs/sophon.md": true, "docs/data.md": true,
		"LICENSE": true, "LICENSE-DATA": true, "data/versions.json": true,
		"data": true, "docs": true, "docs/versions": true,
	}
	link := regexp.MustCompile(`\]\(([^)\s]*)\)|href="([^"]*)"`)
	for src, body := range f {
		for _, m := range link.FindAllStringSubmatch(body, -1) {
			target := m[1] + m[2]
			if strings.HasPrefix(target, "https://") {
				continue
			}
			file, frag, _ := strings.Cut(target, "#")
			dest := src
			if file != "" {
				dest = path.Clean(path.Join(path.Dir(src), file))
			}
			content, ok := f[dest]
			if !ok && !handWritten[dest] {
				t.Errorf("%s: link %q points at %q, which is not generated or known", src, target, dest)
				continue
			}
			if frag != "" && (!ok || !anchors(content)[frag]) {
				t.Errorf("%s: link %q: no heading with anchor %q in %s", src, target, frag, dest)
			}
		}
	}
}

func TestLinksResolve(t *testing.T) {
	checkLinks(t, render(t, sample()))
	checkLinks(t, realData(t))
}

func TestHandWrittenDocsExist(t *testing.T) {
	for _, p := range []string{"download-guide.md", "sophon.md", "data.md"} {
		if _, err := os.Stat(filepath.Join("..", "..", "docs", p)); err != nil {
			t.Error(err)
		}
	}
}

func TestAnchor(t *testing.T) {
	for in, want := range map[string]string{
		"5.5.0":                  "550",
		"Update from 5.4.0":      "update-from-540",
		"Sophon full build `ab`": "sophon-full-build-ab",
		"Genshin Impact 5.x":     "genshin-impact-5x",
	} {
		if got := anchor(in); got != want {
			t.Errorf("anchor(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestShieldEscapes(t *testing.T) {
	for _, c := range [][4]string{
		{"last checked", "2026-10-26", "success", "https://img.shields.io/badge/last%20checked-2026--10--26-success"},
		{"data", "CC BY-NC-SA 4.0", "lightgrey", "https://img.shields.io/badge/data-CC%20BY--NC--SA%204.0-lightgrey"},
		{"a_b", "1+2", "blue", "https://img.shields.io/badge/a__b-1%2B2-blue"},
	} {
		if got := shield(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("shield(%q, %q) = %q, want %q", c[0], c[1], got, c[3])
		}
	}
}

func TestMajorPageContent(t *testing.T) {
	f := render(t, sample())
	seven := f["docs/versions/7.x.md"]
	for page, wants := range map[string][]string{
		"docs/versions/7.x.md": {
			"# Genshin Impact 7.x\n",
			"[← All versions](../../README.md#versions)",
			"## 7.1.0\n\n**Live version** in the API since 2026-10-26 (UTC).\n",
			"## 7.2.0\n\n**Pre-download** in the API since 2026-10-26 (UTC).\n",
			"### Sophon full build `B`",
			"| Game | 2.00 GiB | 3.00 GiB | 7 | 9 | [manifest](https://example.invalid/manifests/B/manifest_game) |",
			"Chunk URL prefix: `https://example.invalid/chunks/B`",
			"### Sophon update `P` (from 7.0.0, 6.7.0)",
			"| Content | From 7.0.0 | From 6.7.0 | Manifest | Diff ID |",
			"| `10016` |\n",
			"Diff URL prefix: `https://example.invalid/diffs/P/<diff ID>`",
			"| Game | 1.00 GiB | 2.00 GiB |",
			"Last checked: 2026-10-26 (UTC)",
		},
		"docs/versions/6.x.md": {
			"## 6.0.0\n\nNo daily API check saw this version as live or pre-download.\n",
			"Only the full build is recorded",
		},
		"docs/versions/5.x.md": {
			"The game archive `GenshinImpact_5.0.0.zip` is named for an earlier release: upstream lists it as the 5.0.1 game data.",
			"`bbb` (published for 100 bytes; the served file is 90 bytes, so it may not match)",
		},
		"docs/versions/4.x.md": {
			"**Live version** in the API from 2024-07-18 to 2024-08-27 (UTC).",
			"[GenshinImpact_4.8.0.zip.001](https://example.invalid/GenshinImpact_4.8.0.zip.001) | Game, part 1 of 2 | 10.00 GiB | `aaa` |",
			"(dead link, HTTP 404)",
			"[Audio_English(US)_4.8.0.zip](https://example.invalid/Audio_English%28US%29_4.8.0.zip) | English voice-over | 3.00 MiB |",
			"### Update from 4.7.0\n",
			"| Game | unknown |",
		},
	} {
		for _, want := range wants {
			if !strings.Contains(f[page], want) {
				t.Errorf("%s lacks %q", page, want)
			}
		}
	}
	if strings.Index(seven, "## 7.2.0") > strings.Index(seven, "## 7.1.0") {
		t.Error("versions are not newest first")
	}
	if strings.Index(seven, "| Game | 2.00 GiB") > strings.Index(seven, "| Japanese voice-over | 2.00 GiB") {
		t.Error("manifest order")
	}
	for _, body := range f {
		if regexp.MustCompile(`(?m)^#+ \d+\.\d+\.\d+ \(`).MatchString(body) {
			t.Error("a version heading carries a status suffix")
		}
	}
}

func TestWriteRemovesStale(t *testing.T) {
	root := t.TempDir()
	vdir := filepath.Join(root, "docs", "versions")
	if err := os.MkdirAll(vdir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{filepath.Join(vdir, "9.x.md"), filepath.Join(root, "docs", "keep.md")} {
		if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	f := render(t, sample())
	if err := Write(root, f); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(vdir, "9.x.md")); !os.IsNotExist(err) {
		t.Error("stale page 9.x.md was kept")
	}
	if _, err := os.Stat(filepath.Join(root, "docs", "keep.md")); err != nil {
		t.Error("hand-written doc was removed")
	}
	for p, want := range f {
		got, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
		if err != nil || string(got) != want {
			t.Errorf("%s not written as rendered", p)
		}
	}
}

func TestRenderDeterministic(t *testing.T) {
	if !maps.Equal(render(t, sample()), render(t, sample())) {
		t.Error("two renders of the same data differ")
	}
}

func TestDiffDir(t *testing.T) {
	m := func(prefix, suffix string) *store.SophonManifest {
		return &store.SophonManifest{DataURLPrefix: prefix, DataURLSuffix: suffix}
	}
	for _, tc := range []struct {
		ms   []*store.SophonManifest
		want string
	}{
		{[]*store.SophonManifest{m("https://x/d/P/1", ""), m("https://x/d/P/2", "")}, "https://x/d/P"},
		{[]*store.SophonManifest{m("https://x/d/P/1", ""), m("https://x/d/Q/2", "")}, ""},
		{[]*store.SophonManifest{m("https://x/d/P/1", "?s")}, ""},
		{[]*store.SophonManifest{m("https://x/d/P/", "")}, ""},
	} {
		if got := diffDir(tc.ms); got != tc.want {
			t.Errorf("diffDir = %q, want %q", got, tc.want)
		}
	}
}
