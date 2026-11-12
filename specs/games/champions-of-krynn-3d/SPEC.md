---
speck_version: "0.2"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-11T09:45:00Z"
updated_at: "2026-11-12T11:25:00Z"
title: "Campioni di Kream 3D / Champions of Krynn 3D (Party Shooter & ARPG Hybrid)"
tags: [game, 3d, dragonlance, kream, threejs, webgl, quake, diablo, jericho, party-mechanic, plastilina, hunyuan3d, niccolo-game-design]
author: "Ermete Bottazzi & Riccardo Carlesso"
source_of_inspiration: "Champions of Krynn (SSI 1990) + Clive Barker's Jericho + Quake / Hexen + Diablo 3 + CarChess3D Plastilina"
implementation_repo: "palladius/krynn-smandrupume"
---

# 🐉 Campioni di Kream 3D (Champions of Krynn 3D)

> *"Un incrocio viscerale tra Quake, Clive Barker's Jericho e Diablo 3: party di 6 eroi leggendari con rotazione istantanea del personaggio giocabile (Body-Swapping `TAB`/`1-6`), tripla telecamera (`V`: 1st Person Quake, 3rd Person Tomb Raider, Isometric 45° Diablo), e ordate di Dragonzoli balcanici, Goblin, Ettin e Draghi Cromatici da polverizzare a colpi di Fireball, Fiondona e spadone!"*

---

## 1. Vision & Core Concept

Portare il classico GDR fantasy anni '80 **Champions of Krynn** (SSI Gold Box, 1990) nel web gaming 3D moderno in [`krynn-smandrupume`](https://github.com/palladius/krynn-smandrupume), fondendo l'esplorazione delle rovine sommerse (già estratte dai blocchi binari `.DAX` in `src/dungeon3d.js`) con un **Action-RPG / Boomer-Shooter a squadre in tempo reale**.

Il giocatore guida l'iconico party di **6 Compagni** attraverso rovine gotiche, templi ottagonali e caverne invase da ordate di **Dragonzoli** (*Bazić*, *Bozan*, *Kapać*), **Goblin**, **Ettin** (giganti a due teste) e **Draghi Cromatici** (*Drago Nero* e *Drago Blu*).

### Precedenti Illustri nel Game Design
1. **Clive Barker's Jericho (2007)**: Meccanica identica! FPS a squadre dove il giocatore salta istantaneamente nel corpo di 6–7 specialisti diversi (`TAB` / `1-6`), mentre i compagni non controllati combattono come bot IA autonomi al seguito.
2. **Hexen: Beyond Heretic (1995)**: Il capostipite dark-fantasy su motore Doom/Quake con Guerriero, Mago e Chierico che massacrano Ettin e mostri in prima persona.
3. **Diablo 3 (2012)**: Leggibilità tattica isometrica a 45°, AoE esplosivi, gestione della folla (crowd control) e pioggia di loot.
4. **Anachronox (2001, id Tech 2 / Quake II) & FFVII Remake**: Real-time action 3D con party al seguito e cambio al volo dell'eroe attivo.

---

## 2. Architettura Legale & Naming "Sotto il Radar" (SRD 5.1 CC-BY-4.0 + Svolta Italo-Balcanica)

Per rimanere al **100% dentro un perimetro legale blindato e Zero-Royalty**, sfruttiamo la distinzione netta tra **Open Game Content (D&D SRD 5.1 rilasciato sotto Creative Commons `CC-BY-4.0` irrevocabile + folklore pubblico)** e la **Product Identity di Wizards of the Coast** (nomi propri dei romanzi *Dragonlance*).

