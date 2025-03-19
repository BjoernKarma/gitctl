package color

import (
	"strconv"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/fatih/color"
)

var (
	gitSuccess     []string
	gitSuccessTree *tree.Tree
	gitInfos       []string
	gitInfosTree   *tree.Tree
	gitErrors      []string
	gitErrorsTree  *tree.Tree
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
}

func PrintGitStatistics() {
	PrintSubtleInfo("\n============ Statistics ============\n")
	PrintSubtleInfo("Success: " + strconv.Itoa(len(gitSuccess)))
	PrintSubtleInfo("Info: " + strconv.Itoa(len(gitInfos)))
	PrintSubtleInfo("Errors: " + strconv.Itoa(len(gitErrors)))
}
