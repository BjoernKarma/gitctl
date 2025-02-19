package gitrepo

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"
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
	var verbose = viper.GetBool("verbose")
	var repos []GitRepo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error walking the path %q: %v\n", root, err)
			log.Println(err)
			return err
		}
		if info.IsDir() && info.Name() == gitDirToSearch {
			gitDir := filepath.Dir(path)
			if verbose {
				log.Printf("found a git directory: %+v \n", gitDir)
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
	var verbose = viper.GetBool("verbose")
	if gitRepo == nil || gitRepo.path == "" {
		if verbose {
			log.Printf("The repository path is empty. Skipping the git command.\n")
		}
		return nil, nil
	}

	var cmd *exec.Cmd
	switch command {
	case GitPull:
		cmd = exec.Command(gitCommand, pullCommand)
	case GitStatus:
		cmd = exec.Command(gitCommand, statusCommand)
	default:
		cmd = exec.Command(gitCommand, statusCommand)
	}

	cmd.Dir = gitRepo.path
	out, err := cmd.CombinedOutput()
	// Format the output with headers and separators and color
	formattedOutput := FormatOutput(out, gitRepo.path)
	if err != nil {
		return []byte(formattedOutput), err
	} else {
		return []byte(formattedOutput), nil
	}
}

func FormatOutput(output []byte, header string) string {
	coloredOutput, color := ConvertToColoredOutput(string(output))
	return AddHeaderToOutput(coloredOutput, header, color)
}
