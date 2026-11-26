Potremmo creare uno scheletro / piattaforma Ruby on Rails di una app auto-aggiornante con workflows dove poi la specializziamo per fare:
1. sito di news a destra e sinistra (`auto-news`)
2. sito di AI news delle ultime 24h con podcast mattutino e jingle comico (`whats-up-with-ai`)
3. qualunque altra cosa futura.

Decisioni architetturali chiave (emerse da Drensin `/grill-me`):
- Un'unica piattaforma Rails Multi-Site (1 sola Admin UI centrale dei workflow + N siti/frontend serviti dallo stesso backend per dominio o path).
- Esecuzione su Cloud Run scale-to-zero vero: Cloud Scheduler + Cloud Tasks HTTP step-by-step + PostgreSQL (Supabase / Cloud SQL), dove ogni nodo del DAG è una chiamata HTTP idempotente, zero worker in-memory sempre accesi.
- Workflow management debuggabile graficamente da interfaccia Admin: devo vedere che N cron jobs sono partiti, quali sono andati male a metà, cliccare sul sotto-elemento rosso del workflow per vedere log/prompt/costi, e fare Replay deterministico dal nodo fallito ("Retry from failed step") riusando i payload JSON già salvati dei nodi padre.
- Crawler ibrido corazzato contro 403/Paywall: RSS/Atom ufficiali + Google News RSS + Gemini Search Grounding / Jina Reader.

- Nome dell'app/motore: **Doug** (omofono di DAG e personale — *"Diamo questo in pasto a Doug!"*).
