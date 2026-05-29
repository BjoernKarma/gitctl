## MODIFIED Requirements

### Requirement: Supported commands
`gitctl` SHALL support the following subcommands, each mapping to a single `git` invocation per repository.

| Subcommand | Git command executed          |
|------------|-------------------------------|
| `status`   | `git status`                  |
| `pull`     | `git pull`                    |
| `fetch`    | `git fetch`                   |
| `stash`    | `git stash`                   |
| `branch`   | `git branch --show-current`   |

#### Scenario: branch subcommand dispatches git branch --show-current
- **WHEN** the user runs `gitctl branch`
- **THEN** `git branch --show-current` is executed in every discovered repository
- **THEN** results are printed in discovery order using the standard output format
