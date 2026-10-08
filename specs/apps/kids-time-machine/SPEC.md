# 👶 Kids Time Machine (ChronoKids) — Specification

> **Project Code:** `kids-time-machine` (Bambin Time Machine)  
> **Target Audience:** Riccardo, Kate, Alessandro (8yo), Sebastian (6yo), family members.  
> **Source Idea:** Voice dispatch by Riccardo Carlesso (2026-10-05).

---

## 1. Executive Summary & Vision

**Kids Time Machine** is an ultra-fast, responsive web application and computational photography pipeline designed to organize thousands of uncurated family photos of Alessandro and Sebastian from birth to the present day.

The platform solves two distinct challenges:
1. **Interactive Life Timeline & Atlas (Web App):** A snappy, zero-lag infinite mosaic photo browser that outperforms Google Photos by using precomputed multi-tier WebP thumbnails, paired with an interactive world map and an age-based "Time Machine" slider. Full-resolution originals are safely archived and lazy-loaded from Google Cloud Storage (GCS).
2. **Face Morphing / Aging Video Engine (Media Pipeline):** A computer-vision pipeline that extracts, landmark-aligns, and morphs ~100 chronological face crops of each child to create a fluid, continuous time-lapse video showing a newborn infant seamlessly transforming into an 8-year-old boy.

---

## 2. Problem Statement & Motivation

### The Problem with Existing Solutions (Google Photos, Apple Photos)
- **Laggy Scrolling & Throttling:** Scrolling back 8 years in Google Photos or iCloud on the web causes constant network stalls, blank gray placeholders, and high latency while multi-megabyte originals or dynamic thumbnails are generated on the fly.
- **Lost Geo-Context:** Trips are scattered across calendar dates. There is no simple way to open a world map, see pins on Mallorca, Sardinia, Dublin, Zurich, or South Africa, and jump directly to that trip's highlight album.
- **Lack of Age-Normalized Filtering:** While photos have capture dates, parents think in terms of child developmental milestones (*"Show me Ale at 6 months vs Sebi at 6 months"* or *"How did they look at 2 years old?"*).
- **No Native Face-Aging Time-Lapse:** Mainstream photo tools do not provide face-morphing video tools with stabilized eye/nose alignment.

---

## 3. Goals & Non-Goals

### Goals
- ⚡ **Sub-100ms Mosaic Scrolling:** Pre-rendered local/CDN thumbnail grid (300px WebP, <15KB per tile) ensuring instant 60fps scrolling across 10,000+ photos without network hitching.
- ☁️ **Cost-Effective GCS Storage:** High-resolution originals (10MB–50MB RAW/JPEG/HEIC) reside in low-cost Google Cloud Storage buckets, loaded on-demand only when a user clicks a photo into full-screen inspection mode.
- 🗺️ **Interactive World Map:** Leaflet / MapLibre vector map showing geo-clustered trips with hero preview cards (e.g. *"Mallorca 2024: Ale cycling by the Faraglione"*).
- ⏳ **Age-Milestone Time Machine:** Dual age sliders for Ale (born ~2018) and Sebi (~2020) enabling cross-comparison at identical ages (e.g., both at 18 months).
- 🎬 **Automated Face Morphing Pipeline:** Headless Python script taking ~100 face images, detecting 68/468 facial landmarks via MediaPipe/OpenCV, performing affine alignment (eyes leveled, interpupillary distance normalized), and generating morphing video transitions via optical flow or Delaunay warping.

### Non-Goals
- ❌ Replacing Google Cloud Storage with costly commercial SaaS photo hosting.
- ❌ Building an entire native mobile app (v1 is a responsive web application optimized for desktop and iPad).
- ❌ Manual tagging of every single photo (automation via EXIF metadata, folder names, and clustering).

---

## 4. User Journey & Core Features

