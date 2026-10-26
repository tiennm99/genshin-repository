// Package upstream imports the history kept by the projects this repository
// continues: the "Version X.md" link lists of JRSKelvin/GenshinRepository and
// the daily API snapshots of JRSKelvin/Genshin-Repository-List.
package upstream

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tiennm99/genshin-repository/internal/hoyo"
	"github.com/tiennm99/genshin-repository/internal/store"
)

// Pinned upstream commits; re-running the import gives the same result.
const (
	RepoMarkdown  = "JRSKelvin/GenshinRepository"
	CommitMD      = "871fd3c4417789a0c5d5ad8bf094396bdd4cb89a"
	RepoSnapshots = "JRSKelvin/Genshin-Repository-List"
	CommitSnaps   = "01b151331e9ba022852d713a83da8ed665cda3c6"
)

// ZipAPIFrozenFrom is the 5.6 release date. From then on the zip API kept
// answering 5.5.0, so snapshots from that date say nothing about the live
// version.
const ZipAPIFrozenFrom = "2025-05-07"

// TarballURL is the GitHub source tarball of repo at commit.
func TarballURL(repo, commit string) string {
	return "https://codeload.github.com/" + repo + "/tar.gz/" + commit
}

// Download fetches a .tar.gz and returns its regular files keyed by path
// without the top-level directory.
func Download(ctx context.Context, hc *http.Client, u string) (map[string][]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: HTTP %d", u, resp.StatusCode)
	}
	return ReadTarGz(resp.Body)
}

// ReadTarGz reads a gzip-compressed tar stream into memory.
func ReadTarGz(r io.Reader) (map[string][]byte, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	tr := tar.NewReader(gz)
	files := map[string][]byte{}
	for {
		h, err := tr.Next()
		if errors.Is(err, io.EOF) {
			return files, nil
		}
		if err != nil {
			return nil, err
		}
		if h.Typeflag != tar.TypeReg {
			continue
		}
		_, name, ok := strings.Cut(h.Name, "/")
		if !ok {
			continue
		}
		b, err := io.ReadAll(io.LimitReader(tr, 64<<20))
		if err != nil {
			return nil, err
		}
		files[name] = b
	}
}

// ImportMarkdownRepo applies every "Version X.md" file.
func ImportMarkdownRepo(s *store.Store, files map[string][]byte) int {
	n := 0
	for _, name := range sortedKeys(files) {
		if strings.HasPrefix(name, "Version ") && strings.HasSuffix(name, ".md") && !strings.Contains(name, "/") {
			ParseMarkdown(s, name, files[name])
			n++
		}
	}
	return n
}

// ImportSnapshotRepo applies every fetch/original/result-*.json snapshot in
// date order and returns the count per format.
func ImportSnapshotRepo(s *store.Store, files map[string][]byte) (map[string]int, error) {
	counts := map[string]int{}
	for _, name := range sortedKeys(files) {
		base := path.Base(name)
		if !strings.HasPrefix(name, "fetch/original/") || !strings.HasPrefix(base, "result-") || !strings.HasSuffix(base, ".json") {
			continue
		}
		date := strings.TrimPrefix(base, "result-")
		if len(date) < 10 {
			continue
		}
		format, err := ApplySnapshot(s, date[:10], files[name])
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		counts[format]++
	}
	return counts, nil
}

func sortedKeys(m map[string][]byte) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ---------- Version X.md ----------

var (
	linkRe     = regexp.MustCompile(`^\[[^\]]*\]\((https?://\S+)\)$`)
	reFullGame = regexp.MustCompile(`/GenshinImpact_([\d.]+)\.zip(?:\.(\d+))?$`)
	reFullAud  = regexp.MustCompile(`/Audio_([A-Za-z()]+)_([\d.]+)\.zip$`)
	rePatch    = regexp.MustCompile(`/(?:game_|audio_)?(?:([a-z]{2}-[a-z]{2})_)?([\d.]+)_([\d.]+)_h?diff_[^/]+\.zip$`)
	reLauncher = regexp.MustCompile(`/update_\d+[^/]*\.zip$`)
	langFile   = map[string]string{"Chinese": "zh-cn", "English(US)": "en-us", "Japanese": "ja-jp", "Korean": "ko-kr"}
)

