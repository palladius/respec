---
speck_version: "0.2"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-13T20:55:00Z"
updated_at: "2026-11-14T11:34:00+01:00"
title: "Kids Words (Parole in Fuga / Swipe & Shrink Kids Edition)"
tags: [game, word-puzzle, mobile, kids, didactic, multilingual, gemini-tts, puffin-login, family]
author: "Ermete Bottazzi & Riccardo Carlesso"
source_of_inspiration: "Wordscapes / Word Sweeper + Riccardo's Letter Depletion + Storybook Reveal & Gemini 3.8 TTS"
---

# 🔤 Kids Words (Parole in Fuga — Edizione Bimbi)

> *"Un gioco di parole su misura per Ale e Seby dove ruoti il dito sulla griglia e, man mano che trovi le parole, le tessere esaurite scompaiono svelando lo sfondo illustrato con i bimbi visti di spalle! Alla fine ti rimangono solo le lettere dell'ultima parola, e la voce di Gemini 3.8 TTS pronuncia ogni parola — spiegando a voce e in un fumetto quelle più ostiche (oltre A1/A2)!"* — Riccardo Carlesso

---

## 1. Executive Summary & Vision

**Kids Words** (*Parole in Fuga*) è un gioco educativo di anagrammi a scorrimento continuo (*swipe puzzle*) progettato su misura per **Alessandro (8 anni, P3)** e **Sebastian (6 anni, P1)**. Rispetto ai giochi di parole per adulti, introduce cinque pilastri pensati per i bambini:

1. **Esaurimento Progressivo delle Tessere (`Shrink` / Reference Counting)**: Ogni tessera $(r, c)$ conta quante parole ancora da scoprire la attraversano ($D(r,c)$). Quando $D(r,c) = 0$, la tessera esplode in stelline e scompare dal tabellone. Per l'ultima parola rimasta restano sul campo **esattamente e solo** le sue lettere.
2. **Sfondo "Cutie Back-View" a Svelamento Progressivo**: Sotto la griglia c'è un'illustrazione stile *Pixar / Storybook* con **Ale e Seby visti di spalle** che ammirano un paesaggio magico (Lago di Zurigo, bosco incantato, spazio, fondale marino). Man mano che le tessere scompaiono, la griglia libera la visuale e svela l'intero quadro.
3. **Due Tier di Difficoltà Tarati per Età (PuffinLogin)**:
   - **Tier Seby (6yo)**: Griglia $3 \times 3$, $N = 3$ parole corte (3–4 lettere).
   - **Tier Ale (8yo)**: Griglia $4 \times 4$, $N = 4\text{–}5$ parole medie (4–6 lettere).
4. **Voce Narrante `gemini-3.8-flash-tts` + Spiegazione Parole Ostiche (> A1/A2)**:
   - Ogni parola indovinata viene pronunciata con entusiasmo nella lingua del livello.
   - Se una parola necessaria per chiudere l'incrocio della griglia **non è di livello base A1/A2** (es. `ARNIA`, `AIA`, `SEPPIA` invece di `CANE`, `SOLE`, `MARE`), appare subito un **fumetto illustrato** e la voce TTS legge una breve spiegazione adatta ai bambini.
   - A fine livello, la voce festeggia il bambino per nome riepilogando le parole trovate.
5. **Libreria Statica Pre-generata ($0 Costo Runtime) & Rollout Italiano-First**:
   - **Fase 1 (Taratura)**: 30–50 livelli statici curati in **Italiano 🇮🇹** con file audio pre-generati offline tramite Free Tier (`GEMINI_API_KEY_FREE_TTS_STT`).
   - **Fase 2 (Multilingue Tandem IMS)**: Estensione a **Tedesco 🇩🇪**, **Francese 🇫🇷** e **Inglese 🇬🇧** una volta validata la curva di difficoltà.

---

## 2. Core Mechanics & Algoritmo di Esaurimento

