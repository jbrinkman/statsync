## 1. Project Scaffold

- [x] 1.1 Initialize Go module (`github.com/jbrinkman/statsync`)
- [x] 1.2 Add cobra, viper, and valkey-go dependencies
- [x] 1.3 Create main.go with cobra root command and `--config` flag
- [x] 1.4 Set up internal package structure (separate CLI layer from business logic for future TUI)
- [x] 1.5 Create Docker Compose file with Valkey bundle image

## 2. Configuration

- [x] 2.1 Implement viper config loading (flag → config file → env → defaults)
- [x] 2.2 Support `$HOME/.statsync.yaml` as default config path
- [x] 2.3 Support `--config` flag for custom config file path

## 3. Valkey Connection

- [x] 3.1 Implement Valkey client initialization using config-resolved address
- [x] 3.2 Add connection error handling with clear error message

## 4. Status Model

- [x] 4.1 Define canonical status constants and validation function
- [x] 4.2 Implement case-insensitive status normalization
- [x] 4.3 Implement computed project status derivation from work item statuses

## 5. Data Model

- [x] 5.1 Define Project struct (name, slug, type, framework, assignee, workItems, notes, updatedAt)
- [x] 5.2 Define WorkItem struct (label, status, prs)
- [x] 5.3 Implement slug generation from project name

## 6. CRUD Commands

- [x] 6.1 Implement `statsync add` command (create project with optional metadata)
- [x] 6.2 Implement `statsync list` command (display all projects with status summary, optional type filter)
- [x] 6.3 Implement `statsync show` command (display full project details)
- [x] 6.4 Implement `statsync update` command (update metadata, add/update work items, add PR links)
- [x] 6.5 Implement `statsync delete` command (remove project)

## 7. Validation

- [x] 7.1 Add duplicate project detection on add
- [x] 7.2 Add not-found error handling for show/update/delete
- [x] 7.3 Add status validation on work item status changes