// ParseMarkdown applies one "Version X.md" file. The version and language
// of each link come from the URL file name, not the label: upstream labels
// are sometimes wrong (1.5.0 swaps CN and KO, 1.6.1 relists 1.6.0 files).
// A file is filed under the version its URL names.
func ParseMarkdown(s *store.Store, name string, content []byte) {
	fileVer := store.NormalizeVersion(strings.TrimSuffix(strings.TrimPrefix(name, "Version "), ".md"))
	sc := bufio.NewScanner(bytes.NewReader(content))
	for sc.Scan() {
		m := linkRe.FindStringSubmatch(strings.TrimSpace(sc.Text()))
		if m == nil {
			continue
		}
		u := m[1]
		switch {
		case reLauncher.MatchString(u):
			add(s, fileVer, "", &store.Package{Kind: store.KindLauncher, URL: u})
		case reFullGame.MatchString(u):
			g := reFullGame.FindStringSubmatch(u)
			part, _ := strconv.Atoi(g[2])
			add(s, g[1], "", &store.Package{Kind: store.KindGame, Part: part, URL: u})
		case reFullAud.MatchString(u):
			g := reFullAud.FindStringSubmatch(u)
			add(s, g[2], "", &store.Package{Kind: store.KindAudio, Language: langFile[g[1]], URL: u})
		case rePatch.MatchString(u):
			g := rePatch.FindStringSubmatch(u)
			p := &store.Package{Kind: store.KindGame, URL: u}
			if g[1] != "" {
				p.Kind, p.Language = store.KindAudio, g[1]
			}
			add(s, g[3], g[2], p)
		}
	}
}

func add(s *store.Store, ver, from string, p *store.Package) {
	v := s.Get(ver)
	v.AddSource(store.SourceMD)
	v.AddPackage(from, p)
}

// ---------- snapshots ----------

type legacyPkg struct {
	Version         string `json:"version"`
	Path            string `json:"path"`
	Size            string `json:"size"`
	PackageSize     string `json:"package_size"`
	MD5             string `json:"md5"`
	Language        string `json:"language"`
	DecompressedURL string `json:"decompressed_path"`
	Segments        []struct {
		Path        string `json:"path"`
		MD5         string `json:"md5"`
		PackageSize string `json:"package_size"`
	} `json:"segments"`
	VoicePacks []legacyPkg `json:"voice_packs"`
}

type legacyGame struct {
	Latest *legacyPkg  `json:"latest"`
	Diffs  []legacyPkg `json:"diffs"`
}

