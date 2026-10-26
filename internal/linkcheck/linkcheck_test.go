package linkcheck

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tiennm99/genshin-repository/internal/store"
)

func TestCheck(t *testing.T) {
	retryDelay = 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/live", "/revived":
			w.Header().Set("Content-Length", "42")
		case "/gone":
			w.WriteHeader(http.StatusNotFound)
		case "/flaky":
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer srv.Close()

	pkg := func(path string, status int) *store.Package {
		return &store.Package{Kind: store.KindGame, URL: srv.URL + path, Size: 40, HTTPStatus: status}
	}
	live, revived, gone, flaky := pkg("/live", 0), pkg("/revived", 404), pkg("/gone", 200), pkg("/flaky", 200)
	patched := pkg("/gone", 0)
	s := &store.Store{Versions: []*store.Version{{
		Version: "1.0.0",
		Full:    []*store.Package{live, revived, gone, flaky},
		Patches: []*store.Patch{{From: "0.9.0", Packages: []*store.Package{patched}}},
	}}}

	l, d, u := Check(context.Background(), srv.Client(), s, 2)
	if l != 2 || d != 2 || u != 1 {
		t.Errorf("live, dead, unknown = %d, %d, %d; want 2, 2, 1", l, d, u)
	}
	for _, c := range []struct {
		name   string
		p      *store.Package
		status int
	}{
		{"live", live, 200},
		{"revived", revived, 200},
		{"gone", gone, 404},
		{"flaky keeps its last status", flaky, 200},
		{"patch package", patched, 404},
	} {
		if c.p.HTTPStatus != c.status {
			t.Errorf("%s: status %d, want %d", c.name, c.p.HTTPStatus, c.status)
		}
	}
}
