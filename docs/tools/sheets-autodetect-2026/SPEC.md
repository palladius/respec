---
speck_version: "0.1"
mode: manual
created_at: "2026-10-06T16:15:00Z"
model: gemini-3.8-flash
inspiration_repo: "https://github.com/palladius/ruby-sheets-autodetect"
status: proposed
---

# `sheets-autodetect` (2026 Edition)
### Next-Gen Spreadsheet-to-Models Engine: Deterministic Profiling, Gemini Semantic Inference, Relational DAG Resolution, and Zero-Friction Code Generation (Rails 8+, SQL, Prisma, Seeds)

---

## 1. Executive Summary: Overcoming the 2015 Era

In 2015–2020, `ruby-sheets-autodetect` proved a brilliant premise: **treat Google Spreadsheets as a collaborative domain modeling whiteboard, and turn them into functional code with a single command**.

However, the 2015 implementation was full of early-era tech debt, brittle shortcuts, and fragile assumptions.

### The 10 Sins of 2015 vs. The 2026 Architectural Solutions

| # | The 2015 Limitation / Sin | The 2026 Modern Breakthrough |
|---|---|---|
| **1** | **Cascading `rescue` Type Inference**: `Integer(s) rescue Float(s) rescue Date.parse(s)` (dangerously slow, parses `"4"` as Year 4 AD!). | **Grammar-Based Lexer & Finite State Profiler**: Zero exception-driven flow. Deterministic pattern engine with regex trie and sub-millisecond per-column tokenization. |
| **2** | **Tiny 9-Row Sample Size**: `arr.first(9)` meant empty top rows or early anomalies destroyed column classification. | **Full-Column Streaming & Reservoir Profiling**: Profiles 100% of rows (or 10k streaming sample), computing confidence distributions, null-ratios, and outlier detection. |
| **3** | **Row 1 Header Assumption**: Assumed headers are strictly on Row 1, Col 1. Real sheets have title banners, blank rows, and notes. | **Table Boundary & Header Discovery Engine**: Evaluates row entropy, type variance, and Google Sheets frozen rows metadata to automatically detect exact table bounds. |
| **4** | **Currency Bug & 4-Code Hardcoding**: Regex `[CHF\|USD\|EUR\|GBP]` only; `Money.new(100)` became $1.00 instead of $100.00. | **Universal ISO-4217 & Symbol Registry**: Parses `$`, `€`, `£`, `¥`, `CHF`, accounting formats `($50)`, decimals/thousands, generating dual-column Rails money (`price_cents` + `currency`) or `DECIMAL(12,4)`. |
| **5** | **Locale Blindness (EU vs. US)**: Ambiguity between `10/05/2026` (May 10 vs Oct 5) and `1.250,50` vs `1,250.50` caused silent corruption. | **Holistic Column Locale Resolver**: Analyzes entire column: if any value has day `> 12` (e.g. `29/12/2026`), the entire column is resolved to `DD/MM/YYYY`. Same for decimal/comma separators. |
| **6** | **Trivial Foreign Key Naming (`*_id`)**: Only recognized foreign keys if the column was literally called `foo_id`. | **Relational Value-Set Inference**: Performs Jaccard similarity and subset inclusion checks between columns across sheets (e.g., `Bookings.car_model` $\subseteq$ `Cars.model` $\rightarrow$ foreign key relation!). |
| **7** | **Zero AI / Semantic Blindness**: Columns were either `String` or `Text`. Could not understand emails, phone numbers, tax codes, or enums. | **Gemini 2.5 Semantic Layer**: Understands domain context: Italian `codice_fiscale`, Brazilian `CPF`, IBAN, avatar URLs, status enums, and auto-generates ActiveRecord validation rules. |
| **8** | **Single String Output (`rails generate`)**: Just printed a CLI string. No migrations, no models, no associations, no seeds. | **Full Stack Scaffolder**: Emits Rails 8 migrations with real DB constraints (`null: false`, `unique: true`, `foreign_key: true`), models with `belongs_to`/`has_many`, and populated `db/seeds.rb`. |
| **9** | **Intrusive Sheet Pollution**: Hardcoded a `_schema_v1.3_` tab with dark yellow styling and laptop hostname (`Socket.gethostname`). | **Non-Destructive by Default + Modern Dashboard**: Zero writes to source sheet by default. Opt-in `--write-schema-tab` writes a clean, formatted Data Dictionary with ER diagrams. |
| **10** | **Auth Agony (`aj-config.json`)**: Required manual OAuth client ID/secrets and deprecated Drive feed scopes. | **Zero-Config Google ADC + Local Files**: Seamless Application Default Credentials (`gcloud auth application-default login`), service accounts, and native local `.xlsx`/`.csv` support. |

---

## 2. Problem Statement

