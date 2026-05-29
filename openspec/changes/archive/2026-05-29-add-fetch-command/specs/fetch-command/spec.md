## ADDED Requirements

### Requirement: Fetch command runs git fetch across all repositories
The `gitctl fetch` subcommand SHALL run `git fetch` on every discovered repository using the shared worker pool execution model. Results SHALL be collected in discovery order and reported in the standard output format.

#### Scenario: Successful fetch across multiple repos
- **WHEN** `gitctl fetch` is invoked with one or more base directories containing git repositories
- **THEN** `git fetch` SHALL be executed on each repository and output SHALL be printed in discovery order

#### Scenario: Fetch respects dry-run mode
- **WHEN** `run_mode.dry_run` is true
- **THEN** no `git fetch` process SHALL be spawned and a dry-run message SHALL be printed for each repository

#### Scenario: Fetch respects concurrency setting
- **WHEN** `run_mode.concurrency` is set to N > 1
- **THEN** up to N fetch operations SHALL run simultaneously

#### Scenario: Single repository fetch failure does not stop others
- **WHEN** one repository's `git fetch` fails (e.g. no network, remote not configured)
- **THEN** remaining repositories SHALL still be processed and the command SHALL exit non-zero after all are complete
