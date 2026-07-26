# Defaultia / BabelSaga (The Esperanto Multi-Agent SimCity RPG)

A 2D SimCity-style multi-namespace simulation game driven by AI agents, Cloud Run, and Model Context Protocol (MCP).

## Key Mechanics & Architecture
- **Language Prohibition & Esperanto Lingua Franca**: English is strictly forbidden! Players choose a native language, but Esperanto is the universal language for chat commands (`/krii`, `/flustri`, `/helpo`) and inter-player communication.
- **3 Starting Droids**:
  1. Protocol / Translator Droid (Active): Interprets and translates between foreign languages.
  2. Gatherer / Worker Droid (Active): Explores the 2D emoji grid, asks questions, collects resources.
  3. Generic ADK AI Droid (Inactive): Programmed by the player using ADK / MCP and deployed to Cloud Run.
- **Bot Protocol (`babot` / Omega13)**: JSON message format with `rot13`/`base64` payloads and status codes (`VABON`, `Aninpospio`). Droids only speak their assigned language unless asked in Esperanto: *"Kian lingvon vi parolas?"*.
- **Universal Basic Income & Currency**: Daily income of 100 Gemoj (💎) per day (tribute to Gemini!).
- **7-Item Backpack**: 7 slots maximum. Includes default items, 1 huge comic item (e.g. 3-story ladder), and a rule scroll in Esperanto.
- **Multi-Instance & Terraform IaC**: Spin up custom city instances for hackathons (e.g., *Mediolanum 3125* or *Alexandria*) using provided Terraform recipes.
