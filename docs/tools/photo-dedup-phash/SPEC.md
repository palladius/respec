---
speck_version: "0.1"
mode: oneshot
idea_file: input_prompt.md
created_at: "2026-07-08T09:39:25Z"
model: gemini-flash-latest
tokens:
    prompt: 177
    output: 1444
    total: 2688
---

# Perceptual Hash Photo Deduplicator CLI

## Problem Statement

Users accumulating large collections of photos over time often end up with duplicate or near-duplicate images. These duplicates arise from resizing, minor editing, format conversion (e.g., converting a PNG to WebP), or downloading the same image at different resolutions. 

Standard cryptographic hash functions (such as MD5 or SHA-256) are inadequate for detecting these duplicates because even a single altered pixel or a change in file metadata results in a completely different hash. Manual inspection of thousands of photos is tedious and error-prone.

## Goals

- Provide a fast, cross-platform Command Line Interface (CLI) tool to identify identical and visually similar images.
- Support major image formats (JPEG, PNG, WebP, GIF, HEIC).
- Implement perceptual hashing algorithms (such as dHash, pHash, or aHash) to capture visual similarity.
- Allow users to configure a tolerance threshold (Hamming distance) for similarity.
- Implement a local cache (based on file paths, modification times, and file sizes) to avoid expensive re-hashing of unchanged files.
- Offer flexible resolution strategies: interactive terminal review, automated deletion of lower-resolution matches, or moving duplicates to a quarantine folder.

## Non-Goals

- Support for video deduplication (this tool will focus strictly on still images).
- Semantic duplicate detection using heavy machine learning models (e.g., CLIP, ResNet), which require GPUs or massive library dependencies.
- Editing or correcting image content (e.g., color correction, cropping).
- Cloud storage integration (the tool operates exclusively on the local filesystem).

## Technical Plan / Approach

### Language & Tooling
We will implement the CLI in **Rust** to maximize performance, leverage multi-core processors for parallel image decoding, and produce a single self-contained binary.

Key Rust Crates:
- `clap`: For command-line argument parsing.
- `image`: For decoding various image formats.
- `img_hash`: For calculating perceptual hashes (aHash, dHash, pHash).
- `rayon`: For parallel processing of files and hashing.
- `rusqlite` or `redb`: A lightweight embedded database to store hashes and file metadata for caching.
- `indicatif`: For rendering progress bars during scanning and processing.

### Hashing & Matching Algorithm
1. **Scanning**: Recursively traverse targeted directories, filtering for supported image mime-types.
2. **Caching Check**: For each file, check if the path, modification time, and size exist in the local cache database. If matched, retrieve the pre-calculated hash. If not, decode the image, calculate its hash, and save it to the cache.
3. **Clustering**: To find matches within a Hamming distance threshold $d$:
   - For smaller sets, a pairwise $O(N^2)$ comparison can be used.
   - For larger datasets, build a **BK-Tree** (Burkhard-Keller Tree) to allow metric space queries in $O(\log N)$ time, or use a binning heuristic based on sub-segments of the hash.
4. **Grouping**: Group images that fall within the threshold of one another into "duplicate clusters."

### CLI Commands & Interface
- `phash-dedup scan <dir>`: Scans a directory, computes hashes, updates the cache, and prints a summary of duplicates.
- `phash-dedup clean <dir> --strategy <interactive|keep-best|quarantine>`: Performs the scan and acts on the duplicates:
  - `interactive`: Prompts the user with file details (size, dimensions, modified date) to choose which to delete.
  - `keep-best`: Automatically retains the file with the highest resolution/file size and deletes/moves the others.
  - `quarantine --out-dir <path>`: Moves duplicates to a specified directory instead of deleting them.

## Alternatives Considered

- **Python with ImageHash & Pillow**: Easier to write quickly, but packaging Python apps into standalone binaries is fragile. Performance is significantly slower due to the GIL and slower image decoding libraries when processing tens of thousands of raw photos.
- **Machine Learning Embeddings (ONNX/CLIP)**: Excellent at semantic similarity (e.g., a photo of a dog from two different angles), but introduces hundreds of megabytes of dependencies, runs slowly on consumer CPUs without a GPU, and is overkill for identifying exact or resized duplicates.

## Implementation Plan

### Phase 1: CLI Scaffolding & Image Hashing
- Set up the Rust project structure with `clap` and configure basic CLI flags.
- Implement image loading and hash generation logic using the `image` and `img_hash` libraries.
- Benchmarking hashing performance on a sample set of 1,000 mixed images.

### Phase 2: Caching & DB Layer
- Integrate an embedded SQLite or `redb` instance.
- Create a schema containing `file_path`, `file_size`, `modified_time`, and `perceptual_hash`.
- Write logic to verify cache validity before reading/writing hashes.

### Phase 3: Matching & Clustering
- Implement the BK-Tree data structure to quickly query similar hashes.
- Implement clustering logic to group duplicates.
- Add the ability to output duplicate groupings to a JSON report for dry runs.

### Phase 4: Resolution & Interactive Terminal UI
- Build the interactive terminal prompt for file resolution.
- Implement safe file-system operations (deleting to system trash or moving to a quarantine directory rather than hard deletion by default).

## Open Questions

- **HEIC support on Linux/Windows**: The pure-Rust `image` crate has limited or non-existent support for HEIC/HEIF without native system bindings (libheif). Should we bundle or dynamically link native libraries to support modern iPhone photos, or fall back to requesting users pre-convert them?
- **Transitive Duplicates**: If Image A matches Image B ($d=1$), and Image B matches Image C ($d=1$), but Image A does not match Image C ($d=2$, where threshold is $d \le 1$), how should these cliques be resolved in groups? We will likely implement a disjoint-set (Union-Find) algorithm to cluster transitively connected duplicates.
