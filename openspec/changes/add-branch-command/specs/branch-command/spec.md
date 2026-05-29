## ADDED Requirements

### Requirement: gitctl branch subcommand exists
The CLI SHALL expose a `branch` subcommand that runs `git branch --show-current` concurrently across all repositories discovered from the configured base directories.

#### Scenario: Branch shows current branch name
- **WHEN** a repository is on a named branch
- **THEN** `git branch --show-current` outputs the branch name
- **THEN** the result is displayed in the standard per-repository output format

#### Scenario: Branch on detached HEAD exits 0
- **WHEN** a repository is in detached HEAD state
- **THEN** `git branch --show-current` exits 0 with empty stdout
- **THEN** this is treated as success, not an error

#### Scenario: Branch runs concurrently across all repositories
- **WHEN** multiple repositories are discovered
- **THEN** `git branch --show-current` is dispatched to all repositories via the existing worker pool
- **THEN** results are collected and printed in discovery order

### Requirement: git branch command form
The implementation SHALL execute `exec.Command("git", "branch", "--show-current")`.

#### Scenario: Correct command form used
- **WHEN** the branch subcommand is invoked
- **THEN** the process spawned is `git branch --show-current` with no additional arguments

### Requirement: Branch honours dry-run mode
When `--dry-run` is set, the branch subcommand SHALL not spawn any git processes, following the existing dry-run contract.

#### Scenario: Dry-run with branch
- **WHEN** `gitctl branch --dry-run` is invoked
- **THEN** no `git branch` processes are spawned
- **THEN** a message is printed for each repository indicating what would have run
