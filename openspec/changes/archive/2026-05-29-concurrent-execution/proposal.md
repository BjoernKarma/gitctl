## Why

`gitctl` runs git commands on all discovered repositories sequentially, even though the operations are independent and I/O-bound. With many repositories, this is unnecessarily slow. The `concurrency` configuration key and CLI flag already exist as a stub but have no effect.

## What Changes

- Wire `run_mode.concurrency` (and `--concurrency` / `-C` flag) to an actual worker pool that runs git commands in parallel.
- Fix the type of `GetConcurrency()` from `string` to `int` (**BREAKING** for any code calling that function).
- Collect results from all workers and print them in discovery order (not completion order) to keep output stable and readable.
- Update the `command-execution` spec to reflect the new concurrent execution model.

## Capabilities

### New Capabilities

- `concurrent-execution`: A bounded worker pool that runs git commands across repositories in parallel, limited to `run_mode.concurrency` goroutines, with results collected and printed in deterministic discovery order.

### Modified Capabilities

- `command-execution`: Execution model changes from strictly sequential to concurrent. The sequential guarantee is replaced by an ordered-output guarantee.

## Impact

- `config/env.go`: `GetConcurrency()` return type changes from `string` to `int`.
- `gitrepo/gitrepos.go`: `RunGitCommand` gains a goroutine worker pool; result collection replaces the current inline print loop.
- `app/cmd/root.go`: Flag default and bind type updated to match `int`.
- `openspec/specs/command-execution/spec.md`: Updated to describe concurrent execution model.
