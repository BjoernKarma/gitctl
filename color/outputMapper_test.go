package color

import (
	"bytes"
	"log"
	"testing"

	"github.com/fatih/color"
)

func TestMapMessageToStatus_Success(t *testing.T) {
	gitSuccess = []string{}
	text := "Success message"
	MapMessageToStatus(text, color.FgGreen)
	if len(gitSuccess) != 1 || gitSuccess[0] != text {
		t.Errorf("expected success message to be added, got %v", gitSuccess)
	}
}

func TestMapMessageToStatus_Info(t *testing.T) {
	gitInfos = []string{}
	text := "Info message"
	MapMessageToStatus(text, color.FgYellow)
	if len(gitInfos) != 1 || gitInfos[0] != text {
		t.Errorf("expected info message to be added, got %v", gitInfos)
	}
}

func TestMapMessageToStatus_Error(t *testing.T) {
	gitErrors = []string{}
	text := "Error message"
	MapMessageToStatus(text, color.FgRed)
	if len(gitErrors) != 1 || gitErrors[0] != text {
		t.Errorf("expected error message to be added, got %v", gitErrors)
	}
}

func TestMapMessageToStatus_Default(t *testing.T) {
	gitInfos = []string{}
	text := "Default message"
	MapMessageToStatus(text, color.FgBlue)
	if len(gitInfos) != 1 || gitInfos[0] != text {
		t.Errorf("expected default info message to be added, got %v", gitInfos)
	}
}

func TestPrintGitRepoStatus_SuccessMessages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{"Success1", "Success2"}
	gitInfos = []string{}
	gitErrors = []string{}

	PrintGitRepoStatus()

	if !bytes.Contains(buf.Bytes(), []byte("Success1")) || !bytes.Contains(buf.Bytes(), []byte("Success2")) {
		t.Errorf("expected success messages to be printed, got %v", buf.String())
	}
}

func TestPrintGitRepoStatus_InfoMessages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{}
	gitInfos = []string{"Info1", "Info2"}
	gitErrors = []string{}

	PrintGitRepoStatus()

	if !bytes.Contains(buf.Bytes(), []byte("Info1")) || !bytes.Contains(buf.Bytes(), []byte("Info2")) {
		t.Errorf("expected info messages to be printed, got %v", buf.String())
	}
}

func TestPrintGitRepoStatus_ErrorMessages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{}
	gitInfos = []string{}
	gitErrors = []string{"Error1", "Error2"}

	PrintGitRepoStatus()

	if !bytes.Contains(buf.Bytes(), []byte("Error1")) || !bytes.Contains(buf.Bytes(), []byte("Error2")) {
		t.Errorf("expected error messages to be printed, got %v", buf.String())
	}
}

func TestPrintGitRepoStatus_MixedMessages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{"Success1"}
	gitInfos = []string{"Info1"}
	gitErrors = []string{"Error1"}

	PrintGitRepoStatus()

	if !bytes.Contains(buf.Bytes(), []byte("Success1")) || !bytes.Contains(buf.Bytes(), []byte("Info1")) || !bytes.Contains(buf.Bytes(), []byte("Error1")) {
		t.Errorf("expected mixed messages to be printed, got %v", buf.String())
	}
}

func TestPrintGitRepoStatus_NoMessages(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{}
	gitInfos = []string{}
	gitErrors = []string{}

	PrintGitRepoStatus()

	if buf.String() != "" {
		t.Errorf("expected no messages to be printed, got %v", buf.String())
	}
}

func TestPrintGitStatistics(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	gitSuccess = []string{"Success1", "Success2"}
	gitInfos = []string{"Info1"}
	gitErrors = []string{"Error1", "Error2", "Error3"}

	PrintGitStatistics()

	if !bytes.Contains(buf.Bytes(), []byte("Success: 2")) || !bytes.Contains(buf.Bytes(), []byte("Info: 1")) || !bytes.Contains(buf.Bytes(), []byte("Errors: 3")) {
		t.Errorf("expected mixed messages to be printed, got %v", buf.String())
	}
}
