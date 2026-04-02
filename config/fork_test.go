package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeDirHandlesInvalidHomeDir(t *testing.T) {
	t.Setenv("HOME", "/invalid/dir")

	_, err := HomeDir()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to stat home directory")
}
