---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
blueprint_file: blueprint.yaml
created_at: "2026-11-17T15:55:00+01:00"
category: apps
slug: cloud-arch-draw
title: "CloudArchDraw: AI-Native Multi-Cloud Architecture Studio (GCP, AWS, Azure)"
target_platforms:
  - Web (React + TypeScript + @xyflow/react + @viz-js/viz Graphviz WASM)
  - Local / Containerized Backend (Python FastAPI + uv + optional Graphviz CLI)
engine: "Canonical ArchitectureSpec AST + Graphviz DOT + Gemini Structured Output & Image Synthesis"
---

# CloudArchDraw: AI-Native Multi-Cloud Architecture Studio (`GCP`, `AWS`, `Azure`)

> **Implementation Pointer**: Implemented in `~/git/cloud-arch-draw` ([`github.com/palladius/cloud-arch-draw`](https://github.com/palladius/cloud-arch-draw) — see [`blueprint.yaml`](blueprint.yaml)). Companion regulated & foundational workshop scenario specs live in [`scenarios/`](scenarios/).

## Problem Statement

Architects, SREs, and engineering educators face three recurring friction points when turning system ideas or `SPEC.md` documents into cloud architecture diagrams:
1. **Static Text-to-Image Lock-In vs Brittle Code-Only DSLs**: Either an image model emits a pretty raster graphic with misspelled service labels that cannot be edited, or a diagram-as-code DSL breaks on the first manual spatial tweak.
2. **Missing Operational Reality**: Most generated diagrams draw happy-path boxes (`User -> API -> DB`) while hiding **bottlenecks, Single Points of Failure (SPOFs), security/residency perimeters, and FinOps trade-offs**.
3. **Disconnected Artifacts & Single-Cloud Myopia**: Moving a component or swapping a database on a visual canvas immediately desynchronizes the Graphviz `.dot` topology, the written Markdown architecture spec, the multi-cloud translation (`GCP` <-> `Azure` <-> `AWS`), and the presentation-ready visual diagram.

**CloudArchDraw** solves this by placing a **Canonical Architecture Graph (`ArchitectureSpec` JSON AST)** at the center of an interactive web studio. From a natural-language prompt (or structured scenario `SPEC.md`) and a Target Cloud selector (`GCP`, `AWS`, or `Azure`), the studio synthesizes a structured topology that simultaneously drives:
1. **An Interactive Visual Lego Canvas** (card & nested boundary drag-and-drop, 500+ multi-cloud icon palette, right-hand inspector for node docs and bottlenecks);
2. **A Deterministic Graphviz / DOT Diagram** with service cards inside nested cluster perimeters and **bottlenecks/SPOFs/trade-offs tethered directly to affected vertices**;
3. **A Synchronized Markdown Specification (`MD`)** documenting components, data/control flows, cross-cloud Rosetta equivalents, and mitigation tables;
4. **An On-Demand Presentation Renderer (`NanoBanana -> Recreate`)** that turns the user's *current edited canvas layout* + AST into a slide-ready visual diagram.

---

## Goals & Non-Goals

### Goals
- **GCP-First, Multi-Cloud Parity (`GCP`, `AWS`, `Azure`)**: First-class Google Cloud (`GCP`) support (cards, pastel zone palettes, regional boundaries), plus seamless switching and `1:1` split parity views for `Microsoft Azure` and `AWS`.
- **True Graphical Lego Editing**: Drag service cards and nested boundary groups (`VPC`, `Region / Sovereign Perimeter`, `Kubernetes Cluster`, `CI/CD Pipeline`), wire/delete edges, drop pieces from `assets/{gcp,aws,azure,generic}/`, and edit per-node documentation and findings directly in the UI.
- **Vertex-Attached Bottlenecks, SPOFs & FinOps Guardrails**: Architectural findings (`bottleneck`, `spof`, `tradeoff`, `security`) attach directly to graph vertices/edges and toggle cleanly as red X-ray overlays on the canvas, dashed callout notes in Graphviz `DOT`, and rows in the Markdown mitigation table.
- **1-Click `NanoBanana -> Recreate`**: After rearranging or editing components graphically, clicking `NanoBanana -> Recreate` captures the live canvas snapshot + canonical `ArchitectureSpec` and synthesizes a high-definition visual diagram via Gemini image generation.
- **100% Offline Classroom Resilience**: Built-in scenario presets (see [`scenarios/`](scenarios/)), visual editing, Graphviz WASM/CLI compilation, and deterministic compilation work 100% offline even when `GEMINI_API_KEY` is unset.
- **Optional Workshop Contrast (`YOLO Prompt` vs `Harness-Engineered`)**: Contrast naive 1-line prompts (missing ACL filters, synchronous LLM HTTP blocks, no token kill-switches) against structured `SPEC.md`-driven architectures scored against a deterministic readiness checklist.

### Non-Goals
- Real-time multi-user CRDT cursor collaboration (single-architect local/web studio is sufficient for v1).
- Executing live `terraform apply` directly from the canvas (exporting `.dot`, `.svg`, `.md`, `.json`, and `.png` is the v1 boundary).

---

## User Journey & Core Studio Workflow

1. **Prompt / Scenario Preset & Cloud Selection**:
   - Load one of the bundled reference scenarios ([`scenarios/01..05_*.md`](scenarios/)) or type a custom system description, and pick the Target Cloud (`GCP` full view, `AWS`, `Azure`, or `Azure | AWS` split parity view).
2. **Initial Tri-Modal Synthesis**:
   - Backend calls Gemini (`gemini-2.5-flash` structured JSON output grounded against `assets/catalog.json` valid `icon_id`s) to emit an `ArchitectureSpec`, auto-layouts left-to-right tiers (`0..5`) and nested group bounding boxes, and compiles both `dot_source` and `markdown_doc`.
3. **Interactive Visual Refinement ("Lego Editing" + Red X-Ray)**:
   - **Left Sidebar (Lego Palette)**: Search 500+ cloud icons (`assets/gcp/`, `assets/aws/`, `assets/azure/`, `assets/generic/`); click or drag-and-drop services or `+ Group` containers onto the canvas.
   - **Center Workspace**: Switch between `Split (Canvas + Graphviz DOT)`, `Full Visual Canvas`, `Graphviz / DOT`, `Markdown Description`, and `NanoBanana Diagram`, plus a `?xray=1` toggle highlighting SPOFs and bottlenecks in red.
   - **Right Inspector**: Click any card, group, or connector to edit its label, service icon, parent group, cross-cloud Rosetta equivalents, Markdown documentation, or **Bottlenecks, SPOFs & Trade-offs**. Every edit debounces `POST /api/compile` (`< 50ms`) so `.dot` and `.md` stay 100% synchronized.
4. **Visual Re-Synthesis (`NanoBanana -> Recreate`)**:
   - When the canvas has been modified, `NanoBanana -> Recreate` highlights dirty state. Clicking it captures a PNG snapshot of the canvas (`html-to-image`) plus the `ArchitectureSpec` and calls `POST /api/nanobanana`, returning a slide-grade diagram faithful to the user's spatial layout.

---

## Canonical Data Model (`ArchitectureSpec`)

All views read from and write to a single canonical schema:

```json
{
  "id": "gcp-canary-cicd",
  "title": "Canary Delivery Pipeline on GKE + Cloud Deploy",
  "cloud": "gcp",
  "region": "europe-west6 (Zurich)",
  "prompt": "Original user prompt or scenario SPEC...",
  "summary": "Executive summary...",
  "groups": [
    {
      "id": "grp_cicd",
      "label": "CI/CD Build & Release Plane",
      "group_type": "pipeline",
      "parent_id": null,
      "background_color": "#E8F0FE",
      "border_color": "#1A73E8",
      "position": { "x": 400, "y": 56 },
      "width": 292,
      "height": 340
    }
  ],
  "nodes": [
    {
      "id": "cloud_build",
      "label": "Build Trigger",
      "service": "Cloud Build",
      "icon_id": "gcp/cloud-build",
      "group_id": "grp_cicd",
      "tier": 2,
      "subtitle": "Test, containerize & attest",
      "documentation": "Runs unit tests and builds immutable container digests.",
      "rosetta": {
        "gcp": "Cloud Build",
        "azure": "GitHub Actions / Azure Pipelines",
        "aws": "AWS CodeBuild"
      },
      "position": { "x": 436, "y": 110 },
      "findings": [
        {
          "id": "f_cache",
          "kind": "bottleneck",
          "severity": "medium",
          "title": "Cold layer cache on monorepo builds",
          "detail": "Rebuilding base layers adds 4-6 min queue latency.",
          "mitigation": "Enable remote Artifact Registry layer caching."
        }
      ]
    }
  ],
  "edges": [
    {
      "id": "e1",
      "source": "cloud_build",
      "target": "artifact_reg",
      "label": "Push OCI digest",
      "protocol": "HTTPS",
      "style": "solid",
      "is_bottleneck": false,
      "bottleneck_note": ""
    }
  ],
  "dot_source": "digraph CloudArchitecture { ... }",
  "markdown_doc": "# Canary Delivery Pipeline...",
  "nanobanana_url": "/generated/nb_gcp-canary-cicd.png"
}
```

---

## Bundled Multi-Cloud Reference Scenarios ([`scenarios/`](scenarios/))

1. **[`01_insurance_claims_disparity.md`](scenarios/01_insurance_claims_disparity.md)** — Event-driven multimodal hail/car-damage photo Vision + PDF repair-invoice OCR disparity workflow (`< CHF 2,000` straight-through payout vs async human adjuster queue).
2. **[`02_cantonal_bank_mortgage_rag.md`](scenarios/02_cantonal_bank_mortgage_rag.md)** — Sovereign banking mortgage dossier RAG with OAuth 2.0 On-Behalf-Of (`RFC 8693`) pre-retrieval ACL filtering and governed read-only `MCP` tools.
3. **[`03_canton_public_admin_citizen_portal.md`](scenarios/03_canton_public_admin_citizen_portal.md)** — Multilingual (`DE/FR/IT/EN`) civic intake with deterministic pre-prompt national-ID/PII redaction and civil-servant decision queue.
4. **[`04_rag_on_vms_vs_serverless_managed.md`](scenarios/04_rag_on_vms_vs_serverless_managed.md)** — Architectural & FinOps trade-off comparing self-hosted open-weight LLMs (`Apertus` / `Gemma` on GPU Kubernetes + `vLLM`) against Serverless Managed RAG.
5. **[`05_redundant_lamp_to_cloud_native.md`](scenarios/05_redundant_lamp_to_cloud_native.md)** — Legacy 2-VM sticky LAMP stack modernized to Multi-AZ stateless containers, managed HA relational DB, Redis, Object Storage, and CDN/WAF.

---

## Free Multi-Cloud Architecture Icons Guide (`assets/{gcp,aws,azure,generic}/`)

To avoid licensing friction or broken hotlinks, **CloudArchDraw** uses a clean directory contract (`assets/gcp/`, `assets/aws/`, `assets/azure/`, `assets/generic/` indexed by `assets/catalog.json`) populated via free official and open-source icon packs:

1. **Official Vendor Free Icon ZIP Portals**:
   - **Google Cloud (`GCP`)**: `https://cloud.google.com/icons`
   - **Amazon Web Services (`AWS`)**: `https://aws.amazon.com/architecture/icons/`
   - **Microsoft Azure (`Azure`)**: `https://learn.microsoft.com/en-us/azure/architecture/icons/`
2. **Automated 1-Command Provisioning (`just fetch-icons`)**:
   - Curated offline 64x64 fallback SVGs for core & GenAI services (`gcp/cloud-run`, `gcp/kubernetes-engine`, `gcp/alloydb`, `gcp/vertex-ai`, `aws/bedrock`, `azure/ai-foundry`, `generic/users`, `generic/git-repo`).
   - Automated extraction of 480+ provider PNGs from `mingrammer/diagrams` (MIT) into `assets/{gcp,aws,azure,generic}/`, compatible with both HTML `<img>` and Graphviz `dot` `<IMG SRC="..."/>`.
   - Optional pure-SVG frontend libraries: `@likec4/icons` (npm) and `jgraph/drawio` (Apache-2.0).

---

## Reference System Architecture & Graphviz Compilation Rules

```text
+-----------------------------------------------------------------------------------+
|                        React + TypeScript Studio Frontend                         |
|  [Lego Icon Palette]  |  [@xyflow/react Canvas] <-> [@viz-js/viz WASM]  | [Inspector] |
|                       |  [Markdown Spec Tab]    <-> [NanoBanana View]   |             |
+-----------------------------------------+-----------------------------------------+
                                          | REST JSON (/api/*)
+-----------------------------------------v-----------------------------------------+
|                         Python FastAPI Backend (managed via uv)                   |
|  - POST /api/generate    -> Gemini 2.5 Flash (Structured JSON) + Offline Fallback |
|  - POST /api/compile     -> Tiered Auto-Layout + Graphviz DOT + Markdown Compiler |
|  - POST /api/nanobanana  -> Gemini Image (Canvas Snapshot + Topology -> PNG)      |
|  - GET  /api/catalog     -> Serves 500+ indexed icons from assets/catalog.json    |
+-----------------------------------------------------------------------------------+
```

- Use `digraph CloudArchitecture` with `rankdir=LR`, `splines=ortho`, `compound=true`.
- Map every `ArchGroup` to `subgraph cluster_<id>` with rounded pastel fill (`#E8F0FE` blue, `#CEEAD6` green, `#FEEFC3` yellow, `#FAD2CF` red).
- Render every `ArchNode` as a rounded HTML-like `<TABLE>` displaying `label`, `service`, and `subtitle`.
- Tether every `NodeFinding` (`show_findings=true`) via `[style=dashed, arrowhead=dot, constraint=false]` so bottleneck callouts never distort left-to-right tier ranking.

---

## Acceptance Criteria & TDD Verification Gate

1. **Icon Catalog Gate**: `GET /api/catalog` returns `>= 40` icons (500+ after `just fetch-icons`) covering `gcp` (sorted first), `aws`, `azure`, and `generic`, with every `local_path` present on disk.
2. **Graphviz Syntax Gate (`dot -Tsvg`)**: Every bundled scenario preset across `GCP`, `AWS`, and `Azure` compiles cleanly through `dot -Tsvg` with exit code `0` and zero stderr errors, including vertex-attached bottleneck callouts.
3. **Bidirectional Edit Sync**: Modifying a node's position, documentation, or `findings` via `POST /api/compile` deterministically updates both `dot_source` and `markdown_doc`.
4. **NanoBanana Recreate Resilience**: `POST /api/nanobanana` produces a high-definition PNG when `GEMINI_API_KEY` is set, and falls back gracefully to a rendered SVG artifact offline.
5. **Zero-Overflow Single-Viewport UI**: `npm run build` exits `0`, and loading `/` in headless Chrome at `1280x800` renders `Palette | Split Canvas + DOT | Inspector` with `0` console errors, `0px` overflow, and URL-persisted state across `F5` reload.
