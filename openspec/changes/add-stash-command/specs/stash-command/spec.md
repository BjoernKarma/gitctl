## ADDED Requirements

### Requirement: gitctl stash subcommand exists
The CLI SHALL expose a `stash` subcommand that runs `git stash` concurrently across all repositories discovered from the configured base directories.

#### Scenario: Stash succeeds in repository with changes
- **WHEN** a repository has uncommitted changes
- **THEN** `git stash` runs successfully, the working tree is cleaned, and the result is displayed in the standard per-repository output format

#### Scenario: Stash on clean repository exits 0
- **WHEN** a repository has no uncommitted changes
- **THEN** `git stash` exits 0 and emits "No local changes to save"
- **THEN** this is treated as success, not an error

#### Scenario: Stash runs concurrently across all repositories
- **WHEN** multiple repositories are discovered
- **THEN** `git stash` is dispatched to all repositories via the existing worker pool
- **THEN** results are collected and printed in discovery order

### Requirement: git stash command form
The implementation SHALL execute `exec.Command("git", "stash")` (not `exec.Command("git", "stash", "push")`).

#### Scenario: Short form used
- **WHEN** the stash subcommand is invoked
- **THEN** the process spawned is `git stash` with no additional sub-subcommand argument

### Requirement: Stash honours dry-run mode
When `--dry-run` is set, the stash subcommand SHALL not spawn any git processes, following the existing dry-run contract.

#### Scenario: Dry-run with stash
- **WHEN** `gitctl stash --dry-run` is invoked
- **THEN** no `git stash` processes are spawned
- **THEN** a message is printed for each repository indicating what would have run
