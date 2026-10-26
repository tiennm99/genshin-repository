// Package render generates README.md and the per-major pages under
// docs/versions/ from the store. The output depends only on the store, so
// unchanged data gives byte-identical files.
package render

import (
	_ "embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/tiennm99/genshin-repository/internal/store"
)

// VersionsDir holds the generated per-major pages, relative to the repo root.
const VersionsDir = "docs/versions"

// Files maps a repo-relative path to its generated content.
type Files map[string]string

//go:embed readme.md.tmpl
var readmeTmpl string

//go:embed major.md.tmpl
var majorTmpl string

var tmpl = template.Must(template.New("readme").Funcs(template.FuncMap{
	"shield":       shield,
	"anchor":       anchor,
	"versionIndex": indexTable,
	"detail":       detail,
}).Parse(readmeTmpl))

var tmplMajor = template.Must(template.Must(tmpl.Clone()).New("major").Parse(majorTmpl))

// page is everything the templates need, computed once per render.
type page struct {
	Live, Pre, Zip *store.Version
	Oldest         string
	Count          int
	Date           string // newest date in the store
	Coverage       string // date ranges the daily checks cover
	Majors         []*major
	Latest         []latestRow
}

// latestRow is one row of the README's Latest table.
type latestRow struct {
	Label, Version, Dates, Link string
}

// major is one major version, newest version first.
type major struct {
	N        string
	Path     string // repo-relative page path
	Versions []*store.Version
	Rows     []row
	Range    string // "7.0.0 to 7.1.0 (2 versions)"
}

// majorPage is the data of one docs/versions page.
type majorPage struct {
	P *page
	M *major
}

// row is one line of a version index.
type row struct {
	Version, Dates, Zip, Sophon, Notes string
}

// Render returns README.md and one docs/versions/<N>.x.md page per major.
func Render(s *store.Store) (Files, error) {
	s.Normalize()
	p := newPage(s)
	f := Files{}
	var b strings.Builder
	if err := tmpl.ExecuteTemplate(&b, "readme", p); err != nil {
		return nil, err
	}
	f["README.md"] = finish(b.String())
	for _, m := range p.Majors {
		b.Reset()
		if err := tmplMajor.ExecuteTemplate(&b, "major", majorPage{p, m}); err != nil {
			return nil, err
		}
		f[m.Path] = finish(b.String())
	}
	return f, nil
}

// finish trims surplus blank lines left by template actions and ends the
// file with exactly one newline.
func finish(s string) string {
	for strings.Contains(s, "\n\n\n") {
		s = strings.ReplaceAll(s, "\n\n\n", "\n\n")
	}
	return strings.TrimSpace(s) + "\n"
}

