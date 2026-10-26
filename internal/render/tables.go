package render

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/tiennm99/genshin-repository/internal/store"
)

var langName = map[string]string{
	"zh-cn": "Chinese",
	"en-us": "English",
	"ja-jp": "Japanese",
	"ko-kr": "Korean",
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

// shield is a static shields.io badge URL. Its text comes from the store, so
// the Markdown stays deterministic.
func shield(label, msg, color string) string {
	esc := func(s string) string {
		s = strings.NewReplacer("-", "--", "_", "__").Replace(s)
		return strings.ReplaceAll(url.PathEscape(s), "+", "%2B")
	}
	return "https://img.shields.io/badge/" + esc(label) + "-" + esc(msg) + "-" + color
}

// writeIndex writes the one-row-per-version table; prefix is the page the
// version anchors live on ("" for the same page).
func writeIndex(b *strings.Builder, rows []row, prefix string) {
	b.WriteString("| Version | Live in API (UTC) | Game zip | Sophon | Notes |\n| --- | --- | --- | --- | --- |\n")
	for _, r := range rows {
		fmt.Fprintf(b, "| [%s](%s#%s) | %s | %s | %s | %s |\n", r.Version, prefix, anchor(r.Version), r.Dates, r.Zip, r.Sophon, r.Notes)
	}
}

// writeVersion writes the full detail section of one version.
func writeVersion(b *strings.Builder, v *store.Version, p *page) {
	fmt.Fprintf(b, "## %s\n\n%s\n\n", v.Version, status(v, p))
	if !v.HasZip() && v.Sophon == nil {
		b.WriteString("No download data recorded.\n")
		return
	}
	if len(v.Full) > 0 {
		b.WriteString("### Full install\n\n")
		writePackages(b, v.Full)
		for _, pk := range v.Full {
			if pk.Kind == store.KindGame && !strings.Contains(path.Base(pk.URL), "_"+v.Version+".") {
				fmt.Fprintf(b, "\nThe game archive `%s` is named for an earlier release: upstream lists it as the %s game data.\n", path.Base(pk.URL), v.Version)
				break
			}
		}
		if v.ResListURL != "" {
			fmt.Fprintf(b, "\nScattered files: `%s` (file list in `pkg_version`).\n", v.ResListURL)
		}
	}
	for _, pt := range v.Patches {
		fmt.Fprintf(b, "\n### Update from %s\n\n", pt.From)
		writePackages(b, pt.Packages)
	}
	if v.Sophon != nil {
		if v.HasZip() {
			b.WriteString("\n")
		}
		writeSophon(b, v)
	}
}

// status is the line under a version heading: what the daily API checks
// saw. The heading itself stays the bare version, so its anchor never moves.
func status(v *store.Version, p *page) string {
	var parts []string
	switch {
	case v == p.Live:
		parts = append(parts, "**Live version** in the API since "+v.MainFirstSeen)
	case v.MainFirstSeen != "":
		parts = append(parts, "**Live version** in the API from "+v.MainFirstSeen+" to "+v.MainLastSeen)
	}
	switch {
	case v == p.Pre:
		parts = append(parts, "**Pre-download** in the API since "+v.PreDownloadFirstSeen)
	case v.PreDownloadFirstSeen != "" && len(parts) == 0:
		parts = append(parts, "**Pre-download** in the API from "+v.PreDownloadFirstSeen+" to "+v.PreDownloadLastSeen)
	case v.PreDownloadFirstSeen != "":
		parts = append(parts, "pre-download from "+v.PreDownloadFirstSeen)
	}
	if len(parts) == 0 {
		return "No daily API check saw this version as live or pre-download."
	}
	return strings.Join(parts, "; ") + " (UTC)."
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
		if dead(p) {
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

// dead reports whether the last link check got an answer other than 200.
func dead(p *store.Package) bool { return p.HTTPStatus != 0 && p.HTTPStatus != 200 }

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
	fmt.Fprintf(b, "### Sophon full build `%s`\n\n", sb.BuildID)
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
	fmt.Fprintf(b, "\n### Sophon update `%s` (from %s)\n\n", sp.PatchID, strings.Join(sp.DiffTags, ", "))
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
