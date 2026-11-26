# 🎩 Doug Edgeworth (`doug`)

> **Target Implementation Repository & Living Spec**: `~/git/doug` ([`github.com/palladius/doug`](https://github.com/palladius/doug)) — see [`~/git/doug/docs/SPEC.md`](file:///usr/local/google/home/ricc/git/doug/docs/SPEC.md) and [`blueprint.yaml`](blueprint.yaml).  
> Per `respec` rules, this directory hosts **specifications only** (zero application code).

## Overview

**Doug Edgeworth** (`🎩` — *The Directed Acyclic Gentleman*) is a single-container **TypeScript + DBOS + Hono + React/Vite** durable agentic DAG platform and multi-app host engineered for **Google Cloud Run (`min-instances=0`, `$0` idle)**. It runs as a lean Cloud Run Service (`doug-web`) paired with a Cloud Run Job (`doug-runner` with `FFmpeg`), featuring an interactive **`@xyflow/react` Visual DAG Control Tower (`/admin`)**, 1-click **`DBOS.forkWorkflow` deterministic replay**, and **Zero-Idle HITL Suspend & Resume**.

## Contents

- **[`SPEC.md`](SPEC.md)** — Pointer and evolution summary referencing the authoritative living spec at `~/git/doug/docs/SPEC.md`.
- **[`input_prompt.md`](input_prompt.md)** — Original prompt and Drensin `/grill-me` architectural decisions.
- **[`blueprint.yaml`](blueprint.yaml)** — Target repository and vertical specialization pointers.

## Launch Specializations

1. **[`specs/apps/auto-news`](../../apps/auto-news/SPEC.md)** — *Riccardo vs Riccardo*: Satirical Left-Wing vs Right-Wing Live Political & World Newspaper (US 🇺🇸, IT 🇮🇹, CH 🇨🇭, World 🌍) with a sacred factual core and tactile 3-position **"La Rotella"**.
2. **[`specs/apps/whats-up-with-ai`](../../apps/whats-up-with-ai/SPEC.md)** — *What's Up With AI*: Rolling 24h AI Lab Radar + 5–7 min Italian 2-Host Morning Podcast (`gemini-3.8-flash-tts` + Lyria 3 synth-funk jingle + `FFmpeg`) and private Podcast RSS feed.
3. **`sre-email-watchdog`** — Agentic SRE investigator monitoring alert emails (GCP SLO violations, broken GitLab/GHP/Cloudflare deploys) with **Zero-Idle HITL** approval gates in `/admin`.
