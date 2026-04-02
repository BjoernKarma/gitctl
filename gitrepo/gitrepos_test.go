package gitrepo

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/bjoernkarma/gitctl/config"
)

func TestRunGitStatus(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	command := GitStatus
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(command, baseDirs)
	assert.NoError(t, err)
}

func TestRunGitDefaultCommand(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	command := "hello"
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(command, baseDirs)
	assert.NoError(t, err)
}

func TestRunGitStatusInvalidBaseDirs(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)

	command := GitStatus
	baseDirs := []string{invalidPath}

	err := RunGitCommand(command, baseDirs)
	assert.Error(t, err)
}

func TestRunGitPull(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	command := GitPull
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(command, baseDirs)
	assert.NoError(t, err)
}

func TestFindGitReposInBaseDirs(t *testing.T) {
	// Mock inputs
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	// Call the function under test
	repos, err := findGitReposInBaseDirs(baseDirs)

	// Assert that there was no error and the result is as expected
	assert.NoError(t, err)
	// Without more information, it's hard to say what the expected result is.
	// Here's an example where we just check that the result is not nil.
	assert.NotNil(t, repos)
}

func TestFindGitReposInvalidBaseDirs(t *testing.T) {
	// Mock inputs
	baseDirs := []string{invalidPath}

	// Call the function under test
	repos, err := findGitReposInBaseDirs(baseDirs)

	// Assert that there was an error and the result is nil
	assert.Error(t, err)
	assert.Nil(t, repos)
}
