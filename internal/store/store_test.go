package store

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestAddPackageMergesByURL(t *testing.T) {
	v := &Version{Version: "5.5.0"}
	v.AddPackage("", &Package{Kind: KindGame, Part: 1, URL: "u1", MD5: "ABC", Size: 10})
	v.AddPackage("", &Package{Kind: KindGame, URL: "u1", Size: 0, HTTPStatus: 200})
	if len(v.Full) != 1 {
		t.Fatalf("full = %d", len(v.Full))
	}
	p := v.Full[0]
	if p.MD5 != "abc" || p.Size != 10 || p.Part != 1 || p.HTTPStatus != 200 {
		t.Errorf("merged = %+v", p)
	}
	v.AddPackage("", &Package{Kind: KindGame, URL: "u1", Size: 12})
	if p.Size != 10 {
		t.Errorf("snapshot size replaced a Content-Length size: %d", p.Size)
	}
	v.AddPackage("", &Package{Kind: KindGame, URL: "u1", Size: 14, HTTPStatus: 200})
	if p.Size != 14 {
		t.Errorf("new Content-Length size not applied: %d", p.Size)
	}
	w := &Version{}
	w.AddPackage("", &Package{Kind: KindGame, URL: "u2", Size: 1})
	w.AddPackage("", &Package{Kind: KindGame, URL: "u2", Size: 2})
	if w.Full[0].Size != 2 {
		t.Errorf("unchecked size not updated: %d", w.Full[0].Size)
	}
}

func TestAddPackageKeepsDroppedPatches(t *testing.T) {
	s := &Store{}
	v := s.Get("5.5")
	v.AddPackage("5.3", &Package{Kind: KindGame, URL: "p53"})
	v.AddPackage("5.4.0", &Package{Kind: KindGame, URL: "p54"})
	// A later fetch only lists the 5.4.0 patch; 5.3.0 must stay.
	s.Get("5.5.0").AddPackage("5.4.0", &Package{Kind: KindGame, URL: "p54", MD5: "x"})
	if len(s.Versions) != 1 || len(v.Patches) != 2 {
		t.Fatalf("versions %d, patches %d", len(s.Versions), len(v.Patches))
	}
	s.Normalize()
	if v.Patches[0].From != "5.4.0" || v.Patches[1].From != "5.3.0" {
		t.Errorf("patch order = %s, %s", v.Patches[0].From, v.Patches[1].From)
	}
}

func TestSetSophonIgnoresEmpty(t *testing.T) {
	v := &Version{}
	b := &SophonBuild{BuildID: "a", Manifests: []*SophonManifest{{MatchingField: "game"}}}
	v.SetSophon(b)
	v.SetSophon(&SophonBuild{BuildID: "empty"})
	v.SetSophon(nil)
	v.SetSophonPatch(&SophonPatch{PatchID: "empty"})
	if v.Sophon != b || v.SophonPatch != nil {
		t.Errorf("sophon = %+v, patch = %+v", v.Sophon, v.SophonPatch)
	}
}

func TestSeenDates(t *testing.T) {
	v := &Version{}
	for _, d := range []string{"2025-03-27", "2025-03-26", "2025-04-01"} {
		v.SeenMain(d)
	}
	v.SeenPreDownload("2025-03-24")
	if v.MainFirstSeen != "2025-03-26" || v.MainLastSeen != "2025-04-01" || v.PreDownloadFirstSeen != "2025-03-24" {
		t.Errorf("dates = %+v", v)
	}
	s := &Store{Versions: []*Version{v}}
	if s.LatestDate() != "2025-04-01" {
		t.Errorf("latest = %s", s.LatestDate())
	}
}