### 2.1 Struttura del Livello e Reference Counting
Ogni livello è definito da:
- **Griglia Iniziale $G$**: $M \times M$ tessere ($3 \times 3$ per Seby, $4 \times 4$ per Ale).
- **Target Words $W = \{w_1, w_2, \dots, w_N\}$**: Insieme delle parole bersaglio da trovare, ciascuna associata a una sequenza ordinata di coordinate adiacenti $\text{Path}(w)$.
- **Grado di Dipendenza Residuo $D(r, c)$**:
  $$D(r, c) = \sum_{w \in W_{\text{unsolved}}} [ (r, c) \in \text{Path}(w) ]$$
  - Se $D(r, c) > 0$: La tessera rimane attiva, opaca e selezionabile.
  - Non appena $D(r, c) = 0$: **Trigger Dissolvenza**! La tessera fa un piccolo salto (*anticipation jump*), esplode in stelline sonore (*poof!*) e diventa completamente trasparente, lasciando visibile il pezzo corrispondente dell'illustrazione di Ale e Seby.

### 2.2 Regole di Connessione Swipe (Solo Orizzontale e Verticale — NO Diagonale)
- **Adiacenza Ortogonale a 4 direzioni ($\Delta r + \Delta c = 1$)**: Ci si può muovere **solo in orizzontale e verticale** tra caselle contigue (niente diagonali, per massima chiarezza visiva e giocabilità per i bimbi).
- **No Riutilizzo nello Stesso Swipe**: La stessa coordinata $(r, c)$ non può essere usata due volte nella stessa parola.
- **Feedback Tattile & Visivo**:
  - Linea morbida luminosa che segue il dito ad angoli retti (orizzontale/verticale) e unisce le tessere.
  - Fumetto grande ad alta leggibilità sopra il dito con le lettere selezionate.
  - Se il bambino traccia una parola valida del dizionario non prevista tra i target, finisce nel **Puffin Cookie Jar** regalando stelline bonus senza alterare $D(r, c)$.

---

## 3. Taratura Difficoltà & Architettura Offline ($0 Costo Runtime)

### 3.1 Due Tier per Profilo Giocatore (8–10 Parole Coese per Livello)

| Parametro | 🏎️ Tier Seby (6 anni — P1) | 🐉 Tier Ale (8 anni — P3) |
| :--- | :--- | :--- |
| **Griglia** | $3 \times 3$ (9 lettere ad alta densità) | $4 \times 4$ (16 lettere) |
| **Parole per livello ($N$)** | **$8\text{–}10$ parole** ortogonali coese | **$8\text{–}10$ parole** ortogonali coese |
| **Lunghezza parole** | $3\text{–}5$ lettere | $3\text{–}6$ lettere |
| **Lessico prevalente** | Base A1/A2 (`APE`, `PANE`, `RANA`, `SOLE`) + parole d'incrocio spiegate (`ARNIA`, `AIA`, `ARA`) | Base A1/A2 + parole curiose `B1+` spiegate (`SEPPIA`, `ASTRO`, `ELMO`, `LEVA`) |
| **Aiuto visivo (Hint)** | Illumina la $1^\text{a}$ e $2^\text{a}$ lettera con freccia | Illumina la $1^\text{a}$ lettera della parola più corta |

### 3.2 Generazione Offline vs Runtime (Implicazioni Costi & AI)
Per garantire **zero latenza**, **zero costi API runtime** e **100% determinismo** (evitando griglie irrisolvibili o parole inappropriate generate al volo):
1. **Generatore Offline CLI (`scripts/generate_levels.py`)**:
   - Costruisce o valida i livelli JSON assicurando che ogni parola abbia un $\text{Path}(w)$ valido sulla griglia e che l'ultima parola lasci esattamente le sue lettere.
   - Classifica automaticamente ogni parola per livello CEFR (`A1`, `A2`, `B1+`).
   - Per ogni parola non-A1/A2 (`cefr: "B1+"` o `needs_explanation: true`), genera una spiegazione di 1 frase adatta a un bimbo di 6–8 anni.
