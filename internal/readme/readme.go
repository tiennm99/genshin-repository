// Package readme renders README.md from the store. The output depends only
// on the store, so unchanged data gives a byte-identical README.
package readme

import (
	_ "embed"
	"fmt"
	"path"
	"slices"
	"strconv"
	"strings"

	"github.com/tiennm99/genshin-repository/internal/store"
)

//go:embed intro.md
var intro string

//go:embed guide.md
var guide string

var langName = map[string]string{
	"zh-cn": "Chinese",
	"en-us": "English",
	"ja-jp": "Japanese",
	"ko-kr": "Korean",
}

// Render returns the README for s.
func Render(s *store.Store) string {
	s.Normalize()
	var b strings.Builder
	b.WriteString(intro)
	writeLatest(&b, s)
	b.WriteString(guide)
	writeVersions(&b, s)
	writeFooter(&b, s)
	return b.String()
}

func writeLatest(b *strings.Builder, s *store.Store) {
	live, pre, zip := latest(s)
	b.WriteString("\n## Latest\n\n| | Version |\n| --- | --- |\n")
	row := func(label, v string) {
		if v == "" {
			v = "none"
		}
		fmt.Fprintf(b, "| %s | %s |\n", label, v)
	}
	row("Live version", live)
	row("Pre-download", pre)
	row("Newest version with zip archives", zip)
}

// latest returns the live version, the open pre-download and the newest
// version that has full zip archives.
func latest(s *store.Store) (live, pre, zip string) {
	var liveSeen string
	for _, v := range s.Versions {
		if v.MainLastSeen != "" && v.MainLastSeen >= liveSeen {
			live, liveSeen = v.Version, v.MainLastSeen
		}
		if len(v.Full) > 0 {
			zip = v.Version
		}
	}
	today := s.LatestDate()
	for _, v := range s.Versions {
		if v.PreDownloadLastSeen == today && v.MainLastSeen < today {
			pre = v.Version
		}
	}
	return live, pre, zip
}

func writeVersions(b *strings.Builder, s *store.Store) {
	b.WriteString("\n## Versions\n\nNewest first. Expand a major version to see its releases and links. Dates are the days a\nversion was seen in the API: up to 5.5.0 from the upstream daily snapshots, and from\n7.1.0 on from this repository's daily check.\n")
	var majors []string
	byMajor := map[string][]*store.Version{}
	for _, v := range s.Versions {
		m := store.Major(v.Version)
		if _, ok := byMajor[m]; !ok {
			majors = append(majors, m)
		}
		byMajor[m] = append(byMajor[m], v)
	}
	slices.SortFunc(majors, func(a, c string) int { return store.CompareVersions(c, a) })
	live, pre, _ := latest(s)
	for _, m := range majors {
		vs := byMajor[m]
		lo, hi := vs[0].Version, vs[len(vs)-1].Version
		count := "1 version"
		if len(vs) > 1 {
			count = strconv.Itoa(len(vs)) + " versions"
		}
		rng := lo
		if lo != hi {
			rng = lo + " to " + hi
		}
		fmt.Fprintf(b, "\n<details>\n<summary><b>Version %s.x</b>: %s (%s)</summary>\n", m, rng, count)
		if len(vs) > 1 {
			b.WriteString("\nJump to:")
			for i := len(vs) - 1; i >= 0; i-- {
				sep := " ·"
				if i == len(vs)-1 {
					sep = ""
				}
				fmt.Fprintf(b, "%s [%s](#%s)", sep, vs[i].Version, anchor(versionTitle(vs[i], live, pre)))
			}
			b.WriteString("\n")
		}
		for i := len(vs) - 1; i >= 0; i-- {
			writeVersion(b, vs[i], live, pre)
		}
		b.WriteString("\n</details>\n")
	}
}

func versionTitle(v *store.Version, live, pre string) string {
	switch v.Version {
	case live:
		return v.Version + " (live)"
	case pre:
		return v.Version + " (pre-download)"
	}
	return v.Version
}

