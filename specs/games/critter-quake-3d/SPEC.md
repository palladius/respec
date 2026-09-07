---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-09-07T06:13:00Z"
model: gemini-3.8-flash
tokens:
    prompt: 352
    output: 2063
    total: 3277
---

# Critter Quake: Kid-Friendly Retro 3D FPS for Web and Terminal

## Problem Statement
Classic 3D retro first-person movement shooters (such as Quake) offer fluid navigation, vertical arena design, and secrets, but their aesthetics are predominantly grim, dark, and gore-heavy—making them inappropriate for younger audiences (ages 6–8). Simultaneously, modern games rarely bridge the gap between rich graphical platforms (Web/WebGL) and lightweight, text-mode environments (raw Linux terminal without an X server). Furthermore, procedural and editable level pipelines often lack rigorous, automated validation (such as reachability and spanning-tree traversability checks) to guarantee that procedural secrets and objectives can actually be reached without softlocks.

Critter Quake solves this by delivering an energetic, bright, non-violent 3D arena-action game rendered both in modern web browsers and directly inside ANSI/VT terminal emulators, paired with a deterministic level validator and headless simulation harness.

## Goals
- **Target Demographics & Theme**: Visually appealing, vibrant aesthetic tailored for 6–8 year olds. Features colorful animal allies and comedic monsters, saturated primary lights (Google-palette: red, blue, green, yellow), and non-violent, funny weaponry (e.g., Slime Splatters, Bubble Blasters, Carrot Launchers, Confetti Poppers).
- **Hybrid Quake & Minecraft Mechanics**: Snappy 3D movement (strafe running, bounce pads, ramps) combined with editable blocky/voxel geometric arenas, breakable blocks, and hidden secret alcoves.
- **Dual Rendering Target**:
  - Web (WebGL/WebGPU) with smooth lighting, particle bursts, and cartoon-shaded assets.
  - CLI/Terminal mode (headless Linux / raw TTY) via ANSI 24-bit truecolor half-block (`▀`) rasterization at playable framerates (>= 20 FPS on standard modern terminals).
- **Editable Persistent Schemas**: Levels generated initially from seeds or schemas, editable by players in-game, with state serialization (save/load) supported in both web local storage and CLI flat files.
- **Automated Validation & Test Engine**:
  - Headless execution engine that computes level spanning trees, ensures reachability of all keys, objectives, and secrets, and detects navigation dead ends.
  - Automated asset verification (missing textures, geometry manifold checks, bounding box collisions).

## Non-Goals
- Realistic physics simulation, bullet drop, or realistic ballistic damage models.
- Gore, dismemberment, or realistic firearms.
- Real-time networked multiplayer (multiplayer is deferred to post-v1; initial focus is polished single-player/offline play).
- In-browser complex 3D CAD level modeling; user editing is restricted to Minecraft-style placing and removing discrete voxel blocks.

## Technical Plan / Approach

### 1. Engine Architecture & Core Language
The core engine will be implemented in **Rust**, compiled to two standalone targets:
1. **Web (WASM + WebGL2/WebGPU)** via `wasm-bindgen` and `wgpu`.
2. **Native CLI Binary** compiled for Linux/macOS/Windows, outputting directly to stdout using raw terminal mode (`termios` / `crossterm`) with a software rasterizer.

Shared core modules:
- `cq-math`: Vector3, AABB, matrix transformations, raycasting.
- `cq-world`: Voxel chunks, level schema representation, lighting grid, entity component system (ECS via `hecs`).
- `cq-physics`: Simplified Quake movement (accelerate, friction, air-control, jump, bounce pad impulse) with swept AABB voxel collision.
- `cq-save`: Portable JSON / CBOR level and player state serialization.
- `cq-validator`: Headless simulation graph, reachability spanning tree generator, and playability agent.

### 2. Dual-Mode Rendering Pipeline
- **Web Renderer**: Leverages WebGL2/WebGPU. Renders stylized voxel meshes using a cartoon cel-shader, dynamic point lights (red, green, blue, yellow light falloff), squash-and-stretch procedural skeletal animations, and particle emitters for splats/bubbles/confetti.
- **CLI Software Renderer**: Downsamples rendering to a terminal resolution (e.g., 80x48 or 160x96 cells). Utilizes a custom software Z-buffer rasterizer rendering into a dual-pixel vertical buffer printed via ANSI truecolor escape codes using the unicode half-block character `▀` (where foreground color sets the top half-pixel and background color sets the bottom half-pixel). Audio in CLI is mapped to raw ALSA/PulseAudio on Linux or terminal bell beeps as a fallback.