// ApplySnapshot applies one Genshin-Repository-List snapshot taken on date
// and returns its format: "legacy" (sdk-os-static) or "hyp".
func ApplySnapshot(s *store.Store, date string, content []byte) (string, error) {
	var probe struct {
		Data map[string]json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(content, &probe); err != nil {
		return "", err
	}
	if raw, ok := probe.Data["game_packages"]; ok {
		var gps []hoyo.GamePackages
		if err := json.Unmarshal(raw, &gps); err != nil {
			return "", err
		}
		for _, gp := range gps {
			if gp.Game.Biz != "" && gp.Game.Biz != hoyo.GameBiz {
				continue
			}
			if v := hoyo.ApplyBranch(s, gp.Main, store.SourceSnapshotHyp); v != "" && date < ZipAPIFrozenFrom {
				s.Get(v).SeenMain(date)
			}
			if v := hoyo.ApplyBranch(s, gp.PreDownload, store.SourceSnapshotHyp); v != "" {
				s.Get(v).SeenPreDownload(date)
			}
		}
		return "hyp", nil
	}
	if raw, ok := probe.Data["game"]; ok {
		var game, pre *legacyGame
		if err := json.Unmarshal(raw, &game); err != nil {
			return "", err
		}
		if rawPre, ok := probe.Data["pre_download_game"]; ok {
			if err := json.Unmarshal(rawPre, &pre); err != nil {
				return "", err
			}
		}
		if v := applyLegacy(s, game); v != nil {
			v.SeenMain(date)
		}
		if v := applyLegacy(s, pre); v != nil {
			v.SeenPreDownload(date)
		}
		return "legacy", nil
	}
	return "unknown", nil
}

func applyLegacy(s *store.Store, g *legacyGame) *store.Version {
	if g == nil || g.Latest == nil || g.Latest.Version == "" {
		return nil
	}
	l := g.Latest
	v := s.Get(l.Version)
	v.AddSource(store.SourceSnapshotLegacy)
	if l.DecompressedURL != "" {
		v.ResListURL = l.DecompressedURL
	}
	// Legacy sizes: package_size = download bytes, size = unpacked bytes.
	if l.Path != "" {
		v.AddPackage("", &store.Package{Kind: store.KindGame, URL: l.Path, MD5: l.MD5,
			Size: hoyo.ParseSize(l.PackageSize), DecompressedSize: hoyo.ParseSize(l.Size)})
	}
	for i, seg := range l.Segments {
		v.AddPackage("", &store.Package{Kind: store.KindGame, Part: i + 1, URL: seg.Path, MD5: seg.MD5,
			Size: hoyo.ParseSize(seg.PackageSize)})
	}
	for _, vp := range l.VoicePacks {
		v.AddPackage("", legacyAudio(vp))
	}
	for _, d := range g.Diffs {
		v.AddPackage(d.Version, &store.Package{Kind: store.KindGame, URL: d.Path, MD5: d.MD5,
			Size: hoyo.ParseSize(d.PackageSize), DecompressedSize: hoyo.ParseSize(d.Size)})
		for _, vp := range d.VoicePacks {
			v.AddPackage(d.Version, legacyAudio(vp))
		}
	}
	return v
}

func legacyAudio(vp legacyPkg) *store.Package {
	return &store.Package{Kind: store.KindAudio, Language: vp.Language, URL: vp.Path, MD5: vp.MD5,
		Size: hoyo.ParseSize(vp.PackageSize), DecompressedSize: hoyo.ParseSize(vp.Size)}
}

// ---------- link check ----------

// CheckLinks sends a HEAD request to every archive URL and records 200, 404
// and 410 answers. Other answers (the CDN sometimes answers 400 for archives
// over 50 GB) leave the previous status alone and count as unknown. For a
// live link, Content-Length replaces the recorded size: legacy snapshot sizes
// are a few MB off for some archives.
func CheckLinks(ctx context.Context, hc *http.Client, s *store.Store, workers int) (live, dead, unknown int) {
	var pkgs []*store.Package
	for _, v := range s.Versions {
		pkgs = append(pkgs, v.Full...)
		for _, p := range v.Patches {
			pkgs = append(pkgs, p.Packages...)
		}
	}
	jobs := make(chan *store.Package)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for p := range jobs {
				status, length := head(ctx, hc, p.URL)
				mu.Lock()
				switch status {
				case http.StatusOK:
					live++
					p.HTTPStatus = status
					if length > 0 {
						p.Size = length
					}
				case http.StatusNotFound, http.StatusGone:
					dead++
					p.HTTPStatus = status
				default:
					unknown++
				}
				mu.Unlock()
			}
		})
	}
	for _, p := range pkgs {
		jobs <- p
	}
	close(jobs)
	wg.Wait()
	return live, dead, unknown
}

func head(ctx context.Context, hc *http.Client, u string) (status int, length int64) {
	for attempt := range 4 {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second)
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)
		if err != nil {
			return 0, 0
		}
		resp, err := hc.Do(req)
		if err != nil {
			continue
		}
		_ = resp.Body.Close()
		status, length = resp.StatusCode, resp.ContentLength
		// Only 200 and "gone" answers are final; retry anything else.
		if status == http.StatusOK || status == http.StatusNotFound || status == http.StatusGone {
			return status, length
		}
	}
	return status, length
}
