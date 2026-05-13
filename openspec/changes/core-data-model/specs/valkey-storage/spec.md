## ADDED Requirements

### Requirement: JSON document storage
The system SHALL store each project as a JSON document in Valkey using the JSON module.

#### Scenario: Project stored as JSON
- **WHEN** a project is created or updated
- **THEN** the project data is persisted at key `project:<slug>` as a JSON document

#### Scenario: Project JSON structure
- **WHEN** a project is stored
- **THEN** the JSON document SHALL contain fields: name, slug, type, framework, assignee, workItems, notes, updatedAt

### Requirement: Valkey connection configuration
The system SHALL connect to Valkey using an address resolved from the following sources in priority order: CLI flag, config file, environment variable, hardcoded default.

#### Scenario: Default connection
- **WHEN** no configuration is provided
- **THEN** the system connects to `localhost:6379`

#### Scenario: Config file
- **WHEN** `valkey.addr` is set in `$HOME/.statsync.yaml`
- **THEN** the system connects to the specified address

#### Scenario: Environment variable override
- **WHEN** the `STATSYNC_VALKEY_ADDR` environment variable is set
- **THEN** it takes precedence over the config file value

#### Scenario: CLI flag override
- **WHEN** the `--valkey-addr` flag is provided
- **THEN** it takes precedence over all other sources

#### Scenario: Custom config file path
- **WHEN** the `--config` flag is provided with a file path
- **THEN** the system reads configuration from that file instead of the default location

### Requirement: Connection error handling
The system SHALL display a clear error message when Valkey is unreachable.

#### Scenario: Valkey unavailable
- **WHEN** the system cannot connect to Valkey
- **THEN** the system displays an error message indicating Valkey is unreachable and exits with a non-zero code
