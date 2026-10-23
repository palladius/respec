---
trigger: always_on
---

# Re-SPEC Project Rules (`respec`) 🦖

## 1. Repository Boundary: SPECS ONLY HERE! 🛑
- **This repository hosts EXCLUSIVELY specifications (`SPEC.md`), prompts (`input_prompt.md`), notes, metadata, and at most initial draft assets (screenshots, starting sketches).**
- **NEVER implement apps or games inside this repository**: implementations must always live in other dedicated git repositories or external sandboxes!
- If the user or a prompt asks to implement an app specified in `respec`, politely decline to code it inside this directory tree and propose creating or switching to a dedicated repository (e.g. using the `spec-to-code-tdd` skill).

## 2. Adherence to `speck` Specifications
- Development of the `speck` CLI tool itself (in Go) must strictly follow `docs/META-SPECS.md` and `docs/SPECS.md`.
- Do not modify specifications unilaterally without explicit user confirmation. When needed, propose concise changes (1-2 lines).
- Promptly report any observed "drift" between code and specifications to the user.

## 3. Code Quality & Versioning
- Before concluding any modification to `speck` Go code, always run `just test` (or `go test ./...`).
- Keep `CHANGELOG.md` updated consistently with every feature or fix.

## 4. Security & Protected Files
- **NEVER modify, create, or touch `.env` files** under any circumstances without explicit permission.
- **NEVER touch the `.riccardo/` directory** under any circumstances (strictly human-only read-only domain).
- Never use destructive git commands such as `git reset --hard` or `git clean -fd`.
