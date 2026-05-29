## MODIFIED Requirements

### Requirement: Supported commands
`gitctl` SHALL support the following subcommands, each mapping to a single `git` invocation per repository.

| Subcommand | Git command executed |
|------------|----------------------|
| `status`   | `git status`         |
| `pull`     | `git pull`           |
| `fetch`    | `git fetch`          |
| `stash`    | `git stash`          |

#### Scenario: stash subcommand dispatches git stash
- **WHEN** the user runs `gitctl stash`
- **THEN** `git stash` is executed in every discovered repository
- **THEN** results are printed in discovery order using the standard output format
