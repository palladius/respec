---
speck_version: "0.1"
mode: oneshot
created_at: "2026-10-28T09:35:00Z"
category: tools
slug: carlessian-dashboard
title: Carlessian Observability Suite (Grafana + Prometheus)
---

# Carlessian Observability Suite 📊🚛

Una suite open source di osservabilità unificata (basata sullo stack Prometheus + Grafana + Blackbox Prober) per monitorare tutti i servizi, workstation/server personali e applicazioni del micromondo Carlessiano.

---

## 🎯 Obiettivi (Goals)

1. **Dashboard Unica ("Pannello di Controllo Carlessiano"):**
   - Raccogliere e visualizzare lo stato di salute di tutti i computer (MacBook, `mini-lobby`, Linux `derek`, Synology NAS).
   - Raccogliere e visualizzare lo stato di tutte le web app e servizi personali (Septober, Family Mahjong 3D, Chess 3D, bot Hermes/Lobby, siti web, API Cloud Run).

2. **Sintetici & Integration Testing Applicativo:**
   - Non solo ping HTTP statico (200 OK), ma controlli sintetici funzionali:
     - *Playability check:* verificare che l'interfaccia di un gioco (es. Chess 3D, Mahjong) risponda e si carichi senza crash JavaScript.
     - *API verification:* verificare endpoint chiave (es. API di Septober o backend Rails/Python).

3. **Deploy Freshness Assertion (Anti-Incagliamento):**
   - Asserire che la versione/commit deployata (esposta nel footer o su `/.well-known/version.json`) coincida esattamente con l'ultimo release tag / commit SHA su GitHub (`main`).
   - Evitare deploy "incagliati" (dove GitHub Actions fallisce silenziosamente o Vercel/Cloudflare/Cloud Run non aggiorna l'artefatto).

4. **Standard & Skill di Deployment ("Chiudere il Cerchio"):**
   - Definire un contratto minimo comune che tutte le app carlessiane devono rispettare (`/healthz`, meta tag di versione, convenzioni di build).
   - Realizzare una skill AI (`carlessian-app-contract`) per garantire che qualsiasi nuovo progetto creato dagli agenti aderisca automaticamente al monitoraggio.

---

## 🚫 Non-Goals

- Non vogliamo ricreare Datadog commerciale o infrastrutture complesse da enterprise multimilionaria.
- Nessun over-engineering: configurazione minima tramite `docker-compose`, config file dichiarativi in YAML e dashboard Grafana pre-confezionate.

---

## 🏗️ Architettura & Scelte Tecnologiche

### 1. Componenti Core
* **Prometheus:** Scrape delle metriche di sistema (`node_exporter`) e probe HTTP/TCP.
* **Prometheus Blackbox Exporter:** Probe per HTTP 200, certificati SSL in scadenza, tempi di risposta e DNS.
* **Grafana:** Dashboard visive semplici con alert configurabili (notifiche via Telegram / Discord / Hermes webhook).
* **Synthetic Prober / Runner:** Piccolo worker containerizzato (es. script Node/Playwright o Python sintetico) che esegue test d'integrazione leggeri end-to-end e pubblica metriche custom a Prometheus.

### 2. Hosting: Cloud vs Synology NAS (Decision Matrix)
* **Opzione A — Synology NAS (Consigliata per v1):**
  - Esecuzione continua 24/7 su Docker/Container Manager locale a costo zero.
  - Può fare polling sia della rete locale (LAN: stampanti, NAS, host casalinghi) sia degli endpoint pubblici.
* **Opzione B — Cloud (Cloud Run / VPS):**
  - Raggiungibile ovunque senza dipendere dalla connettività casalinga.
  - Costi minimi o free-tier, ma richiede tunneling sicuro (es. Tailscale) per vedere gli host interni.

---

## 📋 Contratto Carlessiano per le Applicazioni (`carlessian-app-contract`)

Ogni applicazione (web game, API, tool) espone:
1. **Endpoint di Sanità:** `GET /healthz` (restituisce `{"status": "ok", "uptime": ...}`).
2. **Endpoint di Versione:** `GET /version.json` (oppure attributi HTML standard nel footer):
   ```json
   {
     "repo": "palladius/family-mahjong-3d",
     "version": "1.0.4",
     "commit_sha": "a1b2c3d",
     "deployed_at": "2026-10-28T08:00:00Z"
   }
   ```
3. **Playability Hook:** Un selettore DOM o flag (`window.__APP_READY__ === true`) che conferma l'avvenuta inizializzazione del gioco/app per il prober sintetico.

---

## 🚀 Milestones di Implementazione

- [ ] **M1 — Repository & Compose Starter:**
  - Creare repository dedicato (es. `palladius/carlessian-dashboard` o `carlessian-watchtower`).
  - `docker-compose.yml` con Prometheus, Blackbox Exporter e Grafana con configurazione provisioned automatica.
- [ ] **M2 — Blackbox & Application Ping:**
  - Configurare Prometheus per monitorare gli endpoint correnti (Septober, DevFest, giochi web, blog).
  - Alert su certificati SSL in scadenza (<14 giorni) e downtime (>2 min).
- [ ] **M3 — Synthetic & Freshness Checker:**
  - Script prober leggero che confronta `git rev-parse HEAD` su GitHub con `deployed_at` / `commit_sha` dell'app live.
  - Alert su "Deployment incagliato".
- [ ] **M4 — Skill AI (`carlessian-app-contract`):**
  - Skill per Hermes/OpenClaw che inserisce automaticamente `/version.json` e `/healthz` nei nuovi template e progetti scafoldati.
