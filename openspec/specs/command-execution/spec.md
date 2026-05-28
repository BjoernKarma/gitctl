# Spec: Command Execution

## Overview

After discovering repositories, `gitctl` runs a Git command against each one using a bounded worker pool. Currently supported commands are `status` and `pull`.

## Supported Commands

| Subcommand | Git command executed |
|------------|----------------------|
| `status`   | `git status`         |
| `pull`     | `git pull`           |

## Execution Model

Repositories are executed using a bounded goroutine worker pool controlled by `run_mode.concurrency` (default: `1`). When concurrency is `1`, behaviour is identical to sequential execution.

For each repository a worker:
1. Runs `git <command>` in the repository's directory.
2. Captures combined stdout+stderr output.
3. Stores the result (output + error) at the repository's discovery index.
4. If the command fails, records the failure and continues to the next job.

After all workers complete, results are iterated in discovery order: output is printed and errors are aggregated.

## Dry-Run Mode

When `run_mode.dry_run` is true:
- No `git` process is spawned.
- A message is printed for each repository indicating what would have run.
- No errors are produced.

## Concurrency

The `run_mode.concurrency` setting (default: `1`) and `--concurrency` / `-C` flag control the worker pool size. Values less than `1` are clamped to `1`.

See the [concurrent-execution spec](../concurrent-execution/spec.md) for the full worker pool specification.

## Error Handling

- A failure in one repository does not stop others from being processed.
- Exit code is non-zero if any repository command failed.
- The error message per failure is extracted from git output: explicit `fatal:` / `error:` lines take priority; otherwise the first non-empty output line is used.

## Output Order

Results are always printed in discovery order, regardless of the order in which workers complete.

## Empty Repository List

If no repositories are discovered (all base dirs empty or skipped), no commands are run and the tool exits with code 0.
