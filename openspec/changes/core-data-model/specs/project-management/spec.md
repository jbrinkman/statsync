## ADDED Requirements

### Requirement: Add a project
The system SHALL allow the user to create a new project with a name and optional type, framework, and assignee.

#### Scenario: Create a minimal project
- **WHEN** user runs `statsync add "DB-GPT"`
- **THEN** a project is created with name "DB-GPT", slug "db-gpt", and empty work items list

#### Scenario: Create a project with metadata
- **WHEN** user runs `statsync add "DB-GPT" --type integration --framework "DB-GPT" --assignee jbrinkman`
- **THEN** a project is created with the specified type, framework, and assignee fields populated

#### Scenario: Reject duplicate project name
- **WHEN** user runs `statsync add "DB-GPT"` and a project with slug "db-gpt" already exists
- **THEN** the system SHALL display an error and not overwrite the existing project

### Requirement: List projects
The system SHALL display all tracked projects with their current status summary.

#### Scenario: List all projects
- **WHEN** user runs `statsync list`
- **THEN** the system displays each project's name, type, and a count of work items by status

#### Scenario: Filter by type
- **WHEN** user runs `statsync list --type integration`
- **THEN** only projects with type "integration" are displayed

#### Scenario: Empty project list
- **WHEN** user runs `statsync list` and no projects exist
- **THEN** the system displays a message indicating no projects are tracked

### Requirement: Show project details
The system SHALL display full details of a single project including all work items and their statuses.

#### Scenario: Show existing project
- **WHEN** user runs `statsync show db-gpt`
- **THEN** the system displays the project name, type, framework, assignee, and all work items with their labels, statuses, and PR links

#### Scenario: Show non-existent project
- **WHEN** user runs `statsync show nonexistent`
- **THEN** the system displays an error indicating the project was not found

### Requirement: Update a project
The system SHALL allow the user to modify project metadata and work item statuses.

#### Scenario: Update project metadata
- **WHEN** user runs `statsync update db-gpt --assignee someone-else`
- **THEN** the project's assignee field is updated

#### Scenario: Add a work item
- **WHEN** user runs `statsync update db-gpt --add-item "vector-store"`
- **THEN** a new work item with label "vector-store" and status "Not Started" is added to the project

#### Scenario: Update work item status
- **WHEN** user runs `statsync update db-gpt --item "vector-store" --status "In Progress"`
- **THEN** the work item's status is changed to "In Progress"

#### Scenario: Add PR link to work item
- **WHEN** user runs `statsync update db-gpt --item "vector-store" --add-pr "https://github.com/..."`
- **THEN** the PR URL is appended to the work item's PR links list

#### Scenario: Reject invalid status
- **WHEN** user runs `statsync update db-gpt --item "vector-store" --status "Invalid"`
- **THEN** the system displays an error listing valid statuses

### Requirement: Delete a project
The system SHALL allow the user to remove a project from tracking.

#### Scenario: Delete existing project
- **WHEN** user runs `statsync delete db-gpt`
- **THEN** the project is removed from Valkey

#### Scenario: Delete non-existent project
- **WHEN** user runs `statsync delete nonexistent`
- **THEN** the system displays an error indicating the project was not found
