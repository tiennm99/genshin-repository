// Package store holds the persisted catalogue of Genshin Impact versions and
// the rules for merging new observations into it.
//
// The store only grows: a version, package or patch that once appeared is
// never removed, so links that drop out of the HoYoverse API stay listed.
package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Package kinds.
const (
	KindGame     = "game"
	KindAudio    = "audio"
	KindLauncher = "launcher"
)

// Version sources.
const (
	SourceMD             = "upstream-md"       // JRSKelvin/GenshinRepository "Version X.md"
	SourceSnapshotLegacy = "upstream-legacy"   // Genshin-Repository-List, sdk-os-static format
	SourceSnapshotHyp    = "upstream-hyp"      // Genshin-Repository-List, getGamePackages format
	SourceZipAPI         = "api-game-packages" // live getGamePackages
	SourceSophonAPI      = "api-sophon"        // live getGameBranches + getBuild/getPatchBuild
)

// Store is the whole catalogue, written to data/versions.json.
type Store struct {
	Versions []*Version `json:"versions"` // ascending by version
}

// Version is everything known about one game version.
type Version struct {
	Version              string       `json:"version"`
	Sources              []string     `json:"sources"`
	MainFirstSeen        string       `json:"main_first_seen,omitempty"` // YYYY-MM-DD, UTC
	MainLastSeen         string       `json:"main_last_seen,omitempty"`
	PreDownloadFirstSeen string       `json:"pre_download_first_seen,omitempty"`
	PreDownloadLastSeen  string       `json:"pre_download_last_seen,omitempty"`
	ResListURL           string       `json:"res_list_url,omitempty"`
	Full                 []*Package   `json:"full,omitempty"`
	Patches              []*Patch     `json:"patches,omitempty"` // newest "from" first
	Sophon               *SophonBuild `json:"sophon,omitempty"`
	SophonPatch          *SophonPatch `json:"sophon_patch,omitempty"`
}

// Package is one downloadable archive.
type Package struct {
	Kind             string `json:"kind"`
	Language         string `json:"language,omitempty"` // audio only: zh-cn, en-us, ja-jp, ko-kr
	Part             int    `json:"part,omitempty"`     // 1-based part of a split game archive; 0 = single archive
	URL              string `json:"url"`
	MD5              string `json:"md5,omitempty"`
	Size             int64  `json:"size,omitempty"`              // download bytes
	DecompressedSize int64  `json:"decompressed_size,omitempty"` // unpacked bytes
	HTTPStatus       int    `json:"http_status,omitempty"`       // last link check, 0 = never checked
}

// Patch holds the update archives from one older version.
type Patch struct {
	From     string     `json:"from"`
	Packages []*Package `json:"packages"`
}

// Stats are Sophon size and count totals for one category.
type Stats struct {
	CompressedSize   int64 `json:"compressed_size"`
	UncompressedSize int64 `json:"uncompressed_size"`
	FileCount        int64 `json:"file_count"`
	ChunkCount       int64 `json:"chunk_count"`
}

// SophonBuild is a full Sophon build (getBuild).
type SophonBuild struct {
	BuildID   string            `json:"build_id"`
	Manifests []*SophonManifest `json:"manifests"`
}

// SophonManifest is one category (game data or one voice-over language).
type SophonManifest struct {
	CategoryID        string `json:"category_id"`
	MatchingField     string `json:"matching_field"`
	ManifestID        string `json:"manifest_id"`
	ManifestChecksum  string `json:"manifest_checksum"`
	ManifestURLPrefix string `json:"manifest_url_prefix"`
	ManifestURLSuffix string `json:"manifest_url_suffix,omitempty"`
	// DataURLPrefix is chunk_download for a build and diff_download for a patch.
	DataURLPrefix string `json:"data_url_prefix"`
	DataURLSuffix string `json:"data_url_suffix,omitempty"`
	// Stats is set for builds; StatsByFrom (keyed by from-version) for patches.
	Stats       *Stats            `json:"stats,omitempty"`
	StatsByFrom map[string]*Stats `json:"stats_by_from,omitempty"`
}

