# Meta-spec: the interview philosophy behind `speck chat`

This project's interactive mode (`speck chat`) is deliberately designed around
ideas from Dave Resnin's essay **"Elephants, Goldfish, and the New Golden Age
of Software Engineering"**:
https://drensin.medium.com/elephants-goldfish-and-the-new-golden-age-of-software-engineering-c33641a48874

Credit for the underlying philosophy goes to that essay. This document
records which parts of it `speck` v1 adopts, and which parts are deliberately
deferred to a later version.

## Adopted in v1

- **Interrogate, don't just transcribe.** `speck chat` keeps asking clarifying
  questions instead of accepting the first answer as "good enough."
- **Resist false agreement.** The system prompt explicitly tells the model
  not to be sycophantic — it should challenge vague, underspecified, or
  self-contradictory answers rather than complimenting them.
- **Two-pass structure.** The interview runs in two distinct phases:
  1. **Problem** — what are we building, for whom, and why.
  2. **Acceptance criteria** — how will we know it's done / working.
  These are kept separate rather than blended into one free-for-all chat.
- **The human can end the interview explicitly.** Typing `/done` at any
  prompt forces early finalization of the current phase — mirroring
  "keep asking clarifying questions until I tell you to stop," in either
  direction (the model can also decide it's ready on its own).
- **Preserve reasoning artifacts.** The full Q&A transcript is written
  alongside the generated `SPEC.md` (as `speck_transcript.md`), not
  discarded — durable memory of *why* the spec looks the way it does.
- **Design doc structure.** Generated specs mirror Resnin's design-doc
  shape: Problem Statement, Goals, Non-Goals, Technical Plan / Approach,
  Alternatives Considered, Implementation Plan, Open Questions.

## Deliberately deferred (not in v1)

The essay's full "Elephant-Goldfish Model" describes a multi-session
review pipeline beyond a single CLI invocation:

- **Goldfish validation** — testing a design doc against a fresh,
  memoryless AI session to confirm it's self-contained.
- **Critic Review** — a skeptical, adversarial reviewer pass on the draft
  spec.
- **Implementation Readiness check** — a separate pass confirming there's
  enough detail to start implementation.

These are natural candidates for a future `speck review` or `speck critic`
subcommand once the core idea → spec flow is solid, but they're out of
scope for the first version of this tool.