### Feature 1: Zero-Lag Mosaic Timeline (Infinite Scroll)
- **Grid Layout:** Masonry or justified responsive grid rendered with virtualized DOM scrolling (`tanstack-virtual` or CSS content-visibility).
- **Multi-Tier Thumbnail Hierarchy:**
  1. *Micro-Blur LQIP (Low Quality Image Placeholder):* 20x20 base64 embedded directly in metadata JSON for instant zero-blank-screen feedback.
  2. *Grid Thumbnail:* 300px WebP compressed (~10-15KB) served from fast static cache.
  3. *Full-Res Master:* Fetched on-demand from GCS signed URL only when clicked into lightbox.
- **Fast Date Scrubber:** Vertical timeline bar on the right side allowing quick jumps to any year, season, or month.

### Feature 2: Interactive Travel Atlas (World Map)
- **Cluster Markers:** Map pins grouped by geographical bounding boxes (e.g. Mallorca, Sardinia, Zurich, Dublin, Cape Town).
- **Hero Highlight Card:** Each cluster features the curated "Hero Photo" of the trip (e.g. *Ale pedaling on the cliff overlooking the sea*), with date range and photo count.
- **One-Click Album View:** Clicking the hero card opens a filtered sub-view containing all photos from that specific trip, chronologically ordered.

### Feature 3: The "Time Machine" Age Slider
- **Dual Perspective:** Switch between:
  - *Chronological Date:* Real-world calendar time (e.g., Summer 2021).
  - *Age Milestones:* Filter by developmental age (e.g., "6 months", "1 year", "2 years", "5 years", "Current").
- **Side-by-Side Comparison:** Option to compare Alessandro at age $X$ alongside Sebastian at age $X$.

### Feature 4: Automated Face Morphing / Aging Video Engine
- **Input:** Curated sequence of ~100 chronological portraits of Ale or Sebi.
- **Landmark Detection:** MediaPipe Face Mesh detects pupils, nose tip, and corners of mouth.
- **Affine Normalization:** Rotates and scales every crop so eye centers are aligned horizontally at fixed coordinates $(x_1, y_1)$ and $(x_2, y_2)$ across all 100 frames.
- **Morphing Engine:**
  - *Phase A (Warp & Blend):* Delaunay triangulation morphing between adjacent keyframes $N$ and $N+1$ over $T$ transition frames.
  - *Phase B (AI Frame Interpolation - optional):* FILM (Frame Interpolation for Large Motion) or RIFE for ultra-smooth 60fps transitions.
- **Output:** 4K / 1080p MP4 time-lapse showing 8 years of life morphing in 60-90 seconds, accompanied by a nostalgic soundtrack.

---

## 5. System Architecture & Tech Stack

```text
┌─────────────────────────────────────────────────────────────┐
│                      INGESTION ENGINE                       │
│  Google Photos / Local Dump / SD Cards                      │
│                           │                                 │
│  [ Python / Node.js Ingestion Worker ]                      │
│   ├── 1. Read EXIF (Date, GPS, Camera, Orientation)        │
│   ├── 2. Compute Age from Birthdates (Ale / Sebi)           │
│   ├── 3. Generate Multi-Tier WebP Thumbnails (Sharp / VIPS) │
│   └── 4. Upload Originals to GCS (gs://ricc-family-photos)  │
└───────────────────────────┬─────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                      METADATA CATALOG                       │
│  SQLite / DuckDB / static catalog.json                      │
│  - photo_id, gcs_uri, date, age_ale_months, age_sebi_months │
│  - lat, lon, location_name, album_tag, is_hero              │
└───────────────────────────┬─────────────────────────────────┘
                            │
           ┌────────────────┴────────────────┐
           ▼                                 ▼
┌───────────────────────┐        ┌───────────────────────┐
│     WEB FRONTEND      │        │  FACE MORPH PIPELINE  │
│  - Node.js / Vite     │        │  - Python 3 / OpenCV  │
│  - Vue / React / HTML5│        │  - MediaPipe Landmark │
│  - Leaflet / MapLibre │        │  - Delaunay / RIFE    │
│  - Virtualized Mosaic │        │  - FFmpeg 60fps MP4   │
└───────────────────────┘        └───────────────────────┘
```

