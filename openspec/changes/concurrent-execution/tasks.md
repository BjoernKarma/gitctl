## 1. Fix Concurrency Configuration Type

- [ ] 1.1 Change `GetConcurrency()` in `config/env.go` to return `int` using `viper.GetInt`
- [ ] 1.2 Update `concurrency` variable in `app/cmd/root.go` from `string` to `int`
- [ ] 1.3 Change the CLI flag registration from `StringVarP` to `IntVarP` with default value `1`
- [ ] 1.4 Verify config file, env var (`GITCTL_RUN_MODE_CONCURRENCY`), and flag all bind correctly as integers

## 2. Implement Worker Pool in gitrepos.go

- [ ] 2.1 Define a `repoResult` struct holding `index int`, `output []byte`, and `err error`
- [ ] 2.2 Replace the sequential `for` loop in `RunGitCommand` with a job channel fed by the discovered repo slice
- [ ] 2.3 Spawn `GetConcurrency()` workers (clamped to minimum 1) as goroutines; each reads from the job channel and writes to a pre-allocated `[]repoResult` slice at the repo's discovery index
- [ ] 2.4 Use a `sync.WaitGroup` to wait for all workers to finish before proceeding
- [ ] 2.5 After all workers complete, iterate `[]repoResult` in index order: print output, collect errors, and register failures with the color package

## 3. Update Specs

- [ ] 3.1 Update `openspec/specs/command-execution/spec.md` to reflect the concurrent execution model (remove the "sequential" guarantee, add ordered-output guarantee)
- [ ] 3.2 Add `openspec/specs/concurrent-execution/spec.md` as the permanent spec for the worker pool capability (promote from change specs)

## 4. Tests

- [ ] 4.1 Add a unit test in `gitrepo/gitrepos_test.go` verifying that results are printed in discovery order when concurrency > 1
- [ ] 4.2 Add a unit test verifying that a concurrency value of `0` or negative is clamped to `1`
- [ ] 4.3 Update or add config tests in `config/env_test.go` to verify `GetConcurrency()` returns an `int` correctly from viper

## 5. Verification

- [ ] 5.1 Run `go build ./...` — no compile errors
- [ ] 5.2 Run `go test ./...` — all tests pass
- [ ] 5.3 Run `gitctl --concurrency 3 status` against a directory with multiple repos and confirm output is in consistent order across runs
