## ADDED Requirements

### Requirement: Canonical status set
The system SHALL recognize exactly 12 statuses for work item lifecycle tracking, organized into synonym groups that represent the same progression stage.

#### Scenario: Valid statuses accepted
- **WHEN** user sets a work item status to any of: "Not Started", "In Progress", "In Review", "Submitted", "Awaiting Merge", "Merged", "Done", "Blocked", "Paused", "Closed", "Not a Bug", "Dropped"
- **THEN** the status is accepted and stored

#### Scenario: Invalid status rejected
- **WHEN** user sets a work item status to a value not in the canonical set
- **THEN** the system SHALL reject the input and display the list of valid statuses

### Requirement: Status synonym groups
The system SHALL treat the following statuses as equivalent stages for progression ordering: "Submitted" and "Awaiting Merge" represent the same stage; "Done", "Closed", "Not a Bug", and "Dropped" represent the same terminal stage.

#### Scenario: Synonyms at same progression level
- **WHEN** computing project status with one item at "Submitted" and another at "Awaiting Merge"
- **THEN** the system treats both as the same progression level

### Requirement: Status progression ordering
The system SHALL use the following progression ordering for status comparison: Not Started (1) < In Progress (2) < In Review (3) < Submitted/Awaiting Merge (4) < Merged (5) < Done/Closed/Not a Bug (6). Merged is NOT a terminal state; only Done/Closed/Not a Bug are terminal.

#### Scenario: Merged is not terminal
- **WHEN** all work items have status "Merged"
- **THEN** the project status is "Merged", not "Done"

### Requirement: Default status assignment
The system SHALL assign "Not Started" as the default status when a new work item is created without an explicit status.

#### Scenario: Work item created without status
- **WHEN** a work item is added without specifying a status
- **THEN** the work item's status is set to "Not Started"

### Requirement: Computed project status
The system SHALL derive a project's status from its work items' statuses using the following precedence rules.

#### Scenario: Any item blocked
- **WHEN** any work item has status "Blocked"
- **THEN** the project status is "Blocked"

#### Scenario: Any item paused (none blocked)
- **WHEN** any work item has status "Paused" and no item is "Blocked"
- **THEN** the project status is "Paused"

#### Scenario: All items not started
- **WHEN** all work items have status "Not Started"
- **THEN** the project status is "Not Started"

#### Scenario: Some items progressed
- **WHEN** at least one work item has moved past "Not Started" and none are Blocked or Paused
- **THEN** the project status is the lowest status among non-"Not Started" items, using the ordering: In Progress (2) < In Review (3) < Submitted/Awaiting Merge (4) < Merged (5) < Done/Closed/Not a Bug (6)

### Requirement: Case-insensitive status input
The system SHALL accept status values in any case and normalize them to the canonical form.

#### Scenario: Lowercase input normalized
- **WHEN** user provides status as "in progress"
- **THEN** the system stores it as "In Progress"

#### Scenario: Mixed case input normalized
- **WHEN** user provides status as "awaiting merge"
- **THEN** the system stores it as "Awaiting Merge"
