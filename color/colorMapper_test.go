package color

import (
	"testing"

	"github.com/fatih/color"
)

func TestMapMessageToColor_NothingToCommit(t *testing.T) {
	message := "nothing to commit"
	expected := color.FgGreen
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_ChangesToBeCommitted(t *testing.T) {
	message := "Changes to be committed"
	expected := color.FgYellow
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_UntrackedFilesPresent(t *testing.T) {
	message := "nothing added to commit but untracked files present"
	expected := color.FgYellow
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_ChangesNotStagedForCommit(t *testing.T) {
	message := "Changes not staged for commit"
	expected := color.FgRed
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_AlreadyUpToDate(t *testing.T) {
	message := "Already up to date."
	expected := color.FgGreen
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_IsUpToDate(t *testing.T) {
	message := "is up to date."
	expected := color.FgGreen
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_FastForward(t *testing.T) {
	message := "Fast-forward"
	expected := color.FgYellow
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_UnstagedChanges(t *testing.T) {
	message := "cannot pull with rebase: You have unstaged changes"
	expected := color.FgRed
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMapMessageToColor_NoMatch(t *testing.T) {
	message := "some random message"
	expected := color.Reset
	result := MapMessageToColor(message)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
