## Context

`gitctl` provides uniform fan-out of git commands across multiple repositories via a shared worker pool in `gitrepo/gitrepos.go`. Adding a new command follows a well-established three-step pattern: add a constant + switch case in `gitrepo/gitrepo.go`, create a cobra command file in `app/cmd/`, and register it in `root.go`.

`git branch --show-current` prints the name of the currently checked-out branch and exits 0. In detached HEAD state it prints an empty string and still exits 0.

## Goals / Non-Goals

**Goals:**
- Add `gitctl branch` that runs `git branch --show-current` across all discovered repositories.
- Show the current branch name per repo using the standard output format.
- Keep implementation structurally identical to `gitfetch.go` / `gitstash.go`.

**Non-Goals:**
- Listing all branches (`git branch --all`) — that is a distinct use case for a future change.
- Switching branches across repos — out of scope.
- Filtering or grouping repos by branch name — deferred to a future change.
- Special-casing detached HEAD with a human-friendly label — acceptable edge case; empty output is unambiguous to git users.

## Decisions

### Decision 1: Use `git branch --show-current` (not `git rev-parse --abbrev-ref HEAD`)

**Chosen**: `exec.Command("git", "branch", "--show-current")`.

`--show-current` was introduced in git 2.22 (2019) and is the canonical modern form. It is clean, predictable, and unambiguous. `git rev-parse --abbrev-ref HEAD` is a workaround for older gits and returns the literal string `"HEAD"` in detached state rather than empty — less clean.

**Trade-off**: Requires git ≥ 2.22. gitctl has no explicit minimum git version today; this is acceptable given git 2.22 is over 5 years old.

### Decision 2: Pass `--show-current` as a second argument in the switch case

**Chosen**: The `runRaw` switch handles `GitBranch` with `exec.Command(gitCommand, branchCommand, "--show-current")`. The `branchCommand` constant is `"branch"` (consistent with other single-word constants). The flag is passed inline in the case — no additional constant needed.

**Why not a separate constant?** The flag is not reused elsewhere; inlining it keeps the code readable without adding noisy constants.

### Decision 3: Detached HEAD emits empty output — no special handling

**Chosen**: `git branch --show-current` exits 0 with empty stdout in detached HEAD state. The existing output formatter will display an empty body for that repo. No special-casing is added.

**Why not add a "(detached HEAD)" label?** Would require touching the `color` package global state or adding logic to `FormatOutput`, adding complexity for a rare edge case. A future `output-enhancements` change can address this.

## Risks / Trade-offs

- **git < 2.22** — `--show-current` is not recognised; git exits non-zero and reports an error. → Documented in command description; acceptable given the age of git 2.22.
- **Empty output for detached HEAD** — may confuse users who expect to see "HEAD" or similar. → Acceptable for v1; document in short description or a future follow-on.
