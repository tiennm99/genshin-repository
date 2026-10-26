package upstream

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
	"testing"

	"github.com/tiennm99/genshin-repository/internal/store"
)

func read(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// Upstream 1.6.1 relists 1.6.0 files; they belong to 1.6.0.
func TestParseMarkdownFilesLinksByURLVersion(t *testing.T) {
	s := &store.Store{}
	ParseMarkdown(s, "Version 1.6.1.md", read(t, "version-1.6.1.md"))
	s.Normalize()

	v161 := s.Find("1.6.1")
	if v161 == nil || len(v161.Patches) != 0 {
		t.Fatalf("1.6.1 = %+v", v161)
	}
	kinds := ""
	for _, p := range v161.Full {
		kinds += p.Kind + ":" + p.Language + " "
	}
	if kinds != "launcher: audio:zh-cn audio:en-us " {
		t.Errorf("1.6.1 full = %q", kinds)
	}

	v160 := s.Find("1.6.0")
	if v160 == nil || len(v160.Full) != 1 || v160.Full[0].Kind != store.KindGame || v160.Full[0].Part != 0 {
		t.Fatalf("1.6.0 full = %+v", v160)
	}
	if len(v160.Patches) != 1 || v160.Patches[0].From != "1.5.1" || len(v160.Patches[0].Packages) != 2 {
		t.Fatalf("1.6.0 patches = %+v", v160.Patches)
	}
	if v160.Patches[0].Packages[1].Language != "ko-kr" {
		t.Errorf("audio language from URL = %q", v160.Patches[0].Packages[1].Language)
	}
}

func TestParseMarkdownCRLFAndParts(t *testing.T) {
	s := &store.Store{}
	ParseMarkdown(s, "Version 5.5.0.md", read(t, "version-5.5.0-crlf.md"))
	v := s.Find("5.5.0")
	if v == nil || len(v.Full) != 2 || len(v.Patches) != 1 {
		t.Fatalf("5.5.0 = %+v", v)
	}
	s.Normalize()
	if v.Full[0].Part != 2 || v.Full[1].Language != "ko-kr" {
		t.Errorf("full = %+v %+v", v.Full[0], v.Full[1])
	}
	if p := v.Patches[0]; p.From != "5.4.0" || p.Packages[0].Language != "zh-cn" {
		t.Errorf("patch = %+v", p)
	}
	if v.Sources[0] != store.SourceMD {
		t.Errorf("sources = %v", v.Sources)
	}
}

func TestApplyLegacySnapshot(t *testing.T) {
	s := &store.Store{}
	format, err := ApplySnapshot(s, "2023-04-10", read(t, "legacy-snapshot.json"))
	if err != nil || format != "legacy" {
		t.Fatalf("format %q, err %v", format, err)
	}
	v := s.Find("3.6.0")
	if v.MainFirstSeen != "2023-04-10" || v.ResListURL == "" {
		t.Errorf("3.6.0 = %+v", v)
	}
	if len(v.Full) != 2 || v.Full[0].MD5 != "aa11" || v.Full[0].Part != 1 {
		t.Errorf("full = %+v", v.Full)
	}
	audio := v.Full[1]
	if audio.Size != 100 || audio.DecompressedSize != 200 {
		t.Errorf("legacy sizes not swapped to download/unpacked: %+v", audio)
	}
	if len(v.Patches) != 1 || v.Patches[0].From != "3.5.0" || len(v.Patches[0].Packages) != 2 {
		t.Errorf("patches = %+v", v.Patches)
	}
	pre := s.Find("3.7.0")
	if pre.PreDownloadFirstSeen != "2023-04-10" || pre.MainFirstSeen != "" || len(pre.Full) != 1 || pre.Full[0].Part != 0 {
		t.Errorf("pre-download 3.7.0 = %+v", pre)
	}
}

func TestApplyHypSnapshotIgnoresFrozenDates(t *testing.T) {
	s := &store.Store{}
	for _, d := range []string{"2025-03-26", "2025-05-06", ZipAPIFrozenFrom, "2026-10-25"} {
		if f, err := ApplySnapshot(s, d, read(t, "hyp-snapshot.json")); err != nil || f != "hyp" {
			t.Fatalf("format %q, err %v", f, err)
		}
	}
	v := s.Find("5.5.0")
	if v.MainFirstSeen != "2025-03-26" || v.MainLastSeen != "2025-05-06" {
		t.Errorf("seen %s..%s; snapshots after the zip freeze must not count", v.MainFirstSeen, v.MainLastSeen)
	}
	if len(v.Full) != 2 || v.Full[1].Part != 2 || len(s.Versions) != 1 {
		t.Errorf("full = %+v", v.Full)
	}
}

func TestReadTarGzAndImport(t *testing.T) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	add := func(name string, body []byte) {
		if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o644, Size: int64(len(body)), Typeflag: tar.TypeReg}); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write(body); err != nil {
			t.Fatal(err)
		}
	}
	add("repo-abc/Version 1.6.1.md", read(t, "version-1.6.1.md"))
	add("repo-abc/README.md", []byte("# readme"))
	add("repo-abc/fetch/original/result-2023-04-10-17-10-17.json", read(t, "legacy-snapshot.json"))
	add("repo-abc/fetch/result/result-2023-04-10-17-10-17.json", []byte("not json"))
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	files, err := ReadTarGz(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["Version 1.6.1.md"]; !ok {
		t.Fatalf("top directory not stripped: %v", files)
	}
	s := &store.Store{}
	if n := ImportMarkdownRepo(s, files); n != 1 {
		t.Errorf("md files = %d", n)
	}
	counts, err := ImportSnapshotRepo(s, files)
	if err != nil || counts["legacy"] != 1 || len(counts) != 1 {
		t.Errorf("counts = %v, err %v", counts, err)
	}
}
