---
speck_version: "0.1"
mode: interactive
title: "Pupurabbu 3D: Quake-Kids OSS Specification"
author: "Riccardo Carlesso (Ricchi1) & Antigravity (Pickle Rick Engine)"
created_at: "2026-09-11T08:27:00Z"
target_platforms:
  - Linux (Ubuntu/Debian)
  - macOS
  - Mobile (iOS / Android)
engine: "TypeScript + Three.js + Cannon-es + Vite"
---

# SPEC-01: Pupurabbu 3D (Quake-Kids OSS)

## 1. Executive Summary & Vision

**Pupurabbu 3D** is an open-source, kid-friendly 3D first-person/third-person shooter and world explorer engineered specifically for **Alessandro (age 8)**, **Sebastian (age 6)**, **Riccardo**, and family guests. 

It fuses the responsive, high-speed movement and weapon feel of classic retro shooters (*Quake*, *Doom*) with the playful charm of Nintendo classics (*Super Mario*, *Splatoon*), high-tech Google energy spheres, comic food power-ups (levitating Salmon Nigiri 🍣), and a trading-card collection victory condition inspired by *Pokémon*, *Magic: The Gathering*, and *Saint Seiya* (Gold Saint armored fluffy animals).

The architecture is strictly multi-platform, targeting 30+ FPS (smooth, battery-efficient for kids playing on iPads, phones, and laptops) in standard modern web browsers on desktop (Mac/Linux) and mobile devices (iPad, iPhone, Android), installable as a Progressive Web App (PWA) or wrapped with Capacitor/Electron.

---

## 2. Target Platforms & Input Architecture

### 2.1 Hardware & Runtime Targets
* **Linux & macOS**: Chromium-based browsers, Safari, Firefox, or Electron shell.
* **Mobile (iOS / iPadOS / Android)**: Safari Mobile & Chrome Mobile with full touch optimization.
* **Local LAN Zero-Friction Hosting**: The host machine (`pupurabbux`) runs the Vite development/preview server; family members connect directly over Wi-Fi (`http://pupurabbux.local:5173`).

### 2.2 Multi-Input Scheme
The engine provides an abstract `InputManager` supporting four concurrent control methods:
1. **Desktop Mouse + Keyboard**:
   * Mouse: Smooth 3D look with Pointer Lock API.
   * `W, A, S, D` / Arrows: Move and strafe.
   * `Space`: Jump & Mario Double-Jump.
   * `Hold Space / F`: Anti-Gravity Flight (when Blue Sphere is active).
   * `Left Mouse Button`: Primary fire.
   * `Q / E` or `1 - 4`: Blue Shirt Guy weapon cycle.
   * `V`: Camera toggle (1st-Person $\leftrightarrow$ 3rd-Person).
2. **Gamepad (Xbox / DualSense / Bluetooth)**:
   * Left Stick: Directional movement.
   * Right Stick: Smooth camera orbit/look.
   * Button `A`: Jump & Double-Jump.
   * Button `RT` (Right Trigger): Primary fire.
   * Button `Y` / `LB`: Flight thruster.
   * Button `RB`: Next weapon.
   * Button `Back / Select`: Camera perspective toggle.
3. **Mobile Dual-Touch HUD (Diablo Style)**:
   * Left Thumb: Virtual floating analog joystick / 4-way D-pad.
   * Right Screen Area: Drag swipe touchpad for looking around.
   * Action Buttons:
     * 🍌 **Shoot**
     * 🦘 **Double Jump**
     * 🔄 **Weapon Cycle**
     * 🚀 **Flight Thruster**
     * 📷 **1P/3P Camera Toggle**
4. **AI Bot Virtual Input**:
   * Programmatic state injection without DOM dependencies for automated testing and headless simulations.

---

## 3. Camera System: Hybrid 1P / 3P Perspective

* **First-Person (1P)**:
  * Default Quake-style view with animated weapon mesh held in lower right.
  * Direct crosshair alignment for precision aiming.
