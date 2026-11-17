# Scenario 01: Multimodal Insurance Claims & Invoice Disparity Workflow
id: 01_insurance_claims_disparity
sector: Property & Motor Insurance
region_policy: Strict Swiss Residency (europe-west6 / Switzerland North / eu-central-2)

## 1. Problem & Workload
After an Alpine hailstorm, 8,000 policyholders upload smartphone car/roof photos and garage PDF repair quotes within 4 hours.
Adjusters cannot manually cross-check whether a CHF 6,400 bumper + panoramic-roof invoice matches a minor dent photo.

## 2. Functional Architecture & Data Entities
- Mobile & Web Intake: Direct signed upload to sovereign Object Storage; async event enqueue (never block HTTP threads).
- Branch A (Physical Evidence): Multimodal Vision model inspects hail/collision photos, counting affected panels and severity.
- Branch B (Financial Claim): Document OCR parser extracts line items, labor hours, part SKUs, and total CHF from PDF invoices.
- Disparity & Fraud Agent: Compares physical damage vs billed line items and queries Core Policy DB.
- Decision Router:
  * Clean match AND total < CHF 2,000 -> Straight-through payout approval + immutable ledger entry.
  * Disparity > 20% OR total >= CHF 2,000 -> Async Human-in-the-Loop (HITL) Adjuster Queue with side-by-side `A2UI` visual diff.
- Core Entities: `claims(id, policy_id, status, chf_claimed)`, `damage_findings(panel, severity)`, `invoice_lines(sku, chf)`.

## 3. Non-Functional, Security & FinOps Guardrails
- Compliance: Swiss revDSG Art. 60, Zero Data Retention (ZDR), zero training on customer media, CMEK encryption at rest.
- Latency & Resilience: Intake ACK < 300 ms; async triage < 15 s P95; backpressure queue absorbs 50x hailstorm spikes.
- Token FinOps: Tiered routing (fast multimodal model for 85% clear claims; reasoning escalation + CHF/day budget kill-switch).

## 4. Clarifying Questions & CI/CD Evaluation Gate
- Clarifying Questions: Synchronous mobile wait or async push notification? Auto-payout ceiling CHF 2,000? Dual-region Swiss DR needed?
- Evaluation Gate: >= 97% disparity recall on padded-invoice golden set; 0 cross-claim PII leaks; 0 synchronous HTTP 504 timeouts.
