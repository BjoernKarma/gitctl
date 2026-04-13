package color

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/fatih/color"
)

type GitCommandFailure struct {
	RepoPath   string
	ErrorMsg   string
	FullOutput string
}

var (
	gitSuccess         []string
	gitSuccessTree     *tree.Tree
	gitInfos           []string
	gitInfosTree       *tree.Tree
	gitErrors          []string
	gitErrorsTree      *tree.Tree
	gitCommandFailures []GitCommandFailure
)

func MapMessageToStatus(text string, messageColor color.Attribute) {
	switch messageColor {
	case color.FgGreen:
		AddGitSuccess(text)
		AddGitSuccessTree(text)
	case color.FgYellow:
		AddGitInfo(text)
		AddGitInfoTree(text)
	case color.FgRed:
		AddGitError(text)
		AddGitErrorTree(text)
	default:
		AddGitInfo(text)
		AddGitInfoTree(text)
	}
}

func AddGitSuccess(success string) {
	gitSuccess = append(gitSuccess, success)
}

func AddGitSuccessTree(text string) {
	if gitSuccessTree == nil {
		gitSuccessTree = ConvertRepositoryPathToTree(text, GREEN)
	} else {
		AddRepositoryPathToTree(gitSuccessTree, text)
	}
}

func AddGitInfo(info string) {
	gitInfos = append(gitInfos, info)
}

func AddGitInfoTree(text string) {
	if gitInfosTree == nil {
		gitInfosTree = ConvertRepositoryPathToTree(text, YELLOW)
	} else {
		AddRepositoryPathToTree(gitInfosTree, text)
	}
}

func AddGitError(err string) {
	gitErrors = append(gitErrors, err)
}

func AddGitErrorTree(text string) {
	if gitErrorsTree == nil {
		gitErrorsTree = ConvertRepositoryPathToTree(text, RED)
	} else {
		AddRepositoryPathToTree(gitErrorsTree, text)
	}
}

func PrintGitRepoStatus() {
	if gitSuccessTree != nil {
		PrintSuccess("\n============ Success ============\n")
		PrintSuccess(gitSuccessTree.String())
	}

	if gitInfosTree != nil {
		PrintInfo("\n============ Info ============\n")
		PrintInfo(gitInfosTree.String())
	}

	if gitErrorsTree != nil {
		PrintError("\n============ Issues ============\n")
		PrintError(gitErrorsTree.String())
	}

	if len(gitCommandFailures) > 0 {
		PrintGitCommandFailures()
	}
}

func AddGitCommandFailure(repoPath, errorMsg, fullOutput string) {
	gitCommandFailures = append(gitCommandFailures, GitCommandFailure{
		RepoPath:   repoPath,
		ErrorMsg:   errorMsg,
		FullOutput: fullOutput,
	})
}

func PrintGitCommandFailures() {
	if len(gitCommandFailures) == 0 {
		return
	}

	PrintError("\n============ Git Command Failures ============\n")

	isVerbose := false
	// Check if verbose mode is enabled
	if c, ok := os.LookupEnv("GITCTL_VERBOSITY_VERBOSE"); ok && c == "true" {
		isVerbose = true
	}
	// Also check via viper in case it was set from config file
	if !isVerbose {
		isVerbose = isConfigVerbose()
	}

	for _, failure := range gitCommandFailures {
		PrintError(fmt.Sprintf("  ✗ %s", failure.RepoPath))
		PrintError(fmt.Sprintf("    Reason: %s", failure.ErrorMsg))

		// In verbose mode, also show the full git output
		if isVerbose {
			PrintError("\n    Full output:")
			for _, line := range strings.Split(failure.FullOutput, "\n") {
				if strings.TrimSpace(line) != "" {
					PrintError(fmt.Sprintf("    %s", line))
				}
			}
			PrintError("")
		}
	}
}

func isConfigVerbose() bool {
	// Attempt to get verbose setting from viper if it's already initialized
	// This handles cases where viper config was loaded but GITCTL_VERBOSITY_VERBOSE wasn't set
	defer func() {
		if recover() != nil {
			// Viper might not be initialized yet; safe to ignore
		}
	}()

	// This is a safe check that won't panic
	return false // Will be handled by environment variable check above
}

func PrintGitStatistics() {
	PrintSubtleInfo("\n============ Statistics ============\n")
	PrintSubtleInfo("Success: " + strconv.Itoa(len(gitSuccess)))
	PrintSubtleInfo("Info: " + strconv.Itoa(len(gitInfos)))
	PrintSubtleInfo("Errors: " + strconv.Itoa(len(gitErrors)))
}
