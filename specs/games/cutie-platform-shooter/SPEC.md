---
speck_version: "0.1"
mode: manual
idea_file: input_prompt.md
created_at: "2026-11-10T09:15:00Z"
category: games
slug: cutie-platform-shooter
title: "Cutie Platform Shooter: Fuga dalla Torre delle Verdure (N-Stage Save Game)"
tags:
  - games
  - platformer
  - shooter
  - roguelite
  - kids
  - family
  - puffin-login
  - phaser
  - html5
---

# Cutie Platform Shooter: Fuga dalla Torre delle Verdure 🎮🧸
*(2D Platform Spara-Tutto a Piani, Ironico e Adorabile con N Schemi e Salvataggio Puffin)*

## 1. Problem Statement & Vision

I tipici giochi spara-tutto per mobile (specie quelli visti nelle pubblicità a fette verticali di bunker con zombie e scale a pioli, vedi `assets/`) sono progettati con estetiche cupe o splatter, trappole punitive e livelli procedurali infiniti frustranti per bambini di prima e terza elementare.

**Cutie Platform Shooter** adatta questo format in chiave **ironica, spensierata e "cutie"**:
- **Stile Cutie & Comico**: Grafica colorata e cartoonesca. Niente sangue o orrori: i nemici sono mostriciattoli adorabili (broccoli ribelli, sveglie che suonano, calzini fluttuanti) che quando finiscono i punti vita svaniscono in un soffio di bolle di sapone o coriandoli (*"poof!"*).
- **Struttura a $N$ Schemi Finiti**: Niente mondo procedurale dispersivo o mappa gigantesca ingestibile. Il gioco è suddiviso in una sequenza chiara di **$N$ schemi discreti** (es. Schemi 1–10). Ogni schema ha dimensioni contenute (2-3 piani verticali, 1-2 scale a pioli) con un obiettivo chiaro: sconfiggere i nemici e raggiungere la porta/ascensore di uscita.
- **Sistema Upgrade a Scelta (3 Carte Roguelite)**: A fine ondata o al passaggio di livello, compare la schermata in stile "Choose Your Upgrade" (vedi `assets/screenshot_upgrade_selection.jpg`) con tre carte potenziamento buffe.
- **Salvataggio & Ripresa con Puffin Login**: Integrazione diretta con il componente `<puffin-login>` della famiglia Carlesso. Con il proprio profilo (es. Alessandro PIN `2018`, Sebastian PIN `2020`), il gioco memorizza lo schema esatto raggiunto e gli upgrade collezionati, consentendo di spegnere il tablet e riprendere esattamente da dove si era rimasti.

---

## 2. Personaggi & Meccaniche di Combattimento

### 2.1. Punti Vita (HP) & Combattimento
Sia i protagonisti che i nemici dispongono di **Punti Vita (HP)** numerici visualizzati con piccoli cuoricini o barre morbide sopra la testa:
- **Eroi**: Iniziano con 3-5 Cuori (ogni colpo nemico toglie 1 Cuore o frazioni).
- **Nemici**: Hanno da 2 a 10 HP. Quando colpiti lampeggiano di bianco con una simpatica espressione di sorpresa (stile *Cuphead* / *Kirby*). Quando i loro HP raggiungono lo zero, fanno un piccolo balzo all'indietro ed esplodono in un fumetto di stelline e rilasciano **Monete-Biscotto 🍪** o **Stelline ⭐**.

### 2.2. Protagonisti Selezionabili
1. **Alessandro (8 anni, Classe P3 - Il Cecchino Tattico)**:
   - *HP Base*: 4 Cuori.
   - *Velocità*: Media.
   - *Arma Base*: **Gommino-Blaster** (spara proiettili fluorescenti a linea retta con buon raggio).
   - *Tratto Speciale*: Precisione aumentata e ricarica rapida.
2. **Sebastian (6 anni, Classe P1 - Il Demolitore Vivace)**:
   - *HP Base*: 5 Cuori.
   - *Velocità*: Elevata (+15%).
   - *Arma Base*: **Spara-Macchinine** (spara proiettili pesanti a parabola che rimbalzano sui pavimenti).
   - *Tratto Speciale*: Resiste a un urto senza perdere cuori (scudo peluche iniziale).

---