### Technology Choices
- **Frontend App:** Node.js + Vite + Vanilla/Vue/React with Tailwind CSS and Lucide icons.
- **Mapping:** Leaflet.js with OpenStreetMap or MapLibre GL for smooth globe panning.
- **Image Processing Engine:** Node.js `sharp` (libvips) for ultra-fast thumbnail generation; `exiftool-vendored` for comprehensive EXIF extraction.
- **Storage:** Google Cloud Storage (`gs://`) with public-read or signed URLs for full masters; local/CDN caching for thumbnails.
- **Face Morphing Pipeline:** Python 3 + `mediapipe` + `opencv-python` + `scipy.spatial` (Delaunay triangulation) + `ffmpeg`.

---

## 6. Data Model & Schema (Catalog)

```json
{
  "id": "photo_20240715_093211_01",
  "filename": "IMG_4921.HEIC",
  "capture_date": "2024-07-15T09:32:11Z",
  "age_milestones": {
    "alessandro_months": 75,
    "sebastian_months": 51
  },
  "geo": {
    "lat": 39.8142,
    "lon": 3.1215,
    "place": "Mallorca, Faraglione",
    "country": "ES"
  },
  "album": {
    "trip_id": "mallorca-2024",
    "trip_name": "Vacanze a Mallorca 2024",
    "is_hero": true,
    "caption": "Ale pedala verso il faraglione con la bici da corsa"
  },
  "assets": {
    "lqip_base64": "data:image/webp;base64,UklGR...",
    "thumb_url": "/thumbs/2024/07/IMG_4921_300.webp",
    "gcs_master_uri": "gs://ricc-family-photos-archive/2024/07/IMG_4921.HEIC"
  },
  "faces": [
    { "person": "alessandro", "bbox": [120, 340, 260, 480], "landmarks_detected": true }
  ]
}
```

---

## 7. Implementation Roadmap & Milestones

### Milestone 1: Ingestion & Metadata Pipeline (CLI)
- Script to scan local folders or GCS buckets, parse EXIF dates and GPS coordinates.
- Calculate age in months for Alessandro and Sebastian based on fixed birthdates.
- Output a single consolidated `catalog.json` and generate local 300px WebP thumbnails.

### Milestone 2: Responsive Web Mosaic & Time Machine
- Node.js / Vite web application loading `catalog.json`.
- Virtualized infinite scroll grid displaying thumbnails instantly without layout shift.
- Time Machine slider filtering photos by child age or year.

### Milestone 3: Interactive World Map & Trip Albums
- Embed Leaflet map grouping photos by GPS coordinates into clusters.
- Design Trip Album view with hero cover photos (e.g. Mallorca Faraglione) and sub-galleries.

### Milestone 4: Face Morphing / Aging Time-Lapse Script
- Script `scripts/generate_aging_timelapse.py`:
  - Input: folder of 50-100 cropped portrait photos sorted chronologically.
  - Step 1: Detect eye/nose landmarks with MediaPipe.
  - Step 2: Warp and align all faces to normalized template.
  - Step 3: Compute cross-dissolve with Delaunay morphing (24 frames per transition).
  - Step 4: Render finalized MP4 video with background music.

---

## 8. Open Questions & Future Considerations
1. **Google Photos API vs. Google Takeout / GCS:** Google Photos API historically downscales images and enforces quota limits. Batch exporting via Google Takeout or direct SD card sync to GCS is significantly more robust.
2. **Face Recognition Clustering:** Can we use local embeddings (e.g. FaceNet or insightface) to automatically classify whether a photo contains Ale, Sebi, or both?
3. **Privacy & Access:** Since this contains family photos of minors, the frontend must support simple token-based or password protection (or local LAN deployment on `pupurabbux` / `mini-lobby`).
