---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
created_at: "2026-10-16T18:05:00Z"
category: games
slug: al-tris
title: "Al Tris: Il Tic-Tac-Toe Ferrarese (Coppietta vs Salama da Sugo)"
tags:
  - games
  - ferrarese
  - 3d-glb
  - web
  - threejs
  - goliardia
---

# Al Tris: Il Tic-Tac-Toe Ferrarese 🥖🥩

## Problem Statement
I classici giochi di Tris (Tic-Tac-Toe / Noughts and Crosses) su mobile e web sono noiosi, asettici e privi di identità culturale e culinaria. Mancano di mordente, di goliardia e della tipica anima gastronomica emiliano-romagnola. 

*Al Tris* reinventa il gioco universale della griglia 3x3 sostituendo i banali simboli geometrici con due icone sacre della tradizione ferrarese:
- **La Croce (`X`)** $\rightarrow$ **La Coppia Ferrarese IGP** (la *Ciupèta*), il tipico pane croccante a forma di croce con i quattro cornetti arrotolati.
- **Il Cerchio (`O`)** $\rightarrow$ **La Salama da Sugo IGP**, il mitico insaccato sferico legato con lo spago a spicchi d'aglio e vino rosso.

Il gameplay è immediato e scattante in 2D, ma quando un giocatore allinea tre pezzi e vince la partita, il gioco si trasforma in un trionfo celebrativo 3D ad alto impatto.

---

## Goals & Features

1. **Griglia di Gioco 2D Fluida & Reattiva**:
   - Classica scacchiera 3x3 stile tovaglia a quadretti da trattoria ferrarese o lavagna rustica.
   - Giocatore 1: **Ciupèta** (Coppietta ferrarese).
   - Giocatore 2: **Salama da Sugo** (o CPU con IA goliardica a insulti dialettali).
   - Segnalazione immediata dei turni con animazioni succose di posizionamento pezzo (crumb particles per il pane, schizzi di sugo per la salama).

2. **Victory Screen 3D GLB & Animazione Trionfale**:
   - Tutta la partita è in 2D pulito e leggero.
   - Appena si verifica un tris vincente:
     - La linea vincente si illumina con raggi dorati.
     - L'icona del vincitore si stacca dalla griglia e si trasforma in un **modello 3D (.glb)** rotante in primo piano (Three.js / WebGL / `@google/model-viewer`).
     - Se vince la Coppietta: la ciupèta 3D compie una rotazione gloriosa su se stessa con luce celestiale e briciole dorate.
     - Se vince la Salama: la salama da sugo 3D balla, gira a 360° grondando sugo glorioso sul purè di patate virtuale.

3. **Audio Vocale Dialettale Autentico**:
   - Audio finale registrato da Riccardo: grido trionfale in ferrarese:  
     ***«Ma io... AL TRIS! MA IO AL TRIS!»*** (con eco epico da sagra di paese).
   - Sound FX: croccantezza "crunch" quando si piazza la ciupèta, "splotch" unto quando si piazza la salama da sugo.

4. **Multiplayer Locale & Solo vs Bot**:
   - **Pass and Play**: 2 giocatori sullo stesso schermo (telefono o tablet).
   - **Contro la Zia di Bondeno (Bot)**: IA euristica con 3 livelli:
     - *Insonnolita dal lambrusco* (mosse a caso).
     - *Normale*.
     - *Implacabile* (Minimax imbattibile con commenti in dialetto ferrarese).

---

## Technical Architecture

- **Piattaforma**: Single-page Web App (PWA) / Canvas HTML5 + Three.js per i render 3D.
- **Rendering 2D**: Canvas HTML5 o SVG/CSS grid per rendering istantaneo a 60fps su qualsiasi smartphone.
- **Modelli 3D (.glb)**:
  - `assets/models/coppietta_ferrarese.glb` (croce di pane a 4 cornetti, texture dorata e croccante).
  - `assets/models/salama_da_sugo.glb` (sfera legata con spago e picciolo, texture unta e speziata).
- **Audio Engine**: Web Audio API con pre-buffering degli asset audio (`assets/audio/al_tris_win.mp3`).
- **Dimensione Massima**: Totale bundle $< 5\text{ MB}$ per caricamento fulmineo anche sotto 3G/4G in bicicletta sugli argini del Po.

---

## Assets & TODO Backlog

- [ ] **Modello 3D GLB**: Modellare o scansionare (via fotogrammetria / NeRF) la ciupèta ferrarese e la salama da sugo.
- [ ] **Audio Voice Pack**: Riccardo deve registrare il sample audio ufficiale: *"Ma io al tris! Ma io al tris!"*.
- [ ] **Sound Design**: Campionare suoni di pane spezzato e sfoggio di sugo.
- [ ] **Repository Implementativo**: Creare repo standalone `~/git/al-tris` (o `palladius/al-tris`).
