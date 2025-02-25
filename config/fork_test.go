package config

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// Run a fork test that may crash using os.exit.
func RunForkTest(testName string) (string, string, error) {
	// G204: Subprocess launched with a potential tainted input or cmd arguments (gosec)
	//nolint:gosec
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%v", testName))
	cmd.Env = append(os.Environ(), "FORK=1")

	var stdoutB, stderrB bytes.Buffer
	cmd.Stdout = &stdoutB
	cmd.Stderr = &stderrB

	err := cmd.Run()

	return stdoutB.String(), stderrB.String(), err
}

func TestHomeDirHandlesInvalidHomeDir(t *testing.T) {
	if os.Getenv("FORK") == "1" {
		_ = os.Setenv("HOME", "/invalid/dir")
		HomeDir()
		return
	}

	stdout, stderr, err := RunForkTest("TestHomeDirHandlesInvalidHomeDir")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "exit status 1")
	assert.Contains(t, stderr, "The path /invalid/dir is not a valid directory")
	assert.Contains(t, stdout, "")

	// Verify ExitCode
	var e *exec.ExitError
	if errors.As(err, &e) && !e.Success() {
		if e.ExitCode() != 1 {
			t.Fatalf("process ran with err %v, want exit status 1", err)
		}
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