2. **Pre-generazione Audio (`scripts/generate_tts_assets.py`)**:
   - Usa `gemini-3.8-flash-tts` con la chiave gratuita (`GEMINI_API_KEY_FREE_TTS_STT`) per generare gli `.mp3` statici prima del deploy.
   - Gli asset audio vengono salvati in `public/audio/<lang>/` e serviti staticamente da GitHub Pages / Cloud Run a costo `$0.00`.

---

## 4. Pipeline Audio `gemini-3.8-flash-tts` & Spiegazione Parole Ostiche

Ogni parola nel livello JSON possiede un flag `cefr` (`"A1"`, `"A2"`, `"B1+"`):

1. **Parola Base (`A1` / `A2` — es. `CANE`, `SOLE`, `MARE`)**:
   - Appena indovinata, parte subito l'audio entusiasta: *"SOLE!"*
2. **Parola Ostica / Meno Comune (`B1+` — es. `ARNIA`, `AIA`, `SEPPIA`)**:
   - Spesso negli incroci su griglia escono parole corte ma rare per un bambino (es. `AIA` incrociando `APE` e `MIA`, oppure `ARNIA`, `SEPPIA`, `ORMA`).
   - Appena indovinata:
     - **Fumetto Scritto (Didactic Toast / Comic Bubble)**: Appare un fumetto illustrato con icona/emoji e spiegazione semplice in stampatello chiaro (es. *"🐝 **ARNIA**: È la casetta delle api dove producono il miele!"*).
     - **Audio Parlato (`gemini-3.8-flash-tts`)**: La voce pronuncia la parola seguita subito dalla spiegazione parlata con tono caldo e curioso.
3. **Audio Finale di Fine Livello**:
   - Quando l'ultima tessera scompare e lo sfondo è completamente svelato, parte il clip celebrativo personalizzato per il profilo attivo:
     - Esempio Seby: *"Evviva Seby! Hai scoperto tutte le parole: APE, MIELE, ARNIA!"*
     - Esempio Ale: *"Grande Ale! Hai liberato il quadro trovando MARE, VELA, SEPPIA, RETE!"*

---

## 5. Direzione Artistica: "Cutie Bimbi di Spalle" & UI Reveal

1. **Illustrazioni "Back-View" (Ale & Seby di Spalle)**:
   - Ogni livello (o pacchetto di 5 livelli) ha come sfondo una scena tenerissima stile Pixar / acquerello fiabesco in cui **Alessandro e Sebastian sono ritratti di spalle** (riconoscibili dai capelli, felpe Rosso/Blu o zainetti, accompagnati da Ghigna 🐉 o Pupurabbu 🐧) mentre guardano verso l'orizzonte:
     - *Livello 1*: Ale e Seby di spalle sul pontile del Lago di Zurigo con i cigni e le Alpi.
     - *Livello 2*: Ale e Seby di spalle davanti a un grande acquario luminoso (tema Mare: `MARE`, `RETE`, `SEPPIA`).
     - *Livello 3*: Ale e Seby di spalle in un prato fiorito con un alveare (tema Natura: `APE`, `FIORE`, `ARNIA`).
2. **Effetto Svelamento**:
   - All'inizio del livello la griglia di tessere copre il centro dell'illustrazione con un leggero effetto vetro satinato (`backdrop-filter: blur(4px)` dietro ogni singola tessera attiva).
   - Ogni volta che una tessera raggiunge $D(r,c) = 0$ e scompare, quella porzione di quadro diventa nitida e luminosa.
   - A fine livello la griglia sparisce del tutto per 4 secondi lasciando ammirare la foto intera con un effetto di stelline.

---

## 6. Roadmap Multilingue (IT $\rightarrow$ DE, FR, EN)

