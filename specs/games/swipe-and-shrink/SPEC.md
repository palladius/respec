---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-13T20:55:00Z"
title: "Swipe & Shrink (Parole in Fuga / Word Melter)"
tags: [game, word-puzzle, mobile, kids, didactic, multilingual, p5js, puffin-login, family]
author: "Ermete Bottazzi & Riccardo Carlesso"
source_of_inspiration: "Wordscapes / Word Sweeper playable ad + Riccardo's Letter Depletion concept"
---

# 🔤 Swipe & Shrink (Parole in Fuga / Word Melter)

> *"Un gioco di parole dove ruoti le dita sulla griglia e, man mano che trovi le parole, le lettere esaurite scompaiono! Finché una lettera serve ancora per altre parole rimane lì, ma alla fine per l'ultima parola da 3 lettere ti rimangono solo quelle 3 lettere esatte sul tabellone!"* — Riccardo Carlesso

---

## 1. Executive Summary & Vision

**Swipe & Shrink** è un gioco educativo di anagrammi e tracciamento continuo a scorrimento (swipe puzzle) progettato per bambini e famiglie. Ispirato ai moderni giochi mobile di parole (stile *Wordscapes* e *Boggle*), introduce una meccanica innovativa di **esaurimento progressivo delle tessere (Letter Depletion)**:

1. **Il Tabellone Iniziale**: Una griglia di lettere (es. 5×3 o 4×4) ricca di possibilità combinatorie, dove le parole bersaglio condividono lettere tra loro.
2. **Tracciamento Fluido (Continuous Swipe)**: Il giocatore trascina il dito senza staccarlo tra tessere adiacenti (orizzontali, verticali e diagonali) per comporre parole.
3. **Meccanica "Shrink" (Reference Counting)**: Ogni tessera sulla griglia ha un contatore invisibile di dipendenza: quante parole ancora da indovinare necessitano di quella specifica tessera? Quando tutte le parole che usavano quella lettera sono state scoperte, **la tessera si dissolve in polvere di stelle e scompare per sempre dal tabellone**!
4. **Endgame a Risoluzione Perfetta**: Man mano che il livello procede, il tabellone si svuota e si semplifica progressivamente. Per l'ultima parola rimasta (es. una parola da 3 lettere), sul tabellone rimarranno **esattamente e solo quelle 3 lettere**, regalando al bambino un momento trionfale di chiarezza e gratificazione cognitiva.
5. **Poliglotta Nativo**: Sviluppato su misura per Alessandro (8 anni) e Sebastian (6 anni), alunni della scuola bilingue/multilingue *Tandem IMS* di Zurigo, con dizionari dedicati per **Italiano 🇮🇹**, **Inglese 🇬🇧**, **Tedesco 🇩🇪** e **Francese 🇫🇷**.

---

## 2. Didactic Motivation (Perché per i Bambini?)

Nei tradizionali giochi di anagrammi a cerchio o a griglia statica:
* I bambini piccoli (6-8 anni) si sentono sopraffatti dalla sovrabbondanza visiva (*cognitive overload*): anche quando mancano solo una o due parole facili, il disordine delle lettere già usate continua a distrarli.
* **La soluzione Swipe & Shrink**: Eliminando visivamente il "rumore" delle lettere già esaurite, il gioco guida naturalmente l'attenzione visiva sulle lettere residue. Riduce la frustrazione, rinforza l'ortografia e trasforma la ricerca delle parole in un'avvincente caccia al tesoro a eliminazione.

---

## 3. Core Mechanics & Algoritmo di Esaurimento