### 2.1 Cosa possiamo usare COL VERO NOME D&D (100% Libero bajo CC-BY-4.0 SRD)
* **Mostri Classici D&D**: **Ettin** (gigante bicipite, dall'antico inglese *eoten*), **Goblin**, **Hobgoblin**, **Orc**, **Ogre**, **Troll**, **Kobold**, **Skeleton**, **Zombie**, **Lich**, **Golem** (Iron, Stone, Clay, Flesh), **Wyvern**, **Black Dragon**, **Blue Dragon**, **Red Dragon**, **Green Dragon**, **White Dragon**.
* **Razze & Classi**: **Halfling**, **Gnome** (*Rock/Tinker Gnome*), **Dwarf**, **Elf**, **Half-Elf**, **Fighter**, **Paladin**, **Wizard**, **Cleric**, **Rogue**, **Ranger**.
* **Incantesimi & Meccaniche**: **Fireball**, **Magic Missile**, **Burning Hands**, **Lightning Bolt**, **Chain Lightning**, **Cure Wounds**, **Bless**, **Protection from Evil**, **Hold Person**, **Armor Class (AC)**, **Hit Points (HP)**, **Saving Throws**, **d20**.

### 2.2 Nomi Propri "Sotto il Radar" (I 6 Compagni della Briscola & I Dragonzoli Slavi)
Tutti i nomi protetti da *Product Identity* Dragonlance vengono sostituiti nel modo **Sanitized ("Campioni di Kream")** da una localizzazione **italo-rustica per gli eroi** e **slavo-balcanica per i Dragonzoli** (con toggle opzionale in-game per i puristi):

| Archetipo Originale | Nome "Campioni di Kream" (Sotto il Radar) | Razza / Classe (SRD-Safe) | Arma & Attacco Base (`LMB`) | Abilità Speciale (`RMB` / `Spazio`) |
| :--- | :--- | :--- | :--- | :--- |
| **Caramon Majere** | **Carmine "Cremona" Maggioni** | Umano Guerriero Tank | **Spadone a due mani** (Fendente Cleave ad arco 120°) | **Urlo di Cremona & Colpo di Scudo** (Taunt + Knockback + Stun 1.5s) |
| **Raistlin Majere** | **Rastellino "Raisi" Maggioni** | Umano Mago (Veste Nera/Rossa) | **Bastone di Magio** (*Magic Missile* a ricerca, comando *"Scirocco!"*) | **Fireball (Palla di Fuoco)** AoE esplosiva / *Burning Hands* |
| **Tanis Half-Elven** | **Tano Mezzabarba** | Half-Elf Ranger / Leader | **Arco Lungo** (Frecce veloci a lungo raggio) | **Pioggia di Frecce Perforanti** (Multishot a ventaglio + Rally Buff) |
| **Crysania** | **Cristiana da Taranto** | Umana Cleric (Reverenda) | **Mazzapicchio Benedetto** + dardo sacro | **Protection from Evil & Cura AoE** (Scudo dorato + *Mass Cure Wounds*) |
| **Tasslehoff / Gnosh** | **Tassilo Piedipelosi** *(o **Mastro Gnosh**)* | **Halfling Scout** *(o **Gnome Inventor**)* | **La Fiondona** (bastone-fionda ferrato a ripetizione) / Balestra a Vapore | ** Sbeffeggio Rumoroso & Granata Alchemica** (Confonde i mob + Backstab) |
| **Sturm Brightblade** | **Sturmino Lamachiara** | Umano Paladin (Cavaliere di Solagna) | **Spada Antica d'Onore** (*Lamachiara*) | **Carica di Solagna** (Dash frontale con onda d'urto stordente) |

