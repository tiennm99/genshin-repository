package hoyo

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/tiennm99/genshin-repository/internal/store"
)

// Options controls one Sync run.
type Options struct {
	// Date is the observation date (YYYY-MM-DD, UTC) recorded as seen.
	Date string
	// Backfill probes getBuild for every X.Y.0 tag below the live version,
	// and every recorded version, that has no Sophon build yet. The Sophon
	// API also serves some zip-era versions.
	Backfill bool
	// Logf receives progress lines; nil discards them.
	Logf func(format string, args ...any)
}

// Sync queries every endpoint and merges the answers into s. It returns an
// error, leaving s partly updated, when any required call fails; callers
// should then not save s.
func Sync(ctx context.Context, c *Client, s *store.Store, o Options) error {
	logf := o.Logf
	if logf == nil {
		logf = func(string, ...any) {}
	}

	pk, err := c.GetGamePackages(ctx)
	if err != nil {
		return err
	}
	// The zip API's "main" stays frozen at 5.5.0, so it only contributes
	// archives, never a live-version date.
	if v := ApplyBranch(s, pk.Main, store.SourceZipAPI); v != "" {
		logf("zip main: %s", v)
	}
	if v := ApplyBranch(s, pk.PreDownload, store.SourceZipAPI); v != "" {
		s.Get(v).SeenPreDownload(o.Date)
		logf("zip pre-download: %s", v)
	}

	br, err := c.GetGameBranches(ctx)
	if err != nil {
		return err
	}
	if br.Main == nil || br.Main.Tag == "" {
		return fmt.Errorf("getGameBranches: no main branch tag")
	}
	if err := syncBranch(ctx, c, s, br.Main, false, o.Date, logf); err != nil {
		return err
	}
	if br.PreDownload != nil && br.PreDownload.Tag != "" {
		if err := syncBranch(ctx, c, s, br.PreDownload, true, o.Date, logf); err != nil {
			return err
		}
	}

	tags := slices.Clone(br.Main.DiffTags)
	if o.Backfill {
		tags = append(tags, candidateTags(br.Main.Tag, s)...)
	}
	for _, tag := range tags {
		if err := backfill(ctx, c, s, br.Main, tag, logf); err != nil {
			return err
		}
	}
	return nil
}

func syncBranch(ctx context.Context, c *Client, s *store.Store, b *Branch, pre bool, date string, logf func(string, ...any)) error {
	build, err := c.GetBuild(ctx, b, b.Tag)
	if err != nil {
		return err
	}
	v := s.Get(b.Tag)
	v.AddSource(store.SourceSophonAPI)
	v.SetSophon(build.Build)
	if pre {
		v.SeenPreDownload(date)
	} else {
		v.SeenMain(date)
	}
	logf("sophon %s: %s build %s", b.Branch, b.Tag, build.Build.BuildID)

	patch, err := c.GetPatchBuild(ctx, b)
	if IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if store.NormalizeVersion(patch.Tag) != v.Version {
		logf("sophon %s: patch tag %s != %s, ignored", b.Branch, patch.Tag, b.Tag)
		return nil
	}
	for _, t := range b.DiffTags {
		if !slices.Contains(patch.Patch.DiffTags, t) {
			patch.Patch.DiffTags = append(patch.Patch.DiffTags, t)
		}
	}
	v.SetSophonPatch(patch.Patch)
	logf("sophon %s: %s patch %s from %v", b.Branch, b.Tag, patch.Patch.PatchID, patch.Patch.DiffTags)
	return nil
}

// backfill fetches the full build of an older tag through the main branch.
func backfill(ctx context.Context, c *Client, s *store.Store, b *Branch, tag string, logf func(string, ...any)) error {
	tag = store.NormalizeVersion(tag)
	if v := s.Find(tag); v != nil && v.Sophon != nil {
		return nil
	}
	build, err := c.GetBuild(ctx, b, tag)
	if IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if store.NormalizeVersion(build.Tag) != tag {
		logf("sophon backfill %s: answered tag %s, ignored", tag, build.Tag)
		return nil
	}
	v := s.Get(tag)
	v.AddSource(store.SourceSophonAPI)
	v.SetSophon(build.Build)
	logf("sophon backfill %s: build %s", tag, build.Build.BuildID)
	return nil
}

// candidateTags lists X.Y.0 from 1.0.0 and every version recorded in s,
// ascending and below live.
func candidateTags(live string, s *store.Store) []string {
	maxMajor, _ := strconv.Atoi(store.Major(live))
	var out []string
	for m := 1; m <= maxMajor; m++ {
		for n := range 10 {
			out = append(out, fmt.Sprintf("%d.%d.0", m, n))
		}
	}
	for _, v := range s.Versions {
		out = append(out, v.Version)
	}
	out = slices.DeleteFunc(out, func(t string) bool { return store.CompareVersions(t, live) >= 0 })
	slices.SortFunc(out, store.CompareVersions)
	return slices.Compact(out)
}