### 3.1 Struttura del Livello
Ogni livello è definito da:
- **Griglia Iniziale $G$**: $M \times N$ tessere (es. $5 \times 3 = 15$ tessere).
- **Target Words $W = \{w_1, w_2, \dots, w_k\}$**: Insieme delle parole bersaglio da trovare (di lunghezza variabile, da 3 a 7 lettere).
- **Mappatura Lettera-Parola (Reference Mapping)**:
  Ogni coordinata $(r, c)$ della griglia contiene una lettera $L_{(r,c)}$.
  Definiamo il grado di dipendenza residuo $D(r, c)$:
  $$D(r, c) = \sum_{w \in W_{\text{unsolved}}} [ (r, c) \in \text{Path}(w) ]$$
  * Dove $\text{Path}(w)$ è il percorso di coordinate utilizzato per comporre la parola $w$.
  * Se $D(r, c) > 0$: La tessera rimane attiva, cliccabile e visibile.
  * Non appena $D(r, c) == 0$: **Trigger Dissolvenza**! La tessera attiva l'animazione di "shrink" (rimpicciolimento, pop sonoro, particelle scintillanti) e lascia uno spazio vuoto (o fa collassare le tessere vicine).

### 3.2 Regole di Connessione Swipe
- **Adiacenza a 8 direzioni**: Ci si può muovere in orizzontale, verticale e diagonale tra caselle contigue.
- **No Riutilizzo nello Stesso Swipe**: All'interno di una singola parola, la stessa coordinata $(r, c)$ non può essere attraversata due volte (regola standard *Boggle*).
- **Feedback Tattile & Visivo**: Durante il trascinamento del dito:
  - Una linea laser luminosa (spline curva morbida) unisce le tessere selezionate.
  - La parola corrente appare in tempo reale in un fumetto fluttuante sopra il dito.
  - Vibrazione aptica leggera (*haptic tick*) a ogni nuova lettera agganciata.

### 3.3 Gestione delle Parole Bonus
Se il bambino traccia una parola valida di vocabolario che non fa parte delle *Target Words* del livello:
- La parola finisce nel **Vaso delle Parole Segrete (Cookie Jar / Puffin Bank)**.
- Rilascia stelline bonus o monete per sbloccare temi visivi, senza alterare il contatore delle tessere bersaglio.

---

## 4. Architettura Multilingue (IT / EN / DE / FR)

Il gioco include un selettore di lingua immediato con bandiere ad accesso rapido:

| Lingua | Bandiera | Target Didattico per Ale (P3) & Sebi (P1) | Esempio Livello Starter |
| :--- | :---: | :--- | :--- |
| **Italiano** | 🇮🇹 | Ortografia, doppie, parole di famiglia e animali | `CANE`, `GATTO`, `MARE`, `SOLE`, `LUNA` |
| **Inglese** | 🇬🇧 | Phonics, sight words, rime (cat/hat/bat), vocabolario Tandem | `HOW`, `WORSE`, `THUNDER`, `OWNER`, `STORM` |
| **Tedesco** | 🇩🇪 | Parole scolastiche zurighesi, composti semplici, animali | `HAUS`, `BAUM`, `KATZE`, `SONNE`, `WINTER` |
| **Francese** | 🇫🇷 | Vocabolario di base P1/P3, suoni dolci | `EAU`, `CHAT`, `BLEU`, `PETIT`, `POMME` |

*Nota di pronuncia & sintesi vocale:* Al completamento di una parola, il sistema legge ad alta voce la parola corretta con pronuncia nativa limpida per rinforzare l'apprendimento uditivo.

---

## 5. UI / UX Design & "Juiciness"

Riferimento visivo dagli screenshot in `assets/`:
1. `assets/01-swipe-level-3-ace.jpg`: Mostra la schermata con le parole già scoperte sulla destra (**OUNCE**, **PEDAL**, **CHANCE**, **POUNCE**, **CHANNEL**) e la manina che guida lo swipe per **ACE**.
2. `assets/02-swipe-level-2-thunder.jpg`: Mostra la griglia 5×3 con le parole orizzontali e verticali incrociate (**HOW**, **WORSE**, **THUNDER**).

### Componenti dell'Interfaccia
1. **Area Superiore (Target Slots)**:
   - Caselle in stile cruciverba / *Wordscapes* con il numero di lettere ben visibile.
   - Le parole scoperte si colorano con la palette del profilo giocatore (es. Rosso Fuoco per Ale, Blu Turbo per Sebi).
