## MODIFIED Requirements

### Requirement: Execution Model
`gitctl` SHALL execute git commands across all discovered repositories using a bounded worker pool. The number of concurrent workers SHALL be controlled by `run_mode.concurrency`. Results SHALL be collected and printed in discovery order after all repositories have been processed.

For each repository a worker SHALL:
1. Run `git <command>` in the repository's directory.
2. Capture combined stdout+stderr output.
3. Store the result (output + error) at the repository's discovery index.
4. If the command fails, record the failure; continue to the next job.

After all workers complete, the main goroutine SHALL iterate results in discovery order, print output, and aggregate errors.

#### Scenario: All repos processed with concurrency > 1
- **WHEN** `run_mode.concurrency` is 3 and 10 repositories are discovered
- **THEN** all 10 repositories SHALL be processed and their output printed in discovery order

#### Scenario: Single failure does not stop other repos
- **WHEN** one repository's git command fails
- **THEN** remaining repositories in the pool SHALL still be processed

#### Scenario: All errors collected and returned
- **WHEN** multiple repositories fail
- **THEN** all errors SHALL be collected and the command SHALL exit with a non-zero exit code after all repositories are processed
