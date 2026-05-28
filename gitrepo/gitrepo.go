package gitrepo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bjoernkarma/gitctl/color"
	"github.com/bjoernkarma/gitctl/config"
)

const (
	gitDirToSearch = ".git"
	gitCommand     = "git"
	pullCommand    = "pull"
	statusCommand  = "status"
)

// GitRepo represents a git repository defined with an absolute file path.
type GitRepo struct {
	// The absolute file path for this git repository.
	path string
}

const (
	GitPull   = "pull"
	GitStatus = "status"
)

func FindGitRepos(root string) ([]GitRepo, error) {
	verbose := config.IsVerbose()
	var repos []GitRepo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			color.PrintError(fmt.Sprintf("error walking the path %q: %v\n", root, err))
			log.Println(err)
			return err
		}
		if info.IsDir() && info.Name() == gitDirToSearch {
			gitDir := filepath.Dir(path)
			if verbose {
				color.PrintSubtleInfo(fmt.Sprintf("found a git directory: %+v", gitDir))
			}
			repos = append(repos, GitRepo{path: gitDir})
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return repos, nil
	}
}

// runRaw executes the git command and returns raw combined output without any
// color formatting or global state mutations. Safe to call from goroutines.
func (gitRepo *GitRepo) runRaw(command string) ([]byte, error) {
	verbose := config.IsVerbose()
	dryRun := config.IsDryRun()
	repoPath := ""
	if gitRepo != nil {
		repoPath = gitRepo.path
	}

	if dryRun {
		message := fmt.Sprintf("Dry run enabled. Skipping git %s for repository %s", command, repoPath)
		color.PrintSubtleInfo(message)
		return nil, nil
	}

	if gitRepo == nil || repoPath == "" {
		if verbose {
			color.PrintInfo("The repository path is empty. Skipping the git command.")
		}
		return nil, nil
	}

	var gitCmd *exec.Cmd
	switch command {
	case GitPull:
		gitCmd = exec.Command(gitCommand, pullCommand)
	case GitStatus:
		gitCmd = exec.Command(gitCommand, statusCommand)
	default:
		gitCmd = exec.Command(gitCommand, statusCommand)
	}

	gitCmd.Dir = repoPath
	out, err := gitCmd.CombinedOutput()
	if err != nil {
		return out, fmt.Errorf("git %s failed for %s: %w", command, repoPath, err)
	}
	return out, nil
}

// RunGitCommand executes the git command and returns color-formatted output.
// Not safe to call from concurrent goroutines (mutates global color state).
func (gitRepo *GitRepo) RunGitCommand(command string) ([]byte, error) {
	repoPath := ""
	if gitRepo != nil {
		repoPath = gitRepo.path
	}
	raw, err := gitRepo.runRaw(command)
	if raw == nil && err == nil {
		return nil, nil
	}
	formattedOutput := FormatOutput(repoPath, raw)
	if err != nil {
		return []byte(formattedOutput), err
	}
	return []byte(formattedOutput), nil
}

func FormatOutput(header string, output []byte) string {
	return color.ConvertToColoredMessage(header, string(output))
}
