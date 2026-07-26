---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-26T16:35:43Z"
model: gemini-flash-latest
tokens:
    prompt: 541
    output: 1307
    total: 2963
---

# BabelSaga: Esperanto Multi-Agent SimCity RPG

## Problem Statement
Multi-agent AI simulations often suffer from unconstrained communication, homogenous English-centric agent interactions, and a lack of clear game loops or deployment structures for developer hackathons. BabelSaga (Defaultia) solves this by combining a 2D emoji-based city-building RPG with strict linguistic constraints: English is forbidden, Esperanto serves as the universal lingua franca, and multi-agent droids communicate using a specialized protocol over Model Context Protocol (MCP).

## Goals
- Deliver a 2D emoji-grid SimCity RPG engine capable of rendering multi-namespace city instances provisioned via Terraform.
- Enforce Esperanto as the primary user command set (`/krii` for shout, `/flustri` for whisper, `/helpo` for help) and inter-player communication standard.
- Provide a starting set of 3 droids: active Translator Droid, active Gatherer Droid, and an inactive ADK / MCP droid customizable by the player and deployable to Cloud Run.
- Implement the `babot` / Omega13 JSON wire protocol utilizing `rot13`/`base64` obfuscated payloads and specialized status codes (`VABON`, `Aninpospio`).
- Implement core game mechanics: Universal Basic Income (100 Gemoj 💎/day), 7-item inventory system (including mandatory novelty/oversized items and an Esperanto rule scroll).
- Provide reusable IaC (Terraform) scripts to spin up city environments like *Mediolanum 3125* or *Alexandria* for hackathon sandboxes.

## Non-Goals
- Real-time 3D graphics or high-frame-rate action combat (focus is strictly on grid-based simulation and text/agent protocol interactions).
- Integration with external cryptocurrency or blockchain networks for Gemoj currency.
- Supporting raw English input in system interactions or unmediated LLM dialogue.

## Technical Plan / Approach
### Core System Architecture
1. **Frontend & Grid Engine**: Web-based tile grid renderer using WebGL or HTML Canvas, rendering terrain, player avatars, structures, and resource nodes as emoji sprites.
2. **Backend Engine & State Store**: Node.js/TypeScript or Go backend hosted on Google Cloud Run, backed by Cloud Firestore for persistent grid world state and Redis Pub/Sub for real-time player spatial position updates.
3. **MCP Integration**: Model Context Protocol (MCP) server endpoints exposed per city instance, allowing AI agents to query local tile vision, inventory, and neighboring droids.

### Language & Protocol Specifications
- **Esperanto Command Set**: All chat commands strictly match Esperanto syntax:
  - `/krii [mesaĝo]` - Broadcasts message across a 15-tile radius.
  - `/flustri [celo] [mesaĝo]` - Direct message to a nearby player or droid.
  - `/helpo` - Displays available actions and command descriptions in Esperanto.
- **`babot` / Omega13 Protocol**: Droid-to-droid messaging envelope:
  ```json
  {
    "sender": "droid_01",
    "recipient": "droid_02",
    "encoding": "rot13",
    "payload": "VABON_ORPBAR_IREM1",
    "status_code": "VABON"
  }
  ```
  Droids strictly refuse non-Esperanto direct queries unless challenged with `"Kian lingvon vi parolas?"`.

### Starter Droid Trio & Inventory Mechanics
- **Translator Droid (Active)**: Interfaces with MCP translation tool definitions to translate incoming foreign dialect payloads into Esperanto.
- **Gatherer Droid (Active)**: Automates pathfinding on the emoji grid to collect resources (e.g. 🌲, 💎, 🧱) and query tile attributes.
- **ADK Custom Droid (Inactive)**: Includes a starter Google Agent Development Kit (ADK) template. Players implement custom MCP tools and deploy to Google Cloud Run to activate.
- **7-Item Backpack**: Strict array length limit of 7 elements containing starting default items: Rule Scroll (`rulo.txt`), Translator Battery, Gatherer Tool, 3-Story Ladder (`triloga_stetaro`), and initial 100 Gemoj stipend.

## Alternatives Considered
- **Standard REST API for Agent Communication**: Rejected in favor of MCP (Model Context Protocol) to ensure standardized tool-calling interfaces between custom LLM agents and the city grid.
- **Standard English Command Set**: Rejected as the primary core mechanic relies on language discovery and restricted communication protocols via Esperanto.

## Implementation Plan
- **Phase 1: Foundation & Grid Engine**: Implement 2D emoji grid renderer, spatial movement backend, and 7-slot inventory system.
- **Phase 2: Esperanto Command & Protocol Middleware**: Build command parser (`/krii`, `/flustri`, `/helpo`) and `babot` JSON encoder/decoder with `rot13`/`base64` processing.
- **Phase 3: Droid Lifecycle & MCP Server**: Deploy built-in Translator and Gatherer droids; expose MCP server interface for custom ADK Cloud Run droid attachments.
- **Phase 4: Economy & IaC Sandbox Templates**: Add daily 100 Gemoj distribution cron job and write Terraform recipes for provisioning custom city instances (*Mediolanum 3125*, *Alexandria*).

## Open Questions
- What is the max player/droid concurrency limit per Cloud Run instance before grid state pub/sub requires sharding?
- How should rate-limiting be enforced for broadcast command (`/krii`) when multiple LLM-backed active droids generate simultaneous responses?
