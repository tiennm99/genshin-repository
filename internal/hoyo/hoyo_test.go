package hoyo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/tiennm99/genshin-repository/internal/store"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func data(t *testing.T, name string) []byte {
	t.Helper()
	var env envelope
	if err := json.Unmarshal(fixture(t, name), &env); err != nil {
		t.Fatal(err)
	}
	return env.Data
}

func TestDecodeBuild(t *testing.T) {
	b, err := DecodeBuild(data(t, "get_build.json"))
	if err != nil {
		t.Fatal(err)
	}
	if b.Tag != "7.1.0" || b.Build.BuildID != "sMXGW2ll3Fuu" || len(b.Build.Manifests) != 5 {
		t.Fatalf("got tag %s build %s, %d manifests", b.Tag, b.Build.BuildID, len(b.Build.Manifests))
	}
	var game *store.SophonManifest
	for _, m := range b.Build.Manifests {
		if m.MatchingField == "game" {
			game = m
		}
	}
	if game == nil || game.Stats == nil {
		t.Fatal("no game manifest with stats")
	}
	if game.Stats.CompressedSize != 128310073410 || game.Stats.FileCount != 2744 {
		t.Errorf("game stats = %+v", game.Stats)
	}
	if !strings.Contains(game.DataURLPrefix, "/sophon/chunks/") {
		t.Errorf("chunk prefix = %s", game.DataURLPrefix)
	}
	want := "https://autopatchhk.yuanshen.com/client_app/sophon/manifests/cxhpq4g4rgg0/sMXGW2ll3Fuu/" + game.ManifestID
	if game.ManifestURL() != want {
		t.Errorf("manifest URL = %s", game.ManifestURL())
	}
}

func TestDecodePatch(t *testing.T) {
	p, err := DecodePatch(data(t, "get_patch_build.json"))
	if err != nil {
		t.Fatal(err)
	}
	if p.Tag != "7.1.0" || p.Patch.PatchID != "t663grrK4Of4" {
		t.Fatalf("got %+v", p)
	}
	slices.Sort(p.Patch.DiffTags)
	if !slices.Equal(p.Patch.DiffTags, []string{"6.7.0", "7.0.0"}) {
		t.Errorf("diff tags = %v", p.Patch.DiffTags)
	}
	for _, m := range p.Patch.Manifests {
		if m.Stats != nil || len(m.StatsByFrom) != 2 {
			t.Errorf("%s: stats %v, by-from %d", m.MatchingField, m.Stats, len(m.StatsByFrom))
		}
		if !strings.Contains(m.DataURLPrefix, "/sophon/diffs/") {
			t.Errorf("%s: diff prefix = %s", m.MatchingField, m.DataURLPrefix)
		}
	}
}

func TestApplyBranchZip(t *testing.T) {
	var d struct {
		GamePackages []GamePackages `json:"game_packages"`
	}
	if err := json.Unmarshal(data(t, "get_game_packages.json"), &d); err != nil {
		t.Fatal(err)
	}
	s := &store.Store{}
	if v := ApplyBranch(s, d.GamePackages[0].Main, store.SourceZipAPI); v != "5.5.0" {
		t.Fatalf("main = %q", v)
	}
	if v := ApplyBranch(s, d.GamePackages[0].PreDownload, store.SourceZipAPI); v != "" {
		t.Fatalf("empty pre_download gave %q", v)
	}
	v := s.Find("5.5.0")
	if len(v.Full) != 12 {
		t.Fatalf("full = %d, want 8 parts + 4 audio", len(v.Full))
	}
	s.Normalize()
	if v.Full[0].Part != 1 || v.Full[7].Part != 8 || v.Full[0].Size != 10737418240 || v.Full[0].DecompressedSize != 21485322240 {
		t.Errorf("first/last part = %+v / %+v", v.Full[0], v.Full[7])
	}
	if v.Full[8].Language != "zh-cn" || v.Full[11].Language != "ko-kr" {
		t.Errorf("audio order = %s..%s", v.Full[8].Language, v.Full[11].Language)
	}
	if len(v.Patches) != 2 || v.Patches[0].From != "5.4.0" || len(v.Patches[0].Packages) != 5 {
		t.Errorf("patches = %+v", v.Patches)
	}
	if !strings.HasSuffix(v.ResListURL, "/ScatteredFiles") {
		t.Errorf("res_list_url = %s", v.ResListURL)
	}
}

