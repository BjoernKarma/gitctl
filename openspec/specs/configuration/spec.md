# Spec: Configuration

## Overview

`gitctl` resolves its configuration from three sources in a defined precedence order. All settings are available as CLI flags, environment variables, and config file keys.

## Config File

`gitctl` searches for `gitctl.yaml` in order:

1. Explicit path via `--config` flag (skips auto-search)
2. Current working directory (`./gitctl.yaml`)
3. `~/.config/gitctl/gitctl.yaml`

The first file found is used. If no file is found, built-in defaults apply and a log message is emitted.

### YAML structure

```yaml
verbosity:
  quiet: false       # suppress all output
  verbose: false     # show per-repo git output
  debug: false       # log config and internal state

run_mode:
  local: false       # use CWD as the only base directory
  dry_run: false     # skip git execution, print what would run
  concurrency: 1     # number of concurrent operations (not yet implemented)

output:
  color: true        # colorize terminal output

base_dirs:
  - "/path/to/projects"
```

## Environment Variables

All config keys are readable from environment variables using:
- Prefix: `GITCTL_`
- Key mapping: dots replaced with underscores, uppercased

| Config key              | Environment variable          |
|-------------------------|-------------------------------|
| `verbosity.quiet`       | `GITCTL_VERBOSITY_QUIET`      |
| `verbosity.verbose`     | `GITCTL_VERBOSITY_VERBOSE`    |
| `verbosity.debug`       | `GITCTL_VERBOSITY_DEBUG`      |
| `run_mode.local`        | `GITCTL_RUN_MODE_LOCAL`       |
| `run_mode.dry_run`      | `GITCTL_RUN_MODE_DRY_RUN`     |
| `run_mode.concurrency`  | `GITCTL_RUN_MODE_CONCURRENCY` |
| `output.color`          | `GITCTL_OUTPUT_COLOR`         |

## CLI Flags

All flags are persistent (available to all subcommands):

| Flag              | Short | Default | Description                            |
|-------------------|-------|---------|----------------------------------------|
| `--config`        |       | —       | Explicit config file path              |
| `--quiet`         | `-q`  | false   | Suppress output                        |
| `--verbose`       | `-v`  | false   | Show per-repo git output               |
| `--debug`         | `-d`  | false   | Show debug/config logging              |
| `--local`         | `-l`  | false   | Use CWD as sole base directory         |
| `--dryRun`        | `-D`  | false   | Skip execution, print what would run   |
| `--color`         | `-c`  | true    | Colorize output                        |
| `--concurrency`   | `-C`  | 1       | Number of concurrent operations        |
| `--base.dirs`     |       | []      | Base directories (comma-separated)     |

## Precedence

Highest to lowest:

```
CLI flags  >  Environment variables  >  Config file  >  Built-in defaults
```

## Base Directory Resolution

- Paths in `base_dirs` are resolved to absolute paths.
- Non-existent or invalid paths are silently skipped.
- When `--local` is set, `base_dirs` is ignored and the current working directory is used.
