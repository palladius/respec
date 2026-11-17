# Scenario 03: Multilingual Public-Sector Citizen Portal & Dossier Triage
id: 03_canton_public_admin_citizen_portal
sector: Cantonal & Municipal Public Administration
region_policy: Swiss Public-Sector Sovereignty (europe-west6 / Switzerland North / eu-central-2)

## 1. Problem & Workload
Residents and international expats submit building permits, relocation notices, and tax deferral requests in DE, FR, IT, and EN.
Civil clerks spend 60% of their week checking missing attachments, translating inquiries, and routing dossiers across 9 offices.

## 2. Functional Architecture & Data Entities
- Accessible Civic Portal: WCAG 2.1 AA web intake authenticated via Swiss eID / AGOV, accepting structured forms and PDF scans.
- Pre-Prompt Privacy Shield: Deterministic DLP / PII redaction strips national ID (AHV/AVS) numbers and health notes before LLM calls.
- Multilingual Triage Agent: Summarizes citizen intent in German for clerks while drafting polite replies in the citizen's language.
- Statutory Grounding: Retrieval over versioned cantonal ordinances and municipal zoning rules (strictly verified public corpus).
- Clerk Decision Cockpit: AI never issues legally binding administrative rejections; every draft enters a human civil-servant queue.
- Core Entities: `submissions(id, agov_sub, lang, dept_id)`, `redacted_artifacts(id, pii_vault_ref)`, `ordinance_chunks(id, law_ref)`.

## 3. Non-Functional, Security & FinOps Guardrails
- Legal & Sovereignty: Cantonal data protection + Swiss revDSG, strict Swiss ZDR, customer-managed encryption keys (CMEK).
- Anti-Hallucination: Refuse uncited legal advice; fall back to deterministic checklist + office routing when confidence < 0.85.
- Cost Predictability: Tax-season burst autoscaling to zero overnight; per-department monthly CHF token quota and prompt caching.

## 4. Clarifying Questions & CI/CD Evaluation Gate
- Clarifying Questions: Reversible PII vault tokenization for clerk view? Strict cantonal-only ordinance grounding? eID guest fallback?
- Evaluation Gate: Zero unredacted AHV numbers in model telemetry; >= 94% routing accuracy across DE/FR/IT/EN; 0 autonomous denials.
