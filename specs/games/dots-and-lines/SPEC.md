---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-08T11:27:00Z"
title: "Dots & Lines Deluxe (Pipopipette Neon Dual-Puffin)"
tags: [game, p5js, html5, puffin-login, mobile, family, 2-player]
author: "Ermete Bottazzi & Riccardo Carlesso"
source_of_inspiration: "Paolo's Dots & Lines mobile game + PuffinLogin spec"
---

# 🔴🔵 Dots & Lines Deluxe (Il Gioco dei Puntini, Linee e Conquiste)

> *"Fallo più bello di quello di Paolo, con Dual-Puffin Login per sfide infuocate tra Ale, Sebi e Papino!"* — Riccardo

## 1. Problem Statement

Il classico gioco di carta e penna dei **puntini e quadratini** (*Dots and Boxes*, storicamente noto in Italia anche come *Pipopipette*) è una pietra miliare della logica combinatoria e un perfetto passatempo a due giocatori faccia a faccia. L'implementazione mobile creata dall'amico **Paolo** ha una grafica dark minimalista elegante (neon ciano e magenta su nero profondo, vedi screenshot in `assets/`), ma manca di:
1. **Identità e continuità familiare**: Nessuna integrazione con i profili dei bambini (Ale, Sebi, Papino, Mamma Kate, Nonna Lucilla).
2. **Dual Login Simultaneo (Player 1 vs Player 2)**: Giocando sullo stesso tablet o smartphone condiviso (tavolo, divano, treno), i due giocatori non possono loggarsi insieme per registrare scontri diretti personalizzati (es. *Ale 🐉 [Rosso] vs Sebi 🏎️ [Blu]*, oppure *Ale vs Papino 👑 [Giallo]*).
3. **Juiciness e Game Feel Premium**: Effetti particellari all'esplosione dei box chiusi, sound design dinamico con feedback vocale (stile Mario Kart / arcade party), animazioni al completamento delle catene ("combo streak") e trofei salvati nel cloud di famiglia.

---

## 2. Obiettivi & Core Features (Più bello di quello di Paolo!)

* **🕹️ Gameplay Classico Impeccabile**:
  - Griglia regolabile di punti: **Small (3x3 box / 4x4 punti)**, **Classic (4x4 box / 5x5 punti)**, **Large (5x5 box / 6x6 punti)**, **Epic (6x6 box)**.
  - Regola aurea: Chi traccia il quarto lato e chiude un quadrato guadagna **1 punto** e **gioca immediatamente un turno extra**.
  - Riconoscimento swipe intuitivo tra due nodi vicini oppure tap secco sullo spazio tra due punti.
* **🐧 Dual PuffinLogin Integrato (Must-Have)**:
  - **Player 1 (Sinistra)** e **Player 2 (Destra)** hanno ciascuno il proprio badge PuffinLogin (`pvt-puffin-login`).
  - Supporto per tutti i profili canonici Carlesso:
    * `ale` 🐉 (Rosso `#EF4444`, PIN candy `2018`)
    * `seby` 🏎️ (Blu `#3B82F6`, PIN candy `2020`)
    * `riccardo` 👑 (Giallo `#FACC15`, Password)
    * `kate` 💜 (Viola `#A855F7`, Password)
    * `lucilla` 🌸 (Rosa `#EC4899`, Password)
    * `guest` 🐧 (Ospite rapido con nickname custom)
  - La palette dei segmenti, delle luci neon e delle caselle conquistate **eredita dinamicamente il colore e l'avatar Pixar del profilo loggato**! (Se Ale gioca contro Sebi, il gioco è letteralmente Rosso vs Blu con avatar 3D viventi che esultano a ogni box).
