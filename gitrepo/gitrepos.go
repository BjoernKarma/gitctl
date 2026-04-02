package gitrepo

import (
	"errors"
	"fmt"
	"log"

	"github.com/bjoernkarma/gitctl/color"
	"github.com/bjoernkarma/gitctl/config"
)

func RunGitCommand(command string, baseDirs []string) error {
	allGitRepos, findErr := findGitReposInBaseDirs(baseDirs)
	if findErr != nil {
		log.Println(findErr)
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
			log.Println(err)
			commandErrors = append(commandErrors, err)
		}
		if isVerbose && !isQuiet {
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
			log.Println(err)
			findErrors = append(findErrors, fmt.Errorf("failed to find repositories in %s: %w", baseDir, err))
			continue
		}
		color.PrintSuccess(fmt.Sprintf("Found %d git directories in %s \n", len(repos), baseDir))
		allGitRepos = append(allGitRepos, repos...)
	}

	return allGitRepos, errors.Join(findErrors...)
}
