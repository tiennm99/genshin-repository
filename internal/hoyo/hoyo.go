// Package hoyo calls the public HoYoPlay (hyp) and Sophon APIs for
// Genshin Impact (global) and converts the answers into store records.
package hoyo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tiennm99/genshin-repository/internal/store"
)

// Identifiers of Genshin Impact (global) in the HoYoPlay launcher API.
const (
	GameID     = "gopR6Cufr3"
	GameBiz    = "hk4e_global"
	LauncherID = "VYTpXlbWo8"
)

// Default API hosts.
const (
	DefaultHypBase    = "https://sg-hyp-api.hoyoverse.com/hyp/hyp-connect/api"
	DefaultSophonBase = "https://sg-public-api.hoyoverse.com/downloader/sophon_chunk/api"
)

// RetcodeNotFound is the Sophon answer for an unknown tag.
const RetcodeNotFound = -202

// APIError is a non-zero retcode in the response envelope.
type APIError struct {
	Endpoint string
	Retcode  int
	Message  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: retcode %d: %s", e.Endpoint, e.Retcode, e.Message)
}

// IsNotFound reports whether err is a Sophon "not found" answer.
func IsNotFound(err error) bool {
	var ae *APIError
	return errors.As(err, &ae) && ae.Retcode == RetcodeNotFound
}

// Client talks to the APIs. The zero value is not usable; call New.
type Client struct {
	HTTP       *http.Client
	HypBase    string
	SophonBase string
}

// New returns a client for the live APIs.
func New() *Client {
	return &Client{
		HTTP:       &http.Client{Timeout: 60 * time.Second},
		HypBase:    DefaultHypBase,
		SophonBase: DefaultSophonBase,
	}
}

