## ADDED Requirements

### Requirement: Bounded worker pool execution
`gitctl` SHALL execute git commands across repositories using a bounded goroutine worker pool. The pool size SHALL be determined by `run_mode.concurrency`. When concurrency is `1`, behaviour SHALL be identical to sequential execution.

#### Scenario: Concurrent execution with pool size > 1
- **WHEN** `run_mode.concurrency` is set to `N` (N > 1) and there are M repositories (M > N)
- **THEN** at most N git commands SHALL run simultaneously at any given time

#### Scenario: Default sequential behaviour preserved
- **WHEN** `run_mode.concurrency` is `1` (the default)
- **THEN** repositories SHALL be processed one at a time, matching previous sequential behaviour

#### Scenario: Concurrency value below minimum is clamped
- **WHEN** `run_mode.concurrency` is set to `0` or a negative value
- **THEN** the tool SHALL clamp the pool size to `1` and proceed without error

### Requirement: Ordered output regardless of completion order
Results SHALL be printed in the original discovery order, regardless of the order in which repositories complete.

#### Scenario: Output order matches discovery order
- **WHEN** multiple repositories are processed concurrently and complete in arbitrary order
- **THEN** their output SHALL be printed in the same order as they were discovered (deterministic)

### Requirement: Concurrency configuration type is integer
The `run_mode.concurrency` setting and its corresponding `--concurrency` / `-C` CLI flag SHALL be typed as an integer throughout the configuration layer.

#### Scenario: Integer value accepted from config file
- **WHEN** `gitctl.yaml` contains `run_mode: {concurrency: 4}`
- **THEN** the pool SHALL use 4 workers

#### Scenario: Integer value accepted from environment variable
- **WHEN** `GITCTL_RUN_MODE_CONCURRENCY=4` is set
- **THEN** the pool SHALL use 4 workers

#### Scenario: Integer value accepted from CLI flag
- **WHEN** `--concurrency 4` or `-C 4` is passed
- **THEN** the pool SHALL use 4 workers
