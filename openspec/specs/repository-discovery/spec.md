# Spec: Repository Discovery

## Overview

Given a list of base directories, `gitctl` recursively discovers all Git repositories within them. A Git repository is any directory that directly contains a `.git` subdirectory.

## Inputs

- `base_dirs`: a list of absolute, verified directory paths (see [configuration spec](../configuration/spec.md))

## Discovery Algorithm

For each base directory:

1. Walk the directory tree recursively.
2. When a directory named `.git` is encountered, its parent is recorded as a Git repository.
3. The subtree under `.git` is not walked further (no nested repo traversal).
4. If a filesystem error occurs during the walk, an error is printed and the walk for that subtree is aborted.

Discovery across multiple base directories is independent. All found repositories are collected into a single flat list for execution.

## Output

An ordered list of `GitRepo` values, each holding an absolute path to a repository root.

## Edge Cases

| Situation                         | Behaviour                                  |
|-----------------------------------|--------------------------------------------|
| Base dir does not exist           | Silently skipped before discovery begins   |
| Base dir is a git repo itself     | That repo is included                      |
| Nested repos (repo inside repo)   | Only the outermost repo is found           |
| Permission error during walk      | Error printed; affected subtree skipped    |
| Base dir contains no repos        | Empty result; no error                     |

## Verbose Mode

When `verbosity.verbose` is true, each discovered repository path is printed as it is found.

A summary line is always printed after each base directory is scanned:

```
Found <N> git directories in <base_dir>
```
