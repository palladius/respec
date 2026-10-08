---
speck_version: "0.1"
mode: manual
created_at: 2026-09-16
updated_at: 2026-09-16
---

# Hummerli: Autonomous ADK German/Language Oral Tutor, Bidirectional Barge-In Engine & Multi-Eval Datalake

A dedicated, standalone AI tutoring infrastructure built on Google's **Agent Development Kit (ADK)**. Designed for real-time oral immersion with **bidirectional voice streaming and interruption (barge-in)**, a **Goethe B1 Spaced Repetition (SRS) engine**, a **multi-engine LLM replay/evaluation datalake**, and **multi-profile family support** (Riccardo's Swiss naturalization B1 exam prep + Alessandro & Sebastian's multi-lingual school curricula).

---

## Problem Statement

To obtain Swiss citizenship in the Canton of Zurich, applicants must achieve a certified **B1 level in oral German (Hören & Sprechen)** alongside A2 in reading and writing. Riccardo recently took the FIDE exam and fell just shy of the B1 oral bar (scoring A2 in Hören), requiring targeted oral training before re-taking the module and submitting a fresh application (*neues Gesuch*).

While an earlier Hermes/OpenClaw prototype existed (`learn-german-hummerli`), it suffered from four critical limitations:
1. **Turn-based conversational lag & no barge-in**: Traditional STT → LLM → TTS pipelines introduce jarring turn-taking pauses and prevent the learner from naturally interrupting (*"Wie bitte?"*, *"Non ho capito"*) while the tutor is speaking.
2. **Lack of a curriculum & spaced repetition**: The agent conversed aimlessly without tracking which specific B1 vocabulary words had been acquired, actively used, or neglected.
3. **No historical replay or model decoupling**: Audio and evaluations vanished. There was no way to re-evaluate past conversations with stronger models or hot-swap the evaluation engine while keeping the user experience intact.
4. **Single-user silo**: Riccardo's children (Alessandro, P3, learning German/French A1/English, and Sebastian, P1, learning German/English) require a similar oral sparring partner, but with child-tailored curricula, speeds, and safety guardrails.

---

## Goals

1. **Pure Google ADK Runtime**: Lean microservice leveraging Google ADK in Python. No Hermes runtime dependency, maximizing speed and operational independence.
2. **Bidirectional Streaming & Interruption (Barge-In)**:
   - Built on the **Gemini Multimodal Live API** (bidirectional WebSockets with native audio-in/audio-out) and Google Cloud TTS streaming.
   - The user can **interrupt the tutor mid-sentence** at any moment. Voice Activity Detection (VAD) cuts tutor playback instantly to listen to the user.
3. **Dual-Persona Pedagogical Orchestration**:
   - **Frau Leonie Blücher 🐎 (The German Tutor)**: High-quality Google German Neural voice (`de-DE-Neural2-F` or Gemini native voice), patient, speaking clear B1 High German (*Hochdeutsch*).
   - **The Coach / Translator (Rijckard)**: Warm Italian voice (`it-IT-Neural2-C`) intervening exclusively when the user signals confusion (*"Non ho capito"*, switches to Italian, or asks for grammar clarification).
4. **Goethe B1 Wortliste SRS Engine**:
   - Complete database of the ~1,000–1,500 target lemmata bridging A2 to B1.
   - Modernized **SuperMemo SM-2 / Leitner** algorithm: mastered lemmata are scheduled for 7, 15, 30, and 90 days; failed lemmata are reinforced in subsequent sessions (*"la lingua batte dove il dente duole"*).
   - Hummerli organically embeds or elicits 1–3 target lemmata per session in Zurich civic, cultural, or daily-life scenarios.
5. **Comprehensive Telemetry & Datalake Analytics**:
   - Daily interaction counts (sessions per day, week-over-week trends).
   - Spoken audio minutes (user speech time vs. tutor speech time).
   - Audio message/turn counts.
   - OpenClaw-style datalake time-series visualized in the dashboard.
