---
speck_version: 0.1
mode: manual
created_at: 2026-11-28
authors: Riccardo Carlesso & Antigravity
status: draft
tags: [agents, harness, antigravity-sdk, rust, jev, system-one, cli, threads]
---

# 🦀🚦 agj — a stronger Antigravity: Rust CLI harness with programmable threads and a Jev traffic-light gate

## Problem Statement

Riccardo runs several agent harnesses on the same machine. Two pain points:

1. **Threads are not programmable in Antigravity.** In T3 Code an agent can open or archive threads (`t3_thread_organize`). Antigravity Desktop exposes no API, MCP tool, CLI command or `<agent-embed>` bridge to list, open or archive conversations (verified 2026-11-28: the `chumbia` "Archivia Chat" button was POSTing to a non-existent endpoint). From inside a chat you cannot say *"open me 3 chats in this project"* or *"archive this chat"*.
2. **Command approval is binary and dumb.** Either every `run_command` is confirmed by hand (slow) or everything is allowed (unsafe). There is no fast, calibrated, automatic risk judgement.

The Google Antigravity SDK (`google-antigravity`, Python) lets you build your *own* harness, with its own conversations in its own `save_dir`, fully separate from the Desktop app. That separation is an advantage: in our own harness, threads are ours to manage.

## Goals

- A real **CLI binary** `agj` (like `agy`) with a REPL, one-shot mode and a background daemon.
- **Programmable threads**, available both to humans (CLI) and to the agent (in-chat tools):
  - *"aprimi 3 chat in questo progetto"* → 3 new threads in the current project.
  - *"archivia questa chat"* → the current thread is archived at the end of the turn.
  - Plus list, resume, restore, pin, rename and cross-thread messaging.
- **Jev System One risk gate** on every executing or mutating tool call. It uses calibrated `P_safe` in ~100–200 ms, with a 5-colour traffic light:
  - ⚪ **White**: read-only, decided by a deterministic allowlist with no Jev call → run silently.
  - 🟢 **Green**: `P_safe ≥ 0.80` → **run without asking**.
  - 🟡 **Yellow**: `0.50 ≤ P_safe < 0.80` → ask `[y/N]`.
  - 🟠 **Orange**: `0.20 ≤ P_safe < 0.50` → ask with friction (shows the risk reason; the user types `yes`).
  - 🔴 **Red**: `P_safe < 0.20` or a hard-deny rule matched → **block**, and the agent must propose an alternative.
- The usual agent loop: ask a question, read the file system, edit files, run code, browse the web (SDK built-ins).
- Fast and pleasant: Rust front-end, coloured streaming output, emoji per tool call, `--json` everywhere, `--ai-help`.
- Full audit trail of every gate decision, plus a golden eval set (Eval-Driven Development).

## Non-Goals

- Driving, reading-for-write or archiving **Antigravity Desktop app** conversations. Without an official API this is out of scope, and we never write to the app's internal state.
- A GUI (a TUI is fine). Multi-user or remote access.
- Re-implementing the agent loop. The Antigravity SDK *is* the agent core.
- Deleting threads in v1 (archive only; `purge` comes later as a 🔴-gated command).

## Technical Plan / Approach

### Hard constraint → "Rust outside, Python inside"
The SDK is Python-only, so:

```
agj (Rust binary)
├─ client mode:  CLI + REPL/TUI  ──unix socket ~/.agj/agj.sock (JSON-RPC, 0600)──┐
└─ daemon mode:  `agj daemon` (auto-started by the client)  ◄────────────────────┘
     ├─ SQLite ~/.agj/state.db         (threads, gate_decisions)
     ├─ Jev Gate (Rust)                 → Jev (jev_choice), 500 ms timeout
     └─ supervises N Python workers     (one per *live* thread, stdio JSON-RPC)
          worker = google-antigravity Agent + persistence(save_dir, conversation_id)
                   + hook pre_tool_call_decide → asks daemon for a verdict
                   + custom tools thread_* → forwarded to daemon
```