| Fase | Lingua | Stato | Obiettivo Didattico |
| :--- | :--- | :---: | :--- |
| **Fase 1 (MVP)** | 🇮🇹 **Italiano** | 🎯 **Attiva** | Tarare griglie $3\times 3$ (Seby) e $4\times 4$ (Ale), testare lo svuotamento $D(r,c)$ e il fumetto + TTS per le parole `B1+`. |
| **Fase 2** | 🇩🇪 **Tedesco** | ⏳ Next | Vocabolario scolastico zurighese (Tandem IMS), sostantivi con articolo colorato (*der/die/das*) nel fumetto. |
| **Fase 3** | 🇫🇷 **Francese** | ⏳ Planned | Pronuncia dei dittonghi (`EAU`, `OU`, `OI`) dove sentire l'audio subito dopo lo swipe aiuta moltissimo la lettura. |
| **Fase 4** | 🇬🇧 **Inglese** | ⏳ Planned | Sight words e phonics (`NIGHT`, `LIGHT`, `STORM`). |

---

## 7. Schema Dati JSON (Livelli + Spiegazioni `B1+` + Audio)

```json
{
  "id": "it_seby_003",
  "language": "it",
  "tier": "seby",
  "theme": "Il Prato delle Api",
  "background_image": "assets/bg/ale_seby_back_meadow.webp",
  "completion_audio": {
    "seby": "audio/it/complete_it_seby_003_seby.mp3",
    "ale": "audio/it/complete_it_seby_003_ale.mp3"
  },
  "grid": {
    "rows": 3,
    "cols": 3,
    "initial_letters": [
      ["A", "P", "E"],
      ["R", "N", "I"],
      ["M", "I", "A"]
    ]
  },
  "targets": [
    {
      "word": "APE",
      "cefr": "A1",
      "needs_explanation": false,
      "path": [[0,0], [0,1], [0,2]],
      "audio_word": "audio/it/words/ape.mp3"
    },
    {
      "word": "MIA",
      "cefr": "A1",
      "needs_explanation": false,
      "path": [[2,0], [2,1], [2,2]],
      "audio_word": "audio/it/words/mia.mp3"
    },
    {
      "word": "ARNIA",
      "cefr": "B1+",
      "needs_explanation": true,
      "emoji": "🐝",
      "explanation_text": "L'arnia è la casetta di legno dove vivono le api e preparano il miele!",
      "path": [[0,0], [1,0], [1,1], [1,2], [2,2]],
      "audio_word": "audio/it/words/arnia.mp3",
      "audio_explanation": "audio/it/explanations/arnia_explain.mp3"
    }
  ],
  "bonus_words": ["ERA", "PIA", "IRA", "RANA"]
}
```

---

## 8. Family & PuffinLogin Integration

Seguendo lo standard `puffin-login`:
- **Selezione Rapida Profilo**: All'avvio, cliccando su `Seby` 🏎️ il gioco seleziona automaticamente i livelli `tier: "seby"` ($3\times 3$) e usa i file audio con il suo nome; cliccando su `Ale` 🐉 carica il `tier: "ale"` ($4\times 4$).
- **Album delle Figurine Svelate**: Ogni livello completato sblocca la relativa illustrazione di Ale e Seby di spalle nella "Galleria dei Ricordi", dove i bimbi possono anche riascoltare le parole difficili imparate.

---

## 9. Criteri di Accettazione & Verifica (Test Checklist)

- [ ] **Validazione Invariante `D(r,c)`**: Ogni cella non vuota della griglia iniziale appartiene ad almeno una parola target ($D(r,c) \ge 1$ a inizio livello) e scende a $0$ solo quando tutte le parole che la usano sono state risolte.
- [ ] **Endgame Pulito**: Quando manca una sola parola di $K$ lettere, restano visibili esattamente $K$ tessere.
- [ ] **Spiegazione Automatica Parole Ostiche (`needs_explanation: true`)**: Scoprendo una parola oltre `A1`/`A2` (es. `ARNIA`, `AIA`, `SEPPIA`), si apre il fumetto illustrato e viene riprodotto in coda `audio_word` + `audio_explanation`.
- [ ] **Reveal Sfondo Cutie**: La scomparsa delle tessere rimuove progressivamente l'offuscamento sullo sfondo di Ale e Seby di spalle fino alla vista completa a fine livello.
- [ ] **Zero Chiamate AI a Runtime**: Tutti i livelli e gli audio `.mp3` di `gemini-3.8-flash-tts` funzionano offline/staticamente senza latenza.
