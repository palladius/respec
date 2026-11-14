# Input Prompt & Inspiration

**Date:** 2026-11-13  
**Source:** Riccardo Carlesso (Voice message on Telegram)  
**Game Concept:** Swipe & Shrink (Word Anagram & Depletion Grid)

## Audio Transcription (Verbatim)

> *"Ermete ho trovato questo nuovo gioco in cui tu puoi ruotare le dita e man mano che le lettere si esauriscono poi le parole si poi dopo si cancellano quelle lettere quindi rimani magari con una parola da tre e solo con tre lettere non so se capisci finché quella lettera contribuisce più di una parola c'è se no no crea in respec uno speck.md di questo gioco che domani provo a implementarlo per i bambini in italiano e in inglese e in tedesco e in francese amico mio ok grazie"*

## Key Gameplay Insights from Screenshots

1. **Grid & Swipe**: Bottom tile grid (e.g. 5x3 or 4x4) where players drag/swipe continuously through connected adjacent letters to spell words (like *Wordscapes* / *Scrabble* / *Boggle*).
2. **Target Crossword / Slots**: Top slots display discovered and undiscovered words organized by length (3-letter, 4-letter, 5-letter, 6-letter, 7-letter words).
3. **The "Shrink" / Letter Depletion Mechanic**: 
   - Each tile tracks a reference count of how many remaining target words still require it.
   - When all target words containing a specific letter tile have been discovered, that tile dissolves/pops and disappears from the grid!
   - As the puzzle progresses, the grid empties and simplifies.
   - In the endgame, for the last remaining 3-letter word, only those 3 letters remain on screen, creating an intuitive, satisfying finish.
4. **Target Audience**: Multilingual kids app for Alessandro (8) and Sebastian (6) attending Tandem IMS (supporting **IT**, **EN**, **DE**, **FR**).

---

## Follow-up Refinement (2026-11-14 — T3 Code Session)

> *"ok ora creiamo un new folder called kids-words e diaogli un taglio piu' bambinico. Foto cutie dei bimbi di dietro, parole piu' facii e quando le indovini ti da l'audio di gemini38tts che le pronuncia nella lingua (it, de, fr, en che sono le 4lingue dei miei bimbi). Le parole dovrebbero essere piu' facili del gioco per audlti se possibile. Magari cominciamo solo con l'italiano e tariamo il gioco per l'italiano epoi facciamo le altre lingue una volta che abbiamo figured out la difficolta, cosa fare di N parole, e se gli schemi sono N statici o generati al momento. Ovviamente queste cose hanno implicazioni a livello di costo e generazione AI"*

### Decisions Locked In (`/grill-me`):
1. **Single Canonical Spec**: Moved `specs/games/swipe-and-shrink` $\rightarrow$ `specs/games/kids-words`.
2. **Offline Pre-generated Static Library ($0 Runtime Cost)**: 30–50 curated levels in JSON + pre-generated `.mp3`/`.wav` audio assets via `gemini-3.8-flash-tts` (`GEMINI_API_KEY_FREE_TTS_STT`). Zero runtime LLM/TTS latency or billing.
3. **Two Age-Calibrated Profile Tiers**:
   - **Seby Tier (6yo / Primary 1)**: $3 \times 3$ grid, $N = 3$ target words per level, 3–4 letters each.
   - **Ale Tier (8yo / Primary 3)**: $4 \times 4$ grid, $N = 4\text{–}5$ target words per level, 4–6 letters each.
   - **Phase 1 Focus**: Calibrate and perfect **Italian 🇮🇹 first**, then expand to **DE 🇩🇪, FR 🇫🇷, EN 🇬🇧**.
4. **Cutie Pixar/Storybook Back-View Reveal**: Background illustrations show Ale and Seby seen from behind (over-the-shoulder / back view) looking out at magical landscapes (Zurich lake, enchanted forest, space, ocean). As letter tiles reach $D(r,c) = 0$ and vanish, the unobstructed artwork is progressively revealed.
5. **Enthusiastic Gemini 3.8 Flash TTS**:
   - Immediate enthusiastic native pronunciation on each discovered word.
   - Personalized level-completion celebratory audio clip per profile (e.g., *"Bravo Seby! Hai trovato CANE, SOLE, MARE!"*).
6. **Hard/Uncommon Word Mini-Explainer (Beyond CEFR A1/A2)**:
   - Basic words (`CANE`, `SOLE`, `MARE` — tier `A1`/`A2`) only play the enthusiastic word pronunciation.
   - Harder or less common words inevitably needed to make dense grids work (`ARNIA`, `AIA`, `SEPPIA` — tier `B1+` / `hard: true`) immediately trigger both a **kid-friendly comic speech bubble** and a **spoken Gemini 3.8 Flash TTS mini-explanation** as soon as the word is solved (e.g., *"ARNIA! L'arnia è la casetta dove vivono le api e fanno il miele!"*).
