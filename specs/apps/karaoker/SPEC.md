# 🎤 Palladius Karaoker (K-A-R-A-O-K-A-R) — Spec

> **Puns & OKR Alignment:** *"Karaoker"* (spelled `K-A-R-A-O-K-A-R` to align with Q4 OKRs). A seamless, zero-friction AI song studio and interactive synchronized karaoke player.

---

## 1. Executive Summary & Objective

**Palladius Karaoker** is a web-first application that takes a natural language song idea, generates the full music track via AI, streams immediate audio playback, concurrently paints a stylized CD-ROM jewel case / album cover, computes word- and line-level subtitle alignments (`.srt` / `.vtt`), and renders a full-screen interactive `karaoke.html` interface with synchronized scrolling lyrics and bouncy tempo markers.

### Core Philosophy: Progressive Time-to-Value
1. **Immediate Playback:** The music track is generated and immediately loaded into the audio player so the user hears music within seconds without waiting for downstream visual assets.
2. **Background Asset Enrichment:** Album cover art and lyrics transcription/alignment run asynchronously in parallel.
3. **Synchronized Sing-Along Experience:** Once timestamps are aligned, the UI transitions to an interactive `karaoke.html` view with jumping/highlighted lyrics, vocal cues, and playlist management.

---

## 2. User Journey & Core Features

### Step 1: Input & Configuration
- **API Key Provisioning:** Client-side BYOK (Bring Your Own Key) input for `GEMINI_API_KEY` (persisted in `localStorage` or session vault).
- **Prompt Input:**
  - Topic / Idea (e.g. *"Voglio una canzone punk rock energica sui culurgiones e le vacanze in Sardegna"*).
  - Language selector (Italian, English, French, etc.).
  - Genre / Mood pills (80s Cartoon, Pop-Punk, Operatic Rock, Ballad, Forró, Synth-Wave).
  - Duration selector (30s preview vs. 2-3 minute full track).

### Step 2: Immediate Generation & Play
- Orchestrator triggers the Google GenAI / Lyria music generation backend.
- As soon as the `.mp3` / stream buffer arrives, the player starts playback automatically with a waveform visualizer.

### Step 3: Concurrent Cover Generation (CD Jewel Case)
- In parallel with audio initialization, an image generation prompt is formulated from the song theme.
- Generates a square album cover art (Pixar 3D style, retro vinyl, or 90s CD jewel case).
- Displays the spinning vinyl/CD thumbnail in the player.

### Step 4: Lyrics Alignment & SRT Pipeline
- Extracts or predicts sung lyrics with timestamps (via Lyria structured timestamps or downstream STT alignment like Whisper / Gemini Audio).
- Outputs standard `.srt` / `.vtt` format and JSON timecode cues:
  ```json
  [
    { "start": 0.0, "end": 2.5, "text": "Siamo partiti da Olbia col furgone" },
    { "start": 2.5, "end": 5.8, "text": "Dritti a Cala Gonone a mangiar culurgiones!" }
  ]
  ```

### Step 5: Interactive `karaoke.html` View
- Based on the proven Campobasso DevFest / Canzoniere `karaoke.html` template.
- Features:
  - Bouncing ball / syllable-level or line-level glowing gradient highlight.
  - Fullscreen TV / party mode.
  - Playlist sidebar: browse historical generated songs, filter by tag, re-listen, export `.mp3`, `.srt`, or standalone `karaoke.html` zip bundle.

---

## 3. System Architecture & Tech Stack

```text
[ Browser / Frontend (Vue / React / Vanilla Tailwind + WebAudio) ]
                     │
         ┌───────────┴───────────┐
         ▼                       ▼
  [ Audio Generation ]   [ Visuals & Subtitles ]
   - Lyria 3 / 3.5        - Nano Banana / Imagen 3 (Cover Art)
   - Immediate MP3 stream - Gemini Audio / STT (Timestamp Alignment)
         │                       │
         └───────────┬───────────┘
                     ▼
         [ Interactive Karaoke Engine ]
         - Synchronized HTML5 Audio + WebVTT / JS highlighter
         - Playlist local storage / Cloud Storage sync
```

### Proposed Stack:
- **Frontend:** Lightweight single-page app (Tailwind CSS, Alpine.js / Svelte / Vanilla ES6 for zero bloat, Web Audio API).
- **Backend / API Bridge (optional or serverless):** Fast Python (FastAPI / Cloud Run) or pure client-side calls using `@google/genai` JS SDK if endpoints permit direct browser calls.
- **Storage:** Local IndexedDB for offline song caching, with optional export to Google Cloud Storage or Jekyll Canzoniere structure.

---

## 4. Verification & Milestones (Q4 OKR Deliverables)

- [ ] **M1: Minimal Audio Pipeline:** Input prompt + API key -> MP3 generated and played immediately in browser.
- [ ] **M2: CD Cover Generator:** Concurrent cover art creation rendered in responsive jewel-case card.
- [ ] **M3: Timestamp Alignment & SRT:** Automated lyric timing extraction to VTT/SRT.
- [ ] **M4: Karaoke HTML Renderer:** Seamless integration of Campobasso-style bouncing/highlighting lyric player.
- [ ] **M5: Playlist Management:** Save, load, export, and search through generated tracks.
