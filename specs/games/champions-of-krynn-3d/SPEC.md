---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-11T09:45:00Z"
title: "Champions of Krynn 3D (Dragonlance Party Shooter & ARPG Hybrid)"
tags: [game, 3d, dragonlance, threejs, webgl, quake, diablo, party-mechanic, tripo3d]
author: "Ermete Bottazzi & Riccardo Carlesso"
source_of_inspiration: "Champions of Krynn (SSI 1990) + Clive Barker's Jericho + Quake / Hexen + Diablo"
---

# 🐉 Champions of Krynn 3D

> *"Un incrocio tra Quake e Diablo nell'universo Dragonlance: party di 6 eroi leggendari con rotazione istantanea del personaggio giocabile (body-swapping), telecamera 1st/3rd person, e ordate di draconici, ettin e draghi da polverizzare a colpi di fireball e spadone!"*

---

## 1. Vision & Core Concept

Portare il classico GDR fantasy anni '80 **Champions of Krynn** (SSI Gold Box) nel web gaming 3D moderno via vibe-coding.

Il giocatore guida l'iconico party dei **Compagni della Lancia** (Caramon, Raistlin, Tanis, Crysania, Tasslehoff, Sturm) attraverso dungeon e rovine invase da ordate di draconici Baaz/Bozak, goblin, giganti a due teste (ettin) e draghi cromatici.

### Precedenti Illustri nel Game Design
Questa formula non solo è fattibile, ma ha illustri predecessori di culto:
1. **Clive Barker's Jericho (2007)**: Meccanica identica! FPS a squadre dove il giocatore salta istantaneamente nel corpo di 7 specialisti diversi (cecchino telecinetico, mago del sangue, soldato pesante), mentre i compagni non controllati combattono come bot IA autonomi.
2. **Anachronox (2001, id Tech 2 / motore di Quake II)**: RPG 3D con party esplorativo in terza persona ed eroi al seguito.
3. **Hexen: Beyond Heretic (1995)**: Il capostipite fantasy su motore Doom/Quake con Guerriero, Mago e Chierico che massacrano ettin e centauri in prima persona.
4. **Final Fantasy VII Remake / Dragon Age**: Real-time action 3D con telecamera alle spalle, un tasto (D-pad/TAB) per cambiare al volo eroe controllato, e companion AI attiva sul resto del gruppo.

---

## 2. Dual-Mode Architecture (Le Due Opzioni di Spec)

Come richiesto, la spec formalizza entrambe le strade esplorate da Riccardo:

### Modalità A: Isometric ARPG (Stile *Diablo* / 2.5D aumentato)
- **Visuale**: Telecamera ortografica / prospettica fissa a 45° dall'alto (*top-down isometric*).
- **Controllo**: Movimento con click del mouse o WASD; selezione eroi via hotbar 1-6 o box selection stile RTS.
- **Vantaggi**: Facile da implementare, ottima leggibilità tattica del campo di battaglia per 6 eroi simultanei, collisioni semplificate su piano 2D.
- **Limiti**: Meno viscerale, manca l'impatto claustrofobico e cinetico dell'azione in prima/terza persona.

### Modalità B: 3D Party Shooter (Stile *Quake / Jericho / Lara Croft*) — ⭐ **Target Principale**
- **Visuale**: Tasto `V` commuta tra **1st Person (Soggettiva)** e **3rd Person Over-the-Shoulder (Oggettiva alla Tomb Raider)**.
- **Controllo & Body-Swapping**:
  - Il giocatore muove con WASD + Mouse l'eroe attivo.
  - Tasto `TAB` o numeri `1-6`: **Hot-Swap istantaneo**. La telecamera fa un lerp rapido sul compagno selezionato. Quello appena abbandonato passa immediatamente alla Companion AI autonoma.
  - Gli altri $N-1$ membri del party seguono il leader in formazione a cuneo (flocking) e ingaggiano i nemici a raggio con le loro abilità native.
- **Vantaggi**: Esperienza adrenalinica da sparatutto arcade boomer, feedback audiovisivo devastante quando si passa da un fendente a una palla di fuoco.

---

## 3. Il Party dei Compagni (6 Archetipi)

