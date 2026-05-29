## 1. Core Implementation

- [x] 1.1 Add `GitFetch = "fetch"` constant and `fetchCommand = "fetch"` in `gitrepo/gitrepo.go`
- [x] 1.2 Add `case GitFetch` to the `runRaw` command switch in `gitrepo/gitrepo.go`
- [x] 1.3 Create `app/cmd/gitfetch.go` defining `fetchCmd` following the same pattern as `gitpull.go`
- [x] 1.4 Register `fetchCmd` with `rootCmd.AddCommand(fetchCmd)` in `app/cmd/root.go`

## 2. Tests

- [x] 2.1 Add `TestRunGitFetch` in `app/cmd/gitfetch_test.go` (dry-run mode, mirrors `gitpull_test.go`)
- [x] 2.2 Add `TestGitRepoRunGitFetch` in `gitrepo/gitrepo_test.go` verifying `GitFetch` executes against a valid repo path
- [x] 2.3 Add `TestRunGitFetchCommand` in `gitrepo/gitrepos_test.go` verifying `RunGitCommand(GitFetch, baseDirs)` succeeds in dry-run

## 3. Spec Promotion

- [x] 3.1 Add `fetch` to the supported commands table in `openspec/specs/command-execution/spec.md`
- [x] 3.2 Create `openspec/specs/fetch-command/spec.md` (promote from change specs)

## 4. Verification

- [x] 4.1 Run `go build ./...` — no compile errors
- [x] 4.2 Run `go test -race ./...` — all tests pass, no races
- [x] 4.3 Run `go fmt ./...` on changed files
- [x] 4.4 Verify `gitctl fetch --help` shows the correct usage description
