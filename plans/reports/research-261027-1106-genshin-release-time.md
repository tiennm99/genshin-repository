# Genshin Impact release time and the update cron

Researched 2026-10-27 11:06 (Asia/Saigon).

## Answer

Version updates go live after a maintenance window that starts at **06:00 UTC+8 on a Wednesday** and is announced to last about five hours, ending around **11:00 UTC+8**. In UTC that is 22:00 Tuesday to 03:00 Wednesday. One hour after the window ends is **04:00 UTC** (11:00 Asia/Saigon), so the daily update job now runs at `0 4 * * *`. The previous schedule, 01:30 UTC, fell inside the maintenance window.

## Evidence

- HoYoverse's 7.1 notice, as quoted by guides: maintenance from 2026-09-23 06:00 UTC+8, about 5 hours; the start is one global instant (22:00 UTC on 2026-09-22). The same 06:00-11:00 UTC+8 window held for 6.7 (2026-07-01) and for older versions (2.5, 2.7, 3.8, 5.0).
- Cadence: versions ship every 42 days on a Wednesday. 7.2 is projected for 2026-11-04 but was not confirmed when this was written.
- Our own data agrees: daily API snapshots taken at 17:10 UTC first saw 5.0.0-5.5.0 as live on their release Wednesdays, and saw each pre-download on the Monday two days earlier.

## Notes

- Maintenance can overrun; HoYoverse announces extensions and pays compensation. A run that misses an overrun catches the version the next day.
- GitHub Actions starts scheduled jobs late under load, more often at the top of the hour. A delay only moves the run further past the window, so it is harmless here.
- Pre-downloads open about two days before the release; any daily time catches them within a day.

## Sources

- [Genshin Impact 7.1 maintenance start and end times (allthings.how)](https://allthings.how/genshin-impact-7-1-maintenance-start-and-end-times-by-region/)
- [Genshin 7.1 maintenance time (timesaver.gg)](https://timesaver.gg/blog/genshin-7-1-maintenance-time)
- [Genshin Impact 7.1 maintenance and update times (playnews.gg)](https://www.playnews.gg/en/release-times/genshin-impact-7-1-maintenance-and-update-dates-and-times)
- [Genshin Impact 6.7 maintenance countdown (KeenGamer)](https://www.keengamer.com/articles/news/genshin-impact-6-7-maintenance-countdown-luna-viii-update-times/)
- [Genshin 7.2 release date projection (timesaver.gg)](https://timesaver.gg/blog/genshin-7-2-release-date)
- [Genshin Impact maintenance time in UTC+8 and other zones (Sportskeeda)](https://sportskeeda.com/esports/genshin-impact-server-maintenance-time-utc-8-est-pst-mst-cst-gmt-cet-ist)

## Unresolved questions

- Whether the HoYoPlay API switches to the new version at the start or the end of maintenance is not documented; running after the end is safe either way.