| Eroe | Classe | Ruolo Primario | Attacco Base (LMB) | Abilità Speciale (RMB / Spazio) |
| :--- | :--- | :--- | :--- | :--- |
| **Caramon Majere** | Guerriero Tank | Mischia / Scudo | Fendente Spadone a due mani (cleave) | Scudo rotante / Taunt (attira i draconici) |
| **Raistlin Majere** | Mago Nero/Rosso | DPS Ranged AoE | Dardi Incantati (Magic Missile) | *Palla di Fuoco (Fireball)* / *Burning Hands* |
| **Tanis Mezzelfo** | Arciere / Leader | DPS Singolo Target | Frecce veloci a lungo raggio | Pioggia di frecce perforanti / Rally buff |
| **Crysania di Tarso** | Chierica Reverenda | Support / Healer | Mazzapicchio benedetto | *Protection from Evil* (scudo protettivo) / Cura AoE |
| **Tasslehoff Burrfoot** | Kender Ladro | Scout / Harass | Fionda Hoopuk a ripetizione | Taunt verbale rumoroso (confonde i mob) / Backstab |
| **Sturm Brightblade** | Cavaliere di Solamnia | Melee Bruiser | Spada solamnica d'onore | Carica d'impatto con stordimento nemici |

---

## 4. Bestiario & Minacce di Krynn

1. **Draconici Baaz**: Fanti alati di basso rango. *Tratto peculiare*: Quando muoiono, si pietrificano all'istante (se li colpisci in mischia la spada si blocca momentaneamente nella pietra!).
2. **Draconici Bozak**: Maghi draconici con incantesimi a distanza. *Tratto peculiare*: Quando muoiono esplodono in schegge d'ossa e scintille (danno AoE ai vicini!).
3. **Goblin & Orchetti dei Boschi**: Carne da macello veloce che sciami in branchi numerosi da spazzare via a grappoli.
4. **Mini-Boss & Mostri Elite**:
   - **Ettin**: Gigante a due teste con clava chiodata, immune agli stordimenti.
   - **Catoblepa**: Bestia con lo sguardo pestifero che rallenta e avvelena il party.
5. **Boss Finale**: **Drago Rosso di Krynn (Ember / Pyros)**: Volo sul perimetro dell'arena, getto di fiamme ad area, atterraggio sismico e coda spazzatutto.

---

## 5. Architettura Tecnica & Pipeline Vibe-Coding

* **Engine**: **Three.js** con vanilla JavaScript/TypeScript o HTML5 Canvas WebGL (nessuna dipendenza pesante, eseguibile su qualsiasi browser anche mobile/laptop con zero installazione).
* **Pipeline Asset 3D**:
  - Modelli 3D low-poly generati con **TRIPO 3D** (da prompt testuali o concept 2D generati con Nano Banana / Midjourney).
  - Esportazione in formato standard `.glb` / `.gltf` con animazioni base (Idle, Walk/Run, Attack, Death).
  - Texture retro-palette pixelate o toon shading per preservare il sapore fantasy anni '80.
* **Companion AI Loop**:
  - Se `distance(companion, active_hero) > FOLLOW_RADIUS`: naviga verso l'eroe leader.
  - Se `nearest_enemy < AGGRO_RADIUS`: punta il nemico e lancia l'attacco primario con cooldown periodico.
* **Audio & Suoni**:
  - Suoni metallici di spade, sibilo di frecce, boato delle fiamme e ruggiti mostruosi generati con sintesi web audio / procedurali.

---

## 6. Roadmap & Milestones

- **Milestone 1 (Arena Prototipo 3D)**: Ambiente Three.js con pavimento a piastrelle, 1 eroe (Caramon), 3 draconici nemici con semplice pathfinding verso il giocatore, fendente di spada e proiettili.
- **Milestone 2 (Body-Swapping a 3 Eroi)**: Caramon + Raistlin + Crysania. Tasto `TAB` per scambiare corpo. Companion AI funzionante per i membri non selezionati.
- **Milestone 3 (Full Party & Bestiario Tripo3D)**: Tutti e 6 i campioni con modelli GLTF da Tripo 3D, ondate di draconici Baaz/Bozak, ettin gigante.
- **Milestone 4 (Dragon Arena Boss Fight)**: Scontro col Drago Rosso in un tempio in rovina.
- **Milestone 5 (V2 Multiplayer Remote)**: WebRTC / WebSocket per permettere a un amico di unirsi da remoto e prendere il controllo di un membro del party (es. *"Io sono Raistlin, tu fai Caramon!"*).
