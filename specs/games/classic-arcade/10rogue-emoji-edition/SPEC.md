---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-02T14:37:11Z"
model: gemini-flash-latest
tokens:
    prompt: 289
    output: 1659
    total: 3125
---

# 10Rogue Emoji Edition

## Problem Statement

Classic roguelikes like 10rogue offer fast-paced, highly tactical dungeon crawling. However, many original implementations are platform-dependent (e.g., Windows executables like `10rogue.exe`), lack responsive mobile support, and rely on archaic ASCII or bland tilesets that don't appeal to casual modern web players.

There is a need for a modern, web-native 10rogue clone that retains the signature fast-paced, 10-floor progression and progressive fog-of-war discovery, while utilizing vibrant emoji graphics for rich visual representation (monsters, items, environments) without requiring heavy graphical asset downloads.

## Goals

- **Web Accessibility**: Playable immediately in any modern mobile or desktop browser with zero install and ultra-fast load times (minimal dependencies).
- **Visual Appeal**: High-quality rendering using expressive emojis (e.g., 👹, 🐉, 🗡️) coupled with fluid canvas-based lighting overlays.
- **True 10Rogue Light/FOV Mechanics**: Real-time line-of-sight (LOS) shadowcasting. Unexplored areas are pitch black, currently visible tiles are fully lit and animated, and previously explored but currently out-of-sight areas are dimmed/desaturated.
- **Universal Inputs**: 100% playable via keyboard (WASD, arrow keys, or Numpad) and seamlessly playable on mobile devices via a responsive, non-obtrusive virtual D-pad and tap-to-target system.
- **Progression Loop**: 10 distinct, procedurally generated dungeon levels scaling rapidly in difficulty, culminating in a final boss encounter.

## Non-Goals

- **Infinite Scaling/Endless Mode**: The scope is strictly locked to a tight, high-replayability 10-floor victory condition loop.
- **Multiplayer**: This is a purely single-player, turn-based experience.
- **Complex Asset Pipeline**: No loading of external sprite sheets, PNG packs, or 3D models. All visual elements must be rendered using standard Unicode Emojis and standard CSS/Canvas styling.
- **Persistent Database**: No server-side accounts or online multiplayer leaderboards. Local storage is sufficient for high scores.

## Technical Plan / Approach

### Architecture & Technology Stack
- **Framework**: Vanilla TypeScript/JavaScript with a lightweight rendering framework, or directly using HTML5 Canvas to keep bundle sizes under 100KB.
- **Rendering Layer**: HTML5 Canvas API. Canvas allows us to easily compute and draw custom alpha masks for the lighting effects, and offers superior performance for grid rendering over DOM-based solutions.
- **State Management**: A modular, single-directional state machine representing game phases (Main Menu, Exploration, Combat, Inventory, Game Over, Victory).

### Lighting & FOV (Field of View)
Each tile in the 2D grid map array will contain exploration states:
1. `unexplored` (binary flag: `false`): Rendered as pitch-black (`rgba(0,0,0,1)` overlay).
2. `visible` (dynamically computed each turn): Rendered in full color. Computed using a **Shadowcasting** or **Bresenham's Line** algorithm from the player's coordinate with a fixed radius (e.g., 5-6 tiles).
3. `remembered` (binary flag: `true`, but out of current FOV): Rendered with a semi-transparent dark overlay (`rgba(0,0,0,0.6)`) to signify it is mapped but out of sight. Emojis representing dynamic entities (monsters, moving hazards) are hidden in this state.

### Grid & Performance Optimization
To prevent emoji scaling and layout issues across different browsers:
- Render using a monospace grid layout calculated directly inside the canvas using font-metrics (e.g., `font = "24px Arial"` or system emoji fallback arrays).
- Scale the canvas buffer size dynamically to match the device pixel ratio (DPR) for high-DPI/Retina screens.

### Combat & Mechanics
- **Bumping Combat**: Attacking is executed by moving directly into an adjacent monster.
- **Turn-Based Tick**: Game time only progresses when the player acts (moves, uses an item, or waits a turn).
- **Progression**: Enemies scale from basic slimes and goblins on floor 1 to dragons and liches on floor 10.

## Alternatives Considered

- **DOM/CSS-Grid Rendering**: Representing each tile as an HTML element (like a `<div>`). While easier to construct, DOM-based approaches suffer from performance degradation when applying real-time dynamic lighting masks and scaling animations on mobile browsers.
- **Phaser.js / Pixi.js**: Third-party 2D rendering libraries. These would streamline rendering, but add 200KB-1MB of dependency overhead, violating the "minimal dependencies" core requirement.
- **ASCII Text-Only (no emojis)**: Authentic to classic Roguelikes, but text lacks immediate legibility and vibrant visual appeal for casual modern audiences compared to emojis.

## Implementation Plan

### Phase 1: Core Engine & Boilerplate (Days 1-2)
- Initialize build system (Vite + TypeScript).
- Construct responsive HTML5 Canvas container auto-scaling to desktop or mobile screen constraints.
- Implement keyboard input listeners and virtual mobile D-pad overlays.

### Phase 2: Generation & FOV (Days 3-4)
- Implement a Room-and-Corridor BSP (Binary Space Partitioning) dungeon generator.
- Build the Shadowcasting FOV algorithm.
- Program the three-state fog overlay system (`unexplored`, `visible`, `remembered`).

### Phase 3: Entities & Turn System (Days 5-6)
- Define Entity classes (Player, Monster, Item) with emoji graphics assigned.
- Implement turn-based state queue (Player turn -> Enemy AI turns -> Render loop).
- Implement basic pathfinding (A* or simple tracking) for hostile monsters.

### Phase 4: Systems, Scaling, & Balance (Days 7-8)
- Design standard RPG loop: Experience points, Level-ups, Health/Mana systems.
- Create item spawning logic (potions, weapons, armor) and quick-access inventory slots.
- Configure progressive monster spawning based on current floor level (1-10).

### Phase 5: Polish & Polish UI (Days 9-10)
- Add floaty combat text (damage numbers drifting upwards above hit entities).
- Integrate local-storage high score tracking.
- Test and optimize touch interaction latency and canvas frame rates on older mobile models.

## Open Questions

1. **Emoji Cross-Platform Consistency**: Emojis render differently on iOS, Android, and Windows. Should we embed a custom web font (like Twemoji) to guarantee uniform visuals, or accept platform-specific emoji rendering to keep bundle sizes absolute-minimum?
2. **Mobile UI Layout**: Does a portrait orientation (canvas on top, large d-pad and stats on bottom) offer the best user experience over landscape, or should we support responsive orientation swapping?