* **Third-Person (3P Over-The-Shoulder)**:
  * Spherical camera arm following 2.5 meters behind and 0.8 meters above the player.
  * Essential for 6-year-olds to maintain spatial orientation and visually enjoy their hero costume and trailing companion (Pupurabbu, Dragon, Lobster).
  * Raycast occlusion test pushes camera forward if walls intervene.
* **Toggle**: Smooth instant switch via `V` key or on-screen camera icon.

---

## 4. Playable Characters & Companions

Players choose their hero at the start screen:

| Hero | Icon | Trailing Companion | Companion Role & Lore |
| :--- | :--- | :--- | :--- |
| **Alessandro** (Ale) | 🐉 | **Emerald Baby Dragon** | Flaps wings, sniffs out hidden Dragon Gems, does mid-air somersaults. |
| **Sebastian** (Seby) | 🧸 | **Pupurabbu** | The mythical giant gentle purple furry monster. Hugs enemies away, points to Salmon Nigiri. |
| **Riccardo** (Dad) | 🦞 | **Carlobster** | Agile red lobster drone with espresso jets and snappy claws. |
| **Guest Heroine** | ⭐ | **Star Sprite** | Cosmic glowing star companion radiating soft starlight. |

### Companion Physics & Mechanics
* Trailing tether using spring-damper physics (`F = -k * x - c * v`).
* Idle animation: Gentle vertical sinusoidal bobbing.
* Wayfinding Assistance: If the player remains motionless for $> 5$ seconds, the companion floats forward and gestures towards the nearest uncollected card or objective.
* Power-up Interaction: If the player has 100% full health, walking into a Salmon Nigiri lets the companion eat it and trigger a cheer animation.

---

## 5. Movement & Core Physics

* **Gravity & Snap**: Tuned to $-22\ \text{m/s}^2$ for snappy, responsive arcade jumping without floatiness.
* **Mario Double-Jump**:
  * Jump 1: Launches from ground with $+9.0\ \text{m/s}$ impulse.
  * Jump 2: Usable anytime in mid-air before landing; fires a $+8.5\ \text{m/s}$ boost with an expanding cyan energy ring effect.
  * Counter resets immediately upon touching any walkable surface.
* **Anti-Gravity Flight Mode**:
  * Activated by collecting a **Blue Google Sphere**.
  * Grants a 10-second flight gauge displayed on the HUD.
  * Holding Fly activates vertical thrust ($+18\ \text{m/s}^2$ acceleration up to terminal velocity of $10\ \text{m/s}$).
  * Drains fuel gauge over 10 seconds.
* **Ramps & Step Climbing**: Collision sphere physics naturally glides over stairs and 45-degree ramps without snagging.

---

## 6. Combat, Defeat & Death Protocol

### 6.1 Visceral Enemy Defeat ("Voxel Gib Splat")
* Enemies have distinct HP pools (40–100 HP).
* When reduced to $\le 0$ HP, the enemy detonates into **20–30 colorful 3D voxel cubes**:
  * Voxels inherit explosive velocity in a spherical hemisphere with gravity and bounce restitution.
  * Accompanied by retro arcade pop/splat sound effects and floating confetti stars.
  * Voxels shrink and despawn over 3 seconds.
  * Satisfies the Quake-style visceral explosion feel while maintaining playful cartoon aesthetics.

### 6.2 Player "Death" & Tantrum-Proof Respawn
* When player HP reaches 0:
  * The camera stays upright while the player avatar spins 360 degrees and bursts into stars.
  * Comedic sound effect plays (e.g. funny spring boing / cartoon slip).
  * The player **instantly respawns** at the nearest checkpoint or level spawn pad.
  * **Zero loss of collected gems, nigiri, or cards**.
  * Grants a **3-second Green Shield ("Pupurabbu Hug")** with invulnerability to prevent spawn camping.

---

## 7. Weapon Arsenal & Schema-Dependent Loadouts

Weapons are cycled via the *Blue Shirt Guy* weapon selector (`Q / E`, numbers `1 - 4`, or wheel):

```
[ 1: 🍌 Banana Launcher ] ──> [ 2: 💎 Gem Blaster ] ──> [ 3: 🏎️ Hotwheels Roller ]
```

