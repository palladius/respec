---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-09-05T12:49:36Z"
model: gemini-pro-latest
tokens:
    prompt: 304
    output: 1218
    total: 3608
---

# TUI Vim Snake

## Problem Statement
Raspberry Pi users running headless Linux environments lack lightweight, terminal-native entertainment options that also help them build muscle memory for standard CLI tools. A terminal-based Snake game using Vim keybindings can provide a fun, nostalgic experience while subtly training users to navigate their keyboards using H, J, K, and L.

## Goals
- Create a TUI (Terminal User Interface) Snake game compatible with headless Linux (specifically Raspberry Pi).
- Use Vim keyboard bindings (H=left, J=down, K=up, L=right) for movement controls.
- Render food items using colored circle emojis: blue 🔵, red 🔴, yellow 🟡, green 🟢, and a super rare purple 🟣.
- Ensure a colorful, nostalgic visual style reminiscent of classic Nokia Snake but modernized with terminal emojis.
- Provide a standalone script easily runnable from a standard Linux terminal without complex GUI dependencies.

## Non-Goals
- Support for Windows or macOS (strict Linux/Raspberry Pi focus).
- Graphical User Interface (GUI) or mouse support.
- Arrow key or WASD control schemes.
- Online multiplayer or cloud-based leaderboards.
- Complex audio or sound effects.

## Technical Plan / Approach
**Language & Libraries:** The game will be written in Python using the built-in `curses` library, which is natively available, highly compatible with Linux terminal environments, and perfect for a Raspberry Pi.

**Game State & Loop:**
- **Grid System:** The playable area will be mapped to a terminal window coordinate system (rows and columns).
- **Snake Representation:** The snake's body will be stored as a list or `collections.deque` of `(y, x)` coordinate tuples. The head is the first element.
- **Game Loop:** A non-blocking `while` loop will manage the game tick. The tick speed will decrease (causing the game to speed up) as the snake grows.

**Emoji and Rendering:**
- The snake's body will be rendered using colorful ASCII blocks (e.g., `█`) utilizing curses color pairs to maintain a classic, high-contrast look.
- Food will be rendered as emojis. Because terminal emulators often treat emojis as double-width characters, the grid mapping will treat the X-axis carefully (e.g., advancing by 2 columns for horizontal movement if necessary, depending on standard terminal behavior).
- **Food Spawning Logic:** When food is eaten, a new coordinate is chosen randomly. The color is determined by a random number generator:
  - 92% chance to spawn a standard Google color (23% each for 🔵, 🔴, 🟡, 🟢).
  - 8% chance to spawn the rare purple ball (🟣).

**Input Handling:**
- `nodelay(True)` will be used on the curses window to ensure the game doesn't pause while waiting for user input.
- Key presses are mapped strictly to `h`, `j`, `k`, and `l` (accepting both lower and upper case variants).

## Alternatives Considered
- **C++ with ncurses:** Considered for maximum performance, but Python is more than fast enough for a terminal snake game and allows for much faster iteration, easier dependency management, and simpler script distribution on a Raspberry Pi.
- **Using WASD or Arrow Keys:** Explicitly rejected to enforce Vim keybinding training, which is a core feature of the request.
- **Standard ASCII food (e.g., `*` or `@`):** Rejected in favor of emojis to meet the specific colorful, modern-classic requirement outlined in the pitch.

## Implementation Plan
1. **Environment Setup:** Create the basic Python project structure and initialize the `curses` wrapper to handle terminal state cleanly (hiding the cursor, disabling input echo).
2. **Core Rendering:** Draw the game boundaries (walls) and implement a simple UI banner that prints the current score.
3. **Snake Mechanics:** Implement the snake state, initial drawing on the grid, and continuous movement logic based on a fixed tick rate.
4. **Vim Inputs:** Hook up `h`, `j`, `k`, and `l` to update the snake's current direction vector. Add logic to prevent the snake from reversing direction directly into itself.
5. **Food & Emoji Logic:** Implement the RNG for spawning the 🔵, 🔴, 🟡, 🟢, and 🟣 emojis. Implement spatial collision detection between the snake's head and the food coordinates.
6. **Scoring & Growth:** Grow the snake's tail when food is eaten. Assign 10 base points to standard Google-colored balls and 50 bonus points to the purple ball.
7. **Collision & Game Over:** Implement collision detection for the grid walls and the snake's own body. Build a simple "Game Over" screen that displays the final score and prompts to press 'q' to quit or 'r' to restart.

## Open Questions
- Should we enforce a fixed terminal board size (e.g., 40x20) and center it, or should we dynamically scale the game board to match the user's maximum terminal dimensions?
- Since older Raspberry Pi terminal emulators sometimes struggle with rendering double-width emoji characters correctly, should we include a fallback command-line flag (e.g., `--no-emoji`) that reverts the food to standard colored ASCII text?
