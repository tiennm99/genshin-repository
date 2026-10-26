package readme

import (
	"regexp"
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

func TestRenderMajorsCollapsedNewestFirst(t *testing.T) {
	out := Render(sample())
	if strings.Contains(out, "<details open") {
		t.Error("a major is expanded by default")
	}
	summaries := regexp.MustCompile(`(?m)^<details>\n<summary><b>Version (\d+)\.x</b>: ([^<]+)</summary>\n(?:\nJump to:[^\n]*\n)?\n### `).FindAllStringSubmatch(out, -1)
	var got []string
	for _, m := range summaries {
		got = append(got, m[1]+" "+m[2])
	}
	want := []string{"7 7.1.0 to 7.2.0 (2 versions)", "6 6.0.0 (1 version)", "5 5.0.0 to 5.0.1 (2 versions)", "4 4.8.0 (1 version)"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("summaries = %q, want %q", got, want)
	}
	if strings.Count(out, "<details>") != 4 || strings.Count(out, "\n</details>\n") != 4 {
		t.Error("details blocks not balanced")
	}
	// Blank line before </details> so the Markdown inside renders.
	if strings.Contains(out, "|\n</details>") {
		t.Error("table runs into </details>")
	}
	if strings.Index(out, "### 7.2.0") > strings.Index(out, "### 7.1.0") {
		t.Error("versions inside a major are not newest first")
	}
}

func TestRenderVersionContent(t *testing.T) {
	out := Render(sample())
	for _, want := range []string{
		"| Live version | 7.1.0 |",
		"| Pre-download | 7.2.0 |",
		"| Newest version with zip archives | 5.0.1 |",
		"Jump to: [7.2.0](#720-pre-download) · [7.1.0](#710-live)\n",
		"The game archive `GenshinImpact_5.0.0.zip` is named for an earlier release: upstream lists it as the 5.0.1 game data.",
		"`bbb` (published for 100 bytes; the served file is 90 bytes, so it may not match)",
		"### 7.1.0 (live)",
		"### 7.2.0 (pre-download)",
		"[GenshinImpact_4.8.0.zip.001](https://example.invalid/GenshinImpact_4.8.0.zip.001) | Game, part 1 of 2 | 10.00 GiB | `aaa` |",
		"(dead link, HTTP 404)",
		"[Audio_English(US)_4.8.0.zip](https://example.invalid/Audio_English%28US%29_4.8.0.zip) | English voice-over | 3.00 MiB |",
		"#### Update from 4.7.0\n",
		"| Game | unknown |",
		"live version from 2024-07-18 to 2024-08-27",
		"#### Sophon full build `B`",
		"| Game | 2.00 GiB | 3.00 GiB | 7 | 9 | [manifest](https://example.invalid/manifests/B/manifest_game) |",
		"Chunk URL prefix: `https://example.invalid/chunks/B`",
		"#### Sophon update `P` (from 7.0.0, 6.7.0)",
		"| Content | From 7.0.0 | From 6.7.0 | Manifest | Diff URL prefix |",
		"| Game | 1.00 GiB | 2.00 GiB |",
		"Only the full build is recorded",
		"## How to download 5.6 and later",
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
	// Game is listed before voice-over packs.
	if strings.Index(out, "| Game | 2.00 GiB") > strings.Index(out, "| Japanese voice-over | 2.00 GiB") {
		t.Error("manifest order")
	}
}

func TestRenderDeterministic(t *testing.T) {
	first, second := Render(sample()), Render(sample())
	if first != second {
		t.Error("two renders of the same data differ")
	}
}
