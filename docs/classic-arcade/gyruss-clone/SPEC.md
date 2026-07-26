---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-02T14:09:55Z"
model: gemini-flash-latest
tokens:
    prompt: 267
    output: 1645
    total: 3340
---

# Gyruss Clone (p5.js Tube Shooter)

## Problem Statement

Classic arcade games like *Gyruss* (1983) offered a unique pseudo-3D perspective through "tube-shooting" mechanics. Unlike traditional vertical or horizontal scrollers, the player moves in a circular orbit around the outer edge of the screen, shooting inward toward the center, while enemies spawn from the center and spiral outward.

Developing a browser-accessible clone of this classic game requires solving several spatial and computational challenges: translating polar coordinates to Cartesian space, simulating three-dimensional depth scaling with 2D rendering APIs, handling complex enemy curving paths, and ensuring smooth performance on a variety of client devices without relying on backend infrastructure.

## Goals

* **Fully Client-Side Architecture**: Build the game to run entirely in the browser using static HTML5, CSS, and `p5.js` with no backend database or server runtime required.
* **Faithful Tube-Shooter Mechanics**: Implement 360-degree circular movement for the player ship, inward-directed projectile physics, and enemies scaling up as they travel outward from a central horizon point.
* **Retro Vector Aesthetics**: Use stylized vector-based shapes, wireframes, and particle effects in `p5.js` to capture the classic 1980s neon/arcade look without requiring large external image assets.
* **Smooth 60 FPS Performance**: Maintain fluid gameplay loop, input handling, and collision detection for dozens of on-screen entities.
* **Interactive Audio**: Generate sound effects (laser fire, explosions, wave start) procedurally using the Web Audio API or standard oscillators to avoid loading external audio files.

## Non-Goals

* **1-to-1 Rom Emulation**: The game is a *clone* designed to replicate core gameplay feel and aesthetics, not a perfect cycle-accurate emulation of the original arcade hardware memory and bugs.
* **Backend Multiplayer/Leaderboards**: High scores will be stored locally in the browser's `localStorage` rather than synced to a global database server.
* **3D Engine Usage**: Standard 3D engines (e.g., Three.js) will not be used. Depth simulation must be achieved mathematically via 2.5D projection inside a 2D canvas context.

## Technical Plan / Approach

### Game Engine and Rendering

The game will be built using the `p5.js` library in global/instance mode. The canvas coordinates will be offset to center $(0,0)$ in the middle of the viewport using `translate(width/2, height/2)` to simplify radial math.

### Coordinate Systems and Movement Physics

Rather than storing positions solely in standard $(x, y)$ Cartesian coordinates, entities (Player, Bullets, Enemies) will be modeled using polar coordinates: $(\theta, r)$ where:
* $\theta$ is the angle in radians $[0, 2\pi]$.
* $r$ is the radial distance from the screen center (origin).

Conversion to Cartesian coordinates for rendering and physical distance checks is calculated as:
* $x = r \cdot \cos(\theta)$
* $y = r \cdot \sin(\theta)$

#### Player Ship
* Constrained to a fixed maximum radius $R_{player}$ near the screen border.
* Movement is controlled by updating $\theta$ left or right based on Arrow/A-D keys.

#### Bullets
* Spawn at $R_{player}$ with angle $\theta_{player}$.
* Move inward by decreasing $r$ towards $0$. Bullets disappear when $r \le R_{threshold}$.

#### Enemies
* Spawn at a near-zero radius $r \approx 0$ with a scale factor of $0$.
* Travel along calculated parametric paths (e.g., spiral trajectories $r(\theta) = a + b\theta$ or bezier curves) expanding towards the outer perimeter.
* Scale linearly or exponentially with $r$ to simulate a 3D scaling effect: $\text{scale} = r / R_{player}$.

### Starfield Simulation

To simulate moving forward through space, a 3D starfield effect will be rendered. Star objects will store $(\theta, r, v_r)$ where velocity $v_r$ represents outward speed. As stars reach $r > R_{max}$, they reset to $r = 0$ with a randomized trajectory angle $\theta$.

### Audio Generation

To ensure zero asset-loading overhead, a lightweight synthesizer class using the browser's native `AudioContext` will be constructed. Simple tone sweeps (pitch slides) will generate laser firing, ship explosions, and level-transition fanfares.

## Alternatives Considered

* **Using Three.js**: While Three.js would handle depth and perspective automatically, it increases initial bundle size and introduces unnecessary overhead for a game that relies on simple concentric circles and flat scaling sprites.
* **Vanilla Canvas API**: While functional, using native 2D Canvas requires verbose boilerplate for coordinate transformations, rotation, keyboard event hooks, and loop timings. `p5.js` provides excellent mathematical utilities (e.g., `p5.Vector`) and intuitive drawing operations that speed up game creation.

## Implementation Plan

### Phase 1: Engine Initialization & Starfield
* Create a static `index.html` loading `p5.js` via CDN.
* Implement the core game loop structure: `setup()`, `draw()`, and window-resize handlers.
* Build the dynamic outward-radiating 3D starfield to establish depth.

### Phase 2: Player Controller & Bullet Physics
* Render the player ship on the outer circular ring.
* Map keyboard inputs to circular rotation along the ring with acceleration and friction coefficients.
* Implement firing mechanics, generating laser projectiles that travel inward from the ship's current angle towards the origin $(0,0)$.

### Phase 3: Enemy Spawning and Movement Curves
* Implement the enemy class utilizing polar-based coordinate calculations.
* Design wave paths: enemies must spawn tiny near the center, travel outwards along spiral trajectories, and temporarily enter circular orbits near the player.
* Render vector-line enemy shapes that scale in size relative to their current radius.

### Phase 4: Collisions and Game State Management
* Implement bounding circle collision checks in Cartesian space: $d = \sqrt{(x_1 - x_2)^2 + (y_1 - y_2)^2}$.
* Establish a basic state machine: Title Screen, Playing State, Player Death, Game Over, and Level Cleared.
* Add scoring, life tracking, and progressive difficulty settings (faster enemies, more frequent firing).

### Phase 5: Audio, Polish, and Particle Effects
* Write a simple Synthesizer utility using Web Audio API oscillators for sound effects.
* Add particle explosion effects when enemies or player ships are destroyed.
* Integrate high-score persistence using the browser's local storage.

## Open Questions

* **Collision Accuracy**: Does checking collisions in Cartesian coordinates cause artifacts for objects moving rapidly towards or away from the center? *Answer: Standard Cartesian distance checks remain mathematically correct, but bullet step-size must be tuned so fast-moving projectiles do not skip past enemy boundaries.*