// fakeAPI serves the fixtures. getPatchBuild answers only POST, like the
// real endpoint; unknown getBuild tags answer retcode -202.
func fakeAPI(t *testing.T, branches []byte) *Client {
	t.Helper()
	builds := map[string][]byte{
		"":      fixture(t, "get_build.json"),
		"7.1.0": fixture(t, "get_build.json"),
		"7.0.0": fixture(t, "get_build_7.0.0.json"),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hyp/getGamePackages", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("game_ids[]") != GameID {
			http.Error(w, "missing game_ids[]", http.StatusBadRequest)
			return
		}
		_, _ = w.Write(fixture(t, "get_game_packages.json"))
	})
	mux.HandleFunc("/hyp/getGameBranches", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(branches)
	})
	mux.HandleFunc("/sophon/getBuild", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		b, ok := builds[r.URL.Query().Get("tag")]
		if !ok {
			_, _ = w.Write([]byte(`{"retcode":-202,"message":"not found","data":null}`))
			return
		}
		_, _ = w.Write(b)
	})
	mux.HandleFunc("/sophon/getPatchBuild", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		_, _ = w.Write(fixture(t, "get_patch_build.json"))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return &Client{HTTP: srv.Client(), HypBase: srv.URL + "/hyp", SophonBase: srv.URL + "/sophon"}
}

func TestSync(t *testing.T) {
	c := fakeAPI(t, fixture(t, "get_game_branches.json"))
	s := &store.Store{}
	if err := Sync(context.Background(), c, s, Options{Date: "2026-10-26", Backfill: true}); err != nil {
		t.Fatal(err)
	}
	zip := s.Find("5.5.0")
	if zip == nil || len(zip.Full) != 12 {
		t.Fatal("5.5.0 zip archives missing")
	}
	if zip.MainFirstSeen != "" {
		t.Errorf("frozen zip API must not mark 5.5.0 live, got %s", zip.MainFirstSeen)
	}
	live := s.Find("7.1.0")
	if live == nil || live.Sophon == nil || live.SophonPatch == nil {
		t.Fatalf("7.1.0 sophon data missing: %+v", live)
	}
	if live.MainFirstSeen != "2026-10-26" || live.MainLastSeen != "2026-10-26" {
		t.Errorf("7.1.0 seen = %s..%s", live.MainFirstSeen, live.MainLastSeen)
	}
	if !slices.Contains(live.SophonPatch.DiffTags, "7.0.0") || !slices.Contains(live.SophonPatch.DiffTags, "6.7.0") {
		t.Errorf("diff tags = %v", live.SophonPatch.DiffTags)
	}
	// 7.0.0 is a diff tag with a build; 6.7.0 and the probed X.Y.0 tags are
	// unknown to the fake API and must not create empty versions.
	if v := s.Find("7.0.0"); v == nil || v.Sophon == nil || v.Sophon.BuildID != "3XnPviMfoILi" {
		t.Errorf("7.0.0 backfill = %+v", v)
	}
	if v := s.Find("6.7.0"); v != nil {
		t.Errorf("unknown tag created %+v", v)
	}
	if len(s.Versions) != 3 {
		t.Errorf("versions = %d, want 5.5.0, 7.0.0, 7.1.0", len(s.Versions))
	}
}

func TestSyncPreDownload(t *testing.T) {
	var br map[string]any
	if err := json.Unmarshal(fixture(t, "get_game_branches.json"), &br); err != nil {
		t.Fatal(err)
	}
	gb := br["data"].(map[string]any)["game_branches"].([]any)[0].(map[string]any)
	main := gb["main"].(map[string]any)
	// The patch fixture is for 7.1.0, so pretend 7.1.0 is the pre-download
	// on top of a live 7.0.0.
	pre := map[string]any{}
	for k, v := range main {
		pre[k] = v
	}
	pre["branch"] = "predownload"
	main["tag"] = "7.0.0"
	main["diff_tags"] = []string{}
	gb["pre_download"] = pre
	raw, _ := json.Marshal(br)

	s := &store.Store{}
	if err := Sync(context.Background(), fakeAPI(t, raw), s, Options{Date: "2026-11-01"}); err != nil {
		t.Fatal(err)
	}
	v := s.Find("7.1.0")
	if v == nil || v.PreDownloadFirstSeen != "2026-11-01" || v.MainFirstSeen != "" || v.SophonPatch == nil {
		t.Fatalf("pre-download 7.1.0 = %+v", v)
	}
	if m := s.Find("7.0.0"); m == nil || m.MainFirstSeen != "2026-11-01" {
		t.Fatalf("main 7.0.0 = %+v", m)
	}
}

func TestAPIError(t *testing.T) {
	err := decodeEnvelope("getBuild", []byte(`{"retcode":-202,"message":"not found","data":null}`), &struct{}{})
	if !IsNotFound(err) {
		t.Fatalf("err = %v", err)
	}
	err = decodeEnvelope("getBuild", []byte(`{"retcode":-201,"message":"invalid params","data":null}`), &struct{}{})
	if err == nil || IsNotFound(err) {
		t.Fatalf("err = %v", err)
	}
}

func TestCandidateTags(t *testing.T) {
	got := candidateTags("6.2.0")
	want := []string{"5.6.0", "5.7.0", "5.8.0", "5.9.0", "6.0.0", "6.1.0"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
