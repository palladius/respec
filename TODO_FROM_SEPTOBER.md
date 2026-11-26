# 📐 Respec & Code Backlog (Travasati da Septober)

> Idee progettuali e spec in `~/git/respec/specs/{apps,tools,games,agents,meta}/<slug>/`.

---

- [ ] 🟠 P1 **[Septober #209]** **Agent on Rails: opinionated framework e CLI convenzioni per agenti AI (ADK, folders, evals)**
  - **Due:** `2026-11-02` | **Where:** `Derek / Zurich` | **Tags:** `#ai #architecture #rails #agents #adk #evals #cli #framework`
  - **Contesto:** Idea architetturale di Riccardo: 'Agent On Rails' (un Ruby on Rails per agenti AI autonomi basato su ADK).

Concetti chiave (Convention over Configuration):
- Struttura directory standard e opinata (app/agents, app/skills, app/tools, app/evals, config/).
- Dove collocare i progetti, le memorie, i test e soprattutto la suite di valutazione (evals standardizzate).
- CLI di scaffolding ('agent new', 'agent g skill ...', 'agent test/eval').
- Da analizzare e progettare insieme lunedì 2 Novembre 2026.

- [ ] 🟠 P1 **[Septober #170]** **Sviluppare Palladius Karaoker (K-A-R-A-O-K-A-R) frontend web**
  - **Due:** `2026-10-02` | **Where:** `Zurich` | **Tags:** `#tech #okr #karaoke #respec #music` | **URL:** file:///Users/riccardo/git/respec/specs/apps/karaoker/SPEC.md
  - **Contesto:** Ideona per OKR di fine anno: Palladius Karaoker (K-A-R-A-O-K-A-R). Frontend con BYOK API key che prende un prompt musicale, fa partire l'audio generato immediatamente via Lyria, genera in background copertina CD e SRT coi timestamp, per poi lanciare il player interattivo karaoke.html (come a Campobasso). Spec salvata in git/respec/specs/apps/karaoker/SPEC.md.

- [ ] 🟠 P1 **[Septober #191]** **Aggiungere asset per Al Tris (3D GLB e Audio)**
  - **Due:** `2026-10-23` | **Where:** `Studio Giallo T27 / Respec` | **Tags:** `#respec #games #ferrara #3d #audio #goliardia` | **URL:** file:///Users/riccardo/git/respec/specs/games/al-tris/SPEC.md
  - **Contesto:** Riccardo deve fornire gli asset per il gioco 'Al Tris': 1) Modelli GLB 3D della Coppietta Ferrarese (croce) e della Salama da Sugo (cerchio). 2) Registrazione vocale urlata 'Ma io... AL TRIS! MA IO AL TRIS!' per la victory screen animata in 3D.

- [ ] 🟡 P2 **[Septober #208]** **Studiare Spec Kit (github/spec-kit) per Spec-Driven Development**
  - **Due:** `2026-11-04` | **Where:** `Codemotion Milan / Derek` | **Tags:** `#slides #respec #spec #github #agents #codemotion #sdd` | **URL:** https://github.com/github/spec-kit
  - **Contesto:** Studiare Spec Kit (Specification Kit - github/spec-kit), il toolkit ufficiale di GitHub per lo Spec-Driven Development (SDD).
Obiettivo: analizzare convenzioni, formati di spec e workflow di orchestrazione con agenti AI, confrontandoli con il nostro framework in respec per arricchire slide, talk e pratiche di engineering.

- [ ] 🟡 P2 **[Septober #205]** **Studiare Pi Durable (earendil-works/pi) per memoria e stato durabile agenti SRE**
  - **Due:** `2026-11-04` | **Where:** `Zurich / Derek` | **Tags:** `#ai #slides #agents #sre #harness #durable` | **URL:** https://github.com/earendil-works/pi/tree/main/packages/durable
  - **Contesto:** Durable agent harness sperimentale di earendil-works/pi: commit atomici su SQLite/JSONL/Cloudflare DO, transcript immutabili, crash-recovery garantito, compaction automatica del contesto e subagent replay-safe. Ottimo spunto architetturale per la gestione di memoria persistente e stato negli agenti SRE da integrare nelle slide.

- [ ] 🟡 P2 **[Septober #189]** **Refactor tassonomia e cartelle Respec**
  - **Due:** `2026-10-23` | **Where:** `Repo git/respec` | **Tags:** `#respec #refactor #taxonomy #code`
  - **Contesto:** La cartella ~/git/respec è disorganizzata con specs/games, giochi/, samples/games, docs/tools, carlesso-family/. Definire tassonomia standard specs/{category}/{slug}/ e ripulire.

- [ ] 🟢 P3 **[Septober #210]** **Studiare Mastra per Evals e confronto con ADK e Langfuse**
  - **Due:** `2026-11-02` | **Where:** `Derek / Zurich` | **Tags:** `#ai #architecture #agents #adk #evals #mastra #langfuse` | **URL:** https://mastra.ai
  - **Contesto:** Valutare Mastra (framework agentico completo in TS/JS, 'tipo un ADK fatto dalle forze del male') con particolare focus sul layer di Evaluation (Evals sopra il framework) e benchmarking rispetto a Langfuse e all'architettura ADK / Agent on Rails.

- [ ] 🟢 P3 **[Septober #161]** **Rails App: bottone per cancellare immagine post e triggerare rigenerazione AI**
  - **Due:** `2026-09-16` | **Where:** `mini-lobby / Rails codebase` | **Tags:** `#ai #rails #ruby #ui #feature`
  - **Contesto:** Aggiungere un bottone nell'app Ruby on Rails per cancellare l'immagine di un post, in modo che alla cancellazione si inneschi automaticamente la rigenerazione con AI.

- [ ] 🟢 P3 **[Septober #177]** **Kids Time Machine: revisione spec e kickoff implementazione**
  - **Due:** `2026-10-12` | **Where:** `host-macbook` | **Tags:** `#ai #kids #spec #coding #photos`
  - **Contesto:** Revisionare la specifica in ~/git/respec/specs/apps/kids-time-machine/SPEC.md: timeline mosaic ultra-veloce su Node.js, mappa viaggi e pipeline face-morphing per Ale e Sebi da neonati a oggi
