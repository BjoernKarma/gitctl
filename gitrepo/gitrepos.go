package gitrepo

import (
	"errors"
	"fmt"

	"github.com/bjoernkarma/gitctl/color"
	"github.com/bjoernkarma/gitctl/config"
)

func RunGitCommand(command string, baseDirs []string) error {
	allGitRepos, findErr := findGitReposInBaseDirs(baseDirs)
	if findErr != nil {
		color.PrintError(fmt.Sprintf("Error finding repositories: %v", findErr))
	}

	isVerbose := config.IsVerbose()
	isQuiet := config.IsQuiet()

	if isVerbose && !isQuiet {
		fmt.Printf("\n============ GIT OUTPUT (VERBOSE) ============\n")
	}
	var commandErrors []error
	if findErr != nil {
		commandErrors = append(commandErrors, findErr)
	}
	for _, gitRepo := range allGitRepos {
		output, err := gitRepo.RunGitCommand(command)
		if err != nil {
			commandErrors = append(commandErrors, err)
			// Always display the formatted git output on failure so the user
			// can see exactly what went wrong, regardless of verbose mode.
			if !isQuiet {
				fmt.Printf("%s", output)
			}
		} else if isVerbose && !isQuiet {
			fmt.Printf("%s", output)
		}

	}
	if isVerbose && !isQuiet {
		fmt.Printf("\n============ GIT OUTPUT END ============\n")
	}

	// Print statistics and git output
	color.PrintGitStatistics()
	color.PrintGitRepoStatus()

	return errors.Join(commandErrors...)
}

func findGitReposInBaseDirs(baseDirs []string) ([]GitRepo, error) {
	var allGitRepos []GitRepo
	var verbose = config.IsVerbose()
	var findErrors []error

	for _, baseDir := range baseDirs {
		if verbose {
			color.PrintInfo(fmt.Sprintf("Searching for git repositories in : %s", baseDir))
		}

		repos, err := FindGitRepos(baseDir)
		if err != nil {
			findErrors = append(findErrors, fmt.Errorf("failed to find repositories in %s: %w", baseDir, err))
			continue
		}
		color.PrintSuccess(fmt.Sprintf("Found %d git directories in %s \n", len(repos), baseDir))
		allGitRepos = append(allGitRepos, repos...)
	}

	return allGitRepos, errors.Join(findErrors...)
}
