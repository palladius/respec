---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-02T14:11:13Z"
model: gemini-flash-latest
tokens:
    prompt: 253
    output: 2010
    total: 3562
---

# p5.js Pac-Man Arcade Clone

## Problem Statement

Recreating the classic 1980 arcade game *Pac-Man* in a modern web environment requires accurate replication of its core mechanics: grid-aligned movement, responsive input buffering, precise ghost behaviors, and state-driven game phases (Chase, Scatter, and Frightened). While many clones simplify these mechanics, a true-to-the-original recreation requires implementing the specific pathfinding rules and distinct AI personalities that made the original game engaging.

This project must be entirely client-side, run directly in a web browser without a backend, and leverage the `p5.js` library for rendering, game loops, and input management.

## Goals

*   **Classic Maze Representation**: Implement the standard 28x36 tile layout, including appropriate corridors, the ghost house, and side escape tunnels.
*   **Accurate Ghost AI**: Replicate the four unique ghost behaviors:
    *   **Blinky (Red)**: Direct shadow; aggressively targets Pac-Man's exact tile.
    *   **Pinky (Pink)**: Ambusher; targets four tiles ahead of Pac-Man's current direction.
    *   **Inky (Cyan)**: Fickle/Tactician; targets a vector based on both Blinky's position and Pac-Man's position.
    *   **Clyde (Orange)**: Feign-coward; targets Pac-Man if far away, but retreats to his scatter corner if within an 8-tile radius.
*   **Global Timer Cycles**: Alternate game phases between Chase and Scatter modes based on classic arcade level timings.
*   **Smooth Input Buffering**: Allow players to queue their next direction change before reaching an intersection, ensuring smooth navigation through tight corridors.
*   **Power Pellets and Frightened Mode**: Eating a power pellet turns ghosts blue, reduces their speed, reverses their direction, and allows Pac-Man to consume them for cascading point multipliers (200, 400, 800, 1600).
*   **Zero-Dependency Deployment**: Run entirely client-side using a single index.html file loading `p5.js` from a CDN, utilizing Web Audio API for synthetic retro sound effects to avoid asset-loading issues.

## Non-Goals

*   **Perfect Level-256 Glitch Emulation**: We will not emulate the original Z80 memory overflow crash at level 256; levels will loop safely.
*   **Online Leaderboards**: High scores will be tracked locally in the browser's `localStorage` instead of using a remote database.
*   **Complex Pixel-Art Assets**: To maintain a lightweight footprint and ensure reliability, visuals will be rendered procedurally using 2D canvas shapes (arcs, lines, circles) rather than loading external sprite sheets.

## Technical Plan / Approach

### 1. Rendering Engine and Coordinate System
*   **Framework**: `p5.js` will handle the game canvas lifecycle (`setup()` and `draw()`).
*   **Grid System**: The game board will be modeled as a grid of 28 cols by 36 rows. Each cell is $W \times W$ pixels (e.g., $16 \times 16$ or $20 \times 20$ based on window dimensions).
*   **Entity Positions**: Characters (Pac-Man and Ghosts) will have continuous floating-point coordinates $(x, y)$ to allow for smooth interpolation between tiles, but their pathfinding decisions will trigger exclusively when their integer position aligns exactly with the center of a tile coordinate.

### 2. Pac-Man Movement & Input Buffering
*   An active direction vector and a buffered/intended direction vector will be maintained.
*   If the user presses an arrow key, the intended direction is cached. 
*   When Pac-Man approaches the center of a tile, the game checks if the intended direction is free of obstacles. If so, it becomes the active direction; otherwise, Pac-Man continues in his current direction until hitting a wall.

### 3. Ghost Pathfinding Engine
*   **Target Selection**: Each ghost calculates a target tile $(T_x, T_y)$ every frame based on their current state (Chase, Scatter, Frightened, Eaten).
*   **Decision Points**: At every intersection, a ghost looks ahead to the next tile. It evaluates the neighboring tiles (excluding the direction it just came from—ghosts cannot turn back unless a state change forces a 180-degree reversal).
*   **Distance Metric**: The ghost selects the neighboring tile that has the shortest straight-line (Euclidean) distance to its target tile:
    $$\text{Distance}^2 = (N_x - T_x)^2 + (N_y - T_y)^2$$