func TestNormalizeOrder(t *testing.T) {
	s := &Store{}
	v := s.Get("10.0.0")
	s.Get("2.0.0")
	s.Get("1.3")
	v.AddPackage("", &Package{Kind: KindAudio, Language: "ko-kr", URL: "k"})
	v.AddPackage("", &Package{Kind: KindAudio, Language: "zh-cn", URL: "c"})
	v.AddPackage("", &Package{Kind: KindGame, Part: 2, URL: "g2"})
	v.AddPackage("", &Package{Kind: KindGame, Part: 1, URL: "g1"})
	v.AddPackage("", &Package{Kind: KindLauncher, URL: "l"})
	v.AddSource(SourceZipAPI)
	v.AddSource(SourceMD)
	v.AddSource(SourceMD)
	s.Normalize()
	var got []string
	for _, x := range s.Versions {
		got = append(got, x.Version)
	}
	if got[0] != "1.3.0" || got[1] != "2.0.0" || got[2] != "10.0.0" {
		t.Errorf("version order = %v", got)
	}
	var urls string
	for _, p := range v.Full {
		urls += p.URL + " "
	}
	if urls != "l g1 g2 c k " {
		t.Errorf("package order = %q", urls)
	}
	if len(v.Sources) != 2 || v.Sources[0] != SourceZipAPI {
		t.Errorf("sources = %v", v.Sources)
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	s := &Store{}
	v := s.Get("5.5.0")
	v.AddSource(SourceZipAPI)
	v.AddPackage("", &Package{Kind: KindGame, Part: 1, URL: "https://example.com/a.zip.001", MD5: "m", Size: 1})
	v.SetSophon(&SophonBuild{BuildID: "b", Manifests: []*SophonManifest{{MatchingField: "game", Stats: &Stats{FileCount: 3}}}})
	path := filepath.Join(t.TempDir(), "versions.json")
	if err := s.Save(path); err != nil {
		t.Fatal(err)
	}
	first, _ := s.Marshal()
	loaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	second, _ := loaded.Marshal()
	if !bytes.Equal(first, second) {
		t.Errorf("round trip changed output:\n%s\n%s", first, second)
	}
	if bytes.Contains(first, []byte(`&`)) {
		t.Error("HTML escaping is on")
	}
	missing, err := Load(filepath.Join(t.TempDir(), "none.json"))
	if err != nil || len(missing.Versions) != 0 {
		t.Errorf("missing file: %v, %v", missing, err)
	}
}

func TestCompareVersions(t *testing.T) {
	cases := []struct {
		a, b string
		sign int
	}{
		{"5.10.0", "5.9.0", 1},
		{"4.0.1", "4.0.0", 1},
		{"1.3.0", "1.3.0", 0},
		{"1.3.2", "1.4.0", -1},
	}
	for _, c := range cases {
		got := CompareVersions(c.a, c.b)
		if (got > 0) != (c.sign > 0) || (got < 0) != (c.sign < 0) {
			t.Errorf("CompareVersions(%s, %s) = %d", c.a, c.b, got)
		}
	}
	if NormalizeVersion("1.3") != "1.3.0" || Major("7.1.0") != "7" {
		t.Error("NormalizeVersion or Major")
	}
}

func TestMD5SizeTracksPublishedSize(t *testing.T) {
	// API first, then a HEAD answer with another size.
	v := &Version{}
	v.AddPackage("", &Package{Kind: KindAudio, URL: "a", MD5: "m1", Size: 100})
	p := v.Full[0]
	p.HTTPStatus = 200
	p.SetServedSize(90)
	if p.Size != 90 || p.MD5Size != 100 {
		t.Errorf("after HEAD: size %d, md5 size %d", p.Size, p.MD5Size)
	}
	// The API repeats the old pair: the mismatch stays recorded.
	v.AddPackage("", &Package{Kind: KindAudio, URL: "a", MD5: "m1", Size: 100})
	if p.Size != 90 || p.MD5Size != 100 {
		t.Errorf("after repeat: size %d, md5 size %d", p.Size, p.MD5Size)
	}
	// A new MD5 published for the served size clears it.
	v.AddPackage("", &Package{Kind: KindAudio, URL: "a", MD5: "m2", Size: 90})
	if p.MD5 != "m2" || p.MD5Size != 0 {
		t.Errorf("after matching MD5: md5 %s, md5 size %d", p.MD5, p.MD5Size)
	}
	// A matching size from HEAD needs no note.
	q := &Package{MD5: "m", Size: 5}
	q.SetServedSize(5)
	if q.MD5Size != 0 {
		t.Errorf("equal sizes set md5 size %d", q.MD5Size)
	}
}