## 3. Struttura del Livello: Il Modello a $N$ Schemi Discreti

Per rendere il gioco semplice, rilassante e pedagogicamente gratificante, il mondo è diviso in **Mondi tematici composti da $N$ schemi brevi (durata: 60-90 secondi a schema)**:

```
+-------------------------------------------------------------+
| PIANO 3: [Porta d'Uscita / Ascensore 🚪]   [Mostro Boss/Loot]|
|          ========================= [Scala a Pioli 🪜] ======|
| PIANO 2: [Broccolo Vagante 🥦]       [Cassa Sorpresa 🎁]     |
|          ====== [Scala a Pioli 🪜] =========================|
| PIANO 1: [Punto di Spawn Eroe 👦]    [Sveglia Zombie ⏰]     |
+-------------------------------------------------------------+
```

### 3.1. Anatomia di uno Schema
1. **Layout a 2 o 3 Piani Verticali**: Piattaforme orizzontali solide collegate da una o due scale a pioli verticali (come negli screenshot di riferimento).
2. **Scala a Pioli (Ladder)**: L'eroe sale e scende premendo Su/Giù. Sulla scala si è vulnerabili ma si può sparare in orizzontale.
3. **Flusso di Gioco**:
   - Fase A: Spawn dell'eroe al piano inferiore.
   - Fase B: Eliminazione dei nemici cutie presenti nelle stanze.
   - Fase C: Attivazione della porta d'uscita (si illumina di verde con freccia animata).
   - Fase D: L'eroe tocca la porta $\rightarrow$ Schermata di vittoria dello schema $\rightarrow$ Selezione Upgrade $\rightarrow$ Auto-Save $\rightarrow$ Passaggio allo Schema successivo.

### 3.2. I Mondi di Gioco
- **Mondo 1: La Casa & La Cameretta (Schemi 1 – 5)**:
  - Nemici: *Broccoli Zombie* (lenti, dicono *"gnam?"*), *Sveglie Saltanti* (saltellano a tempo di tic-tac), *Polverini Fluttuanti*.
- **Mondo 2: La Scuola & I Compiti (Schemi 6 – 10)**:
  - Nemici: *Matite Giganti* (caricano in linea retta), *Gomme da Cancellare* (assorbono i colpi), *Libri di Grammatica Volanti*.
- **Mondo 3: Il Giardino & Il Garage dei Camion 🚛 (Schemi 11 – 15)**:
  - Nemici: *Tubi d'Irrigazione Birichini*, *Pneumatici Rotolanti*, *Mini-Ruspe Giocattolo*.

---

## 4. Roguelite Upgrade System ("Choose Your Upgrade")

Alla fine di ogni schema (o quando si riempie la barra di ESP raccogliendo biscotti), il gioco mette in pausa l'azione e mostra un overlay con **3 Carte Upgrade pescate a caso**:

```
+---------------------------------------------------------------+
|                    SCEGLI UN POTENZIAMENTO!                   |
+---------------------------------------------------------------+
|  [ 🍕 Cappelletto ]    |  [ 👟 Scarpa Turbo ]   |  [ 🔫 Triplo-Shot ] |
|  Recupera 2 Cuori      |  Velocità +25% e       |  L'arma spara 3     |
|  e aumenta la salute   |  doppio salto con      |  proiettili a ventaglio|
|  massima di 1 cuore.   |  scia di fumo 🚛.      |  invece di uno solo.|
|      [ SCEGLI ]        |       [ SCEGLI ]       |      [ SCEGLI ]     |
+---------------------------------------------------------------+
```

### Elenco Potenziamenti Ironici
* **Raffica di Cappelletti**: Aumenta il rateo di fuoco del 30%.
* **La Ciabatta Radiocomandata**: I proiettili inseguono debolmente il nemico più vicino.
* **Scudo di Peluche**: Genera un cricetino orbitante (Ghigna) che blocca un proiettile nemico ogni 15 secondi.
* **Suola Rimbalzante**: Permette di saltare sulla testa dei nemici stile Mario per infliggere danno senza subire colpi.
* **Faro del Camion 🚛**: I nemici colpiti vengono storditi per 1 secondo da una luce abbagliante.

---

## 5. Salvataggio & Ripresa con Puffin Login (`<puffin-login>`)

