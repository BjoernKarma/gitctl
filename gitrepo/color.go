package gitrepo

import (
	"strings"
)

const (
	// Color
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"

	// Separator
	separatorLength = 80
	separatorChar   = "="
	lineBreakChar   = "\n"
)

var (
	mappings = map[string]string{
		// git status messages
		"nothing to commit":                                   colorGreen,
		"Changes to be committed":                             colorYellow,
		"nothing added to commit but untracked files present": colorYellow,
		"Changes not staged for commit":                       colorRed,
		// git pull messages
		"Already up to date.": colorGreen,
		"is up to date.":      colorGreen,
		"Fast-forward":        colorYellow,
		"cannot pull with rebase: You have unstaged changes": colorRed,
	}
)

func GetColor(message string) string {
	for key, color := range mappings {
		if strings.Contains(message, key) {
			return color
		}
	}
	return colorReset
}

func ConvertToColoredOutput(output string) (string, string) {
	var coloredOutput string
	var lastMappedNonDefaultColor = colorReset
	lines := strings.Split(output, lineBreakChar)
	for _, line := range lines {
		color := GetColor(line)
		coloredOutput += color + line + colorReset + lineBreakChar
		if color != colorReset {
			lastMappedNonDefaultColor = color
		}
	}
	return coloredOutput, lastMappedNonDefaultColor
}

func AddHeaderToOutput(output string, header string, color string) string {
	return lineBreakChar + color + header + colorReset + lineBreakChar +
		strings.Repeat(separatorChar, separatorLength) + lineBreakChar + output
}
