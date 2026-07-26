---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-26T16:35:25Z"
model: gemini-flash-latest
tokens:
    prompt: 442
    output: 1367
    total: 3314
---

# Mage & Familiar: Retro GCP Co-Op RPG

## Problem Statement
Standard technical workshops teaching Google Cloud Platform (GCP) concepts like Identity and Access Management (IAM), Cloud SQL, and Cloud Run can often feel dry, solitary, and detached from playful problem solving. Workshop participants lack an engaging, collaborative medium to learn backend REST concepts and cloud infrastructure workflows.

There is a need for a retro 80s CRT-style co-op RPG where a human player ("Mage") and an API/AI agent ("Familiar") must collaborate symbiotically in the Realm of Gicipya (GCP) to collect cloud artifacts, perform alchemy, and trigger actual GCP resource deployments.

## Goals
- Deliver a retro 80s CRT web application featuring glowing phosphor styling, scanlines, and chiptune audio effects.
- Implement a 50/50 symbiotic co-op game system where neither the human Mage (UI controls, movement, Latin spells) nor the API Familiar (REST API calls, sonar pinging, Esperanto incantations) can complete the dungeon alone.
- Build a GCP Alchemy system featuring collectible artifacts: *Potion of IAM* (service account authorization), *Cauldron of Cloud SQL* (3 bat wings + sulfur, brewed with delayed async completion), and *Rune of Cloud Run*.
- Execute a real-world Cloud Run deployment trigger upon completing the final alchemy ritual, deploying a lightweight live container displaying "IT WORKS!".
- Log full session telemetry to Firebase Firestore for real-time syncing and session recording.
- Build a Dual-Trail Visual Replay map displaying the Mage's path in cyan/blue and the Familiar's path in red, with particle fireworks visual effects whenever their trajectories cross.
- Provide a public Global Replay Leaderboard allowing workshop attendees to inspect and interactively play back historical runs step-by-step.

## Non-Goals
- Building a general-purpose 3D game engine or high-framerate real-time action RPG.
- Supporting generic multi-player lobbies with more than 2 entities (1 Mage, 1 Familiar per game session).
- Providing full cloud infrastructure provisioning outside the controlled sandbox environment required for the Cloud Run deployment ritual.

## Technical Plan / Approach
### Frontend & UI Architecture
- **Framework**: React / Next.js with HTML5 Canvas overlay for grid rendering and CRT post-processing filters (scanlines, screen curvature, glow).
- **Audio**: Web Audio API synth generating procedural 8-bit retro sound effects for movement, spell casts, item pickups, and fireworks bursts.

### Core Gameplay Loop & Dual Controls
- **Grid Engine**: Synchronous 2D tile grid representing the Realm of Gicipya.
- **Human Mage Controls**: Keyboard/UI buttons for N/S/E/W grid navigation, physical item interaction, and Latin spells (`Incendio`, `Protego`, `Revelio`).
- **API/AI Familiar Interface**: Exposed REST API endpoints (`/api/v1/sessions/{id}/familiar/act`, `/ping`, `/incantation`). The Familiar performs sonar map pings, scouts hidden tiles, and casts Esperanto incantations (`Lumo`, `Scio`, `Muro`).

### GCP Alchemy Artifacts & Live Ritual
1. **Potion of IAM**: Obtains temporary JWT tokens that unlock service-account-gated doors.
2. **Cauldron of Cloud SQL**: Requires combining ingredients in a tile location; initiates an asynchronous Cloud Function background job with real-time polling to simulate database provisioning delays.
3. **Rune of Cloud Run**: Interacting with the final altar triggers a serverless Cloud Build / Cloud Run Admin API execution that deploys an actual microservice returning "IT WORKS!".

### Telemetry, Dual-Trail Replay & Leaderboard
- **Firestore Schema**: `sessions` collection tracking `mage_trail` (array of `{x, y, timestamp}`), `familiar_trail`, score, duration, and alchemy completion status.
- **Dual-Trail Rendering**: Blue line (`#00F0FF`) rendered for Mage movement history, Red line (`#FF0055`) for Familiar movement history.
- **Intersection Fireworks**: Frame-by-frame path analyzer detects tile coordinate overlaps between Mage and Familiar trails, triggering a particle explosion burst animation on the canvas.
- **Leaderboard Playback UI**: Public view querying Firestore for top scores, providing step-by-step playback controls (Play, Pause, Scrub timeline slider, 1x/2x/4x playback speed).

## Alternatives Considered
- **WebSockets for Familiar interaction**: Considered for bidirectional communication, but standard REST HTTP endpoints were chosen so workshop participants can easily interact using `curl`, Python scripts, Postman, or custom HTTP clients.
- **Client-only simulated deployment**: Considered faking the Cloud Run deployment, but executing a real deployment via GCP Node.js SDK provides higher educational value and satisfaction during live workshops.

## Implementation Plan
- **Phase 1 (Core Engine & Retro Visuals)**: Implement HTML5 Canvas tile map, CRT CSS shaders, basic grid physics, keyboard movement, and Web Audio SFX.
- **Phase 2 (Familiar REST API & Game State Sync)**: Build session management endpoints and REST controllers for Familiar actions (ping, Esperanto spells, movement).
- **Phase 3 (GCP Alchemy & Cloud Run Trigger)**: Implement item interactions, IAM authorization checks, simulated Cloud SQL delay timer, and backend GCP Cloud Run deployment handler.
- **Phase 4 (Firestore Telemetry & Dual-Trail Replay)**: Store turn-by-turn state histories, implement path intersection detection algorithms, and build canvas particle fireworks animations.
- **Phase 5 (Global Leaderboard & Replay Viewer)**: Construct public leaderboard UI, interactive replay player, scrub controls, and workshop deployment documentation.

## Open Questions
- What is the best strategy for handling GCP service quota management during large workshops with 100+ simultaneous Cloud Run deployments?
- Should the AI Familiar be provided as a pre-built Python/Node.js starter script for participants, or driven optionally by a Gemini API agent client?
