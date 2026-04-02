package cmd

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/bjoernkarma/gitctl/config"
)

func TestRootCommandShowsHelp(t *testing.T) {
	var buf bytes.Buffer
	viper.Reset()
	originalLogWriter := log.Writer()
	log.SetOutput(&buf)
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	defer func() {
		log.SetOutput(originalLogWriter)
		configFile = ""
		viper.Reset()
	}()

	rootCmd.SetArgs([]string{"--help"})
	err := rootCmd.Execute()

	expected := "Run git commands on multiple git repositories"
	assert.Contains(t, buf.String(), expected, "expected %v to be contained in %v", expected, buf.String())
	assert.NoError(t, err)
}

func TestCommandReturnsErrorForInvalidConfigFile(t *testing.T) {
	viper.Reset()
	tmpDir := t.TempDir()
	invalidConfig := filepath.Join(tmpDir, "gitctl.yaml")
	err := os.WriteFile(invalidConfig, []byte("verbosity: ["), 0600)
	assert.NoError(t, err)

	rootCmd.SetArgs([]string{"status", "--config", invalidConfig, "--local", "--dryRun"})
	err = rootCmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read configuration file")

	configFile = ""
	viper.Reset()
}

func TestInitConfigReadsUnderscoreEnvVars(t *testing.T) {
	viper.Reset()
	configFile = ""
	t.Setenv("GITCTL_RUN_MODE_LOCAL", "true")

	err := InitConfig()

	assert.NoError(t, err)
	assert.True(t, config.IsLocal())

	configFile = ""
	viper.Reset()
}

func TestInitConfigReadsPrefixedVerbosityEnvVars(t *testing.T) {
	viper.Reset()
	configFile = ""
	t.Setenv("GITCTL_VERBOSITY_VERBOSE", "true")

	err := InitConfig()

	assert.NoError(t, err)
	assert.True(t, config.IsVerbose())

	configFile = ""
	viper.Reset()
}

