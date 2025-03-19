package gitrepo

import (
	"fmt"
	"log"

	"github.com/bjoernkarma/gitctl/color"
	"github.com/bjoernkarma/gitctl/config"
)

func RunGitCommand(command string, baseDirs []string) {
	allGitRepos, err := findGitReposInBaseDirs(baseDirs)
	if err != nil {
		log.Println(err)
	}

	isVerbose := config.IsVerbose()
	isQuiet := config.IsQuiet()

	if isVerbose && !isQuiet {
		fmt.Printf("\n============ GIT OUTPUT (VERBOSE) ============\n")
	}
	for _, gitRepo := range allGitRepos {
		output, err := gitRepo.RunGitCommand(command)
		if err != nil {
			log.Println(err)
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
}

func findGitReposInBaseDirs(baseDirs []string) ([]GitRepo, error) {
	var allGitRepos []GitRepo
	var verbose = config.IsVerbose()

	for _, baseDir := range baseDirs {
		if verbose {
			color.PrintInfo(fmt.Sprintf("Searching for git repositories in : %s", baseDir))
		}

		repos, err := FindGitRepos(baseDir)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		color.PrintSuccess(fmt.Sprintf("Found %d git directories in %s \n", len(repos), baseDir))
		allGitRepos = append(allGitRepos, repos...)
	}

	return allGitRepos, nil
}
