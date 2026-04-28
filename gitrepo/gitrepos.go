package gitrepo

import (
	"errors"
	"fmt"
	"strings"

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
			errorMsg := extractErrorMessage(string(output))
			color.AddGitCommandFailure(gitRepo.path, errorMsg, string(output))
			// In verbose mode, show the full formatted output immediately
			if isVerbose && !isQuiet {
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
	verbose := config.IsVerbose()
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

// extractErrorMessage pulls the most relevant error message from git output.
// Skips the formatted header (with separators) and looks for explicit error patterns.
func extractErrorMessage(output string) string {
	lines := strings.Split(output, "\n")
	var gitOutputLines []string

	// Skip the formatted header section (path + separator line)
	inHeader := true
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if inHeader {
			// Look for the separator line to know when header ends
			if strings.HasPrefix(trimmed, "=") && len(trimmed) > 20 {
				inHeader = false
				continue
			}
		} else {
			gitOutputLines = append(gitOutputLines, line)
		}
	}

	// First pass: look for explicit error/fatal lines in git output
	for _, line := range gitOutputLines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.Contains(trimmed, "fatal:") || strings.Contains(trimmed, "error:") || strings.Contains(trimmed, "ERROR:") {
			return trimmed
		}
	}

	// Second pass: return first non-empty line from git output
	for _, line := range gitOutputLines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		return trimmed
	}

	return "Unknown error"
}
