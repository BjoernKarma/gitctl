## Context

`gitctl` uses a uniform pattern for all git subcommands: a cobra command file calls `config.GetBaseDirs()` and `gitrepo.RunGitCommand(command, baseDirs)`. The worker pool, error handling, output formatting, and concurrency are all handled by the shared execution layer — adding a new command requires no changes to that layer.

`git fetch` updates remote-tracking refs (e.g. `origin/main`) without modifying the working tree or local branches. It is the standard way to check what's changed upstream without committing to a merge.

## Goals / Non-Goals

**Goals:**
- Add a `fetch` subcommand that runs `git fetch` across all discovered repositories using the existing execution infrastructure.
- Keep the implementation as small as a copy of `gitpull.go` with `GitFetch` substituted for `GitPull`.

**Non-Goals:**
- Fetch-specific flags (e.g. `--prune`, `--all`, `--tags`) — plain `git fetch` is sufficient for the first iteration.
- Fetch-then-status combined output.
- Any changes to the worker pool, output formatting, or config system.

## Decisions

### Decision 1: Reuse the existing command pattern verbatim

`gitstatus.go` and `gitpull.go` are identical in structure — one new file `gitfetch.go` follows the same pattern. No abstraction is needed; the pattern is clear and consistent.

### Decision 2: Add `GitFetch` constant and `fetch` case to `runRaw`

The `runRaw` switch in `gitrepo.go` maps command strings to `exec.Cmd`. Adding a `fetchCommand = "fetch"` constant and a `case GitFetch` keeps the pattern consistent with `pull` and `status`.

## Risks / Trade-offs

- **Remote availability** — `git fetch` requires network access. Failures are handled identically to `git pull` failures: recorded, reported, non-zero exit. → No mitigation needed beyond existing error handling.
- **Scope creep** — fetch flags (`--prune`, `--all`) will likely be requested. Keeping this first implementation flag-free makes the scope clear. → Document as a non-goal; handle in a future change.