Developers and product teams prototype data models in spreadsheets because they provide the fastest visual editing experience. But converting 10+ tabs into production-grade database tables, relational associations, validations, migrations, and test seeds takes hours of tedious manual coding. 

Existing open-source tools from the 2010s rely on brittle regular expressions and dumb heuristics that break on real-world dirty data. Developers need a 2026-native tool that combines **deterministic statistical profiling** with **multimodal AI semantic intelligence** to convert any spreadsheet into production-ready software models in seconds.

---

## 3. Goals

- **Zero-Friction Authentication**:
  - Support Google Application Default Credentials (ADC) out of the box (`gcloud auth application-default login` or GCP Workload Identity).
  - Support local files: `.xlsx`, `.csv`, `.tsv`, `.parquet` directly without requiring any Google Cloud account.
  - Support public Google Sheets without credentials.
- **Robust Table Boundary Detection**:
  - Auto-discover table boundaries within worksheets (skip banner headers, notes, empty rows/columns).
  - Support multi-table worksheets (detect multiple independent tables on the same sheet tab).
- **Two-Tier Inference Engine**:
  - **Tier 1 (High-Performance Deterministic Lexer)**:
    - 0-exception control flow; Lexer tokenizes values into candidates.
    - Deep locale-awareness (auto-resolves decimal vs. comma, date format EU vs. US across the dataset).
    - Statistical distribution: detects nullable columns, enums (low cardinality strings), composite currencies, and text vs. string based on length entropy.
    - Dirty data tolerance: if 98% of values are integers and 2% are `"N/A"` or `"TBD"`, correctly infers `integer`, flags the 2 outliers, and coerces them to `nil` with warnings.
  - **Tier 2 (Gemini 2.5 Semantic Layer via `--ai`)**:
    - Discovers semantic domain types: PII, email, phone number, VAT ID, Italian Fiscal Code, address, avatar URL.
    - ActiveStorage & RichText detection: maps image URLs to `has_one_attached :photo` or ActiveStorage attachments, and rich formatted text to `has_rich_text`.
    - Relationship discovery: detects cross-tab relations even with disparate column names.
    - Generates idiomatic Rails validations and docstrings.
- **Relational Graph Solver (DAG)**:
  - Resolves primary keys, foreign keys, junction tables (Many-to-Many), and polymorphic relations.
  - Topologically sorts tables to ensure migrations and seeds run in strict dependency order.
- **Polyglot & Rails 8+ Exporters**:
  - **Rails 8+**:
    - `bin/rails generate scaffold` commands.
    - Migration files (`db/migrate/*.rb`) with foreign key constraints, indexes, null constraints.
    - Models (`app/models/*.rb`) with `belongs_to`, `has_many`, `enum`, `validates`.
    - `db/seeds.rb` populated with type-coerced, sanitized row data.
    - Optional `--apply` flag to build and migrate the Rails app directly.
  - **SQL DDL**: PostgreSQL, SQLite, DuckDB, BigQuery.
  - **Prisma & TypeScript**: `schema.prisma` and `types.ts`.
  - **Frictionless Data**: `tableschema.json`.
- **Interactive Terminal UI (TUI)**:
  - Rich interactive terminal preview (powered by Lipgloss/Bubbles or TTY) showing inferred tables, relations, and types.
  - Allows developer to override types or rename columns before emitting code.

---

## 4. Non-Goals

- Continuous bi-directional ETL synchronization (this is a bootstrapping, code-generation, and schema-migration tool, not Fivetran).
- Silent data loss or destructive mutation of existing Google Sheets.

---

## 5. Technical Architecture & System Design