- **Rust**: `tokio`, `clap` (with completions), `ratatui` + `crossterm` (TUI), `rusqlite`, `serde_json`, `ulid`, `reqwest` (Jev).
- **Python worker**: deliberately thin, pinned in a `uv`-managed venv at `~/.agj/py/`. Default model **Gemini 4 Hydrogen** (hosted, `GEMINI_API_KEY` from env/`.env`); optional local **Gemma 4** through `LiteRTAgentConfig`.
- **Single source of truth = daemon.** Policy is decided in Rust, so a crashed or malicious worker cannot bypass the gate. A dead worker means the thread goes `idle` and is resumed later from its `save_dir`.

### Thread model
| Field | Notes |
|---|---|
| `id` | ULID (short form = first 8 chars) |
| `title` | auto after first turn, editable |
| `project` | absolute path (git root if detected) |
| `state` | `active` · `idle` · `archived` |
| `pinned`, `parent_id` | `parent_id` set when spawned by another thread |
| `model`, `save_dir` | `~/.agj/threads/<id>/` (SDK persistence + `transcript.jsonl`) |

Lifecycle: `new → idle ⇄ active → archived → (restore) idle`. Creating a thread does **not** spawn a worker; the worker starts on the first message.

### Interfaces
**CLI**
```
agj                                   # REPL, new thread in $PWD
agj chat [-t <id>]                    # REPL, resume
agj -p "prompt" [--json] [--yes-yellow]
agj threads ls|new|archive|restore|pin|rename   (new: --project . -n 3 --prompt "...")
agj gate test "<cmd>"                 # dry-run the gate: light, P_safe, rule, latency
agj daemon start|stop|status
agj --ai-help                         # compact machine-readable usage
```
Non-TTY mode never prompts: 🟡/🟠 exit with code 4 and the reason (unless `--yes-yellow`). Lists are bounded (`--limit 20`).

**Agent tools (SDK custom tools in every worker)**
| Tool | Gate |
|---|---|
| `thread_list(project?, state?, limit)` | ⚪ |
| `thread_open(project?, count 1–5, titles?, initial_prompt?)` | 🟢 if count ≤ 3, 🟡 if 4–5 |
| `thread_archive(id="this")` | 🟢 for `this`, 🟡 for other threads; takes effect at end of turn |
| `thread_message(id, text)` | 🟡 |

**REPL**: streaming answers; one line per tool call, e.g. `🟢 run_command  rm -rf build/  P_safe=0.91`. Slash commands: `/threads /new [n] /archive /switch <id> /gate <cmd> /trust yellow /chumbia /quit`.

### Jev gate pipeline
1. **Hard-deny floor (🔴, deterministic, Jev can never override)**: writes to `.env*`, `.riccardo/`, `~/.ssh`, `~/.gnupg`, gcloud credentials; `git clean -fd*`, `git reset --hard`, force-push to `main`/`master`; `rm -rf` outside the project, of `~` or of `/`; `curl … | sh`; `sudo`; `chmod -R 777`; `gcloud … delete`, `gcloud projects …`, `bq rm`, `kubectl delete ns`.
2. **White allowlist (⚪)**: read-only tools and commands (`ls`, `cat`, `rg`, `git status/log/diff`, `read_file` inside the project).
3. **Cache**: same normalized command + context class within 24 h → reuse the light.
4. **Jev call**: one atomic `choice` question, `options=["safe","unsafe"]`. Context is **structured facts computed in Rust** (resolved target paths, inside-project?, git dirty?, tool name, network egress?). Never raw file contents; secrets redacted. `P_safe = p("safe")`.
5. **Mapping** to 🟢🟡🟠🔴 using thresholds from `~/.agj/config.toml` (`[gate] green=0.80 yellow=0.50 orange=0.20`).
6. **Adjustments (stricter only)**:
   - *Ceiling*: `git push`, deploys, package installs and non-allowlisted network egress are at least 🟡.
   - *Low confidence*: top-2 options within 0.10 → bump up one step.
   - *Fail-closed*: a Jev timeout or error gives 🟡, never 🟢.
7. **Session trust**: `/trust yellow` auto-runs 🟡 for the current thread only, logged. There is no override for 🟠 or 🔴.
8. **Audit**: every decision is written to `gate_decisions` (ts, thread, tool, normalized cmd, context, P_safe, light, rule, user answer, latency).