// anchor is the id GitHub gives a heading: lower case, spaces to hyphens,
// other punctuation dropped.
func anchor(title string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '-', r == '_':
			b.WriteRune(r)
		case r == ' ':
			b.WriteRune('-')
		}
	}
	return b.String()
}

func writeVersion(b *strings.Builder, v *store.Version, live, pre string) {
	fmt.Fprintf(b, "\n### %s\n\n", versionTitle(v, live, pre))
	if d := dates(v, v.Version == live); d != "" {
		b.WriteString(d + "\n\n")
	}
	if !v.HasZip() && v.Sophon == nil {
		b.WriteString("No download data recorded.\n")
		return
	}
	if len(v.Full) > 0 {
		b.WriteString("#### Full install\n\n")
		writePackages(b, v.Full)
		for _, p := range v.Full {
			if p.Kind == store.KindGame && !strings.Contains(path.Base(p.URL), "_"+v.Version+".") {
				fmt.Fprintf(b, "\nThe game archive `%s` is named for an earlier release: upstream lists it as the %s game data.\n", path.Base(p.URL), v.Version)
				break
			}
		}
		if v.ResListURL != "" {
			fmt.Fprintf(b, "\nScattered files: `%s` (file list in `pkg_version`).\n", v.ResListURL)
		}
	}
	for _, p := range v.Patches {
		fmt.Fprintf(b, "\n#### Update from %s\n\n", p.From)
		writePackages(b, p.Packages)
	}
	if v.Sophon != nil {
		if v.HasZip() {
			b.WriteString("\n")
		}
		writeSophon(b, v)
	}
}

func dates(v *store.Version, isLive bool) string {
	var parts []string
	switch {
	case v.MainFirstSeen != "" && isLive:
		parts = append(parts, "live version since "+v.MainFirstSeen)
	case v.MainFirstSeen != "":
		parts = append(parts, "live version from "+v.MainFirstSeen+" to "+v.MainLastSeen)
	}
	if v.PreDownloadFirstSeen != "" {
		parts = append(parts, "pre-download from "+v.PreDownloadFirstSeen)
	}
	if len(parts) == 0 {
		return ""
	}
	return "Seen in the API as the " + strings.Join(parts, ", ") + " (UTC)."
}

