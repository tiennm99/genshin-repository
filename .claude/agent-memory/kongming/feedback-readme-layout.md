---
name: feedback-readme-layout
description: How the user wants a public README laid out — data/download links first, credits last, long content split into docs/, few top-level entries
metadata:
  type: feedback
---

Put the things a visitor came for (live version, download/data links) at the top of a README, move credits and licence to the end, and split long reference content into `docs/` so the README keeps few entries.

**Why:** On 2026-10-26 the user asked for the genshin-repository README to be "impressive": "credit should go end", "links for data should be on top to allow user quickly view it", "split many docs into docs, not make readme too many entries". The 200 KB single-file README was replaced by a ~9.5 KB index plus generated `docs/versions/<N>.x.md` pages and hand-written guides, and the user accepted that shape.

**How to apply:** When advising on any README or landing doc for this user, lead with a Latest/quick-links block and a compact index, keep per-item detail in `docs/`, and place Credits and License last. Keep generated and hand-written docs clearly separated.