```
+-------------------------------------------------------------------------+
|                              INPUT LAYER                                |
|  - Google Sheets API v4 (ADC / Token / Public)                          |
|  - Local File Connectors (CSV, TSV, XLSX via Calamine, Parquet)         |
+-------------------------------------------------------------------------+
                                     |
                                     v
+-------------------------------------------------------------------------+
|                  TABLE BOUNDARY & NORMALIZATION ENGINE                  |
|  - Header row detection (Entropy / Type variance)                       |
|  - Coordinate bounding box extraction (x1, y1) -> (x2, y2)              |
|  - Dirty row quarantine & empty cell normalization                      |
+-------------------------------------------------------------------------+
                                     |
                                     v
+-------------------------------------------------------------------------+
|                    TWO-TIER TYPE INFERENCE ENGINE                       |
|                                                                         |
|  [Tier 1: Deterministic Lexer]        [Tier 2: Gemini 2.5 Flash]        |
|  - Zero-exception token matching      - Semantic domain types (PII/Tax) |
|  - Locale resolution (EU/US date/dec) - Cross-tab relation inference    |
|  - Cardinality & Enum extraction      - ActiveStorage / RichText tags   |
|  - Outlier detection & tolerance      - Validation rule generation      |
+-------------------------------------------------------------------------+
                                     |
                                     v
+-------------------------------------------------------------------------+
|                       RELATIONAL GRAPH SOLVER                           |
|  - Primary Key / Foreign Key graph construction                         |
|  - Subset value inclusion analysis (Jaccard similarity)                 |
|  - Topological sorting (DAG migration order)                            |
+-------------------------------------------------------------------------+
                                     |
                                     v
+-------------------------------------------------------------------------+
|                    CANONICAL SCHEMA IR (IN-MEMORY)                      |
|  Tables, Columns, Relationships, Validations, Cleaned Seed Data         |
+-------------------------------------------------------------------------+
                                     |
      +------------------+-----------+--------+--------------------+
      |                  |                    |                    |
      v                  v                    v                    v
[Rails 8 Exporter] [SQL DDL Exporter]  [Prisma / TS]       [Data Hydrator]
- Scaffolds & Migr - PostgreSQL DDL    - schema.prisma     - db/seeds.rb
- Models & Enums   - SQLite / DuckDB   - TypeScript Types  - test/fixtures/*.yml
- Validations      - BigQuery DDL      - Pydantic v2 (Py)  - Outlier CSV report
```

---

### 5.1. The Lexer-Based Type Profiler (No More `rescue` Chains!)

Instead of brute-force Ruby `rescue` blocks, each value passes through a compiled lexer:

```ruby
module SheetsAutodetect
  class Tokenizer
    PATTERNS = {
      null:      /\A(null|nil|n\/a|none|-|\s*)\z/i,
      boolean:   /\A(true|false|yes|no|1|0|si|no|y|n|vero|falso)\z/i,
      integer:   /\A[+-]?\d{1,3}(,\d{3})*\z|\A[+-]?\d+\z/,
      decimal:   /\A[+-]?(?:\d+[\.,]\d+|\d{1,3}(?:[.,]\d{3})*[.,]\d+)\z/,
      currency:  /\A(?:\$|€|£|¥|CHF|EUR|USD|GBP)\s*[-+]?\d+([.,]\d+)?\z|\A[-+]?\d+([.,]\d+)?\s*(?:\$|€|£|¥|CHF|EUR|USD|GBP)\z/i,
      iso_date:  /\A\d{4}-\d{2}-\d{2}(?:[T ]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:Z|[+-]\d{2}:?\d{2})?)?\z/,
      slashed_date: /\A\d{1,2}[\/\-\.]\d{1,2}[\/\-\.]\d{2,4}\z/,
      email:     /\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z/i,
      url:       /\Ahttps?:\/\/[^\s\/$.?#].[^\s]*\z/i,
      uuid:      /\A[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\z/i
    }.freeze

    def self.tokenize(raw_value)
      str = raw_value.to_s.strip
      return :null if str.empty? || str.match?(PATTERNS[:null])

      PATTERNS.each do |type, regex|
        next if type == :null
        return type if str.match?(regex)
      end

      :string
    end
  end
end
```

### 5.2. Relational Value-Set Inference (Automated Foreign Key Discovery)

Even if a column in `Bookings` is called `Vehicle` instead of `car_id`, the engine tests the intersection of sets:

$$\text{Containment}(A, B) = \frac{|V(A) \cap V(B)|}{|V(A)|}$$

If 95%+ of distinct non-null values in `Bookings.Vehicle` exist in `Cars.Name`, the engine marks this as a high-confidence foreign key relationship:
- In Rails: `belongs_to :car, primary_key: :name, foreign_key: :vehicle`
- Migration: `add_foreign_key :bookings, :cars, column: :vehicle, primary_key: :name`

---

### 5.3. Gemini 2.5 Semantic Prompt Architecture

When invoked with `--ai`, the engine bundles table profiles into a single structured prompt for Gemini:

```
You are a Principal Database Architect and Rails 8 Core Contributor.
Analyze this spreadsheet schema profile and sample data:
- Tables: Cars, Users, Rentals
- Cars columns: [id, brand, model, license_plate, daily_rate_chf, photo_url, is_electric]
- Rentals columns: [id, renter_email, car_model, start_date, return_date, total_paid]

Tasks:
1. Identify exact domain semantic types (e.g., photo_url -> ActiveStorage, license_plate -> unique uppercase string).
2. Identify cross-tab foreign keys and cardinality (belongs_to / has_many).
3. Generate idiomatic ActiveRecord validations.
4. Extract enums with canonical keys.
Return strictly valid JSON adhering to SchemaIR specification.
```

---

## 6. CLI Usage & Developer Experience (DX)

