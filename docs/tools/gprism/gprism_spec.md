# Gprism (Git Privatize in Secret Manager) Specification

## Objective
Evolve `git-privatize` to use Google Cloud Secret Manager for storing encrypted environment and configuration files instead of a centralized Git repository.

## Architecture

1. **Configuration & Namespacing**
   - The tool reads configuration from `~/.gprism.yaml`, `~/.config/gprism/config.yaml`, or `.env`.
   - Variables expected: `project_id`, `identity`, and `environment`.
   - Secret Manager secrets are namespaced automatically using `[environment]-[slugified-filename]` to ensure isolation between dev, staging, and prod environments.

2. **File Tracking (`.git-privatize.list`)**
   - Operates on a `.git-privatize.list` file which works similarly to `.gitignore`.
   - Paths matching entries in this list are marked for privatization.

3. **Encryption (AES-256-CBC)**
   - All files are locally encrypted symmetrically before upload.
   - The encryption key is sourced from the `GIT_PRIVATIZE_KEY` environment variable.
   - If missing, it queries macOS Keychain (`security find-generic-password -s "gprism" -a "gprism" -w`).
   - If missing from Keychain, it prompts the user securely and optionally saves it back to the Keychain.
   - Encrypted output is Base64 encoded and prepended with a `Salted__[salt]` header before pushing to GCP.

4. **Secret Naming & Length Limits**
   - GCP Secret Manager imposes a 255 character limit on secret names.
   - Filepaths are slugified (non-alphanumeric chars replaced with `-`).
   - If the combined namespace and filename slug exceeds 255 chars, the slug is truncated and a SHA-256 hash of the filepath is appended to guarantee uniqueness.

5. **Commands**
   - `gprism push --all`: Encrypts and pushes files. Generates `.readme` breadcrumb files pointing to the GCP secret. Deletes local plaintext file. Adds the file to `.gitignore`.
   - `gprism pull --all`: Fetches the encrypted payload from GCP, decrypts it, and restores the original file.
   - `gprism status`: Reviews `.git-privatize.list` and reports the presence of the file locally and in GCP.

## Implementation Details
- Language: Ruby
- Testing: Minitest suite checking end-to-end functionality including GCP mocks and Base64 padding.
- Color formatting outputs for User Experience: GREEN (Success), RED (Error), CYAN (Info/Breadcrumbs).