### 3. Weapons and Entity System
All combat interactions use non-violent state machines:
- **Bubble Blaster**: Fires floating orbs that trap monsters in place; jumping on trapped monsters acts as a temporary bounce pad.
- **Slime Launcher**: Coats surfaces and enemies in slippery, colorful goo, altering physics (friction reduction) and slowing monsters.
- **Carrot Rocket**: Explodes in a spray of confetti and carrots; uncovers hidden cracked walls.
- **Squeaky Hammer**: Close-range tool used for both editing (destroying blocks) and bopping monsters into a dizzy, spinning state.

### 4. Level Schemas and Spanning-Tree Validation
Levels are defined as 3D grid volumes divided into 16x16x16 chunks. 
- **Schema Generation**: Seeded procedural generation constructs rooms, ramps, elevation corridors, locked toy doors, and secret toy chests.
- **Level Graph Construction**: The `cq-validator` constructs a 3D reachability graph where nodes are reachable voxel clusters and edges are traversal transitions (walk, standard jump, rocket/bounce-pad boost).
- **Spanning Tree & Softlock Detection**: A Dijkstra/A* graph traversal simulates virtual player navigation with bounded abilities (jump height, speed). The level is only stamped valid if:
  1. The path from Spawn to Victory is fully connected.
  2. Every key item (e.g., yellow star key) is reachable prior to its corresponding door.
  3. Secret areas possess at least one entrance and cannot irrevocably trap the player without exit options (no zero-jump pits without ladders/pads).

### 5. Automated CI Test Suite
- **Asset Integrity**: Headless test suite parses all JSON schemas, voxel models, sound effects, and ensures bounding boxes match visual boundaries.
- **Playability Regression**: Runs headless bots on 1,000 deterministic seeds in CI; fails the build if any seed generates an unfinishable level or an unreachable secret.
- **Deployment**: Automatic GitHub Actions workflow building WASM artifacts hosted via GitHub Pages and native cross-compiled CLI binaries published as GitHub Releases.

## Alternatives Considered
- **Pure C + Ncurses**: Classic approach for terminal games, but lacks modern memory safety, requires separate porting for Web (Emscripten overhead), and ncurses overhead limits 60 FPS full-color rasterization.
- **Three.js + Node CLI**: Building with JavaScript/TypeScript allows easy web development, but running high-performance software 3D rendering in pure Node.js on a headless Linux CLI yields poor CPU cache utilization and frame drops compared to Rust.
- **Raycasting (Wolfenstein 3D style) instead of True 3D**: Raycasting limits vertical room-over-room layouts, bounce pads, and jumping. True voxel polygon rasterization enables authentic Quake-like vertical movement and Minecraft-like editing.

## Implementation Plan
- **Phase 1 (Core Engine & Math)**: Implement Rust core math, swept AABB physics, Quake movement physics (air-strafe and bounce pads), and voxel chunk memory models.
- **Phase 2 (CLI Software Rasterizer)**: Build native ANSI half-block 3D renderer with Z-buffering and dynamic colored point lights. Connect keyboard input via raw TTY.
- **Phase 3 (Web Port & Visual Assets)**: Compile core to WASM. Implement WebGL renderer, load colorful animal models, comedic monsters, squash-and-stretch animation controllers, and light management.
- **Phase 4 (Weapons, Audio, & Editing)**: Implement the 4 non-violent weapons, block placement/destruction mode, cartoon sound synthesis, and local save/load serialization.
- **Phase 5 (Automated Test Engine & CI)**: Build the graph-based spanning tree reachability tester, headless bot simulation, automated schema validator, and CI release pipeline.

## Open Questions
- **CLI Frame Budget on Slow Terminals**: While modern terminals (Alacritty, Kitty, WezTerm) handle 60 FPS truecolor ANSI output easily, standard Linux virtual consoles (`/dev/tty1`) and older emulators can choke on massive I/O escape sequences. Should we implement an adaptive frame dropper or lower-resolution fallback specifically for raw VT virtual consoles?
- **Mouse Look in CLI**: Raw terminals do not natively capture relative mouse deltas like Web Pointer Lock. Should the CLI version default to keyboard-only look (Doom/classic Duke Nukem style with arrow keys/numpad) or utilize ANSI SGR mouse tracking?
