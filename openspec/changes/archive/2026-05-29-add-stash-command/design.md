## Context

`gitctl` provides uniform fan-out of git commands across multiple repositories via a shared worker pool in `gitrepo/gitrepos.go`. Adding a new command follows a well-established three-step pattern: add a constant + switch case in `gitrepo/gitrepo.go`, create a cobra command file in `app/cmd/`, and register it in `root.go`.

`git stash` (equivalent to `git stash push`) saves uncommitted changes to the stash stack and restores a clean working tree. It exits 0 whether or not there were changes to stash; repos with nothing to stash emit "No local changes to save" to stdout. This is distinct from an error.

## Goals / Non-Goals

**Goals:**
- Add `gitctl stash` that runs `git stash` across all discovered repositories.
- Treat "No local changes to save" (exit 0) as a success, not a failure — consistent with how git itself behaves.
- Keep implementation identical in structure to `gitfetch.go` / `gitpull.go`.

**Non-Goals:**
- `git stash pop` / `git stash apply` / `git stash list` sub-subcommands — these are natural follow-ons but out of scope for this change.
- A `gitctl stash pop` paired shortcut — deferred to a future `stash-subcommands` change.
- Any stash message (`-m`) or include-untracked (`-u`) flag support in this iteration.

## Decisions

### Decision 1: Run `git stash` (not `git stash push`)

**Chosen**: Execute `exec.Command("git", "stash")` rather than `exec.Command("git", "stash", "push")`. Both are equivalent since git 2.13, and the shorter form matches what users type interactively.

### Decision 2: No special-casing of "No local changes to save"

**Chosen**: The worker pool collects raw output and non-zero exit codes as errors. `git stash` exits 0 when there is nothing to stash, so the existing error handling naturally treats this as success. No special logic needed.

**Why not filter the message?** The output formatter already handles the output display. Filtering would add complexity for no user benefit — the message is informative, not alarming.

## Risks / Trade-offs

- **Stash stack pollution** — Running `gitctl stash` across many repos creates one stash entry per repo that had changes. Users must remember to `git stash pop` in each repo. → Document in command short description; a future `gitctl stash pop` command addresses recovery.
- **Untracked files not stashed by default** — Plain `git stash` does not stash untracked files. Users needing that would use `git stash -u` manually. → Acceptable for v1; can add `--include-untracked` flag in a follow-on.
