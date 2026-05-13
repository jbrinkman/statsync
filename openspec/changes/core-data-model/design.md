## Context

StatSync is a new CLI tool for a single user to track integration project status across multiple systems. Currently no code exists — this is a greenfield Go CLI backed by Valkey's JSON module. The user manages hundreds of Valkey integration projects for open source agentic frameworks and needs a canonical store to drive reporting to Confluence, weekly markdown reports, and Quip.

## Goals / Non-Goals

**Goals:**
- Establish the Go CLI scaffold with Valkey connectivity
- Define a flexible project/work-item data model stored as JSON in Valkey
- Implement core CRUD commands (add, list, show, update, delete)
- Define the canonical status enum used across all reporting
- Provide a Docker Compose file for local Valkey development

**Non-Goals:**
- Jira import (separate proposal)
- Export/reporting to external systems (separate proposal)
- Web UI or API server
- Historical status tracking or audit log
- Multi-user access control

## Decisions

### CLI Architecture: Command Layer Separation

Separate command parsing from business logic to allow a TUI frontend to be added later without rewriting core functionality.

**Rationale**: The user wants CLI for MVP but may add a TUI mode later. Keeping business logic in a `pkg/` or `internal/` layer means a TUI can call the same functions without going through cobra.

### Configuration: File with Flag Override

Use a configuration file with the following resolution order (highest priority first):
1. CLI flag (`--config <path>`)
2. Default config file (`$HOME/.statsync.yaml`)
3. Hardcoded defaults

Config file format: YAML. Managed via `github.com/spf13/viper` (pairs naturally with cobra).

**Configurable values**:
- `valkey.addr` (default: `localhost:6379`)
- Additional settings as needed (export formats, status mappings, etc. in future proposals)

**Rationale**: A config file avoids repetitive flags for settings that rarely change. Viper provides file + env + flag merging out of the box. YAML is human-readable and easy to edit.

**Alternatives considered**:
- TOML — slightly less common in Go CLI ecosystem
- JSON — harder to hand-edit with comments
- No config file (env vars only) — insufficient for growing config surface

Use `github.com/spf13/cobra` for CLI structure.

**Rationale**: De facto standard for Go CLIs. Provides subcommand routing, flag parsing, help generation, and shell completions. Well-maintained and widely understood.

**Alternatives considered**:
- `urfave/cli` — simpler but less ecosystem support for completions
- No framework (just `flag`) — too much boilerplate for a multi-command CLI

### Valkey Client: valkey-go

Use `github.com/valkey-io/valkey-go` as the client library.

**Rationale**: Official Valkey client for Go. Supports the JSON module commands natively. Dog-foods the team's own product.

### Data Model: JSON Documents

Store each project as a single JSON document in Valkey using the JSON module (`JSON.SET`, `JSON.GET`).

**Key scheme**: `project:<slug>` where slug is a kebab-case identifier derived from the project name.

**Rationale**: JSON module allows partial reads/updates via JSONPath, keeps the data model flexible (work items vary per project), and avoids the impedance mismatch of mapping to Redis hashes.

**Alternatives considered**:
- Redis hashes — too flat for nested work items
- Separate keys per work item — over-normalized, complicates listing

### Status Enum: String Constants

Statuses are stored as string values matching the canonical set. Validation happens at the CLI layer before writing to Valkey.

**Canonical statuses**: Not Started, In Progress, In Review, Submitted, Awaiting Merge, Merged, Done, Blocked, Paused, Closed, Not a Bug.

### Project Listing: Valkey Key Scan

Use `SCAN` with pattern `project:*` to enumerate projects. For the expected scale (~hundreds), this is sufficient without a secondary index.

**Rationale**: Simple, no additional data structures needed. If scale grows to thousands, we can add a sorted set index later.

## Risks / Trade-offs

- **[No secondary indexes]** → Filtering by status or assignee requires scanning all projects. Acceptable at hundreds of projects. Mitigation: add a status index set later if needed.
- **[Single Valkey instance]** → No replication or persistence guarantees beyond RDB snapshots. Mitigation: acceptable for a personal tool; data can be re-imported from Jira.
- **[Schema flexibility]** → No enforced schema in Valkey means invalid data could be written by other tools. Mitigation: CLI validates on write; add a `JSON.SET` schema validation hook if needed later.
