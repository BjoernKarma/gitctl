## Why

`gitctl` supports `status` and `pull` across multiple repositories, but lacks `fetch` — a safe, read-only operation that updates remote-tracking refs without touching the working tree. This makes it impossible to check whether repos are behind their remotes without also running a full pull.

## What Changes

- Add a `fetch` subcommand that runs `git fetch` on all discovered repositories.
- Register `fetch` alongside `status` and `pull` in the root command.
- Extend the `GitRepo.runRaw` command switch to handle `fetch`.

## Capabilities

### New Capabilities

- `fetch-command`: The `gitctl fetch` subcommand — runs `git fetch` concurrently across all discovered repositories, collects results in discovery order, and reports success/failures in the standard output format.

### Modified Capabilities

- `command-execution`: A new supported command (`fetch`) is added to the command table. No behavioral changes to the execution model itself.

## Impact

- `app/cmd/gitfetch.go`: New file defining the `fetchCmd` cobra command.
- `app/cmd/root.go`: Register `fetchCmd`.
- `gitrepo/gitrepo.go`: Add `GitFetch` constant and handle it in the `runRaw` switch.
- `openspec/specs/command-execution/spec.md`: Add `fetch` to the supported commands table.
