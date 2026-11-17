# Scenario 02: Sovereign Bank Mortgage Dossier RAG & Core Tooling
id: 02_cantonal_bank_mortgage_rag
sector: Regulated Banking & Wealth Credit Desk
region_policy: Sovereign Zurich Only (europe-west6 / Switzerland North / eu-central-2)

## 1. Problem & Workload
650 credit officers analyze 80-page German/French/Italian mortgage dossiers, tax returns, and land-registry extracts.
Analysts need grounded synthesis plus live affordability and LTV simulation from Core Banking APIs without leaking VIP files.

## 2. Functional Architecture & Data Entities
- Identity & Access: Corporate SSO issuing short-lived OAuth 2.0 On-Behalf-Of (`RFC 8693`) tokens carrying branch and desk ACL groups.
- Ingestion Pipeline: PDF layout parsing, semantic chunking, embedding generation, and indexing with mandatory `acl_groups` tags.
- Pre-Retrieval ACL Vector Search: Filter `WHERE acl_groups && user.obo_groups` BEFORE similarity ranking (zero Confused Deputy).
- Governed Tool Calling (`MCP`): Read-only Core Banking tools (`calc_affordability`, `fetch_saron_curve`, `check_ltv_policy`).
- Cited Advisor Workspace: Every rate, ratio, and covenant claim links to exact PDF page bounding boxes; human sign-off > CHF 500k.
- Core Entities: `dossiers(id, client_tier, acl_groups[])`, `dossier_chunks(id, embedding, page_box, acl_groups[])`, `audit_log`.

## 3. Non-Functional, Security & FinOps Guardrails
- Regulatory: Swiss Banking Secrecy (Art. 321 StGB / Art. 47 BankG), FINMA Guidance 08/2024 explainability, revDSG Art. 60.
- AI Firewall: Prompt-injection & PII boundary guardrail; Zero Training & Zero Data Retention (ZDR) contractual enforcement.
- Token FinOps: Context caching on underwriting handbooks; max 6 tool turns per session with hard token kill-switch.

## 4. Clarifying Questions & CI/CD Evaluation Gate
- Clarifying Questions: Read-only Core Banking `MCP` tools or draft write-back? Private-wealth VIP ringfencing? Audit retention 10 years?
- Evaluation Gate: 100.0% ACL isolation across red-team cross-branch probes; >= 95% citation faithfulness on 200 credit committee cases.