func writePackages(b *strings.Builder, ps []*store.Package) {
	parts := 0
	for _, p := range ps {
		if p.Kind == store.KindGame && p.Part > 0 {
			parts++
		}
	}
	b.WriteString("| File | Content | Size | MD5 |\n| --- | --- | --- | --- |\n")
	for _, p := range ps {
		name := path.Base(p.URL)
		link := fmt.Sprintf("[%s](%s)", name, escapeURL(p.URL))
		if p.HTTPStatus != 0 && p.HTTPStatus != 200 {
			link += fmt.Sprintf(" (dead link, HTTP %d)", p.HTTPStatus)
		}
		md5 := ""
		if p.MD5 != "" {
			md5 = "`" + p.MD5 + "`"
		}
		if p.MD5Size != 0 {
			md5 += fmt.Sprintf(" (published for %d bytes; the served file is %d bytes, so it may not match)", p.MD5Size, p.Size)
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n", link, content(p, parts), size(p.Size), md5)
	}
}

func content(p *store.Package, parts int) string {
	switch p.Kind {
	case store.KindLauncher:
		return "Launcher"
	case store.KindAudio:
		return voice(p.Language)
	}
	if p.Part > 0 {
		return fmt.Sprintf("Game, part %d of %d", p.Part, parts)
	}
	if parts > 0 {
		return "Game, single archive"
	}
	return "Game"
}

func voice(lang string) string {
	if n, ok := langName[lang]; ok {
		return n + " voice-over"
	}
	return lang + " voice-over"
}

func category(field string) string {
	if field == store.KindGame {
		return "Game"
	}
	return voice(field)
}

func writeSophon(b *strings.Builder, v *store.Version) {
	sb := v.Sophon
	fmt.Fprintf(b, "#### Sophon full build `%s`\n\n", sb.BuildID)
	b.WriteString("| Content | Download | Installed | Files | Chunks | Manifest |\n| --- | --- | --- | --- | --- | --- |\n")
	prefixes := map[string]bool{}
	for _, m := range sb.Manifests {
		st := m.Stats
		if st == nil {
			st = &store.Stats{}
		}
		prefixes[m.DataURLPrefix+"\x00"+m.DataURLSuffix] = true
		fmt.Fprintf(b, "| %s | %s | %s | %d | %d | [manifest](%s) |\n", category(m.MatchingField),
			size(st.CompressedSize), size(st.UncompressedSize), st.FileCount, st.ChunkCount,
			escapeURL(m.ManifestURL()))
	}
	if len(prefixes) == 1 && sb.Manifests[0].DataURLSuffix == "" {
		fmt.Fprintf(b, "\nChunk URL prefix: `%s`\n", sb.Manifests[0].DataURLPrefix)
	} else {
		b.WriteString("\nChunk URL prefixes:\n\n")
		for _, m := range sb.Manifests {
			fmt.Fprintf(b, "- %s: `%s` (suffix `%s`)\n", category(m.MatchingField), m.DataURLPrefix, m.DataURLSuffix)
		}
	}
	if v.SophonPatch == nil {
		if !v.HasZip() {
			b.WriteString("\nOnly the full build is recorded: the Sophon API serves update data for the current version only.\n")
		}
		return
	}
	sp := v.SophonPatch
	fmt.Fprintf(b, "\n#### Sophon update `%s` (from %s)\n\n", sp.PatchID, strings.Join(sp.DiffTags, ", "))
	b.WriteString("| Content |")
	for _, from := range sp.DiffTags {
		fmt.Fprintf(b, " From %s |", from)
	}
	b.WriteString(" Manifest | Diff URL prefix |\n|" + strings.Repeat(" --- |", len(sp.DiffTags)+3) + "\n")
	for _, m := range sp.Manifests {
		fmt.Fprintf(b, "| %s |", category(m.MatchingField))
		for _, from := range sp.DiffTags {
			if st, ok := m.StatsByFrom[from]; ok {
				fmt.Fprintf(b, " %s |", size(st.CompressedSize))
			} else {
				b.WriteString(" - |")
			}
		}
		fmt.Fprintf(b, " [manifest](%s) | `%s%s` |\n", escapeURL(m.ManifestURL()), m.DataURLPrefix, m.DataURLSuffix)
	}
	b.WriteString("\nSizes are the download size of the diff data for each starting version.\n")
}

// escapeURL keeps parentheses in file names such as Audio_English(US) from
// ending a Markdown link early.
func escapeURL(u string) string {
	return strings.NewReplacer("(", "%28", ")", "%29", " ", "%20").Replace(u)
}

func size(n int64) string {
	const (
		mib = 1 << 20
		gib = 1 << 30
	)
	switch {
	case n <= 0:
		return "unknown"
	case n >= gib:
		return fmt.Sprintf("%.2f GiB", float64(n)/gib)
	case n >= mib:
		return fmt.Sprintf("%.2f MiB", float64(n)/mib)
	default:
		return fmt.Sprintf("%.2f KiB", float64(n)/1024)
	}
}

func writeFooter(b *strings.Builder, s *store.Store) {
	fmt.Fprintf(b, `
## License

The code (the Go fetcher and generator) is licensed under the [Apache License 2.0](LICENSE).
The data in [data/](data/) and the content of this README are licensed under
[CC BY-NC-SA 4.0](LICENSE-DATA), the license of the upstream lists they build on.

---

Last checked: %s (UTC). This file is generated from
[data/versions.json](data/versions.json) by `+"`go run ./cmd/fetch`"+`; do not edit it by hand.
`, s.LatestDate())
}
