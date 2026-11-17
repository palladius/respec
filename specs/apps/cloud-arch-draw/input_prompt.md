Turn an application idea or short `SPEC.md` into a target Multi-Cloud Architecture diagram across **Google Cloud (`GCP`)**, **Amazon Web Services (`AWS`)**, and **Microsoft Azure** following each cloud's official Well-Architected best practices.

Core capabilities:
1. Canonical `ArchitectureSpec` JSON AST at the center so a single model drives:
   - An interactive visual Lego canvas (cards, nested VPC/cluster/region boundaries, drag-and-drop palette of 500+ official multi-cloud icons);
   - A deterministic Graphviz `DOT` diagram (`dot -Tsvg`) with **bottlenecks, Single Points of Failure (SPOFs), security boundaries, and FinOps trade-offs** tethered directly to affected vertices;
   - A live-synchronized Markdown architecture specification;
   - A 1-click `NanoBanana -> Recreate` visual renderer turning the user's edited canvas snapshot + topology into a presentation-grade diagram.
2. Multi-Cloud parity & comparison: support both single-cloud deep views (`GCP` first) and side-by-side `Azure | AWS` parity translation, plus `YOLO blind prompt` vs `Harness-Engineered` (`SPEC.md -> /plan -> CI/CD Eval Gate`) comparison for architecture workshops.
3. Bundled regulated & foundational reference scenarios (`scenarios/01..05_*.md`): Multimodal Insurance Claims Disparity, Sovereign Bank Mortgage RAG with OBO ACLs, Multilingual Public-Sector Citizen Portal, Self-Hosted Open-Weight GPU K8s vs Serverless Managed RAG, and Redundant LAMP-to-Cloud-Native Multi-AZ Modernization.
