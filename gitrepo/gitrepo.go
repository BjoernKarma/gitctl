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
	var verbose = config.IsVerbose()
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

func (gitRepo *GitRepo) RunGitCommand(command string) ([]byte, error) {
	var verbose = config.IsVerbose()
	var dryRun = config.IsDryRun()
	if dryRun {
		message := fmt.Sprintf("Dry run enabled. Skipping git %s for repository %s", command, gitRepo.path)
		color.PrintSubtleInfo(message)
		return nil, nil
	}

	if gitRepo == nil || gitRepo.path == "" {
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

	gitCmd.Dir = gitRepo.path
	out, _ := gitCmd.CombinedOutput()
	// Format the output with headers and separators and color
	formattedOutput := FormatOutput(gitRepo.path, out)
	return []byte(formattedOutput), nil
}

func FormatOutput(header string, output []byte) string {
	return color.ConvertToColoredMessage(header, string(output))
}
