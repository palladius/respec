# Scenario 05: Redundant LAMP Stack to Cloud-Native Multi-AZ Modernization
id: 05_redundant_lamp_to_cloud_modernization
sector: Enterprise Foundation Workload (Industry, Retail & Mid-Market IT)
region_policy: Multi-AZ Swiss Region (europe-west6 / Switzerland North / eu-central-2)

## 1. Problem & Workload
A business-critical PHP/Python web portal and MySQL/PostgreSQL database currently run on two aging VMs with local `/var/www/uploads`
NFS storage, hardcoded `.ini` DB credentials, manual nightly cron mysqldumps, and sticky-session Apache load balancing.

## 2. Target Well-Architected Modernization & Data Entities
- Edge & Ingress: Regional/Global CDN + Web Application Firewall (WAF) + managed TLS in front of a Multi-AZ L7 Load Balancer.
- Stateless Compute Tier: Containerized web/API workers scaling horizontally across 3 Availability Zones with zero local disk state.
- Shared Media & Session Tier: User uploads offloaded to durable Object Storage via signed URLs; session/query cache in managed Redis.
- High-Availability Database Tier: Managed Relational DB with synchronous standby failover across AZs + read replica + PITR backups.
- Zero-Secret Hygiene: Workload Identity / Managed Identity fetching rotated credentials from Cloud Secrets Manager at startup.
- Core Entities: `users(id, email, mfa_state)`, `orders(id, user_id, total_chf, status)`, `media_objects(id, bucket_uri, sha256)`.

## 3. Non-Functional, Resilience & FinOps Guardrails
- Availability & DR: 99.95% SLA; automated Multi-AZ DB failover RTO < 60 s and point-in-time recovery RPO < 5 min.
- Security Posture: Private VPC subnets for DB and cache (zero public IPs), least-privilege IAM roles, container image scanning.
- FinOps Hygiene: Autoscaling min-2 / max-20 stateless instances + coldline lifecycle tiering on historical attachments.

## 4. Clarifying Questions & CI/CD Architecture Gate
- Clarifying Questions: Refactor PHP sessions to Redis or keep container-compatible app runtime? Read-heavy reporting traffic split? RPO < 5m?
- Evaluation Gate: Zero single points of failure (0 SPOFs) under chaos AZ-drop test; 100% parity across GCP, Azure, and AWS checklists.
