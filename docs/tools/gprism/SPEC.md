---
speck_version: "0.1"
mode: oneshot
created_at: "2026-07-12T19:16:00Z"
category: tools
slug: gprism
title: Gprism
---

# Gprism: Git Privatize in Secret Manager (v4)

This plan details the implementation of **`gprism`** (Git Privatize in Secret Manager) with keychain security, namespacing, support for `.git-privatize.list` files, and a comprehensive automated test suite.

No commits will be pushed to remote GitHub until you return and explicitly approve. We will only create the repository and work on it locally.

## User Review Required

> [!IMPORTANT]
> **1. Password Caching (macOS Keychain Integration)**
> To avoid prompting for the password or writing it in plaintext config files:
> - The tool will attempt to read the password from the macOS Keychain under the service name `gprism` and account name `gprism`:
>   `security find-generic-password -s "gprism" -a "gprism" -w`
> - If not found, the tool will prompt the user and offer to automatically save it in the macOS Keychain for future usage via:
>   `security add-generic-password -s "gprism" -a "gprism" -w "PASSWORD" -U`
> 
> **2. Name Length Safety & Nested File Conflicts**
> - Names will be generated using the relative file path to prevent collision (e.g. `path/to/another/.env` maps to the slug `path-to-another-env`).
> - If a generated secret name exceeds 255 characters, the middle section is truncated and appended with a SHA-256 hash slice of the full path to ensure it remains unique and under Secret Manager's limits.
> 
> **3. `.git-privatize.list` & `--all` Command Support**
> - The repository can contain a `.git-privatize.list` file (similar to `.gitignore`) containing patterns of secrets to privatize.
> - Running `gprism push --all` automatically encrypts and pushes all files matching the list.
> - Running `gprism pull --all` restores all files in the list.
> - All files in `.git-privatize.list` are automatically appended to `.gitignore`.
> 
> **4. Visual Styling**
> - The CLI will output clear messages using standard Linux terminal colors (folders in **BLUE**, symlinks/breadcrumbs/readmes in **CYAN**, errors in **RED**, success in **GREEN**).

## Proposed Changes

We will create the repository inside `/Users/ricc/.gemini/antigravity/scratch/gprism`:

### Gprism Repository

#### [NEW] [README.md](file:///Users/ricc/.gemini/antigravity/scratch/gprism/README.md)
Update documentation to reflect `gprism`, `.git-privatize.list`, and Keychain storage.

#### [NEW] [gprism](file:///Users/ricc/.gemini/antigravity/scratch/gprism/bin/gprism)
The core script:
- Parses `.git-privatize.list` and supports `--all` / `-a` flag for pushing and pulling.
- Queries macOS Keychain via `security find-generic-password` to fetch/store the encryption passphrase.
- Implements length safety checks for secret names: if length > 255, hash and truncate to safety limit.
- Applies standard terminal coloring standards (Folders in **BLUE**, Symlinks/Readmes in **CYAN**).

#### [NEW] [test_gprism.rb](file:///Users/ricc/.gemini/antigravity/scratch/gprism/test/test_gprism.rb)
A comprehensive Minitest suite that:
- Initializes a dummy git repo.
- Creates `.git-privatize.list` with multiple secret files.
- Runs `gprism push --all` with a mock key and asserts:
  - Secrets are uploaded to Secret Manager.
  - Secret versions contain the `Salted__` encryption header.
  - `.readme` breadcrumb files are generated.
  - `.gitignore` includes the secrets.
- Verifies `gprism status` reports correct states.
- Deletes the secrets and runs `gprism pull --all` to assert all files are decrypted and restored to original plaintext.
- Verifies decrypt failures on wrong password.

## Verification Plan

### Automated Test Suite Execution
- Run `ruby test/test_gprism.rb` and verify all tests pass.
- Clean up any test secrets created in Secret Manager during testing.
- Run `git status` to ensure no test secret files are tracked.