type envelope struct {
	Retcode int             `json:"retcode"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (c *Client) call(ctx context.Context, method, endpoint string, q url.Values, out any) error {
	u := endpoint + "?" + q.Encode()
	var lastErr error
	for attempt := range 3 {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(attempt) * 2 * time.Second):
			}
		}
		body, retry, err := c.do(ctx, method, u)
		if err != nil {
			lastErr = err
			if retry {
				continue
			}
			return err
		}
		return decodeEnvelope(endpoint, body, out)
	}
	return lastErr
}

func (c *Client) do(ctx context.Context, method, u string) (body []byte, retry bool, err error) {
	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, true, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err = io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if err != nil {
		return nil, true, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode >= 500, fmt.Errorf("%s %s: HTTP %d", method, u, resp.StatusCode)
	}
	return body, false, nil
}

func decodeEnvelope(endpoint string, body []byte, out any) error {
	var env envelope
	if err := json.Unmarshal(body, &env); err != nil {
		return fmt.Errorf("%s: %w", endpoint, err)
	}
	if env.Retcode != 0 {
		return &APIError{Endpoint: endpoint, Retcode: env.Retcode, Message: env.Message}
	}
	if err := json.Unmarshal(env.Data, out); err != nil {
		return fmt.Errorf("%s data: %w", endpoint, err)
	}
	return nil
}

func gameQuery() url.Values {
	return url.Values{"launcher_id": {LauncherID}, "game_ids[]": {GameID}}
}

// ---------- getGamePackages (zip archives, frozen at 5.5.0) ----------

// WireFile is a zip archive in getGamePackages. Sizes are JSON strings.
type WireFile struct {
	Language         string `json:"language"`
	URL              string `json:"url"`
	MD5              string `json:"md5"`
	Size             string `json:"size"`
	DecompressedSize string `json:"decompressed_size"`
}

// WireBuild is a major build or a patch in getGamePackages.
type WireBuild struct {
	Version    string     `json:"version"`
	GamePkgs   []WireFile `json:"game_pkgs"`
	AudioPkgs  []WireFile `json:"audio_pkgs"`
	ResListURL string     `json:"res_list_url"`
}

// WireBranch is main or pre_download in getGamePackages.
type WireBranch struct {
	Major   *WireBuild  `json:"major"`
	Patches []WireBuild `json:"patches"`
}

type wireGame struct {
	ID  string `json:"id"`
	Biz string `json:"biz"`
}

// GamePackages is the Genshin entry of getGamePackages.
type GamePackages struct {
	Game        wireGame   `json:"game"`
	Main        WireBranch `json:"main"`
	PreDownload WireBranch `json:"pre_download"`
}

// GetGamePackages returns the zip-archive listing.
func (c *Client) GetGamePackages(ctx context.Context) (*GamePackages, error) {
	var data struct {
		GamePackages []GamePackages `json:"game_packages"`
	}
	if err := c.call(ctx, http.MethodGet, c.HypBase+"/getGamePackages", gameQuery(), &data); err != nil {
		return nil, err
	}
	for i := range data.GamePackages {
		if data.GamePackages[i].Game.Biz == GameBiz {
			return &data.GamePackages[i], nil
		}
	}
	return nil, fmt.Errorf("getGamePackages: no %s entry", GameBiz)
}

// ParseSize reads a size sent as a JSON string; empty means 0.
func ParseSize(s string) int64 {
	n, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return n
}

// ApplyBranch merges a getGamePackages branch into the store and returns
// the version it describes ("" when the branch is empty).
func ApplyBranch(s *store.Store, b WireBranch, src string) string {
	if b.Major == nil || b.Major.Version == "" {
		return ""
	}
	v := s.Get(b.Major.Version)
	v.AddSource(src)
	if b.Major.ResListURL != "" {
		v.ResListURL = b.Major.ResListURL
	}
	addFiles(v, "", b.Major)
	for i := range b.Patches {
		addFiles(v, b.Patches[i].Version, &b.Patches[i])
	}
	return v.Version
}

func addFiles(v *store.Version, from string, b *WireBuild) {
	split := len(b.GamePkgs) > 1
	for i, f := range b.GamePkgs {
		p := pkg(f, store.KindGame)
		if from == "" && split {
			p.Part = i + 1
		}
		v.AddPackage(from, p)
	}
	for _, f := range b.AudioPkgs {
		v.AddPackage(from, pkg(f, store.KindAudio))
	}
}

func pkg(f WireFile, kind string) *store.Package {
	return &store.Package{
		Kind:             kind,
		Language:         f.Language,
		URL:              f.URL,
		MD5:              f.MD5,
		Size:             ParseSize(f.Size),
		DecompressedSize: ParseSize(f.DecompressedSize),
	}
}

// ---------- getGameBranches (Sophon) ----------

// Branch is main or pre_download in getGameBranches.
type Branch struct {
	PackageID string   `json:"package_id"`
	Branch    string   `json:"branch"`
	Password  string   `json:"password"`
	Tag       string   `json:"tag"`
	DiffTags  []string `json:"diff_tags"`
}

// GameBranches is the Genshin entry of getGameBranches.
type GameBranches struct {
	Game        wireGame `json:"game"`
	Main        *Branch  `json:"main"`
	PreDownload *Branch  `json:"pre_download"`
}

// GetGameBranches returns the Sophon branches.
func (c *Client) GetGameBranches(ctx context.Context) (*GameBranches, error) {
	var data struct {
		GameBranches []GameBranches `json:"game_branches"`
	}
	if err := c.call(ctx, http.MethodGet, c.HypBase+"/getGameBranches", gameQuery(), &data); err != nil {
		return nil, err
	}
	for i := range data.GameBranches {
		if data.GameBranches[i].Game.Biz == GameBiz {
			return &data.GameBranches[i], nil
		}
	}
	return nil, fmt.Errorf("getGameBranches: no %s entry", GameBiz)
}

type wireDownload struct {
	URLPrefix string `json:"url_prefix"`
	URLSuffix string `json:"url_suffix"`
}

type wireStats struct {
	CompressedSize   string `json:"compressed_size"`
	UncompressedSize string `json:"uncompressed_size"`
	FileCount        string `json:"file_count"`
	ChunkCount       string `json:"chunk_count"`
}

func (w wireStats) stats() *store.Stats {
	return &store.Stats{
		CompressedSize:   ParseSize(w.CompressedSize),
		UncompressedSize: ParseSize(w.UncompressedSize),
		FileCount:        ParseSize(w.FileCount),
		ChunkCount:       ParseSize(w.ChunkCount),
	}
}

type wireManifest struct {
	CategoryID    string `json:"category_id"`
	MatchingField string `json:"matching_field"`
	Manifest      struct {
		ID       string `json:"id"`
		Checksum string `json:"checksum"`
	} `json:"manifest"`
	ManifestDownload wireDownload    `json:"manifest_download"`
	ChunkDownload    *wireDownload   `json:"chunk_download"`
	DiffDownload     *wireDownload   `json:"diff_download"`
	Stats            json.RawMessage `json:"stats"`
}

func (w wireManifest) base() *store.SophonManifest {
	m := &store.SophonManifest{
		CategoryID:        w.CategoryID,
		MatchingField:     w.MatchingField,
		ManifestID:        w.Manifest.ID,
		ManifestChecksum:  w.Manifest.Checksum,
		ManifestURLPrefix: w.ManifestDownload.URLPrefix,
		ManifestURLSuffix: w.ManifestDownload.URLSuffix,
	}
	for _, d := range []*wireDownload{w.ChunkDownload, w.DiffDownload} {
		if d != nil {
			m.DataURLPrefix, m.DataURLSuffix = d.URLPrefix, d.URLSuffix
		}
	}
	return m
}

// Build is the decoded getBuild answer.
type Build struct {
	Tag   string
	Build *store.SophonBuild
}

// Patch is the decoded getPatchBuild answer.
type Patch struct {
	Tag   string
	Patch *store.SophonPatch
}

func branchQuery(b *Branch, tag string) url.Values {
	q := url.Values{"branch": {b.Branch}, "package_id": {b.PackageID}, "password": {b.Password}}
	if tag != "" {
		q.Set("tag", tag)
	}
	return q
}

// GetBuild returns the full build of tag on branch b ("" = current tag).
func (c *Client) GetBuild(ctx context.Context, b *Branch, tag string) (*Build, error) {
	var raw json.RawMessage
	if err := c.call(ctx, http.MethodGet, c.SophonBase+"/getBuild", branchQuery(b, tag), &raw); err != nil {
		return nil, err
	}
	return DecodeBuild(raw)
}

// GetPatchBuild returns the current patch build of branch b. The endpoint
// only accepts POST, with the parameters in the query string.
func (c *Client) GetPatchBuild(ctx context.Context, b *Branch) (*Patch, error) {
	var raw json.RawMessage
	if err := c.call(ctx, http.MethodPost, c.SophonBase+"/getPatchBuild", branchQuery(b, b.Tag), &raw); err != nil {
		return nil, err
	}
	return DecodePatch(raw)
}

// DecodeBuild converts the data object of a getBuild answer.
func DecodeBuild(data []byte) (*Build, error) {
	var w struct {
		BuildID   string         `json:"build_id"`
		Tag       string         `json:"tag"`
		Manifests []wireManifest `json:"manifests"`
	}
	if err := json.Unmarshal(data, &w); err != nil {
		return nil, fmt.Errorf("getBuild: %w", err)
	}
	out := &Build{Tag: w.Tag, Build: &store.SophonBuild{BuildID: w.BuildID}}
	for _, wm := range w.Manifests {
		m := wm.base()
		var st wireStats
		if len(wm.Stats) > 0 {
			if err := json.Unmarshal(wm.Stats, &st); err != nil {
				return nil, fmt.Errorf("getBuild stats: %w", err)
			}
			m.Stats = st.stats()
		}
		out.Build.Manifests = append(out.Build.Manifests, m)
	}
	return out, nil
}

// DecodePatch converts the data object of a getPatchBuild answer.
func DecodePatch(data []byte) (*Patch, error) {
	var w struct {
		BuildID   string         `json:"build_id"`
		PatchID   string         `json:"patch_id"`
		Tag       string         `json:"tag"`
		Manifests []wireManifest `json:"manifests"`
	}
	if err := json.Unmarshal(data, &w); err != nil {
		return nil, fmt.Errorf("getPatchBuild: %w", err)
	}
	out := &Patch{Tag: w.Tag, Patch: &store.SophonPatch{PatchID: w.PatchID, BuildID: w.BuildID}}
	from := map[string]bool{}
	for _, wm := range w.Manifests {
		m := wm.base()
		var st map[string]wireStats
		if len(wm.Stats) > 0 {
			if err := json.Unmarshal(wm.Stats, &st); err != nil {
				return nil, fmt.Errorf("getPatchBuild stats: %w", err)
			}
		}
		if len(st) > 0 {
			m.StatsByFrom = map[string]*store.Stats{}
			for k, v := range st {
				m.StatsByFrom[k] = v.stats()
				from[k] = true
			}
		}
		out.Patch.Manifests = append(out.Patch.Manifests, m)
	}
	for k := range from {
		out.Patch.DiffTags = append(out.Patch.DiffTags, k)
	}
	return out, nil
}
