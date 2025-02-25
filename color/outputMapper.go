package color

import (
	"strconv"

	"github.com/fatih/color"
)

var (
	gitSuccess []string
	gitInfos   []string
	gitErrors  []string
)

func MapMessageToStatus(text string, messageColor color.Attribute) {
	switch messageColor {
	case color.FgGreen:
		AddGitSuccess(text)
	case color.FgYellow:
		AddGitInfo(text)
	case color.FgRed:
		AddGitError(text)
	default:
		AddGitInfo(text)
	}
}

func AddGitSuccess(info string) {
	gitSuccess = append(gitSuccess, info)
}

func AddGitInfo(info string) {
	gitInfos = append(gitInfos, info)
}

func AddGitError(err string) {
	gitErrors = append(gitErrors, err)
}

func PrintGitRepoStatus() {
	if len(gitSuccess) > 0 {
		PrintSuccess("\n============ Success ============\n")
		for _, i := range gitSuccess {
			PrintSuccess(i)
		}
	}

	if len(gitInfos) > 0 {
		PrintInfo("\n============ Info ============\n")
		for _, i := range gitInfos {
			PrintInfo(i)
		}
	}

	if len(gitErrors) > 0 {
		PrintError("\n============ Issues ============\n")
		for _, e := range gitErrors {
			PrintError(e)
		}
	}
}

func PrintGitStatistics() {
	PrintInfo("\n============ Statistics ============\n")
	PrintSubtleInfo("Success: " + strconv.Itoa(len(gitSuccess)))
	PrintSubtleInfo("Info: " + strconv.Itoa(len(gitInfos)))
	PrintSubtleInfo("Errors: " + strconv.Itoa(len(gitErrors)))
}
