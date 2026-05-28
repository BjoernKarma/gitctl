# Spec: Command Execution

## Overview

After discovering repositories, `gitctl` runs a Git command against each one. Currently supported commands are `status` and `pull`.

## Supported Commands

| Subcommand | Git command executed |
|------------|----------------------|
| `status`   | `git status`         |
| `pull`     | `git pull`           |

## Execution Model

Repositories are executed **sequentially** in discovery order.

For each repository:
1. Run `git <command>` in the repository's directory.
2. Capture combined stdout+stderr output.
3. Format the output (see [output spec](../output/spec.md)).
4. If the command fails, record the failure; continue to the next repository.

All repository errors are collected. After all repositories are processed, errors are joined and returned as a single error.

## Dry-Run Mode

When `run_mode.dry_run` is true:
- No `git` process is spawned.
- A message is printed for each repository indicating what would have run.
- No errors are produced.

## Concurrency

The configuration key `run_mode.concurrency` (default: `1`) and corresponding `--concurrency` flag exist but **concurrency is not yet implemented**. All execution is currently sequential regardless of this setting.

See the concurrency change proposal for the planned implementation.

## Error Handling

- A failure in one repository does not stop execution for others.
- Exit code is non-zero if any repository command failed.
- The error message per failure is extracted from git output: explicit `fatal:` / `error:` lines take priority; otherwise the first non-empty output line is used.

## Empty Repository List

If no repositories are discovered (all base dirs empty or skipped), no commands are run and the tool exits with code 0.
