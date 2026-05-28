package gitrepo

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/bjoernkarma/gitctl/color"
	"github.com/bjoernkarma/gitctl/config"
)

type repoResult struct {
	rawOutput []byte
	err       error
}

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

	results := runWithWorkerPool(command, allGitRepos)

	var commandErrors []error
	if findErr != nil {
		commandErrors = append(commandErrors, findErr)
	}
	for i, result := range results {
		// FormatOutput mutates global color state — must stay in the main goroutine.
		formattedOutput := FormatOutput(allGitRepos[i].path, result.rawOutput)
		if result.err != nil {
			commandErrors = append(commandErrors, result.err)
			errorMsg := extractErrorMessage(formattedOutput)
			color.AddGitCommandFailure(allGitRepos[i].path, errorMsg, formattedOutput)
			if isVerbose && !isQuiet {
				fmt.Printf("%s", formattedOutput)
			}
		} else if isVerbose && !isQuiet {
			fmt.Printf("%s", formattedOutput)
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

// runWithWorkerPool executes the git command across all repos using a bounded
// goroutine pool. Results are stored at each repo's discovery index so that
// the caller can iterate them in deterministic order. Workers call runRaw to
// avoid concurrent mutations of global color state.
func runWithWorkerPool(command string, repos []GitRepo) []repoResult {
	results := make([]repoResult, len(repos))
	if len(repos) == 0 {
		return results
	}

	concurrency := config.GetConcurrency()
	if concurrency < 1 {
		concurrency = 1
	}

	type job struct {
		index int
		repo  GitRepo
	}

	jobs := make(chan job, len(repos))
	var wg sync.WaitGroup

	for range concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				raw, err := j.repo.runRaw(command)
				results[j.index] = repoResult{rawOutput: raw, err: err}
			}
		}()
	}

	for i, repo := range repos {
		jobs <- job{index: i, repo: repo}
	}
	close(jobs)

	wg.Wait()
	return results
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