2. **Area Centrale (Playfield / Shrinking Grid)**:
   - Tessere a rilievo stile legno liscio o caramella levigata con lettere maiuscole chiare ad alta leggibilità (font sans-serif arrotondato, es. *Fredoka* o *Nunito*).
   - Numeretto o gemme minuscole sulla tessera che indicano (opzionale nei livelli avanzati) quante parole la usano ancora.
   - Animazione di distruzione: Quando una lettera si esaurisce, fa un piccolo balzo (*anticipation jump*), esplode in bolle o stelline (*poof!*) e lascia il campo sgombro.
3. **Pannello Inferiore (Tools & Controlli)**:
   - Tasto **Rimescola (Shuffle)**: Riorganizza le posizioni delle tessere residue (utile se il bambino è bloccato).
   - Tasto **Suggerimento (Hint / Lampadina)**: Illumina la prima lettera della parola più corta rimasta.
   - Tasto **Pronuncia (Audio)**: Riascolta le parole scoperte.

---

## 6. Family & PuffinLogin Integration

Seguendo lo standard delle app Carlesso (`pvt-puffin-login`):
* **Profili Personalizzati**:
  - `ale` 🐉 (Tema Rosso, Draghetto Ghigna come aiutante).
  - `seby` 🏎️ (Tema Blu, Macchina da corsa o Pupurabbu come mascotte).
  - `riccardo` 👑 (Tema Giallo/Oro).
  - `kate` 💜 (Tema Viola).
* **Progressi e Trofei**:
  - Salva i livelli superati in `localStorage` e li sincronizza con Firestore (`ror-goldie`).
  - Statistiche settimanali: quante parole scoperte in tedesco, francese, inglese e italiano.

---

## 7. Schema Dati (JSON Spec per i Livelli)

I livelli sono salvati in formato dichiarativo JSON indipendente dall'engine:

```json
{
  "id": "level_en_002",
  "language": "en",
  "difficulty": "medium",
  "theme": "Weather & Nature",
  "grid": {
    "rows": 3,
    "cols": 5,
    "initial_letters": [
      ["T", "H", "U", "N", "D"],
      ["W", "O", "R", "S", "E"],
      ["O", "W", "N", "E", "R"]
    ]
  },
  "targets": [
    {
      "word": "THUNDER",
      "hint": "Loud noise during a storm",
      "path": [[0,0], [0,1], [0,2], [0,3], [0,4], [1,4], [2,4]]
    },
    {
      "word": "WORSE",
      "hint": "Opposite of better",
      "path": [[1,0], [1,1], [1,2], [1,3], [1,4]]
    },
    {
      "word": "OWNER",
      "hint": "Someone who has something",
      "path": [[2,0], [2,1], [2,2], [2,3], [2,4]]
    },
    {
      "word": "HOW",
      "hint": "Question word",
      "path": [[0,1], [1,1], [1,0]]
    }
  ],
  "bonus_words": ["NOW", "ROW", "ONE", "RED", "WORM"]
}
```

---

## 8. Criteri di Accettazione & Verifica (Test Checklist)

- [ ] **Riconoscimento Coordinate**: Lo swipe unisce correttamente le tessere adiacenti in qualsiasi sequenza orizzontale, verticale o diagonale.
- [ ] **Reference Counter Esatto**:
  - Test case: Data una griglia con 2 parole che condividono la lettera `E` (es. `ACE` e `RED`):
    - Scoprendo solo `ACE`, la lettera `E` **NON** deve scomparire.
    - Scoprendo anche `RED`, la lettera `E` (avendo raggiunto dipendenza residua 0) **DEVE** dissolversi immediatamente.
- [ ] **Endgame Pulito**: Quando rimane una sola parola da $K$ lettere, sul tabellone devono rimanere esattamente $K$ tessere visibili.
- [ ] **Supporto Multilingue**: Il cambio lingua carica istantaneamente i dizionari e i set di livelli corretti per IT, EN, DE, FR.
- [ ] **Performance Mobile**: 60 FPS stabili su browser touch (iOS Safari, Android Chrome) con particelle fluide e zero lag durante lo swipe continuo.
- [ ] **Kid-Safe**: Nessun banner pubblicitario, nessun acquisto in-app, grafica pulita e suoni rassicuranti.
