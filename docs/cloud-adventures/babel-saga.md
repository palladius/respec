# 🌍 BabelSaga (Defaultia)
## Complete Game Design, Lore & AI Agent System

**Version**: 1.0.0

> **A multi-namespace SimCity-style simulation game powered by AI agents and developer-written bots.**
> *Explore a grid-world where the English language is forbidden, Esperanto is the common tongue, and custom AI bots manage economy, resources, and communication via the Model Context Protocol (MCP) and Cloud Run.*

---

## 🌟 1. Executive Summary & Vision

**BabelSaga** (previously *UniBabel* or *Defaultia*) is a 2D SimCity-like simulation game built for hackathons and developers. Players program, deploy, and upgrade AI bots that navigate a simulated city, collect resources, trade property, and communicate with other bots and players.

### Core Philosophy:
* **The No-English Rule**: English is strictly forbidden in Gicipya's cities. Players and bots must use their native tongue (interpreted via translation bots) or the universal common connector language: **Esperanto**.
* **AI-First Integration**: Attendees deploy their bots to Google Cloud Run and connect them to the game engine using the **Model Context Protocol (MCP)** or ADK.
* **Multi-Instance Hackathons**: The game is multi-namespace, allowing custom instances (cities) to be spun up instantly via Terraform for events and competitions.

---

## 🤖 2. The Droid & AI Agent System

Each player starts with a suite of three bots (Droids) that live in their home and help them operate:

1. 🌐 **Protocol/Translator Droid (Active)**: Automatically translates chats, signs, and communications.
2. 🧲 **Gatherer/Worker Droid (Active)**: Automatically explores the grid, interacts with characters, asks questions, and harvests resources.
3. 🧠 **Custom AI Droid (Inactive by default)**: A blank-slate bot that the player programs using the Google Agent Development Kit (ADK) or custom code, and deploys to Cloud Run.

### Bot Communication Protocol (`babot` / Omega13)
To ensure AI security and isolation, bots communicate with each other and the server using a specialized JSON protocol named `babot` (or *Omega13*).
* **Payload Encryption**: Communication content is encoded in `rot13` or `base64` so humans cannot read bot-to-bot secrets directly.
* **Custom Status Strings**: Instead of standard HTTP/JSON statuses, bots use Esperanto terms like `VABON` (Very Good / OK) or `Aninpospio`.
* **Example Payload**:
  ```json
  {
    "Msg": "zrfvdnttvb_va_rot13",
    "status": "VABON"
  }
  ```
* Bots are instructed to **only** speak their designated language and ignore messages in other tongues unless a translator is present. There is a universal query command in Esperanto to ask any bot: *"Kian lingvon vi parolas?"* (What language do you speak?).

### Progression & Leveling
Droids earn XP logaritmic scale (D&D style):
$$\text{Level} = \log\left(\frac{\text{XP}}{2000}\right)$$
As they level up, they learn new skills (e.g., the Translator learns more languages; the Gatherer mines resources faster).

---

## 🎮 3. Game Mechanics & User Experience

### Onboarding & Character Creation
* **Profile**: Players provide their name, optional gender, and a public headline (Default: *"Sono troppo pigro per scrivere una headline"* / *"I'm too lazy to write a headline"*).
* **Bag of Holding (Borsa Conservante)**:
  * Strict capacity limit of **7 slots**.
  * **Starting items**: Standard explorer tools (e.g., compass, hat), one rules document written in Esperanto (*No English, no insults, respect others*), and **one giant comical item** (e.g., a whole pre-fab house or a 3-story ladder) to keep things whimsical.

### Economy & Crowdsourced Real Estate
* **Universal Basic Income (UBI)**: Players receive a daily allowance of **100 Gemoj** (singular: *Gemo*, symbol: 💎) in reference to the Gemini API.
* **Crowdsourced Grid**:
  * The map is randomized on creation. Shop plots are typically $10 \times 10$ meters (100 sqm).
  * Players can buy vacant plots, assign an emoji (e.g., 🍕 for a pizzeria, 🍻 for a tavern), name it, and convert it into a residence, restaurant, store, or warehouse.
* **Storage**: Houses contain a storage chest (the bank) where Gatherer bots automatically deposit collected loot and Gemoj.

---

## 💬 4. Chat & Communication

The chat interface is local (room-based) but supports varying scopes of broadcasting:
* **Text input (no prefix)**: Transmitted locally to the current room.
* **Esperanto Commands**:
  * `/krii [message]` (Shout) – Reaches the entire neighborhood.
  * `/flustri [user] [message]` (Whisper) – Private message to a specific user.
  * `/helpo` (Help) – Displays commands.
* *Onboarding Simplification*: Standard English command aliases (like `/shout` or `/help`) are accepted as transparent wrappers so new players don't get stuck.

---

## 🚀 5. v1.1 Add-on Features & Extensions

### 5.1 3D Isometric / Voxel Rendering
- In addition to the flat 2D emoji grid, the frontend supports a 3D isometric/voxel viewport using Three.js or simple CSS 3D projections.
- Emojis (e.g. 🍕) and buildings are rendered as blocky 3D voxel models on a navigable and zoomable map.

### 5.2 A2A Gem-Barter Protocol & Indirect Cooperation
- The final milestone requires 3 of the same gem type (3 Diamonds, 3 Rubies, or 3 Emeralds).
- Since each player begins with a mixed set (1 Diamond, 1 Ruby, 1 Emerald), the Familiar/Droid must negotiate and trade gems autonomously with other players' bots via Agent-to-Agent (A2A) endpoints.

---

## 🏗️ 6. Technical Architecture & Tech Stack

BabelSaga is architected to run perfectly in a local dev environment while scaling to a production cloud topology:

* **Local Dev**: Can run Frontend and Backend on separate localhost ports. This allows running multiple independent cities locally that can communicate with each other.
* **Deployments**: Spin up a new namespace/city via **Terraform** (e.g., for an hackathon). Each city can be set in a custom epoch and location (e.g., *Mediolanum 3125* or *Antica Alessandria*).
* **Infrastructure**:
  * **IaC**: Terraform.
  * **Serverless Backend**: Google Cloud Run.
  * **Primary DB (High-write)**: PostgreSQL / SQLite (for coordinates, movements, transactions).
  * **Secondary DB (Metadata/Profile)**: Firebase Firestore (for user profiles, inventory bags, and player configuration).
