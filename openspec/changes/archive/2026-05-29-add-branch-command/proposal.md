## Why

When working across many feature branches simultaneously, it's essential to know which branch each repository is on at a glance. `git status` provides this but includes working-tree noise; there's no clean `gitctl` command that shows just the current branch per repo.

## What Changes

- Add a `branch` subcommand that runs `git branch --show-current` on all discovered repositories and reports the current branch name for each.
- Register `branch` alongside `status`, `pull`, `fetch`, and `stash` in the root command.
- Extend the `runRaw` command switch to handle `branch`.
- Add `branch` to the supported commands table in the command-execution spec.

## Capabilities

### New Capabilities

- `branch-command`: The `gitctl branch` subcommand — runs `git branch --show-current` concurrently across all discovered repositories, collects results in discovery order, and reports the current branch name per repo using the standard output format.

### Modified Capabilities

- `command-execution`: A new supported command (`branch`) is added to the command table. No behavioral changes to the execution model itself.

## Impact

- `app/cmd/gitbranch.go`: New file defining the `branchCmd` cobra command.
- `app/cmd/root.go`: Register `branchCmd`.
- `gitrepo/gitrepo.go`: Add `GitBranch` constant and `branchCommand` const (`--show-current` flag passed); handle in the `runRaw` switch.
- `openspec/specs/command-execution/spec.md`: Add `branch` to the supported commands table.
