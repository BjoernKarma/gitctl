## 1. Core Implementation

- [x] 1.1 Add `GitStash = "stash"` constant and `stashCommand = "stash"` in `gitrepo/gitrepo.go`
- [x] 1.2 Add `case GitStash` to the `runRaw` command switch in `gitrepo/gitrepo.go`
- [x] 1.3 Create `app/cmd/gitstash.go` defining `stashCmd` following the same pattern as `gitfetch.go`; set short description to "Execute git stash on multiple git repositories."
- [x] 1.4 Register `stashCmd` with `rootCmd.AddCommand(stashCmd)` in `app/cmd/root.go`

## 2. Tests

- [x] 2.1 Add `TestRunGitStash` in `app/cmd/gitstash_test.go` (dry-run mode, mirrors `gitfetch_test.go`)
- [x] 2.2 Add `TestGitRepoRunGitStash` in `gitrepo/gitrepo_test.go` verifying `GitStash` executes against a valid repo path
- [x] 2.3 Add `TestRunGitStashCommand` in `gitrepo/gitrepos_test.go` verifying `RunGitCommand(GitStash, baseDirs)` succeeds in dry-run

## 3. Spec Promotion

- [x] 3.1 Add `stash` to the supported commands table in `openspec/specs/command-execution/spec.md`
- [x] 3.2 Create `openspec/specs/stash-command/spec.md` (promote from change specs)

## 4. Verification

- [x] 4.1 Run `go build ./...` — no compile errors
- [x] 4.2 Run `go test -race ./...` — all tests pass, no races
- [x] 4.3 Run `go fmt ./...` on changed files
- [x] 4.4 Verify `gitctl stash --help` shows the correct usage description

