# Spec: Output

## Overview

`gitctl` formats all terminal output consistently across commands. Output is color-coded by default and structured around per-repo results followed by a summary.

## Output Modes

| Mode      | Flag  | Behaviour                                                    |
|-----------|-------|--------------------------------------------------------------|
| Default   | —     | Show summary and statistics only                             |
| Verbose   | `-v`  | Show per-repo git output in addition to summary              |
| Quiet     | `-q`  | Suppress all output (exit code still reflects success/fail)  |
| Debug     | `-d`  | Emit internal log messages (config values, file paths, etc.) |

Quiet takes precedence over verbose when both are set.

## Per-Repo Output (verbose mode)

Each repository's output is prefixed with a header showing the repository path, followed by a separator, then the raw git output:

```
/path/to/repo
========================================
<git output>
```

## Summary Output

After all repositories are processed, two blocks are printed:

### Statistics block

Shows counts of repositories found, succeeded, and failed.

### Failure list

If any repository failed, a list is printed showing:
- Repository path
- Short error message (extracted from git output)

Example:
```
✗ /path/to/repo — fatal: not a git repository
```

## Color

When `output.color` is true (default), output uses ANSI colors:
- Repository headers: neutral/subtle
- Success indicators: green
- Failure indicators: red
- Info/verbose messages: cyan or subtle

When color is false (e.g. piped output or `--color=false`), plain text is used.

## Exit Codes

| Condition                   | Exit code |
|-----------------------------|-----------|
| All commands succeeded      | 0         |
| One or more commands failed | 1         |
| Internal/config error       | 1         |
