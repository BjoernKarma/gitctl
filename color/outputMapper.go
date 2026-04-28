package color

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/fatih/color"

	"github.com/bjoernkarma/gitctl/config"
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

	for _, failure := range gitCommandFailures {
		PrintError(fmt.Sprintf("  ✗ %s", failure.RepoPath))
		PrintError(fmt.Sprintf("    Reason: %s", failure.ErrorMsg))

		if config.IsVerbose() {
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

func PrintGitStatistics() {
	PrintSubtleInfo("\n============ Statistics ============\n")
	PrintSubtleInfo("Success: " + strconv.Itoa(len(gitSuccess)))
	PrintSubtleInfo("Info: " + strconv.Itoa(len(gitInfos)))
	PrintSubtleInfo("Errors: " + strconv.Itoa(len(gitErrors)))
}
