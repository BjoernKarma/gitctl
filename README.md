# gitctl

## Description

`gitctl` is a command-line interface (CLI) tool designed to simplify and streamline your GitLab workflow. It is built with Go and is designed to be easy to use and highly efficient.

## Features

- Base Path Config: Define base paths for all your GitLab projects (`gitctl.yaml`)
- Search Git Repositories: `gitctl` will search for Git repositories in the base paths you've defined
- Run Git Commands: `gitctl` will run Git commands on the repositories found in the base paths

Currently, `gitctl` supports the following Git commands:
- `git status`
- `git pull`

## Installation

### Go install

`go install github.com/bjoernkarma/gitctl@latest`

### Download and install manually

Download the latest release for your platform from the [Releases page](https://gitlab.com/ethical-developer/cli/gitctl/-/packages/), 
then extract and move the `gitctl` binary to any place in your `$PATH`.

### Download and install from source
To install `gitctl`, follow these steps:

1. Clone the repository: `git clone git@github.com:BjoernKarma/gitctl.git`
2. Navigate to the cloned directory: `cd gitctl`
3. Build the project: `go build`
4. Install the CLI: `go install`

## Configuration

Add a `gitctl.yaml`file the `.config\gitctl` folder in your home directory (`~/.config\gitctl\gitctl.yaml`) with the following format:

```yaml
# Verbosity settings
verbosity:
  quiet: false
  verbose: false
  debug: true

# Run mode settings
run_mode:
  local: true
  dry_run: false
  concurrency: 3

# Output settings
output:
  color: true

# Base directories for git repositories
base_dirs:
  - "//dev//gitctl"
```

## Usage

Here's how you can use `gitctl`:

```bash
gitctl [command] [arguments]
```

For more information about the commands, use:

```bash
gitctl --help
```

```text
Run git commands on multiple git repositories. 
For example, you can run 'gitctl pull' to pull all the git 
repositories in the base directories.

Usage:
  gitctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  pull        Execute git pull on multiple git repositories.
  status      Execute git status on multiple git repositories.

Flags:
      --config string   config file (default is $HOME/gitctl.yaml)
  -h, --help            help for gitctl
  -v, --verbose         verbose output
      --version         version for gitctl
      --viper           use Viper for configuration (default true)

Use "gitctl [command] --help" for more information about a command.
```

## Contributing
We welcome contributions to `gitctl`. If you'd like to contribute, please submit a merge request with your changes.  

## Support
If you encounter any problems or have any questions, please open an issue on our GitLab page.  

## Project Status
`gitctl` is currently under active development. We're always looking for feedback and suggestions for new features.