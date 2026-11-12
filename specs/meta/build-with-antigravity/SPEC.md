---
speck_version: "0.1"
mode: manual
created_at: 2026-09-04
title: "Build with Antigravity Showcase & Automated Ingestion Pipeline"
category: "antigravity"
status: "draft"
---

# Build with Antigravity: Showcase & App Ingestion Framework

An interactive, Netflix-style showcase portal celebrating applications, agentic systems, and tooling built with **Antigravity**, featuring a dual-mode ingestion pipeline (self-service PRs via YAML/JSON Schema 1.0 + guided web intake) and a dedicated Antigravity skill for automated developer onboarding.

---

## 1. Problem Statement & Motivation
As developers adopt Antigravity across internal teams and the open-source community, there is no centralized, visually stunning gallery showcasing real-world creations. 

Today:
- Projects are scattered across GitHub repositories, internal chats, and ad-hoc demos.
- Discoverability is poor; newcomers cannot easily see what is possible.
- Curating an app showcase manually is tedious and does not scale.

**Goal:** Create **"Build with Antigravity"**, a media-rich showcase portal (modeled after the Netflix carousel / Lobby Private Journal aesthetic) paired with an automated submission pipeline driven by **Schema 1.0** and an Antigravity developer skill.

---

## 2. User Experience & Architecture

### 2.1 The Showcase Web Portal ("Netflix for Antigravity Apps")
- **Visual Design:** High-contrast dark theme, horizontal scrolling carousels sorted by category (e.g., *Trending*, *Enterprise SRE*, *Creative Agents*, *CLI Tools*, *Community Highlights*).
- **Rich Media Cards:**
  - Hero preview image or looping video/GIF.
  - Title, badge pill (e.g., `MIT`, `Apache-2.0`, `Internal`), and one-line punchy elevator pitch.
  - Hover action: quick modal preview with architecture diagram, author bio, LinkedIn profile badge, GitHub star counter, and a direct "Launch Demo" or "View Source" button.

### 2.2 Dual Ingestion Funnel
1. **GitOps Pull Request Workflow (Automated):**
   - Contributors submit their app by placing a single manifest file in `apps/<slug>/app.yaml` (or `app.json`).
   - GitHub Actions runs the Python validator (`validate_submission.py`) on PR creation.
   - If valid and tests pass, the PR can be auto-previewed and merged.
2. **Form / Guided Web Intake:**
   - A lightweight web form that outputs the compliant YAML file directly or triggers a repository submission.

---

## 3. Manifest Specification: Schema 1.0

Every submission must validate against **Schema 1.0**. The schema is intentionally simple for v1.0, with planned strictness evolutions in v1.1+.

```yaml
# Schema Version 1.0 Reference Manifest
schema_version: "1.0"
name: "SRE Incident Bot Benjamin"
slug: "sre-incident-bot-benjamin"
headline: "Autonomous multi-agent incident response team built with Antigravity & Google ADK."
description: |
  Benjamin orchestrates 5 specialized subagents (Incident Commander, Communications, 
  Scribe, Ops, and Postmortem) to resolve production outages with human-in-the-loop steering.

category: "enterprise-sre" # Choices: enterprise-sre, developer-tools, creative, productivity, community
tags:
  - "sre"
  - "gemini"
  - "incident-response"
  - "discord"

author:
  name: "Riccardo Carlesso"
  github: "palladius"
  linkedin: "https://www.linkedin.com/in/riccardocarlesso/"
  role: "Principal Cloud Architect"

repository:
  url: "https://github.com/palladius/adk-sre-benjamin"
  license: "Apache-2.0" # Standard SPDX identifier or "Proprietary/Internal"
  branch: "main"

showcase:
  hero_image: "assets/hero.png" # 16:9 ratio, min 1280x720
  demo_url: "https://benjamin.demo.internal/" # Optional
  video_url: "https://youtu.be/..." # Optional

antigravity:
  version_tested: ">=2.0.0"
  features_used:
    - "skills"
    - "delegation"
    - "mcp"
```

### Validation Rules (Schema 1.0)
- `schema_version`: Must strictly equal `"1.0"`.
- `slug`: Lowercase alphanumeric and hyphens only (`^[a-z0-9-]+$`), max 48 characters.
- `name`: String, 3 to 64 characters.
- `headline`: String, 10 to 140 characters (clean, punchy, no markdown formatting).
- `author.linkedin`: Must be a valid HTTPS URL matching `^https:\/\/(www\.)?linkedin\.com\/in\/[a-zA-Z0-9_-]+\/?$`.
- `repository.license`: Must be a recognized SPDX license tag or `"Internal"`.
- `showcase.hero_image`: Must be a valid URL or path to an existing image file.

---

## 4. Submission Skill (`build-with-antigravity`)

The repository will ship with an official Antigravity skill located at `skills/build-with-antigravity/SKILL.md`.

### Skill Responsibilities:
1. **Interactive Interview:** Asks the developer for their project repo, LinkedIn, description, and hero media.
2. **Scaffolding:** Writes the compliant `app.yaml` file into `apps/<slug>/app.yaml`.
3. **Local Lint & Verification:** Executes `python scripts/validate_submission.py apps/<slug>/app.yaml`.
4. **Git Delivery:** Creates a branch `add-app-<slug>`, commits the file, and offers to open a GitHub PR via `gh pr create`.

---

## 5. Verification Script: `validate_submission.py`

A standalone, zero-dependency Python script validates submissions both locally and in CI:

```bash
# Validate single file
python scripts/validate_submission.py apps/sre-incident-bot-benjamin/app.yaml

# Validate all apps
python scripts/validate_submission.py --all
```

---

## 6. Future Roadmap (Post-v1.0)
- **Schema 1.1:** Add automated repository health checks (verifying repo is public, checking commit activity, scanning for license files).
- **Live Preview Generator:** Auto-deploy a ephemeral static preview site on every PR using Cloud Run or GitHub Pages.
- **Upvoting & Bookmarking:** Enable Antigravity community members to star and comment on showcased apps.
