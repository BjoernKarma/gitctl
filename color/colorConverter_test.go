package color

import (
	"strings"
	"testing"

	"github.com/fatih/color"
)

func TestConvertToColoredMessage_GreenHeader(t *testing.T) {
	header := "Success"
	message := "nothing to commit"
	expectedColor := color.FgGreen
	result := ConvertToColoredMessage(header, message)
	if !strings.Contains(result, color.New(expectedColor).SprintFunc()(header)) {
		t.Errorf("expected header to be colored %v, got %v", expectedColor, result)
	}
}

func TestConvertToColoredMessage_YellowMessage(t *testing.T) {
	header := "Warning"
	message := "Changes to be committed"
	expectedColor := color.FgYellow
	result := ConvertToColoredMessage(header, message)
	if !strings.Contains(result, color.New(expectedColor).SprintFunc()(message)) {
		t.Errorf("expected message to be colored %v, got %v", expectedColor, result)
	}
}

func TestConvertToColoredOutput_MultipleLines(t *testing.T) {
	output := "line1\nline2\nline3"
	result := ConvertToColoredOutput(output)
	lines := strings.Split(result, lineBreakChar)
	if len(lines) != 4 { // 3 lines + 1 empty line at the end
		t.Errorf("expected 4 lines, got %d", len(lines))
	}
}

func TestConvertToColoredHeader_SeparatorIncluded(t *testing.T) {
	header := "Header"
	expectedColor := color.FgGreen
	result := ConvertToColoredHeader(header, expectedColor)
	if !strings.Contains(result, strings.Repeat(separatorChar, separatorLength)) {
		t.Errorf("expected separator to be included, got %v", result)
	}
}

func TestConvertToColoredMessage_RedMessage(t *testing.T) {
	header := "Error"
	message := "Changes not staged for commit"
	expectedColor := color.FgRed
	result := ConvertToColoredMessage(header, message)
	if !strings.Contains(result, color.New(expectedColor).SprintFunc()(message)) {
		t.Errorf("expected message to be colored %v, got %v", expectedColor, result)
	}
}

func TestConvertToColoredOutput_EmptyString(t *testing.T) {
	output := ""
	result := ConvertToColoredOutput(output)
	if result != "\n" {
		t.Errorf("expected empty string to result in a single newline, got %v", result)
	}
}