### 5.1. Integrazione Family Credentials
Il gioco include nella barra superiore il web component `<puffin-login>` con i profili famiglia pre-configurati:
- `ale` (Alessandro, PIN `2018`)
- `seby` (Sebastian, PIN `2020`)
- `riccardo` (Papino, Password)
- `kate` (Mamma, Password)
- `guest` (Ospite / Partita veloce senza credenziali)

### 5.2. Schema Dati del SaveGame (`save_state`)
Ogni volta che si supera uno schema o si sceglie un upgrade, lo stato viene serializzato in JSON:

```json
{
  "user_id": "ale",
  "character": "alessandro",
  "current_stage": 4,
  "highest_stage_unlocked": 4,
  "hearts_current": 4,
  "hearts_max": 5,
  "collected_upgrades": [
    "cappelletto_max_hp",
    "triplo_shot",
    "scarpa_turbo"
  ],
  "total_cookies": 142,
  "last_saved_at": "2026-11-10T10:14:00Z"
}
```

### 5.3. Flusso di Caricamento ("Riprendi da dove hai preso")
1. All'avvio dell'app, l'utente seleziona il proprio profilo (es. Alessandro inserisce `2018`).
2. Il client verifica se esiste un `save_state` associato a `puffin_save_ale` (su Firestore se connesso, o su `localStorage` in locale).
3. Se presente: compare il pulsante **"▶️ Continua Schema 4"** insieme a **"🔄 Nuovo Gioco"**.
4. Cliccando "Continua", l'eroe viene caricato direttamente all'inizio dello Schema 4 con tutti gli upgrade e i cuori sbloccati nella sessione precedente!

---

## 6. Architettura Tecnica & Controlli

### 6.1. Stack Software
- **Piattaforma**: HTML5 Canvas PWA responsive (ottimizzata sia per iPad/tablet in orizzontale 16:9 che per smartphone).
- **Game Engine**: **Phaser 3** (consigliato per gestione fisica arcade nativa, sprite animati a scale a pioli, sound manager e particelle) oppure motore custom Canvas 2D leggero.
- **Backend & Autenticazione**: `<puffin-login>` collegato all'endpoint Cloud Run (`https://puffin-login-932724269481.europe-west1.run.app/v1`).
- **Audio**: Web Audio API / Chiptune sintetico leggero a 140 BPM (zero dipendenze pesanti di terze parti).

### 6.2. Schema dei Controlli
- **Su Touchscreen / iPad**:
  - D-pad virtuale morbido a sinistra (Sinistra, Destra, Su/Giù per le scale a pioli).
  - Tasto **Salto (A)** grande a destra.
  - Tasto **Spara (B)** con opzione di **Autofire / Mira Automatica assistita** (fondamentale per permettere a Sebastian di 6 anni di giocare senza frustrazione).
- **Su Tastiera / Desktop**:
  - Movimento: `Frecce` o `WASD`.
  - Salto: `Barra Spaziatrice` o `W`/`Freccia Su`.
  - Spara: `J` o `Click Sinistro Mouse`.

---

## 7. Criteri di Accettazione & Verifica (DoD)

- [ ] **Grafica & Nemici Cutie**: Nessun elemento splatter; nemici chiaramente identificabili come oggetti comici con occhietti e faccine espressive; animazione di scomparsa a bolle/coriandoli.
- [ ] **HP System**: Cuori visibili per il giocatore; barre/pallini HP sopra i nemici. I colpi tolgono HP fino alla sconfitta.
- [ ] **N Schemi Discreti**: Almeno 5 schemi completabili in sequenza con caricamento istantaneo del successivo.
- [ ] **Upgrade Roguelite**: Schermata a 3 carte funzionante dopo ogni vittoria di schema con applicazione immediata delle statistiche modificate.
- [ ] **Integrazione Puffin Login & Save**:
  - Login con PIN di Ale (`2018`) e Sebi (`2020`).
  - Salvataggio automatico dello schema raggiunto in `localStorage` / backend.
  - Test di ripresa: ricaricando la pagina, il pulsante "Continua" fa ripartire dal livello salvato con gli upgrade intatti.
- [ ] **Responsive & Touch Friendly**: Testato e giocabile sia con d-pad touch su mobile che con tastiera su PC.