1. 🍌 **Banana Launcher**:
   * *Type*: Curved projectile with parabolic gravity.
   * *Damage*: 35 HP.
   * *Special Effect*: Bounces up to 2 times. If an enemy steps on a resting banana peel, they cartoon-slip, spin 360 degrees, and are stunned for 2 seconds.
   * *Ammo*: Infinite.
2. 💎 **Gem Blaster**:
   * *Type*: High-velocity straight beam crystal.
   * *Damage*: 25 HP.
   * *Special Effect*: Pinball ricochet! Reflects off walls up to 3 times before shattering.
   * *Ammo*: Infinite.
3. 🏎️ **Hotwheels Roller**:
   * *Type*: Heavy ground-hugging kinetic rocket car.
   * *Damage*: 60 HP.
   * *Special Effect*: Follows ground contours, rushes up ramps, and causes high knockback.
4. **Schema-Specific Weapon Customization**:
   * Individual schemas (levels) can swap or restrict weapons (e.g., Savanna Level equips a Watermelon Mortar; Toybox Level equips Hotwheels Roller).

---

## 8. Food & Google Power-Up Spheres

### 8.1 🍣 Abundant Levitating Salmon Nigiri
* Plentiful throughout the schema: placed on ramps, floating over platforms, hidden inside burrows, and dropped by defeated mini-bosses. They are consumable food items available in large numbers, NOT a single rare trophy!
* Restores **+35 HP**.
* Grants **"Wasabi Speed Rush"** for 5 seconds:
  * Player movement speed increased by $1.5\times$.
  * Fiery orange particle trail behind player footsteps.
* Companion Feeder: If player HP is 100%, collecting Nigiri feeds the companion, giving $+100$ bonus score.

### 8.2 Google Energy Spheres (4 Primary Colors)
* 🔴 **Google Red (Heart / Hyper Speed)**:
  * Instantly restores $+50$ HP.
  * Grants double movement speed for 10 seconds.
* 🔵 **Google Blue (Anti-Gravity Flight)**:
  * Refills Flight Jetpack Fuel to 100%.
  * Grants 10 seconds of free vertical flight and aerial hovering.
* 🟢 **Google Green (Pupurabbu Shield / Hug)**:
  * Surrounds player in a glowing emerald bubble for 12 seconds.
  * Absorbs 100% of enemy contact damage.
  * Enemies that bump into the bubble get knocked backward with a boing sound.
* 🟡 **Google Yellow (Banana Overdrive)**:
  * Weapon fire rate doubled for 10 seconds.
  * Fires a **3-banana spread shot** per click without cooldown.

---

## 9. Collectible Trading Cards Victory Condition

### 9.1 The Schema Booster Set System
In addition to eliminating enemies, **clearing a schema requires collecting the entire booster set of trading cards** hidden across the level.