### Storage
```
~/.agj/config.toml   state.db   agj.sock   py/   threads/<id>/   logs/agjd.log
```
`agj` never writes secrets or `.env` files.

## Alternatives Considered

- **Pure Python CLI**: simplest, but slow startup and a weaker TUI. Rejected because Riccardo explicitly wants something fast. Python is confined to the SDK worker.
- **PyO3-embedded Python in one Rust binary**: single artifact, but a heavy, fragile build. Deferred; the `uv` venv is simpler for v1.
- **MCP server instead of a binary**: rejected by the user. The goal is a full harness ("an Antigravity CLI"), not a plugin for someone else's harness. A read-only sessions MCP may be extracted later.
- **LLM-as-judge for risk**: seconds of latency and uncalibrated. Rejected in favour of Jev System One (~100–200 ms, calibrated probabilities).
- **Gate inside the Python worker**: rejected. A policy that lives in the process it is policing can be bypassed; the daemon owns policy.
- **Hacking Antigravity Desktop storage to archive chats**: rejected. It relies on undocumented internal state, breaks host boundaries, and risks corruption.

## Implementation Plan

Each phase is TDD and ends green on `just test`. The `justfile` starts with `list` (`just -l`). `VERSION` + `CHANGELOG.md` are bumped every phase.

| # | Deliverable | Done when |
|---|---|---|
| 0 | Repo `agj` (private, `palladius`), Cargo workspace (`crates/agj`, `crates/gate`, `crates/store`), `worker/` uv project, `justfile`, CI | `just test` green with smoke tests |
| 1 | Python worker MVP + stdio JSON-RPC; Rust client one-shot `agj -p` | Ask, read a file, run `ls`, resume via `conversation_id` |
| 2 | Daemon + SQLite + thread lifecycle + CLI `threads *` | `threads new -n 3 --json` < 300 ms with no worker spawned; daemon restart keeps state |
| 3 | Agent tools `thread_*` | "aprimi 3 chat" → 3 threads with `parent_id`; "archivia questa chat" → reply delivered, then archived and worker exited |
| 4 | Jev gate (floor, allowlist, cache, Jev, mapping, adjustments, audit) + `agj gate test` | Acceptance cases below pass, with Jev mocked |
| 5 | REPL/TUI with traffic-light lines and 🟡/🟠 modals | Phases 1–4 scenarios work end-to-end in the TUI |
| 6 | Golden eval set `evals/gate_cases.jsonl` (≥ 50 cases) + `just eval-gate` | ≥ 90% accuracy overall, 100% on expected-🔴 |
| 7 | `/chumbia` native: audit + real archive | Session archived and verified by test |
| 8 | Hardening: `BudgetConfig`, local Gemma 4 mode, packaging (`cargo install` / brew tap) | Works offline on LiteRT |

### Acceptance criteria
1. `agj gate test "ls -la"` → ⚪, no Jev call, < 5 ms.
2. `agj gate test "git reset --hard"` → 🔴 via the floor, even with Jev mocked at `P_safe=0.99`.
3. Jev mocked to time out → 🟡, never 🟢.
4. `rm -rf build/` inside the repo with Jev at `P_safe=0.85` → 🟢, runs, logged.
5. `git push` with Jev at `P_safe=0.95` → 🟡 (ceiling).
6. Non-TTY `agj -p` with a 🟡 action → exit 4 with the reason; never hangs.
7. `kill -9` on a worker mid-turn → thread `idle`; the next message resumes from `save_dir`.
8. `agj chat -t <archived>` → refused with a `threads restore` hint, exit 3.

## Open Questions

- Name: `agj` or something punnier?
- Colour semantics: is ⚪ = read-only/no-Jev and 🟠 = ask-with-friction right, or should ⚪ mean "Jev unavailable"?
- Is the 🟢 threshold 0.80 global, or per-project (stricter in prod/infra repos)?
- Jev access: is the `jev-proxy` MCP / API key already set up on this Mac, or is `install-jev-for-antigravity` step zero?
- v2: interop with T3 Code threads (drive T3 via its MCP from `agj`)?
- Python embedding: stay with a `uv` venv, or move to PyO3 for a single artifact?
