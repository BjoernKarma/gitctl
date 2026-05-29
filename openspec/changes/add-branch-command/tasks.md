## 1. Core Implementation

- [ ] 1.1 Add `GitBranch = "branch"` constant and `branchCommand = "branch"` in `gitrepo/gitrepo.go`
- [ ] 1.2 Add `case GitBranch` to the `runRaw` command switch in `gitrepo/gitrepo.go`, executing `exec.Command(gitCommand, branchCommand, "--show-current")`
- [ ] 1.3 Create `app/cmd/gitbranch.go` defining `branchCmd` following the same pattern as `gitfetch.go`; set short description to "Show current branch for multiple git repositories."
- [ ] 1.4 Register `branchCmd` with `rootCmd.AddCommand(branchCmd)` in `app/cmd/root.go`

## 2. Tests

- [ ] 2.1 Add `TestBranchCommandExecutesGitBranchOnLocalRepo` in `app/cmd/gitbranch_test.go` (dry-run mode, mirrors `gitfetch_test.go`)
- [ ] 2.2 Add `TestGitRepoRunGitBranch` in `gitrepo/gitrepo_test.go` verifying `GitBranch` executes against a valid repo path
- [ ] 2.3 Add `TestRunGitBranchCommand` in `gitrepo/gitrepos_test.go` verifying `RunGitCommand(GitBranch, baseDirs)` succeeds in dry-run

## 3. Spec Promotion

- [ ] 3.1 Add `branch` to the supported commands table in `openspec/specs/command-execution/spec.md`
- [ ] 3.2 Create `openspec/specs/branch-command/spec.md` (promote from change specs)

## 4. Verification

- [ ] 4.1 Run `go build ./...` — no compile errors
- [ ] 4.2 Run `go test -race ./...` — all tests pass, no races
- [ ] 4.3 Run `go fmt ./...` on changed files
- [ ] 4.4 Verify `gitctl branch --help` shows the correct usage description
