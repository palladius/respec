---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-08T11:54:13Z"
model: gemini-flash-latest
tokens:
    prompt: 830
    output: 1994
    total: 4480
---

# Executable Spec Framework (ESF): Tiered Automated Acceptance Tests for AI Code Generation

## Problem Statement

AI coding assistants are highly capable of generating code that compiles, passes syntax checks, and matches static definitions. However, they frequently fail during execution in real-world environments. Common failure modes include:

*   **Cloud Infrastructure:** Terraform configurations compile and pass local dry-runs, but fail during `terraform apply` due to live cloud state drift, hidden organizational policies, service limits, or permission bottlenecks.
*   **External APIs:** Service integrations look syntactically correct but fail to authenticate, or fail when interacting with real endpoints due to undocumented schema quirks or API changes.
*   **Web Applications:** Frontend deployments build without errors but silently fail at runtime due to state initialization crashes, broken login forms, or zero-data screens that traditional unit tests (often mocked) miss.

Without an automated, deterministic way to validate the actual *behavior* of code in simulated or sandbox environments, AI-generated code relies entirely on manual human verification, negating the speed advantages of automated software generation.

## Goals

*   Define a standardized directory and schema specification (`.spec/` folder format) to bundle functional requirements alongside executable test tiers.
*   Establish a three-tiered testing hierarchy that balances execution cost, speed, and real-world assurance:
    *   **Tier 1 (Fixtures/Sandboxes):** Deterministic playground definitions to isolate execution (e.g., ephemeral databases, sandbox cloud environments).
    *   **Tier 2 (Continuous Verification):** Fast, cheap, deterministic integration and unit tests run continuously during agent development cycles.
    *   **Tier 3 (LLM-as-Judge E2E):** Slow, high-fidelity end-to-end evaluations that orchestrate real interactions (e.g., browser-driven Playwright paths, CLI calls against sandboxes) and use LLMs to visually or semantically assert success.
*   Provide an orchestration protocol that instructs AI developers how to sequentially run, evaluate, and self-correct using this framework.

## Non-Goals

*   Building an agentic orchestrator that runs parallel code-generation attempts to pit different LLMs against each other.
*   Creating a proprietary cloud provisioning system or managing cloud infrastructure directly (the framework declares the requirements; infrastructure execution is delegated to tools like Docker, LocalStack, or external sandbox managers).
*   Replacing core continuous integration (CI) engines. The specification is designed to integrate into existing CI/CD runners and local agent run loops alike.

## Technical Plan / Approach

### 1. Specification Directory Structure
An Executable Spec Bundle lives alongside the codebase, structured inside a `.spec/` directory:

```
my-project/
├── SPEC.md                      # Human/AI-readable feature requirements
└── .spec/
    ├── config.json              # Schema definition and metadata for validation runs
    ├── tier1-fixtures/          # Environment configuration & seed state
    │   ├── docker-compose.yml
    │   └── seed.sql
    ├── tier2-tests/             # Fast deterministic test suite
    │   └── auth-integration.test.js
    └── tier3-evals/             # Visual/Semantic automated E2E tests
        ├── login-and-view-dashboard.spec.js  # Playwright script
        └── evaluation-prompt.md # Instructions for LLM judge
```

### 2. Schema Definition (`.spec/config.json`)

```json
{
  "spec_version": "1.0.0",
  "name": "secure-user-login",
  "tier1": {
    "type": "docker-compose",
    "config_path": "./tier1-fixtures/docker-compose.yml",
    "healthcheck_url": "http://localhost:5432",
    "timeout_seconds": 60
  },
  "tier2": {
    "runner": "npm run test:integration",
    "test_directory": "./tier2-tests/",
    "max_duration_seconds": 15
  },
  "tier3": {
    "engine": "playwright",
    "scripts_directory": "./tier3-evals/",
    "capture_artifacts": ["video", "screenshots", "console-logs"],
    "judge": {
      "model": "gpt-4o",
      "prompt_path": "./tier3-evals/evaluation-prompt.md"
    }
  }
}
```

### 3. The Three-Tier Execution Loop

#### Tier 1: Deterministic Fixtures
Before any tests run, the orchestration runner sets up the requested sandbox environment. This ensures the agent is not working against simulated mocks but against tangible dependencies:
*   **Local Sandboxes:** Spins up local dependencies via Docker Compose (e.g., local PostgreSQL seeded with mock data, mock AWS via LocalStack).
*   **Cloud Sandboxes:** Points to an authenticated sandbox or ephemeral namespace (e.g., a dedicated GCP project with predefined resource quotas provided through pre-configured credentials).

#### Tier 2: Continuous Verification
These are fast, automated code tests designed to run on every file write. 
*   Highly deterministic.
*   Fast execution feedback (under 15-30 seconds).
*   Standard testing framework exit codes (0 for pass, non-zero for fail) are fed back to the editing agent for rapid self-correction cycles.

#### Tier 3: LLM-As-Judge E2E Evals
Designed to run pre-merge or on milestone completion. It acts as a final validation gate to check if the application actually *works* when all pieces assemble.
*   **Action Script:** Playwright, Puppeteer, or custom CLI scripts interact directly with the running, built application. 
*   **Artifact Gathering:** Scripts record a video of the browser session, capture page screenshots at critical milestones, and record all runtime logs.
*   **LLM Evaluation:** The captured artifacts (video frames, screenshots, final logs) along with the `evaluation-prompt.md` are sent to a high-reasoning LLM to answer functional questions, e.g.:
  > "Verify that the user was successfully logged in, that the dashboard loaded real metrics instead of error spinners, and that no blank states or broken CSS layouts are visible."
*   **Output:** The LLM returns a structured JSON evaluation block containing a boolean pass status and a descriptive list of visual/functional failures.

## Alternatives Considered

*   **Prose Acceptance Criteria Only (Traditional SPEC.md):** Relies on human developers reading, manually executing, and verifying results. While standard, it breaks down completely when utilizing autonomous code-generation loops.
*   **Pure E2E Integration Suites (Without LLM Judge):** Relying solely on assertions like `expect(page.locator('.title')).toHaveText('Dashboard')` is fragile when layouts change. It misses silent UX breakages, blank canvases, or layout shifts that a visual LLM judge easily catches. Concurrently, coding assertions for every minor element is slower for a developer to write than providing high-level visual/behavioral prompts.

## Implementation Plan

*   **Phase 1: Spec Schema Definition & CLI runner (MVP)**
    *   Formally publish the JSON Schema for `.spec/config.json`.
    *   Develop a lightweight CLI runner (`esf-run`) written in Node or Python that parses the config, boots the Tier 1 docker container, runs Tier 2 integration suites, and captures Tier 3 Playwright outputs.
*   **Phase 2: LLM-Judge Module Integration**
    *   Integrate vision model API endpoints (e.g., OpenAI or Anthropic) into `esf-run` to process captured screenshots and videos against the evaluation prompt, returning exit-codes based on LLM determination.
*   **Phase 3: Agentic Prompting Templates**
    *   Develop system prompt templates that instruct coding agents (e.g., Cursor, Aider, custom langchain loops) on how to look for `.spec/` configurations, run the `esf-run` test suite internally, and iterate until all tiers pass.

## Open Questions

*   **Flakiness Mitigation:** How do we minimize flaky test results from the Tier 3 LLM-as-Judge step? Will we need consensus voting (calling the judge LLM multiple times) or highly constrained output schemas to guarantee deterministic test runs?
*   **Credential Handling:** What is the most secure method for passing ephemeral credentials (e.g., AWS tokens, external API keys) down to the Tier 1 sandbox provisioning layer inside agent execution contexts?
