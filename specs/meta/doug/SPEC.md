---
speck_version: "0.2"
mode: drensin-grill-me
idea_file: input_prompt.md
created_at: "2026-11-24T11:05:00+01:00"
updated_at: "2026-11-25T15:00:00+01:00"
title: "Doug Edgeworth (doug) — Durable Agentic DAG Platform & Multi-App Host on Google Cloud Run"
category: "meta"
slug: "doug"
status: "in-implementation"
living_spec: "~/git/doug/docs/SPEC.md"
---

# 🎩 Doug Edgeworth (`doug`) — Pointer to Living Specification

> **Living Specification**: [`~/git/doug/docs/SPEC.md`](file:///usr/local/google/home/ricc/git/doug/docs/SPEC.md)  
> **Intent (10 lines)**: [`~/git/doug/docs/INTENT.md`](file:///usr/local/google/home/ricc/git/doug/docs/INTENT.md)  
> **Repository**: [`https://github.com/palladius/doug`](https://github.com/palladius/doug) (`~/git/doug`)

---

## Architectural Evolution Summary (v0.1 Rails Draft → v1.0 Living Spec)

Following a critical DAG feasibility review (`DBOS`, `Inngest`, `Prefect 3`, `Vercel WDK`, `Windmill`, `Restate`, `Google Cloud Workflows`, `Vertex AI Pipelines`, `Rails 8 Custom`), **Doug Edgeworth** (`🎩` — *The Directed Acyclic Gentleman*) moved from the initial custom Rails 8 + Cloud Tasks sketch to a **single-container TypeScript + DBOS** architecture:

- **Durable Engine**: **DBOS (`@dbos-inc/dbos-sdk`, MIT OSS)** backed by PostgreSQL (`DATABASE_URL`-agnostic via Drizzle ORM; Docker Postgres 16 in dev; Supabase Session Pooler `:5432` / Neon Direct / Cloud SQL in prod — never PgBouncer Transaction Pooler `:6543`). Native 1-line `DBOS.forkWorkflow(id, startStep)` replay without maintaining a custom DAG lock engine.
- **Web & Control Tower UI**: **Hono (`@hono/node-server`)** + **React/Vite** (Tailwind + **`@xyflow/react`** visual DAG canvas + 2s polling + 90-day rolling cookie via Google Sign-In `openid email` / `ADMIN_SECRET`).
- **Google Cloud Run (`$0` idle)**: 1 Docker image (Node 22 + `FFmpeg`) deployed as **`doug-web`** (Cloud Run Service `512Mi`) + **`doug-runner`** (Cloud Run Job `2Gi` triggered by Cloud Scheduler or `/admin` via Cloud Run Jobs API).
- **Zero-Idle HITL Suspend & Resume**: First-class `WAITING_HITL` gate (`doug_hitl_gates`) that exits the Cloud Run Job container immediately (`$0` idle) and resumes from the exact checkpoint via `DBOS.forkWorkflow` upon human approval in `/admin`.
- **Target Verticals**:
  1. [`specs/apps/auto-news`](../../apps/auto-news/SPEC.md) — *Riccardo vs Riccardo* satirical dual-lens newspaper with 3-position tactile **"La Rotella"** (`🔴 SX` | `⚖️ 0` | `🔵 DX`).
  2. [`specs/apps/whats-up-with-ai`](../../apps/whats-up-with-ai/SPEC.md) — 24h AI Lab Radar + 2-host Italian morning podcast (`gemini-3.8-flash-tts` + Lyria 3 jingle + `FFmpeg` + private RSS).
  3. `sre-email-watchdog` — Read-only Gmail poller for GCP SLO violations & failed GitLab/GHP/Cloudflare deploys → AI investigation → **Zero-Idle HITL gate** → approved remediation.
