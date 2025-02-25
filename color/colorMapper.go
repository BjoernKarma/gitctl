package color

import (
	"strings"

	"github.com/fatih/color"
)

var (
	mappings = map[string]color.Attribute{
		// git status messages
		"nothing to commit":                                   color.FgGreen,
		"Changes to be committed":                             color.FgYellow,
		"nothing added to commit but untracked files present": color.FgYellow,
		"Changes not staged for commit":                       color.FgRed,
		// git pull messages
		"Already up to date.": color.FgGreen,
		"is up to date.":      color.FgGreen,
		"Fast-forward":        color.FgYellow,
		"cannot pull with rebase: You have unstaged changes": color.FgRed,
	}
)

func MapMessageToColor(message string) color.Attribute {
	for key, printColor := range mappings {
		if strings.Contains(message, key) {
			return printColor
		}
	}
	return color.Reset
}
