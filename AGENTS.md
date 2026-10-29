I'm Riccardo usually I code with Agy/Gemini cLI so its my first non-gemini.md in my life! :)


# Auth

Note this claude session is wrapped by my $GIC/bin/claudio-XXXX scripts. 
1. There's one in python to test that model is working and non fa i capricci.
2. Then there's the 'claudio' wrapper which exportes the right ENV variables. Notably it points to a different JSON for ADC.
3. Then there's the authenticate-as-ricc which is needed to create/refresh the JSOn in a non-standard position.

## Specs

0. **Hey! In this repo ONLY specs!** 🛑 Never implement the specified apps or games here: implementations live in OTHER dedicated repositories. This repo hosts exclusively specifications (`SPEC.md`, `input_prompt.md`), design notes, metadata, and at most initial draft assets (screenshots, starting sketches). Other agents will implement the specs elsewhere!
1. ensure you code things by following `docs/META-SPECS.md` and `docs/SPECS.md`
2. Do not change specs alone, always prompt user and ask for explicit confirmation to change those. if You do, do it in concise way (ie add 1-2 lines not a full paragraph or chapter!)
3. If you observe drift betwene code and specs, PLEASE let user know so you can take action.

## GitHub Spec Kit (`specify-cli`) & Skills

We strongly recommend installing [GitHub Spec Kit (`specify-cli`)](https://github.com/github/spec-kit) alongside `speck`:

```bash
uv tool install specify-cli
specify init --here --integration agy   # or --integration gemini / claude
specify extension add assess
specify extension add bug
specify extension add github
```

Installed skills in `.agents/skills/`:
- **In `respec` (Spec Foundry)**: use `/speckit-assess-*` (`intake` -> `research` -> `define` -> `shape` -> `decide`), `/speckit-specify`, `/speckit-clarify`, and `/speckit-checklist` to triage ideas and refine specs.
- **In target implementation repos**: use `spec-to-code-tdd` alongside `/speckit-plan` -> `/speckit-tasks` -> `/speckit-analyze` -> `/speckit-implement` <-> `/speckit-converge` (append-only gap convergence loop).

## code

* Maintain a version somewhere (either in file or inside constants/main code file) and add changes to `CHANGELOG.md` consistently with version changes.
* Do not exit this git repo "cage" for coding. Seek approval if needed.
