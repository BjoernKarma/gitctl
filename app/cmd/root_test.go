package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	// Call the function under test
	err := Execute()

	// Assert that there was no error
	assert.NoError(t, err)
}

func TestInitConfigWithCfgFile(t *testing.T) {
	// Call the function under test
	cfgFile = "../../gitctl.yaml"
	initConfig()

	// Since initConfig doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
}

func TestInitConfig(t *testing.T) {
	// Call the function under test
	cfgFile = ""
	initConfig()

	// Since initConfig doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
}