6. **Cloud Audio Vault & Multi-Engine Evaluation Replay**:
   - 100% of raw user and tutor audio files archived to Google Cloud Storage (`gs://hummerli-audio-archive/{user}/{session_id}/`).
   - Decoupled background evaluation: the live conversation runs at low latency; a separate **LLM Judge Pipeline** evaluates performance in the background.
   - **Offline Replay / Multi-Judge Engine**: Capability to replay archived audio/transcripts through newer, larger models (e.g. Gemini 1.5/2.0 Pro, Claude 3.7) to compare up to 3 evaluation engines concurrently without altering user runtime.
   - Hot-swappable eval logic (*"cambiare il motore sotto il culo"* without downtime).
7. **Multi-User / Family Profiles**:
   - Multi-tenant architecture running on the same backend infrastructure.
   - **Profile `riccardo`**: B1 Swiss naturalization prep (Hören/Sprechen, Zurich politics/geography, Goethe B1 Wortliste).
   - **Profile `ale` (Alessandro, P3)**: Primary 3 German, English, French A1 (Tandem IMS curriculum / *Les Loustics*), kid-friendly tone, 1.25x speed option, zero profanity.
   - **Profile `sebi` (Sebastian, P1)**: Primary 1 early German immersion, phonics, English, playful tone.
8. **Astro / Node.js Heatmap & Analytics Dashboard**:
   - Visual heatmap/grid of the Goethe B1 Wortliste (🟢 Mastered, 🟡 In Progress, 🔴 Needs Review).
   - Interactive drill-down to view all historical attempts, scores, and tutor comments for any word.
   - Datalake charts: daily sessions, minutes spoken, audio file counts, and model eval drift curves.

---

## Non-Goals

1. **Swiss German (Schweizerdeutsch / Züritüütsch)**: The naturalization FIDE exam strictly tests Standard High German (*Hochdeutsch*). High German remains the target language for Riccardo's profile.
2. **Written Worksheets / Pen-and-Paper Drills**: Hummerli is focused on **oral comprehension and production** (*Hören & Sprechen*). Reading and writing are secondary.
3. **General Assistant Workflows**: No calendar modifications, smart-home toggling, or system administration inside Hummerli.
4. **Third-Party Legacy TTS**: No Microsoft Edge / `DiegoNeural` dependencies. Only Google Cloud TTS and Gemini native audio models are used.

---

## Technical Plan / Approach