// Write writes f under root and removes generated pages in VersionsDir that
// f no longer contains. Other files under docs/ are left alone.
func Write(root string, f Files) error {
	for p, c := range f {
		full := filepath.Join(root, filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(full, []byte(c), 0o644); err != nil {
			return err
		}
	}
	dir := filepath.Join(root, filepath.FromSlash(VersionsDir))
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
			continue
		}
		if _, ok := f[VersionsDir+"/"+e.Name()]; ok {
			continue
		}
		if err := os.Remove(filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}

func newPage(s *store.Store) *page {
	p := &page{Count: len(s.Versions), Date: s.LatestDate()}
	if len(s.Versions) > 0 {
		p.Oldest = s.Versions[0].Version
	}
	var liveSeen string
	for _, v := range s.Versions {
		if v.MainLastSeen != "" && v.MainLastSeen >= liveSeen {
			p.Live, liveSeen = v, v.MainLastSeen
		}
		if len(v.Full) > 0 {
			p.Zip = v
		}
	}
	for _, v := range s.Versions {
		if v.PreDownloadLastSeen == p.Date && v.MainLastSeen < p.Date {
			p.Pre = v
		}
	}
	p.Coverage = coverage(s, p.Date)
	p.Latest = p.latest()

	byMajor := map[string]*major{}
	for i := len(s.Versions) - 1; i >= 0; i-- {
		v := s.Versions[i]
		n := store.Major(v.Version)
		m := byMajor[n]
		if m == nil {
			m = &major{N: n, Path: VersionsDir + "/" + n + ".x.md"}
			byMajor[n] = m
			p.Majors = append(p.Majors, m)
		}
		m.Versions = append(m.Versions, v)
		m.Rows = append(m.Rows, p.row(v))
	}
	slices.SortFunc(p.Majors, func(a, b *major) int { return store.CompareVersions(b.N, a.N) })
	for _, m := range p.Majors {
		lo, hi := m.Versions[len(m.Versions)-1].Version, m.Versions[0].Version
		m.Range = lo + " (1 version)"
		if len(m.Versions) > 1 {
			m.Range = lo + " to " + hi + " (" + strconv.Itoa(len(m.Versions)) + " versions)"
		}
	}
	return p
}

func (p *page) row(v *store.Version) row {
	r := row{Version: v.Version, Dates: "—", Zip: "—", Sophon: "—"}
	switch {
	case v == p.Live:
		r.Dates = "since " + v.MainFirstSeen
	case v.MainFirstSeen != "":
		r.Dates = v.MainFirstSeen + " → " + v.MainLastSeen
	case v == p.Pre:
		r.Dates = "pre-download since " + v.PreDownloadFirstSeen
	}
	if n := gameZip(v.Full); n > 0 {
		r.Zip = size(n)
	}
	if v.Sophon != nil {
		for _, m := range v.Sophon.Manifests {
			if m.MatchingField == store.KindGame && m.Stats != nil && m.Stats.CompressedSize > 0 {
				r.Sophon = size(m.Stats.CompressedSize)
			}
		}
	}
	var notes []string
	switch v {
	case p.Live:
		notes = append(notes, "live")
	case p.Pre:
		notes = append(notes, "pre-download")
	}
	notes = append(notes, deadNotes(v)...)
	md5 := false
	for _, pk := range packages(v) {
		md5 = md5 || pk.MD5Size != 0
	}
	if md5 {
		notes = append(notes, "MD5 may not match")
	}
	r.Notes = strings.Join(notes, ", ")
	return r
}

// gameZip is the download size of the game data alone, without voice-over
// packs: the split parts when there are any, else the largest single
// archive. It is 0 when a needed size is unknown.
func gameZip(full []*store.Package) int64 {
	var parts, single int64
	split := false
	for _, pk := range full {
		switch {
		case pk.Kind != store.KindGame:
		case pk.Part > 0:
			if pk.Size <= 0 {
				return 0
			}
			split = true
			parts += pk.Size
		default:
			single = max(single, pk.Size)
		}
	}
	if split {
		return parts
	}
	return single
}

// packages lists every full and update archive of v.
func packages(v *store.Version) []*store.Package {
	all := slices.Clone(v.Full)
	for _, pt := range v.Patches {
		all = append(all, pt.Packages...)
	}
	return all
}

// deadNotes summarises dead links for the index: whether the whole version
// or its full install is gone, else how many links are dead.
func deadNotes(v *store.Version) []string {
	count := func(ps []*store.Package) int {
		n := 0
		for _, pk := range ps {
			if dead(pk) {
				n++
			}
		}
		return n
	}
	plural := func(n int, what string) string {
		if n == 1 {
			return "1 dead " + what
		}
		return strconv.Itoa(n) + " dead " + what + "s"
	}
	all := packages(v)
	total, full := count(all), count(v.Full)
	switch {
	case total == 0:
		return nil
	case total == len(all):
		return []string{"all links dead"}
	case full > 0 && full == len(v.Full):
		if rest := total - full; rest > 0 {
			return []string{"full install dead", plural(rest, "update link")}
		}
		return []string{"full install dead"}
	}
	return []string{plural(total, "link")}
}

// coverage describes the date ranges in which daily checks saw a live
// version, such as "2023-03-17 to 2025-05-06 and since 2026-10-26". A gap of
// more than one day between two live versions splits the ranges.
func coverage(s *store.Store, latest string) string {
	type span struct{ from, to string }
	var spans []span
	for _, v := range s.Versions {
		if v.MainFirstSeen == "" {
			continue
		}
		if n := len(spans); n > 0 && !after(spans[n-1].to, v.MainFirstSeen) {
			spans[n-1].to = max(spans[n-1].to, v.MainLastSeen)
			continue
		}
		spans = append(spans, span{v.MainFirstSeen, v.MainLastSeen})
	}
	var parts []string
	for _, sp := range spans {
		if sp.to == latest {
			parts = append(parts, "every day since "+sp.from)
		} else {
			parts = append(parts, sp.from+" to "+sp.to)
		}
	}
	switch len(parts) {
	case 0:
		return "no dates yet"
	case 1:
		return parts[0]
	}
	return strings.Join(parts[:len(parts)-1], ", ") + ", and " + parts[len(parts)-1]
}

// after reports whether date b is more than one day after date a.
func after(a, b string) bool {
	ta, err1 := time.Parse(time.DateOnly, a)
	tb, err2 := time.Parse(time.DateOnly, b)
	if err1 != nil || err2 != nil {
		return b > a
	}
	return tb.Sub(ta) > 24*time.Hour
}

// indexTable renders the version index of m; prefix is the page the version
// anchors live on, "" for the page itself.
func indexTable(m *major, prefix string) string {
	var b strings.Builder
	writeIndex(&b, m.Rows, prefix)
	return b.String()
}

// detail renders the detail sections of every version in m.
func detail(p *page, m *major) string {
	var b strings.Builder
	for i, v := range m.Versions {
		if i > 0 {
			b.WriteString("\n")
		}
		writeVersion(&b, v, p)
	}
	return b.String()
}

// latest builds the Latest table rows: live version, open pre-download and
// newest version with zip archives.
func (p *page) latest() []latestRow {
	mk := func(label string, v *store.Version, text string) latestRow {
		if v == nil {
			return latestRow{Label: label, Version: "none", Dates: "—", Link: "—"}
		}
		return latestRow{
			Label:   label,
			Version: "**" + v.Version + "**",
			Dates:   p.row(v).Dates,
			Link:    fmt.Sprintf("[%s](%s/%s.x.md#%s)", text, VersionsDir, store.Major(v.Version), anchor(v.Version)),
		}
	}
	return []latestRow{
		mk("Live", p.Live, "Sophon build and update"),
		mk("Pre-download", p.Pre, "Sophon build"),
		mk("Newest zip archives", p.Zip, "Download links"),
	}
}