*(Inoltre: l'arcimago rivale *Dalamar* diventa **Dalmazio il Nero**, la città sommersa di *Xak Tsaroth* diventa **Zagarolo Sommersa**, e la dea *Mishakal* diventa **Misericordia / Myron**).*

---

## 3. Bestiario & Meccaniche Fisiche di Morte in 3D

Il vero colpo di genio del combattimento in tempo reale è trasformare le regole a turni del Gold Box 1990 in **meccaniche fisiche spaziali 3D**:

1. **Dragonzolo Bazić** (*ex-Baaz Draconian* — Fante Scaglioso d'Assalto):
   - Corre in branco verso il party con scimitarra e planata bassa.
   - **Tratto di Morte (Pietrificazione Istantanea)**: A `0 HP`, il modello 3D si trasforma in **statua di pietra grigia scolpita** che rimane sul pavimento per `8 secondi` come **ostacolo fisico e copertura reale** (blocca i proiettili!).
   - **Incastro della Spada (Weapon Stuck)**: Se lo uccidi in mischia con Carmine o Sturmino senza fare *step-back*, la tua lama resta **incastrata nella statua per `1.2s`**. *Trick da Pro*: premi `2` o `TAB` per fare **Body-Swap istantaneo su Rastellino Raisi** e lanciare una *Fireball* mentre il bot di Carmine sfila automaticamente la spada!
2. **Dragonzolo Bozan** (*ex-Bozak Draconian* — Stregone Scaglioso):
   - Bombardiere a distanza con scaglie di bronzo, lancia *Magic Missile* e globi arcani.
   - **Tratto di Morte (Esplosione d'Ossa AoE)**: A `0 HP`, le sue scaglie si prosciugano e lo scheletro **esplode in una deflagrazione viola/arancio** (raggio `3.5m`), infliggendo danno AoE agli eroi (e agli altri nemici!) vicini. Va eliminato a distanza con Tano o Tassilo!
3. **Dragonzolo Kapać** (*ex-Kapak Draconian* — Assassino Veleno):
   - Lame avvelenate; alla morte rilascia una **pozza d'acido verde ribollente** sul pavimento per `6 secondi`.
4. **Goblin & Orcs (SRD)**:
   - Carne da macello veloce che attacca a sciami, perfetta da spazzare via con la *Fireball* di Raisi o il fendente di Cremona.
5. **Mini-Boss Elite — Ettin (Gigante Bicipite SRD)**:
   - Alto `2.80m`, impugna due clave chiodate, immune a stun leggeri, attacco schianto a terra (*Ground Slam*).
6. **Boss Finali — Draghi Cromatici Antichi (già modellati con anatomia v0.3.0 in `src/dungeon3d.js`)**:
   - **Crisante l'Ossidiana (Ancient Black Dragon / *Khisanth*)**: corna taurine ricurve in avanti, soffio di acido smeraldo in linea/cono, pozze corrosive nella caverna sommersa.
   - **Calandro il Fulmine (Ancient Blue Dragon / *Khellendros*)**: corno nasale da rinoceronte, soffio di fulmini tridimensionali ramificati (`TubeGeometry`).

---

## 4. La "Trifecta" delle Telecamere (`V`) + Body-Swapping (`TAB` / `1-6`)

Invece di scegliere tra Modalità A (*Diablo*) e Modalità B (*Quake/Jericho*), in Three.js usiamo **un'unica scena 3D** con un `CameraManager` che commuta istantaneamente tra **3 visuali giocabili** col tasto `V` (più la modalità Museo):

1. **Modalità 1 — `1st Person Boomer Shooter` (Stile *Quake / Hexen*)**:
   - Telecamera ad altezza occhi (`y = 1.65m` per umani, `y = 1.15m` per Tassilo/Gnosh).
   - **Viewmodel 3D dell'arma in primo piano** in basso a destra/centro (Spadone di Cremona, Bastone illuminato di Raisi, Arco di Tano, Mazzapicchio di Cristiana, Fiondona di Tassilo) con oscillazione di camminata (*head-bob*) e animazione di swing/cast al click.
   - Mirino centrale (crosshair) dinamico.
2. **Modalità 2 — `3rd Person Over-the-Shoulder` (Stile *Tomb Raider / Jericho*)**:
   - Telecamera posizionata a `3.8m` dietro e `0.6m` a destra della spalla dell'eroe attivo (`y = 2.4m`).
   - Permette di vedere il proprio eroe per intero, le animazioni del fendente e i 5 compagni che fiancheggiano il leader.
3. **Modalità 3 — `Tactical Isometric 45°` (Stile *Diablo 3*)**:
   - Telecamera prospettica/ortografica inclinata a `45°` dall'alto (`offset = [0, 14, 12]`) ancorata all'eroe attivo.
   - Ideale quando la stanza si riempie di 15 Dragonzoli ed Ettin per piazzare *Fireball* e *Cure AoE* con precisione chirurgica sul terreno.
4. **Modalità 4 — `Museo / Auto-Tour Cinematografico` (Tasto `M` o `0`)**:
   - Preserva al 100% il tour panoramico delle 5 stazioni di Xak Tsaroth / Zagarolo Sommersa già presente in `src/dungeon3d.js`.

### Meccanica di Body-Swapping (`TAB` o Tasti `1-6`)
- Premendo `1..6` (o `TAB` per ciclare al prossimo eroe vivo), il controllo passa istantaneamente al nuovo eroe.
- La telecamera esegue un **Lerp parabolico rapido di `220ms`** dal vecchio corpo al nuovo (con un leggero flash cromatico del colore dell'eroe), rendendo immediata la percezione spaziale di dove ci si trova nell'arena.
- L'eroe appena lasciato passa senza soluzione di continuità sotto il controllo del `CompanionAIManager`.

---

## 5. Pipeline 3D "Plastilina" (`$niccolo-game-design` + `carchess3d`) 🎨➡️🧊

I modelli procedurali fatti di sfere e cilindri in `src/dungeon3d.js` (eccetto i Draghi v0.3.0 che hanno già fauci, zanne, corna e ali articolate) servono solo come **fallback zero-byte**. Per i personaggi e i mostri adottiamo la **Pipeline Plastilina a €0 di budget** sviluppata e validata in `carchess3d` (`tools/model-pipeline/RUNNING_DOC.md`):

```
[Prompt + Reference]
       │
       ▼
1. Nano Banana 2D (`gemini-3-pro-image-preview`)
   Turnaround Sheet (Front 0° + Side 90° + 3/4) su sfondo neutro
   (+ opzionale Face-Grafting biometrico Mode B per dare a Carmine/Raisi facce reali!)
       │
       ▼
2. Generazione Mesh 3D `.glb` su `pupurabbux` (RTX 5080 16GB) o HF ZeroGPU
   ├─ Via Primaria: Tencent Hunyuan3D-2 / 2.1 o TRELLIS (100% legale in Svizzera, €0)
   └─ Via Fallback Locale: Mode A Visual-Hull SDF (`54×74×54`) + Marching Cubes
      + Taubin Smoothing + Seamless `1024×1024` UV Atlas (`reconstruct_mode_a_batch.py`)
       │
       ▼
3. Post-Processing & LOD Decimation in Headless Blender (`blender -b -P ...`)
   ├─ Pulizia: `remove_degenerate`, rimozione briciole flottanti (`< 50 faces`)
   ├─ Grounding rigoroso: `min.y == 0.000m`, scala canonica (`1.80m` eroi, `1.20m` Halfling/Gnomo, `2.80m` Ettin)
   ├─ Multi-Tier LOD (`glb-lod-decimation-blender.md`):
   │    • `<unit>.glb` (High: 15k–22k tris per eroe attivo / ravvicinato)
   │    • `<unit>.lod1.glb` (Medium: ~4k tris per companion e ordate di mob)
   └─ Fist Socket Calibration: attacco modulare per Spadone, Bastone di Magio, Fiondona
       │
       ▼
4. Adversarial VLM + Geometric Critic (`critic_geo.py` + 4-Angle Turntable Sheet)
   Gate automatico: watertight, grounded, `<= 25k tris`, score estetico `>= 8.5/10`
```

### Hot-Swap Trasparente in Three.js (`AssetRegistry`)
Il codice Three.js cerca automaticamente `assets/models/<unit_id>.glb` (con cache-busting `?v=<VERSION>`). Se il `.glb` generato da `pupurabbux` è presente su disco lo carica e gli aggancia l'arma; altrimenti usa il rig procedurale low-poly come fallback istantaneo (così il gioco funziona sempre anche offline o durante lo sviluppo).

---

## 6. Architettura del Codice secondo `$niccolo-game-design` (Zero God-Loop!)

Per evitare il classico anti-pattern di un `update(dt)` monolitico da 2.500 righe, strutturiamo il motore 3D seguendo i pattern di **Robert Nystrom (*Game Programming Patterns*)** e la skill `$niccolo-game-design`:

```
src/
├── dungeon3d.js              # Scena architettonica 3D (Xak Tsaroth / Zagarolo) + Builder modelli fallback
├── core/
│   ├── GameEventBus.js       # Pub/Sub tipizzato (HERO_SWAPPED, ENEMY_PETRIFIED, BOZAN_EXPLODED, WAVE_CLEARED...)
│   └── GameContext.js        # Service Locator / DI container (niente singleton globali sparsi)
├── managers/
│   ├── PartyManager.js       # Possiede i 6 eroi, gestisce il Body-Swap (1-6/TAB), HP/Mana e formazione a cuneo
│   ├── CompanionAIManager.js # Macchina a stati per i 5 eroi non controllati (Follow -> Aggro -> Cast/Attack -> Heal)
│   ├── EnemyWaveManager.js   # Spawna ondate di Bazić, Bozan, Kapać, Goblin, Ettin e Boss Draghi
│   ├── CombatManager.js      # Collisioni, cleave, Fireball AoE, Pietrificazione Bazić (spada bloccata), Esplosione Bozan
│   ├── ProjectilePool.js     # Object Pool pre-allocato per frecce, sassi di Fiondona, Magic Missile, Fireball e particelle
│   ├── CameraManager.js      # Commutazione V (1st Person / 3rd Person / Isometric 45° / Tour) + Lerp di Body-Swap
│   └── DropManager.js        # Pity timer & pseudo-random loot (Pozioni, Pergamene, Monete d'Acciaio)
└── views/
    ├── HudView.js            # Ritratti dei 6 eroi (con barra HP/cooldown e tasto 1-6), indicatore arma bloccata, wave counter
    └── AudioView.js          # Listener sugli eventi di combattimento + Synth WebAudio / OPL2 FM
```

### Regole Architetturali (`$niccolo-game-design` Checklist)
- **Single Responsibility**: Ogni Manager possiede un solo dominio di dati e comunica mutazioni agli altri tramite `GameEventBus`.
- **Object Pooling (`ProjectilePool`)**: Zero `new THREE.Mesh()` o `new Projectile()` allocati dentro il loop per frame/sparo — frecce, dardi incantati e schegge d'ossa dei Bozan vengono riciclati da un pool fisso.
- **Seeded RNG**: Tutte le decisioni di spawn e drop usano un PRNG deterministico con seed per garantire test riproducibili.
- **Top Loop Pulito**: Il metodo principale `update(dt)` itera semplicemente la lista ordinata dei Manager (`PartyManager -> CompanionAIManager -> EnemyWaveManager -> ProjectilePool -> CombatManager -> CameraManager`).

---

## 7. Roadmap & Milestones di Implementazione

- **Milestone 1 — Core Party Shooter Loop & Tripla Camera (`V`)**:
  - Scomposizione in Manager (`GameEventBus`, `PartyManager`, `CameraManager`, `CombatManager`).
  - Tutti e **6 i Compagni della Briscola** (*Carmine Cremona, Rastellino Raisi, Tano, Cristiana, Tassilo con Fiondona, Sturmino*) schierati nella Rotonda e nel Colonnato.
  - **Body-Swapping (`1-6` / `TAB`)** con lerp camera e **Viewmodel 1st Person** + **3rd Person** + **Diablo Isometric 45°**.
- **Milestone 2 — Companion AI & Fisica dei Dragonzoli (*Bazić* & *Bozan*)**:
  - `CompanionAIManager`: i 5 compagni seguono l'eroe attivo e usano attacchi base e cure autonomamente.
  - Ondate di **Dragonzoli Bazić** (pietrificazione a 0 HP + blocco temporaneo della spada in mischia + statua come ostacolo) e **Dragonzoli Bozan** (esplosione d'ossa AoE a 0 HP) + **Ettin**.
- **Milestone 3 — Pipeline 3D "Plastilina" su `pupurabbux` (`Nano Banana -> Hunyuan3D/TRELLIS -> GLB LOD`)**:
  - Generazione delle Turnaround Sheet 2D con Nano Banana per i 6 eroi e il bestiario.
  - Conversione 2D$\rightarrow$3D `.glb` su `pupurabbux` (o Mode A SDF fallback), pulizia/decimazione LOD in Blender e caricamento automatico in Three.js.
- **Milestone 4 — Boss Fight nella Caverna di Crisante & Calandro**:
  - Scontro finale contro i due Draghi Cromatici v0.3.0 (getti d'acido, fulmini ramificati, pilastri come copertura e recupero del Tesoro Reale).
- **Milestone 5 — Co-Op Multiplayer Remoto (V2)**:
  - WebRTC / WebSocket per permettere a 2–6 amici di collegarsi e controllare ciascuno uno dei 6 eroi del party in contemporanea.