```bash
# 1. Zero-config from Google Sheets URL
sheets-autodetect https://docs.google.com/spreadsheets/d/1pWqRWW3qhPfzjdmN0_D...

# 2. Complete Rails 8 App Generation with Gemini AI
sheets-autodetect 1pWqRWW3qhPfzjdmN0_D... \
  --ai \
  --target=rails \
  --out=./fleet_manager \
  --apply

# 3. Local Excel Workbook to SQLite & Prisma
sheets-autodetect ./data/inventory.xlsx \
  --target=sqlite,prisma \
  --out=./db_schema

# 4. Interactive review mode (TUI)
sheets-autodetect ./csv_folder/ --interactive
```

### Interactive TUI Preview
```
╔════════════════════════════════════════════════════════════════════════════════╗
║ sheets-autodetect 2026.1 🦖 — Detected 3 Tables (Topologically Sorted)         ║
╠════════════════════════════════════════════════════════════════════════════════╣
║ 1. users (142 rows)                                                            ║
║    ├─ id: Integer (PK, auto)                                                   ║
║    ├─ email: String (Unique, Email, validated)                                 ║
║    └─ role: Enum [:admin, :driver, :guest]                                     ║
║                                                                                ║
║ 2. cars (38 rows)                                                              ║
║    ├─ id: Integer (PK, auto)                                                   ║
║    ├─ owner_id: References -> users.id (99.2% match)                           ║
║    ├─ daily_rate: Currency (CHF, cents: Integer, currency: String)             ║
║    └─ photo: ActiveStorage attachment (image URL detected)                     ║
║                                                                                ║
║ 3. rentals (412 rows)                                                          ║
║    ├─ user_id: References -> users.id                                          ║
║    └─ car_id: References -> cars.id                                            ║
╚════════════════════════════════════════════════════════════════════════════════╝
[Enter] Generate Code  [e] Edit Column Types  [a] Toggle AI  [q] Quit
```

---

## 7. Generated Code Artifacts (Rails 8 Example)

### 7.1. Database Migration (`db/migrate/20261006120002_create_cars.rb`)
```ruby
class CreateCars < ActiveRecord::Migration[8.0]
  def change
    create_table :cars do |t|
      t.references :user, null: false, foreign_key: true
      t.string :brand, null: false
      t.string :model, null: false
      t.string :license_plate, null: false, index: { unique: true }
      t.integer :daily_rate_cents, null: false, default: 0
      t.string :daily_rate_currency, null: false, default: "CHF"
      t.integer :status, null: false, default: 0

      t.timestamps
    end
  end
end
```

### 7.2. ActiveRecord Model (`app/models/car.rb`)
```ruby
class Car < ApplicationRecord
  belongs_to :user
  has_many :rentals, dependent: :destroy
  has_one_attached :photo

  enum :status, { available: 0, rented: 1, in_service: 2 }

  validates :brand, :model, presence: true
  validates :license_plate, presence: true, uniqueness: { case_sensitive: false }
  validates :daily_rate_cents, numericality: { greater_than_or_equal_to: 0 }
end
```

### 7.3. Sanitized Database Seeds (`db/seeds.rb`)
```ruby
# Auto-generated by sheets-autodetect 2026
# Topologically sorted: Users -> Cars -> Rentals

puts "== Seeding Users =="
User.upsert_all([
  { id: 1, email: "ricc@google.com", role: 0 },
  { id: 2, email: "alessandro@example.com", role: 1 }
], unique_by: :id)

puts "== Seeding Cars =="
Car.upsert_all([
  { id: 1, user_id: 1, brand: "Ferrari", model: "Testarossa", license_plate: "MI-12345", daily_rate_cents: 150000, daily_rate_currency: "CHF", status: 0 },
  { id: 2, user_id: 2, brand: "Mercedes", model: "C180", license_plate: "ZH-99887", daily_rate_cents: 35000, daily_rate_currency: "CHF", status: 1 }
], unique_by: :id)
```

---

## 8. Verification & Acceptance Criteria

- [ ] **Zero-Crash Ingestion**: No uncaught exceptions on null, malformed dates, scientific notation, or dirty strings.
- [ ] **Accurate Locale Parsing**: European decimals (`1.250,50 €`) and US decimals (`$1,250.50`) parse to identical numerical quantities without data corruption.
- [ ] **Topological Seed Order**: Seeds execute cleanly in fresh databases with foreign key constraints enabled without ordering failures.
- [ ] **Outlier Resilience**: Columns with $\ge 95\%$ type consistency are not demoted to `String`; anomalies are quarantined in an audit log (`outliers_report.json`).
- [ ] **Zero-Credential Local Run**: Runs flawlessly on local `.xlsx` and `.csv` files completely offline.
- [ ] **Modern Google Auth**: Connects to private Google Sheets via standard ADC without requiring any `aj-config.json` file.