```
+---------------------------------------------------------------------------------------------------+
|                                          CLIENT LAYER                                             |
|                                                                                                   |
|  [ Telegram Bot (Voice Memos) ]          [ Web Client (WebSockets / PTT / Live Barge-In) ]        |
|  - Async voice notes                     - Real-time full-duplex audio                            |
|  - User Profile selector (/profile)      - Push-to-talk + Voice Activity Detection (VAD)         |
+---------------------------------------------------------------------------------------------------+
                                                  |
                                                  v
+---------------------------------------------------------------------------------------------------+
|                                     INGESTION & VAULT LAYER                                       |
|                                                                                                   |
|  - Live Audio Stream / Voice Note Ingestion                                                       |
|  - Cloud Audio Vault: Raw audio saved immediately to GCS:                                         |
|    gs://hummerli-audio-vault/{user_id}/{session_id}/{turn_id}_{speaker}.ogg                       |
|  - Telemetry Logger: records start_time, duration_ms, bytes, user_id, channel                     |
+---------------------------------------------------------------------------------------------------+
                                                  |
                                                  v
+---------------------------------------------------------------------------------------------------+
|                                  HUMMERLI ADK ENGINE CORE                                         |
|                                                                                                   |
|  +-------------------------------------+             +-----------------------------------------+  |
|  |       PROFILE & CONTEXT MANAGER     |             |              SRS SCHEDULER              |  |
|  |  - Active profile (riccardo/ale/sebi|             |  - Loads target curriculum              |  |
|  |  - Language targets (DE/FR/EN/IT)   | ----------> |  - SM-2 queue: selects 1-3 due words    |  |
|  |  - System prompts & safety rules    |             |    (e.g., 'abholen', 'der Antrag')      |  |
|  +-------------------------------------+             +-----------------------------------------+  |
|                                                                       |                           |
|                                                                       v                           |
|  +---------------------------------------------------------------------------------------------+  |
|  |                        GEMINI LIVE STREAMING & DUAL-PERSONA DISPATCH                        |  |
|  |  - Real-time Bidirectional WebSocket (Low latency, Audio-to-Audio)                         |  |
|  |  - Built-in Barge-In: user voice cut-off stops model audio immediately                     |  |
|  |  - Frau Leonie Blücher 🐎: Google de-DE-Neural2-F / Gemini native (Hochdeutsch B1)         |  |
|  |  - Rijckard (Coach): Google it-IT-Neural2-C (Intervenes on "Non ho capito" / IT/EN query)   |  |
|  +---------------------------------------------------------------------------------------------+  |
+---------------------------------------------------------------------------------------------------+
                                                  |
                        +-------------------------+-------------------------+
                        |                                                   |
                        v                                                   v
+-----------------------------------------------+   +-----------------------------------------------+
|          ASYNCHRONOUS EVALUATION ENGINE       |   |             TELEMETRY & DATALAKE              |
|                                               |   |                                               |
|  - Decoupled from live audio loop             |   |  - Daily activity aggregator:                 |
|  - Pluggable multi-judge harness (N judges):  |   |    * Sessions count per day                   |
|    * Judge A (Fast / Real-time): Flash-Latest |   |    * Spoken minutes (User vs Tutor)           |
|    * Judge B (Deep / Shadow): Pro / Ultra     |   |    * Total audio files exchanged              |
|    * Judge C (External / Benchmark): Claude   |   |  - Time-series stored in DuckDB / Parquet     |
|  - Metrics:                                   |   |    and synced to GCS                          |
|    * Target lemma used / form / correct?      |   +-----------------------------------------------+
|    * Grammar, Vocab, Pronunciation (1-10)     |                           |
|    * Pedagogical coach comment (IT)           |                           v
|  - Updates SM-2 interval for the word         |   +-----------------------------------------------+
+-----------------------------------------------+   |           ASTRO ANALYTICS DASHBOARD           |
                        |                           |                                               |
                        v                           |  - Goethe B1 Heatmap (🟢 Mastered,            |
+-----------------------------------------------+   |    🟡 Learning, 🔴 Struggling)               |
|            OFFLINE REPLAY HARNESS             |   |  - Word drill-down modal with full history    |
|                                               |   |  - OpenClaw-style activity metrics charts     |
|  - CLI command: hummerli replay --session ... |   |  - Multi-judge comparison & diff viewer       |
|  - Re-evaluates raw historical audio on GCS   |   |  - Family Profile switcher                    |
|  - Allows hot-swapping judge models cleanly   |   +-----------------------------------------------+
+-----------------------------------------------+
```

### 1. Bidirectional Streaming & Interruption (Barge-In)
* **Web Client**: Implemented using Gemini Multimodal Live API over WebSockets.
  - Full-duplex PCM audio streaming.
  - Automatic Voice Activity Detection (VAD) on client and server: when the user starts speaking, client immediately stops playback buffer and server drops current candidate generation.
* **Telegram Bot**: Operates in fast voice-memo mode:
  - Push-to-talk voice in → high-accuracy transcription → ADK turn → dual-voice TTS → voice bubble out.
  - User can send a voice note at any time to break/redirect the topic.

