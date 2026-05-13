## Why

Integration project status is scattered across 4 systems (Jira, Confluence, weekly markdown reports, Quip) with no single source of truth. Reporting and keeping systems in sync is manual and error-prone. We need a CLI tool backed by Valkey that becomes the canonical store for project status, enabling AI agents to read from Jira and write to other systems.

## What Changes

- Scaffold a Go CLI application (`statsync`) with Valkey connectivity
- Define a project data model with flexible work items per project
- Implement CRUD commands: `add`, `update`, `list`, `show`, `delete`
- Define 11 canonical statuses matching the weekly report status set
- Provide a Docker Compose setup for local Valkey development
- Store projects as JSON documents in Valkey using the JSON module

## Capabilities

### New Capabilities
- `project-management`: CRUD operations for projects and their work items via CLI
- `status-model`: Canonical status enum with definitions for work item lifecycle tracking
- `valkey-storage`: JSON document storage in Valkey for project data

### Modified Capabilities

## Impact

- New Go module at project root with CLI entrypoint
- Dependency on Valkey (via Docker for local dev)
- Dependency on `valkey-go` client library
- Dependency on a CLI framework (e.g., cobra)