* Each level has a numbered deck (e.g., Cards #1 to #8).
* Cards float in 3D space like holographic foils, rotating with sparkling particle borders.
* Touching a card displays a celebratory card-reveal card overlay on screen ("NEW CARD UNLOCKED!").

### 9.2 Card Rarity Tiers & Saint Seiya Mythic Beasts
Cards follow collectible trading card game tiers (Common $\rightarrow$ Rare $\rightarrow$ Mythic Gold):

1. **Common Cards (Bronze Border)**:
   * Local wildlife and objects (e.g., #1 Candy Mushroom, #2 Guinea Pig Burrow, #3 Hotwheels Buggy).
2. **Rare Cards (Silver Holographic)**:
   * Special power items (#4 Wasabi Nigiri Bento, #5 Google Prism Core).
3. **Mythic Gold Cards: Saint Seiya Fluffy Animals (Gold Armored)**:
   * Legendary companions adorned in gleaming golden celestial cloth armor with flaming cosmos auras!
   * **#6 Pegasus Guinea Pig (Gold Saint of the Burrow)** 🐹✨
   * **#7 Dragon Pupurabbu (Gold Cloth of Mount Lu)** 🧸🐉
   * **#8 Phoenix Rhino (Immortal Flame of South Africa)** 🦏🔥
   * **#9 Andromeda Giraffe (Nebula Chain Guardian)** 🦒⚡

### 9.3 Asset Generation Pipeline
Card illustrations are rendered with stylized cartoon/anime aesthetics via Google Imagen / NanoBanana tooling and formatted into standard 3:4 aspect ratio cards.

### 9.4 Schema Victory (True Win Condition)
The sole victory condition to complete a schema is **collecting all trading cards in that schema's booster pack**:
* When the final card is gathered, a victory fanfare plays!
* The golden exit portal archway illuminates at the summit.
* Displays the completed Card Album for that world.

---

## 10. Level 1: Dragon Gem Mountain

* **Ambientazione**: Volcanic candy landscape under a violet starlit sky.
* **Key Landmarks**:
  * *The Candy Caldera*: Base spawn area with bounce pads and glowing crystal spires (Cyan, Ruby, Emerald, Topaz).
  * *Guinea-Pig Burrows*: Low-clearance tunnels under candy rock formations accessible only by crouching or sliding.
  * *High Crystal Altar*: Elevated central platform holding Mythic Card #7 and abundant Salmon Nigiri.
  * *Dragon Archway*: Ancient purple carved pillars that open once the booster pack is complete.
* **Enemies**:
  * *Candy Goblins (Brown)*: Slow-moving melee wanderers.
  * *Mischievous Puffballs (Purple)*: Fast bouncing rolling creatures.
  * *Fire Spitter Mini-Dragons (Red)*: Stationary turrets shooting slow candy sparks.

---

## 11. Difficulty Scaling & Cooperative Multiplayer

### 11.1 Dynamic Player-Count Scaling
To keep gameplay engaging whether played solo or with 4 family members:

$$\text{Enemy HP} = \text{Base HP} \times (1 + 0.35 \times (\text{PlayerCount} - 1))$$
$$\text{Enemy Count} = \text{Base Count} + 3 \times (\text{PlayerCount} - 1)$$

* **Solo Mode (1 Player)**: Relaxed exploration, forgiving damage, lower enemy density.
* **Family Squad (3–4 Players)**: Toughened enemy waves, increased aggressiveness, respawning mini-bosses.

### 11.2 Cooperative Mode (PvE First)
* All players share the Card Album collection progress (if Ale finds Card #3, Seby's album updates too).
* Players can revive each other with a "High Five" / Salmon Nigiri share.

### 11.3 Party PvP Mode ("Banana Tag" - Roadmap)
* Dedicated arena mode where getting hit by a banana turns the player into "IT" (wearing a silly banana hat).

---

## 12. AI Observability, Telemetry & Headless Self-Testing

### 12.1 Automated Headless Runner (`npm run sim`)
The engine is completely decoupled from WebGL/DOM when running under Node.js:
* Uses Three.js mathematical structures and Cannon-es rigidbodies in pure memory.
* Runs at $30\ \text{Hz}$ simulation rate for $N$ seconds (e.g. 900 ticks in $< 2$ seconds real time, with optional 60 Hz high-fidelity mode).
* Evaluates an autonomous `BotPlayer` that hunts enemies, avoids hazards, and grabs power-ups.

### 12.2 Exported Observability Artifacts
At the end of an automated run, the telemetry engine exports:
1. `metrics.json`:
   * `durationSeconds`: Elapsed game time.
   * `kills`: Total enemies eliminated.
   * `damageDealt` & `damageTaken`.
   * `accuracy`: Percentage of shots hitting targets.
   * `nigiriEaten`: Salmon nigiri count.
   * `cardsCollected`: List of unlocked trading cards.
   * `hpTimeSeries`: Per-second sample array `[{ second, hp, flightFuel, kills }]`.
2. `hp_chart.svg`:
   * Standalone vector graphic rendering player HP (red line), Jetpack Fuel (blue dashed line), and kill timeline.

### 12.3 Automated CI/CD Verification
* `npm test`: Runs Vitest suite covering player physics, double-jump state machine, and card collection logic.
* `npm run sim`: Verifies bot survival and telemetry output without regression.