// ManifestURL is the download URL of the manifest file.
func (m *SophonManifest) ManifestURL() string {
	return m.ManifestURLPrefix + "/" + m.ManifestID + m.ManifestURLSuffix
}

// SophonPatch is a Sophon patch build (getPatchBuild).
type SophonPatch struct {
	PatchID   string            `json:"patch_id"`
	BuildID   string            `json:"build_id"`
	DiffTags  []string          `json:"diff_tags,omitempty"`
	Manifests []*SophonManifest `json:"manifests"`
}

// Load reads a store; a missing file yields an empty store.
func Load(path string) (*Store, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return &Store{}, nil
	}
	if err != nil {
		return nil, err
	}
	var s Store
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return &s, nil
}

// Marshal returns the canonical JSON form: sorted, two-space indent,
// trailing newline, so unchanged data gives byte-identical output.
func (s *Store) Marshal() ([]byte, error) {
	s.Normalize()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Save writes the canonical JSON form to path.
func (s *Store) Save(path string) error {
	b, err := s.Marshal()
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// Get returns the record for version v, creating it when absent.
func (s *Store) Get(v string) *Version {
	v = NormalizeVersion(v)
	for _, x := range s.Versions {
		if x.Version == v {
			return x
		}
	}
	x := &Version{Version: v}
	s.Versions = append(s.Versions, x)
	return x
}

// Find returns the record for version v, or nil.
func (s *Store) Find(v string) *Version {
	v = NormalizeVersion(v)
	for _, x := range s.Versions {
		if x.Version == v {
			return x
		}
	}
	return nil
}

// Normalize sorts everything into canonical order.
func (s *Store) Normalize() {
	slices.SortFunc(s.Versions, func(a, b *Version) int { return CompareVersions(a.Version, b.Version) })
	for _, v := range s.Versions {
		slices.Sort(v.Sources)
		sortPackages(v.Full)
		slices.SortFunc(v.Patches, func(a, b *Patch) int { return CompareVersions(b.From, a.From) })
		for _, p := range v.Patches {
			sortPackages(p.Packages)
		}
		if v.Sophon != nil {
			sortManifests(v.Sophon.Manifests)
		}
		if v.SophonPatch != nil {
			sortManifests(v.SophonPatch.Manifests)
			slices.SortFunc(v.SophonPatch.DiffTags, func(a, b string) int { return CompareVersions(b, a) })
		}
	}
}

// AddSource records that a source contributed to v.
func (v *Version) AddSource(src string) {
	if !slices.Contains(v.Sources, src) {
		v.Sources = append(v.Sources, src)
	}
}

// SeenMain records that v was the live version on date (YYYY-MM-DD).
func (v *Version) SeenMain(date string) { seen(&v.MainFirstSeen, &v.MainLastSeen, date) }

// SeenPreDownload records that v was offered as a pre-download on date.
func (v *Version) SeenPreDownload(date string) {
	seen(&v.PreDownloadFirstSeen, &v.PreDownloadLastSeen, date)
}

func seen(first, last *string, date string) {
	if *first == "" || date < *first {
		*first = date
	}
	if date > *last {
		*last = date
	}
}

// AddPackage merges p into the full install (from == "") or into the patch
// from version `from`. Packages are matched by URL. Known values are kept
// unless the new observation carries a value of its own.
func (v *Version) AddPackage(from string, p *Package) {
	p.MD5 = strings.ToLower(p.MD5)
	list := &v.Full
	if from != "" {
		list = &v.patch(from).Packages
	}
	for _, ex := range *list {
		if ex.URL != p.URL {
			continue
		}
		mergePackage(ex, p)
		return
	}
	cp := *p
	*list = append(*list, &cp)
}

func (v *Version) patch(from string) *Patch {
	from = NormalizeVersion(from)
	for _, p := range v.Patches {
		if p.From == from {
			return p
		}
	}
	p := &Patch{From: from}
	v.Patches = append(v.Patches, p)
	return p
}

func mergePackage(ex, p *Package) {
	if p.Kind != "" {
		ex.Kind = p.Kind
	}
	if p.Language != "" {
		ex.Language = p.Language
	}
	if p.Part != 0 {
		ex.Part = p.Part
	}
	if p.MD5 != "" {
		ex.MD5 = p.MD5
	}
	// A size read from a live HEAD answer (status 200) is the real
	// Content-Length; snapshot sizes are sometimes a few MB off, so they do
	// not replace it.
	if p.Size != 0 && (ex.HTTPStatus != 200 || p.HTTPStatus == 200) {
		ex.Size = p.Size
	}
	if p.DecompressedSize != 0 {
		ex.DecompressedSize = p.DecompressedSize
	}
	if p.HTTPStatus != 0 {
		ex.HTTPStatus = p.HTTPStatus
	}
}

// SetSophon replaces the Sophon build when b is non-empty.
func (v *Version) SetSophon(b *SophonBuild) {
	if b != nil && len(b.Manifests) > 0 {
		v.Sophon = b
	}
}

// SetSophonPatch replaces the Sophon patch when p is non-empty.
func (v *Version) SetSophonPatch(p *SophonPatch) {
	if p != nil && len(p.Manifests) > 0 {
		v.SophonPatch = p
	}
}

// HasZip reports whether v has any full or patch archive.
func (v *Version) HasZip() bool { return len(v.Full) > 0 || len(v.Patches) > 0 }

var langOrder = map[string]int{"zh-cn": 1, "en-us": 2, "ja-jp": 3, "ko-kr": 4}

// LanguageRank orders the voice-over languages the way the game lists them.
func LanguageRank(lang string) int {
	if r, ok := langOrder[lang]; ok {
		return r
	}
	return len(langOrder) + 1
}

var kindOrder = map[string]int{KindLauncher: 0, KindGame: 1, KindAudio: 2}

func sortPackages(ps []*Package) {
	slices.SortFunc(ps, func(a, b *Package) int {
		if c := kindOrder[a.Kind] - kindOrder[b.Kind]; c != 0 {
			return c
		}
		if c := a.Part - b.Part; c != 0 {
			return c
		}
		if c := LanguageRank(a.Language) - LanguageRank(b.Language); c != 0 {
			return c
		}
		if c := strings.Compare(a.Language, b.Language); c != 0 {
			return c
		}
		return strings.Compare(a.URL, b.URL)
	})
}

func sortManifests(ms []*SophonManifest) {
	slices.SortFunc(ms, func(a, b *SophonManifest) int {
		ra, rb := 0, 0
		if a.MatchingField != KindGame {
			ra = LanguageRank(a.MatchingField)
		}
		if b.MatchingField != KindGame {
			rb = LanguageRank(b.MatchingField)
		}
		if ra != rb {
			return ra - rb
		}
		return strings.Compare(a.MatchingField+a.CategoryID, b.MatchingField+b.CategoryID)
	})
}

// NormalizeVersion pads a version to three parts: "1.3" -> "1.3.0".
func NormalizeVersion(v string) string {
	parts := strings.Split(strings.TrimSpace(v), ".")
	for len(parts) < 3 {
		parts = append(parts, "0")
	}
	return strings.Join(parts, ".")
}

// CompareVersions compares dotted numeric versions.
func CompareVersions(a, b string) int {
	pa, pb := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < len(pa) && i < len(pb); i++ {
		x, _ := strconv.Atoi(pa[i])
		y, _ := strconv.Atoi(pb[i])
		if x != y {
			return x - y
		}
	}
	return len(pa) - len(pb)
}

// Major returns the major number of a version ("5.5.0" -> "5").
func Major(v string) string {
	m, _, _ := strings.Cut(v, ".")
	return m
}

// LatestDate is the newest date recorded anywhere in the store.
func (s *Store) LatestDate() string {
	var d string
	for _, v := range s.Versions {
		for _, x := range []string{v.MainLastSeen, v.PreDownloadLastSeen} {
			if x > d {
				d = x
			}
		}
	}
	return d
}
