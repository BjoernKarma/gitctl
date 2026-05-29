## ADDED Requirements

### Requirement: Fetch is a supported command
The `fetch` command SHALL be a valid value for the git command executed by the worker pool, alongside `status` and `pull`.

#### Scenario: gitctl fetch maps to git fetch
- **WHEN** the `fetch` subcommand is invoked
- **THEN** `git fetch` SHALL be executed in each discovered repository's directory
