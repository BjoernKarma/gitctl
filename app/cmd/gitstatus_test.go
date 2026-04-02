package cmd

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCommandExecutesGitStatusOnLocalRepo(t *testing.T) {
	var buf bytes.Buffer
	originalLogWriter := log.Writer()
	log.SetOutput(&buf)
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	defer func() {
		log.SetOutput(originalLogWriter)
	}()

	rootCmd.SetArgs([]string{"status", "--local", "--debug", "--verbose"})
	err := rootCmd.Execute()

	expected := "Configuration settings:"
	assert.Contains(t, buf.String(), expected, "expected %v to be contained in %v", expected, buf.String())
	assert.NoError(t, err)
}
