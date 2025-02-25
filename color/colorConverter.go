package color

import (
	"strings"

	"github.com/fatih/color"
)

const (
	// Separator
	separatorLength = 80
	separatorChar   = "="
	lineBreakChar   = "\n"
)

func ConvertToColoredMessage(header, message string) string {
	messageColor := MapMessageToColor(message)
	MapMessageToStatus(header, messageColor)
	colorizedMessage := ConvertToColoredHeader(header, messageColor)
	colorizedMessage += ConvertToColoredOutput(message)
	return colorizedMessage
}

func ConvertToColoredOutput(output string) string {
	var coloredOutput string
	lines := strings.Split(output, lineBreakChar)
	for _, line := range lines {
		// We need to map again for each line, so that each line has a dedicated color
		messageColor := MapMessageToColor(line)
		printColor := color.New(messageColor)
		coloredOutput += printColor.SprintFunc()(line) + lineBreakChar
	}
	return coloredOutput
}

func ConvertToColoredHeader(header string, headerColor color.Attribute) string {
	printColor := color.New(headerColor)
	coloredHeader := printColor.SprintFunc()(header) + lineBreakChar
	separator := strings.Repeat(separatorChar, separatorLength) + lineBreakChar
	return coloredHeader + separator
}