* **✨ Effetti Visivi "Juicy" (Next-Gen Neon Dark Mode)**:
  - Nodi luminosi a impulso reattivo (*breathing dots*).
  - Tracciamento linee con scia laser / neon glow ad alta definizione (WebGL / Canvas).
  - Quando un giocatore chiude un box: **flash di luce**, onda d'urto particellare (*particle burst*), timbro dell'avatar/stemma del giocatore al centro del box e animazione rimbalzo elastico (*squash & stretch*).
  - Se si innesca una catena di scatole consecutive: combo meter moltiplicatore con suoni in scala musicale ascendente (Do-Re-Mi-Fa-Sol!).
* **🔊 Audio & Sound FX Arcade**:
  - Click metallico tattile / haptic al tracciamento linea.
  - Suono "Pop" cristallino alla conquista del box.
  - Trombe da stadio / cori festosi a fine partita con podio animato.
  - Supporto vocale per bambini (senza parolacce, tono entusiasta Ermete/Puffin!).
* **📊 Cloud High Scores & Statistiche Testa a Testa**:
  - Salva i risultati nel backend Firestore di `pvt-puffin-login` (GCP `ror-goldie`):
    - Bilancio vittorie H2H (es. *Ale 14 vittorie - Sebi 11 vittorie - 2 pareggi*).
    - Record di catena più lunga chiusa in un singolo turno (*Longest Combo Streak*).

---

## 3. Mockup & Riferimenti Grafici (L'App di Paolo)

Gli screenshot originali dell'ispirazione sono archiviati in `assets/`:

1. `assets/01-start-rules.jpg`: Regole di ingaggio, logo neon e layout pulsanti.
2. `assets/02-gameplay-p1-moves.jpg`: Visualizzazione turni e tracciamento orizzontale.
3. `assets/03-gameplay-p2-line.jpg`: Tracciamento verticale e glow neon Player 2.
4. `assets/04-gameplay-p1-turn.jpg`: Punteggio 3-2 e saturazione dei box già chiusi.
5. `assets/05-gameplay-box-closed-tie.jpg`: Momento del pareggio 3-3 e layout footer (`Undo`, `New Game`).

---

## 4. Architettura Tecnica Raccomandata

* **Tech Stack**: HTML5 Canvas / Vite / TypeScript oppure PWA leggera (zero dipendenze pesanti, touch-first responsive per mobile Pixel 10, tablet e desktop).
* **Hosting Target**: `palladius.github.io/dots-and-lines` (sotto lo stesso origin delle altre app di famiglia per ereditare la sessione `localStorage['puffin_auth_v1']`).
* **Integrazione PuffinLogin**:
  - Utilizzo del Web Component `<puffin-login>` in modalità **Dual Instance**:
    ```html
    <div class="player-dock player-left">
      <puffin-login player-slot="p1" default-user="ale"></puffin-login>
    </div>
    <div class="player-dock player-right">
      <puffin-login player-slot="p2" default-user="seby"></puffin-login>
    </div>
    ```
* **Game Engine State**:
  - Matrice bidirezionale dei segmenti: `H[rows+1][cols]` e `V[rows][cols+1]`.
  - Matrice dei quadrati: `Box[rows][cols]` contenente `{ owner: 'p1' | 'p2' | null, claimedAtTurn: number }`.
  - Motore di validazione turno puro TypeScript/JavaScript separato dal renderer per testing deterministico (TDD).

---

## 5. Roadmap di Implementazione

1. **Fase 1 (Core Engine & TDD)**: Logica della griglia $N \times M$, calcolo chiusura box, gestione turno extra, undo state stack.
2. **Fase 2 (Canvas Renderer Neon)**: Griglia di nodi, gesture recognizer swipe & tap, interpolazione glow shader/canvas.
3. **Fase 3 (Dual PuffinLogin Hook)**: Integrazione dei due profili, selezione avatar, sincronizzazione colori e nomi.
4. **Fase 4 (Juiciness & Audio)**: Particelle, combo streak sound, celebrazione finale con coriandoli e medaglioni.
5. **Fase 5 (Deploy & Mobile Polish)**: PWA manifest, testing su Google Pixel 10 (verifica anti-overflow) e push su GitHub Pages.