### 2. Multi-Profile Family Engine
Every session is tagged with a `profile_id`:
```yaml
profiles:
  riccardo:
    display_name: "Riccardo"
    target_level: "B1"
    curriculum: "goethe_b1"
    primary_language: "de"
    support_language: "it"
    topics: ["Zürich Einbürgerung", "FIDE Exam", "Swiss Politics", "Cantons", "Daily Life"]
    strictness: 4
    allow_swear_words: false
    speed: "1.0x"
  ale:
    display_name: "Alessandro (P3)"
    target_level: "A1-A2"
    curriculum: "tandem_p3"
    languages: ["de", "fr", "en", "it"]
    topics: ["Stories", "Nature", "Primary 3 School", "Drama Club"]
    strictness: 2
    allow_swear_words: false
    speed: "1.25x"
    kid_friendly: true
  sebi:
    display_name: "Sebastian (P1)"
    target_level: "Early Immersion"
    curriculum: "tandem_p1"
    languages: ["de", "en", "it"]
    topics: ["Animals", "Numbers", "Phonics", "Horse Riding"]
    strictness: 1
    allow_swear_words: false
    speed: "1.25x"
    kid_friendly: true
```

### 3. Spaced Repetition (SRS) Engine
* Tracks lemmata in `data/users/{profile_id}/srs_state.json` (mirrored to GCS):
  - `lemma`: e.g. *abholen*
  - `repetition_count`: $n$
  - `ease_factor`: $EF$ (default 2.5)
  - `interval_days`: $I$
  - `next_review_due`: ISO timestamp
  - `last_score`: $q \in [0, 5]$
* **Organic Elicitation**: Frau Blücher weaves the target lemma into natural queries (e.g. *"Musst du heute die Kinder von der Schule abholen?"* or asks Riccardo what he did yesterday to elicit *abgeholt*).

### 4. Audio Vault & Multi-Judge Replay Pipeline
* **Storage Schema on GCS**:
  ```
  gs://hummerli-audio-vault/
    ├── telemetry/
    │   └── activity_daily.parquet
    ├── profiles/
    │   ├── riccardo/
    │   │   ├── sessions/2026-09-16_143000/
    │   │   │   ├── 001_user.ogg
    │   │   │   ├── 001_tutor_katja.ogg
    │   │   │   ├── 002_user.ogg
    │   │   │   └── turn_metadata.jsonl
    │   │   └── reviews.csv
  ```
* **Pluggable Evaluation Pipeline**:
  - Turns are enqueued to a background task runner.
  - Configurable judges:
    1. `v1-flash`: Fast inline scoring (Gemini Flash).
    2. `v2-pro-shadow`: High-accuracy offline judge (Gemini Pro).
    3. `v3-benchmark`: Experimental/external model.
  - **Replay CLI**:
    ```bash
    hummerli replay --profile riccardo --date-range 2026-09-01..2026-09-15 --judge v2-pro
    ```
    Re-runs historical transcripts and audio through the updated judge prompt and stores shadow evaluation scores without touching the live SRS queue until approved.

### 5. Datalake Telemetry & Astro Dashboard
* **Metrics Captured**:
  - `session_id`, `profile_id`, `timestamp`
  - `turns_count`: total turns in session
  - `user_speech_duration_sec`: audio seconds spoken by user
  - `tutor_speech_duration_sec`: audio seconds synthesized
  - `total_latency_ms`: end-to-end response lag
  - `target_words_attempted`: list of lemmata tested
* **Astro Dashboard Views**:
  - **Datalake Activity View**: GitHub-style activity heatmaps, daily minutes spoken bar charts, cumulative audio count.
  - **Goethe B1 Vocabulary Grid**: Tile matrix of all lemmata colored by mastery status with search, filters (Verben, Nomen, Adjektive), and click-to-view history modal.
  - **Multi-Judge Comparison**: Scatter plot showing score correlation between Judge A and Judge B across past sessions.

---

## Alternatives Considered

