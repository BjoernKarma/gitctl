package cmd

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommandShowsHelp(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	rootCmd.SetArgs([]string{"--help"})
	err := rootCmd.Execute()

	expected := "Run git commands on multiple git repositories"
	assert.Contains(t, buf.String(), expected, "expected %v to be contained in %v", expected, buf.String())
	assert.NoError(t, err)
}
