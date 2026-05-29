## Context

`gitctl` discovers git repositories across one or more base directories and runs a git command (`status` or `pull`) on each. Currently, execution is strictly sequential — one repo at a time. Git operations are I/O-bound (network for pull, filesystem for status), making them natural candidates for concurrency.

The configuration layer already exposes `run_mode.concurrency` (YAML, env, and CLI flag), and `config.GetConcurrency()` exists but returns a `string` and is never called by the execution layer. The groundwork is there; it just needs to be wired up correctly.

## Goals / Non-Goals

**Goals:**
- Implement a bounded goroutine worker pool in `gitrepo/gitrepos.go`, limited by `run_mode.concurrency`.
- Print results in deterministic discovery order regardless of completion order.
- Fix `GetConcurrency()` return type from `string` to `int`.
- Keep `concurrency=1` behaviour identical to current sequential behaviour.

**Non-Goals:**
- Streaming live output as repos complete (deferred; collect-then-print is chosen for stability).
- Progress bars or live status indicators.
- Per-command concurrency limits (single limit applies to all commands).
- Changing how errors are reported or aggregated.

## Decisions

### Decision 1: Worker pool with result collection (vs. fan-out goroutines)

**Chosen**: A fixed-size worker pool reads from a job channel. Each worker writes its result (output + error) to a pre-allocated result slice at the repo's discovery index. After all workers finish, the main goroutine iterates the result slice in order and prints.

**Why not unbounded goroutines (one per repo)?** With hundreds of repos, spawning unlimited goroutines risks exhausting file descriptors or memory. The pool naturally applies backpressure.

**Why not channels for results?** A pre-allocated `[]result` indexed by position gives us ordered output for free, without sorting or a second coordination step.

```
Discovery order:  [repo0, repo1, repo2, repo3, repo4]
                        │
                   job channel
                  ┌──────────┐
    worker 1 ◀───┤          ├───▶ results[0], results[2], results[4]
    worker 2 ◀───┤          ├───▶ results[1], results[3]
                  └──────────┘
                        │
                   (all done)
                        │
               print results[0..4] in order
```

### Decision 2: Fix `GetConcurrency()` to return `int`

**Chosen**: Change signature to `GetConcurrency() int` using `viper.GetInt`. Update the CLI flag binding and default from `"1"` (string) to `1` (int). The flag type changes from `StringVarP` to `IntVarP`.

**Why now?** The type is only used by the new worker pool. Fixing it as part of this change avoids a later migration.

### Decision 3: Minimum concurrency of 1

**Chosen**: If the configured value is `< 1`, clamp to `1`. This prevents deadlocks from a zero-worker pool and makes the behaviour predictable.

## Risks / Trade-offs

- **Interleaved filesystem access** — Multiple `git pull` operations on repos sharing a common remote could hit rate limits on the remote. This is a user-configurable concern; the default of `1` is safe. → _No mitigation needed; document in help text._
- **Type change is breaking** — Any external code importing `config.GetConcurrency()` must update. Since this is a CLI tool (not a library), impact is limited to internal callers. → _Fix all internal call sites as part of this change._
- **Result slice pre-allocation** — Requires knowing the full repo list upfront (already the case). Not a concern for typical repo counts.

## Open Questions

- None. Design is fully resolved for the agreed scope.
