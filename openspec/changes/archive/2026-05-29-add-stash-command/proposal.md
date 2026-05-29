## Why

Before running `gitctl pull` across many repositories, any local modifications must be safely stowed away or the pull will fail or produce conflicts. Doing `git stash` manually in each repo is tedious. A `gitctl stash` command fans out `git stash` across all discovered repositories in one step, making the common "stash → pull → pop" workflow frictionless.

## What Changes

- Add a `stash` subcommand that runs `git stash` (i.e. `git stash push`) on all discovered repositories.
- Register `stash` alongside `status`, `pull`, and `fetch` in the root command.
- Extend the `runRaw` command switch to handle `stash`.
- Add `stash` to the supported commands table in the command-execution spec.

## Capabilities

### New Capabilities

- `stash-command`: The `gitctl stash` subcommand — runs `git stash` concurrently across all discovered repositories, collects results in discovery order, and reports success/failures in the standard output format. Repos with no local changes produce a "No local changes to save" message (exit 0); this is not treated as an error.

### Modified Capabilities

- `command-execution`: A new supported command (`stash`) is added to the command table. No behavioral changes to the execution model itself.

## Impact

- `app/cmd/gitstash.go`: New file defining the `stashCmd` cobra command.
- `app/cmd/root.go`: Register `stashCmd`.
- `gitrepo/gitrepo.go`: Add `GitStash` constant and `stashCommand` const; handle in the `runRaw` switch.
- `openspec/specs/command-execution/spec.md`: Add `stash` to the supported commands table.