*   **Frightened State**: Ghosts generate pseudo-random path decisions at each intersection instead of targeting Pac-Man.
*   **Eaten State**: Once eaten, the ghost becomes a pair of eyes and sets its target tile to the ghost house entrance. Upon arrival, it revives and resumes normal behavior.

### 4. Audio Generation
*   To keep the codebase asset-free, we will use the browser's native `AudioContext` to synthesize classic arcade sounds (the low/high "waka-waka" siren, the power pellet siren, and death fanfares) using sine, triangle, and square oscillators with quick frequency sweeps.

## Alternatives Considered

*   **Phaser 3 Engine**: Considered for its comprehensive tilemap API and arcade physics system. However, it requires a local web server to bypass CORS issues when loading assets, and its bundle size is unnecessarily heavy for a retro clone. `p5.js` offers a more lightweight approach where we can procedurally render elements dynamically.
*   **Vanilla Canvas API**: While extremely lightweight, vanilla canvas lacks the convenient API sugar of `p5.js` for vector operations, scale handling, keystroke event management, and structured state cycles.
*   **Pre-rendered Sprite Sheets**: While visually identical to the original arcade cabinet, downloading sprite sheets introduces CORS bottlenecks and asset management overhead. Procedural rendering ensures instantaneous, offline-ready initialization.

## Implementation Plan

### Phase 1: Grid Setup & Map Rendering
*   Create a 2D array representation of the 28x36 board containing walls, pellets, and power pellets.
*   Write a rendering routine in `p5.js` to draw the blue double-lined walls, small yellow pellets, and flashing power pellets.

### Phase 2: Pac-Man Entity & Movement
*   Implement the Pac-Man entity with properties for tile position, pixel offset, current speed, and direction.
*   Code the input handler using `keyPressed()` to buffer input.
*   Add collision checking against wall grid tiles.
*   Implement pellet eating logic: updating score, removing pellets from the 2D array, and playing a synthesized procedural audio chirp.

### Phase 3: Ghost AI Foundation
*   Implement the base `Ghost` class with physical movement mechanics.
*   Create the pathfinding logic: identifying intersections, preventing backward movements, calculating Euclidean distances, and updating position smoothly towards target tiles.
*   Draw simple procedural ghosts with shifting eyes to indicate target directions.

### Phase 4: Unique AI & Global Timers
*   Code individual subclasses or target-evaluation functions for Blinky, Pinky, Inky, and Clyde.
*   Implement the global game clock to transition ghosts between Chase and Scatter modes (e.g., Scatter 7s, Chase 20s, Scatter 7s, Chase 20s...).

### Phase 5: Power Pellets & Collision Loop
*   Add the Frightened state toggling logic upon Pac-Man eating a power pellet.
*   Implement collision detection between Pac-Man and ghosts. If a ghost is in Frightened state, transition it to Eaten state and award points. If normal, trigger Pac-Man's death sequence (subtract life, reset positions, pause game momentarily).

### Phase 6: Game States, Level Loops & UI Overlay
*   Build the main menu, game over screen, and high score board using standard HTML overlays or text drawn inside the `p5.js` canvas.
*   Manage Level Clear conditions: reset map layout, increment speed slightly, and recreate pellets.

## Open Questions

*   **Classic Corner-Cutting (Speed-up Cornering)**: In the original arcade game, Pac-Man moves slightly faster when turning corners perfectly. Should we emulate this sub-pixel shifting, or is strict tile-centered path navigation sufficient for an enjoyable browser-level feel?
*   **Synthesized Sound Consistency**: Will native Web Audio API oscillators sound sufficiently authentic across different browsers (specifically Safari vs. Chrome)? A safe fallback of muted play or simple gain curves will be coded in.