1. **Microsoft Edge TTS (`DiegoNeural` / `KatjaNeural`)**:
   - *Pros*: Free, easy local python library.
   - *Cons*: High latency, rate limits, no streaming barge-in capability, robotic inflection compared to Google Neural2 / Gemini Live.
   - *Decision*: **Rejected**. Switched to Google Cloud TTS Neural2 and native Gemini Live audio streaming.
2. **Monolithic Live Evaluation (Judge inside live conversational prompt)**:
   - *Pros*: Single LLM call.
   - *Cons*: Bloats live latency, causes awkward pauses, distracts conversational persona with JSON schema adherence.
   - *Decision*: **Rejected**. Decoupled asynchronous background judge keeps conversational response instant while allowing rich metrics and multi-model replays.
3. **Third-Party SaaS LMS / Duolingo Custom Tracks**:
   - *Pros*: Pre-built UI.
   - *Cons*: Cannot do Zurich-specific naturalization context, no audio raw vault, no custom dual-voice logic, no data ownership.
   - *Decision*: **Rejected**.

---

## Implementation Plan

### Phase 1: Curriculum Ingestion & SRS Engine
- [ ] Parse official Goethe-Zertifikat B1 Wortliste into normalized JSON/SQLite schema (`lemma`, `pos`, `level`, `sample_sentences`).
- [ ] Implement SM-2 algorithm class with unit tests for edge cases ($q=0$ to $q=5$).
- [ ] Multi-profile curriculum support (Goethe B1 for `riccardo`, P3 for `ale`, P1 for `sebi`).

### Phase 2: Live Audio Pipeline & Gemini Barge-In
- [ ] Set up Google ADK runtime with WebSocket bidirectional streaming.
- [ ] Implement VAD interruption (barge-in) handler in the audio dispatch loop.
- [ ] Wire Google Cloud TTS (`de-DE-Neural2-F` for Frau Blücher, `it-IT-Neural2-C` for Rijckard) and Gemini Live native audio.

### Phase 3: Audio Vault & Datalake Telemetry
- [ ] Implement automatic GCS uploader streaming raw audio files to `gs://hummerli-audio-vault/`.
- [ ] Build daily telemetry collector (`sessions`, `minutes_spoken`, `audio_files_count`).
- [ ] Export daily summaries to Parquet/JSONL.

### Phase 4: Decoupled Multi-Judge & Replay Harness
- [ ] Implement async background evaluation worker.
- [ ] Define standardized evaluation JSON schema (active/passive usage, grammar, pronunciation estimate, coach feedback).
- [ ] Build `hummerli replay` CLI tool to re-evaluate historical sessions with new model checkpoints.

### Phase 5: Telegram Bot & Web Frontend
- [ ] Telegram bot with profile selection (`/profile riccardo|ale|sebi`) and native voice memo handler.
- [ ] Web audio client supporting full-duplex WebSocket streaming with barge-in button.

### Phase 6: Astro Dashboard
- [ ] Scaffold Astro application with Tailwind CSS.
- [ ] Build Datalake activity charts (daily audio minutes, session frequency, cumulative vocab mastered).
- [ ] Build interactive Goethe B1 vocabulary matrix with drill-down modals.

---

## Open Questions

1. **Barge-In over Telegram**: Telegram voice memos are inherently asynchronous. For Telegram, should we support a "voice-memo barge-in" where sending a new audio cancels any pending voice generation, or keep true real-time barge-in exclusive to the Web client?
2. **Kid Voice Persona for Alessandro & Sebastian**: Should Frau Blücher adapt her tone dynamically for the kids, or should we introduce a third friendly tutor persona (e.g. *Hummerli il Granchietto* 🦞) with simpler vocabulary and enthusiastic encouragement?
3. **Pronunciation Ground Truth**: Without server-side phonetic forced alignment (like Kaldi or Montreal Forced Aligner), how closely does STT transcription confidence score correlate with human pronunciation judgment on Swiss/German vowels (*ö*, *ü*, *ch*)?
