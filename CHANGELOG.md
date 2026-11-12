# Changelog

All notable changes to respec are documented here.

## Unreleased

## 0.3.4

- Added `specs/apps/telegram-voice-agents/` (`SPEC.md`, `README.md`, `blueprint.yaml`, `input_prompt.md`), with implementation pointer to [`palladius/korybantes`](https://github.com/palladius/korybantes): portable Flutter frontend (Google Pixel 10 first) featuring a Telegram-style multi-persona inbox, dual text/voice chat, single-tap hands-free *"Interrompimi"* (barge-in) full-duplex voice mode with `gemini-3.8-flash-tts`, and real-time semantic self-correcting transcription (*"voglio un gelato alla crema, anzi no, al pistacchio"* $\rightarrow$ *"Voglio un gelato al pistacchio"*).
- Consolidated all root-level scattered spec directories (`carlesso-family/`, `samples/`, `giochi/`) into `specs/games/` (`carlesso-family-chess-3d`, `pupurabbu-3d`, `flowconnect-pipe-puzzle`).
- Flattened all nested sub-categories (`specs/games/classic-arcade/*`, `specs/games/cloud-adventures/*`, `specs/meta/antigravity/*`) to the canonical 2-level `specs/<category>/<slug>/SPEC.md` taxonomy.

## 0.3.3

- Added `specs/games/tricolore-chess-3d/SPEC.md` (with `input_prompt.md` and `assets/scacchiera_fisica_riccardo.png`): pragmatic 2D/3D 3-player chess game (*TriBecca / Scacchi Tricolore 3D*, implemented at [`palladius/tribecca`](https://github.com/palladius/tribecca)) replicating Riccardo's 96-square wooden board (`A..L` × `1..12`), the 6-square central Rosette (`D4, D5, I5, I9, E9, E4`) rules (`NMLVaFImzVs`), the 4 elimination modes (`Destroy`, `Protoss Conquer`, `Flotsam`, `Sudden Death`), `PuffinLogin` seat-to-color mapping (`Ale`, `Seby`, `Papino`, `Computer` $\Longleftrightarrow$ `Bianco`, `Verde`, `Rosso`), $120^\circ$ turn camera rotation, and illuminated branching ray trajectories.

## 0.3.2

- Installed GitHub Spec Kit (`specify-cli` v1.2.4, `--integration agy`) with `assess`, `bug`, and `github` extensions under `.agents/skills/speckit-*` and `.specify/`.
- Updated `AGENTS.md` (and symlinked `GEMINI.md -> AGENTS.md`) recommending `specify-cli` installation and documenting the Spec Foundry vs Target Repo workflow.
- chess3d SPEC: 1-line pointer, the living spec is now `carchess3d/docs/spec/SPEC.md` (this copy is a frozen v1 snapshot).
- Moved the last `docs/` specs into `specs/`: `apps/plastilina`,
  `tools/sheets-autodetect-2026`, `tools/workshop-visualizer`.
  `carlesso-family/chess3d/` stays put (linked from `palladius/carchess3d`).
- Moved implementation artefacts (GLBs, Plastilina experiment scripts,
  `RUNNING_DOC.md`) to the new public repo
  [palladius/carchess3d](https://github.com/palladius/carchess3d); left a
  pointer stub. chess3d SPEC: implementation repo open question resolved.
- Added updated chess3d character assets: `knight-seby-{4,5}.jpg` (Gold Saint Sebi with sword/lance) and `queen-kate-{4,5}.jpg` (rejuvenated Queen Kate with bluish-purple gown and single tanzanite ring).
- Added `specs/apps/plastilina/SPEC.md` (image/prompt → game-ready 3D with
  a deterministic + VLM critique/repair loop and skeleton testing, spec only).
- Added `carlesso-family/assets/` chess3d character art
  (`<piece>-<character>[-N].jpg`); `pawn-puffin.jpg` (front + 3/4 view) is
  also the first Plastilina test input.
- Added first free TRELLIS 3D results `carlesso-family/assets/3d/pawn-puffin.glb`
  and `bishop-ale-1.glb`, the Plastilina `RUNNING_DOC.md` lab notebook +
  `experiments/scripts/` spikes; Plastilina SPEC gained 6 lessons learned.
- Added `carlesso-family/chess3d/SPEC.md` (3D family chess, spec only) and
  noted in `AGENTS.md` that this repo hosts only specs (`GEMINI.md` symlinked).
- Configured `.agents/` with project rules (`respec_rules.md`), pre-push test gate hook
  (`hooks.json`), and the `spec-to-code-tdd` skill for implementing specs in external repos.

## 0.3.1

- Migrated all externally generated specs from `docs/` to a dedicated `specs/` folder at the repository root, grouped by category (Games, Apps, Tools, Agents, Meta).
- Added `scripts/list_specs.py` and updated `just list` to display a nice tree-view of all specs with folder emojis.

## 0.3.0

- Renamed the repo and Go module path from `speck` to `respec`
  (`github.com/palladius/respec`) and updated the README title/branding to
  "Re-SPEC" — the CLI command/binary itself stays `speck` for now. The
  original `speck` bacon-pun logo is preserved under `assets/speck/`.
- Added a bonus Ali G quote to the README.

## 0.2.0

- Added a cute, fun, Google-themed logo to the README (`docs/logo.jpg`).
- Default output location is now `./out/` (gitignored), overridable via
  `--output-dir`/`-o` or the `SPECK_OUTPUT_DIR` environment variable.
- `.env` files (in the current directory) are loaded automatically via the
  new `internal/dotenv` package; explicit environment variables still win.
- `--file` gained the `-f` shorthand.
- The exact idea/prompt is now always saved as a sidecar `input_prompt.md`
  next to `SPEC.md`, replacing the old inline-vs-`original_idea.md` split.
- Default model changed from `gemini-2.5-flash` to `gemini-flash-latest`
  (`gemini-pro-latest` available for a smarter/slower model).
- Fixed a bug where some structured JSON responses came back with literal
  `\n`/`\t` escape sequences instead of real newlines/tabs in the spec body.
- Added `--version`.
- Added `samples/games/` with example specs (Gyruss, Pac-Man, and 10Rogue
  clones) generated by speck itself, plus their source ideas under `etc/`.
- Added a `justfile` with `build`, `gen`, `chat`, `test`, `check`, and
  `clean` recipes.

## 0.1.0

- Initial release: `speck oneshot` and `speck chat` (two-phase clarifying
  interview, per `docs/META-SPECS.md`), Gemini API backend, idea resolution
  from arg/`--file`/stdin/`IDEA.md`/interactive prompt, `--source-of-inspiration`
  directory scanning, model-chosen `<category>/<slug>/SPEC.md` output layout,
  YAML frontmatter provenance, shell completion.
