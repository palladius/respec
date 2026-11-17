# Scenario 04: Self-Hosted Sovereign Open-Weight GPU K8s vs Serverless Managed RAG
id: 04_rag_on_vms_vs_serverless_managed
sector: Cross-Industry Architecture Trade-Off (Banking, Pharma & Industry)
region_policy: Zurich Region Parity (europe-west6 / Switzerland North / eu-central-2)

## 1. Problem & Workload
An architecture board must choose between two patterns for an internal engineering & regulatory RAG assistant (400 daily active users):
Pattern A runs open-weight `Apertus` / `Gemma` on dedicated GPU nodes; Pattern B uses fully managed Serverless AI APIs.

## 2. Dual Target Architecture Comparison & Data Entities
- Pattern A (Self-Hosted GPU K8s / VMs): Managed Kubernetes with L4/H100 GPU node pools running `vLLM`, self-managed Vector DB,
  custom Prometheus GPU telemetry, manual weight patching, and 24/7 idle GPU capacity burn during nights and weekends.
- Pattern B (Serverless Managed RAG): Scale-to-zero containers calling regional Enterprise LLM endpoints under ZDR + CMEK,
  managed ACL-aware Vector Search, built-in AI Guardrails, zero CUDA/driver toil, and pure pay-per-token consumption billing.
- Hybrid Escape Hatch: Sovereign AI Gateway routing 90% standard Enterprise ZDR traffic serverlessly and Tier-0 crown-jewel prompts
  to a right-sized open-weight `Apertus` GPU pool.
- Core Entities: `routing_policies(sensitivity_tier, target_backend)`, `corpus_index(doc_id, acl_scope)`, `token_finops_ledger`.

## 3. Architectural & FinOps Trade-Off Matrix
- TCO Crossover: Dedicated 2x H100 HA cluster costs ~CHF 11k-16k/month fixed; Serverless RAG for 400 users costs ~CHF 650-1,400/month.
- Ops Burden: GPU K8s requires ML platform SREs (CUDA upgrades, KV-cache tuning, failover); Serverless focuses teams on evals & data.
- Sovereignty Reality: Enterprise Cloud ZDR + CMEK + Swiss region satisfies 95% of FINMA/revDSG workloads without idle GPU waste.

## 4. Clarifying Questions & CI/CD Evaluation Gate
- Clarifying Questions: Sustained > 60% GPU duty cycle or spiky daytime office hours? Dedicated ML SRE team on call? Tier-0 air-gap mandate?
- Evaluation Gate: Side-by-side P95 latency, monthly CHF TCO at 1x/10x load, answer quality parity rubric, and SPOF score.
