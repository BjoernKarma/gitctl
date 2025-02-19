package gitrepo

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func RunGitCommand(command string, baseDirs []string) {
	allGitRepos, err := findGitReposInBaseDirs(baseDirs)
	if err != nil {
		log.Println(err)
	}

	for _, gitRepo := range allGitRepos {
		output, err := gitRepo.RunGitCommand(command)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", output)
	}
}

func findGitReposInBaseDirs(baseDirs []string) ([]GitRepo, error) {
	var allGitRepos []GitRepo
	var verbose = viper.GetBool("verbose")

	for _, baseDir := range baseDirs {
		if verbose {
			log.Println("Searching for git repositories in : ", baseDir)
		}

		repos, err := FindGitRepos(baseDir)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Printf("Found %d git directories in %s \n", len(repos), baseDir)
		if verbose {
			log.Println("Searching for git repositories in : ", baseDir)
			for _, repo := range repos {
				fmt.Printf("  %s\n", repo.path+",\n  ")
			}
		}
		allGitRepos = append(allGitRepos, repos...)
	}

	return allGitRepos, nil
}
