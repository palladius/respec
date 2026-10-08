---
trigger: always_on
---

# Regole di Progetto Re-SPEC (`respec`) 🦖

## 1. Confini del Repository: QUI SOLO SPECIFICHE! 🛑
- **Questo repository ospita ESCLUSIVAMENTE specifiche (`SPEC.md`), prompt (`input_prompt.md`), note, metadati e al più asset iniziali (screenshot, bozze grafiche).**
- **NON implementare MAI le app o i giochi qui dentro**: le implementazioni vivono sempre in altri repository git dedicati o sandbox esterne!
- Se l'utente o un prompt chiede di implementare un'app specificata in `respec`, rifiutati educatamente di farlo in questo albero di cartelle e proponi di creare o spostarsi in un repository dedicato (ad esempio usando la skill `spec-to-code-tdd`).

## 2. Rispetto delle Specifiche di `speck`
- Lo sviluppo del tool CLI `speck` (in Go) deve rispettare rigorosamente `docs/META-SPECS.md` e `docs/SPECS.md`.
- Non modificare le specifiche da solo senza conferma esplicita dell'utente. Se necessario, proponi modifiche concise (1-2 righe).
- Segnala immediatamente all'utente qualsiasi "drift" osservato tra codice e specifiche.

## 3. Qualità del Codice & Versioning
- Prima di considerare conclusa qualsiasi modifica al codice Go di `speck`, esegui sempre `just test` (o `go test ./...`).
- Mantieni aggiornato `CHANGELOG.md` in modo coerente con ogni novità o fix.

## 4. Sicurezza & File Protetti
- **NON modificare, creare o toccare MAI file `.env`** senza autorizzazione esplicita.
- **NON toccare MAI la cartella `.riccardo/`** sotto alcuna circostanza (è un dominio strettamente umano ad accesso sola-lettura).
- Non utilizzare comandi git distruttivi come `git reset --hard` o `git clean -fd`.
