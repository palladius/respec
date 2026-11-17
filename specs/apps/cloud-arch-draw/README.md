# CloudArchDraw (`cloud-arch-draw`)

> **Implementation Repository**: `~/git/cloud-arch-draw` ([`github.com/palladius/cloud-arch-draw`](https://github.com/palladius/cloud-arch-draw)) — see [`blueprint.yaml`](blueprint.yaml).  
> Per `respec` rules, this directory contains **specifications and workshop scenario presets only** (zero application code).

## Contents

- **[`SPEC.md`](SPEC.md)** — Canonical functional & technical specification (`ArchitectureSpec` AST, 3-pane interactive XYFlow + Graphviz WASM studio, vertex-attached Bottlenecks/SPOFs, `NanoBanana -> Recreate` visual synthesis, Multi-Cloud `GCP` / `AWS` / `Azure` parity, and 500+ free multi-cloud icon provisioning).
- **[`input_prompt.md`](input_prompt.md)** — Original idea prompt.
- **[`blueprint.yaml`](blueprint.yaml)** — Target repository metadata.
- **[`scenarios/`](scenarios/)** — Five self-contained enterprise & workshop target-architecture specs ready to load as presets or live exercises:
  1. [`01_insurance_claims_disparity.md`](scenarios/01_insurance_claims_disparity.md) — Event-driven multimodal car/hail damage photo vs PDF garage invoice disparity workflow.
  2. [`02_cantonal_bank_mortgage_rag.md`](scenarios/02_cantonal_bank_mortgage_rag.md) — Sovereign bank mortgage dossier RAG with OAuth 2.0 On-Behalf-Of (`RFC 8693`) pre-retrieval ACLs and read-only `MCP` tools.
  3. [`03_canton_public_admin_citizen_portal.md`](scenarios/03_canton_public_admin_citizen_portal.md) — Multilingual (`DE/FR/IT/EN`) civic portal with deterministic pre-prompt PII redaction and human clerk queue.
  4. [`04_rag_on_vms_vs_serverless_managed.md`](scenarios/04_rag_on_vms_vs_serverless_managed.md) — Architectural & FinOps trade-off between self-hosted open-weight LLMs (`Apertus` / `Gemma` on GPU K8s + `vLLM`) and Serverless Managed RAG.
  5. [`05_redundant_lamp_to_cloud_native.md`](scenarios/05_redundant_lamp_to_cloud_native.md) — Classic 2-VM sticky LAMP stack modernized to Multi-AZ stateless containers, managed HA DB, Redis, Object Storage, and CDN/WAF.
