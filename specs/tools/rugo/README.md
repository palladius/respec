# Rugo / Goby (`rugo`)

> **Target Implementation Repository**: `~/git/rugo` ([`github.com/palladius/rugo`](https://github.com/palladius/rugo), planned — see [`blueprint.yaml`](blueprint.yaml)).  
> Per `respec` rules, this directory contains **specifications and design metadata only** (zero transpiler implementation code).

## Contents

- **[`SPEC.md`](SPEC.md)** — Canonical functional & technical specification for **Rugo** (*aka Goby*): Ahead-of-Time Ruby-to-Go transpiler architecture (`Prism AST -> Typed HIR -> Go AST`), implicit `@ctx` -> `context.Context` threading, Ruby blocks -> Go closures / `iter.Seq`, symbols -> `const` / `unique.Handle[string]` interning, static elimination of `respond_to?` / `method_missing`, and RSpec -> Go golden-file test harness.
- **[`input_prompt.md`](input_prompt.md)** — Original Septober `#245` idea note and context.
- **[`blueprint.yaml`](blueprint.yaml)** — Target repository and prototype status metadata.
