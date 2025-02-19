package gitrepo

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunGitStatus(t *testing.T) {

	// Mock inputs
	command := GitStatus
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	// Call the function under test
	RunGitCommand(command, baseDirs)

	// Since RunGitCommand doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
}

func TestRunGitDefaultCommand(t *testing.T) {

	// Mock inputs
	command := "hello"
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	// Call the function under test
	RunGitCommand(command, baseDirs)

	// Since RunGitCommand doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
}

func TestRunGitStatusInvalidBaseDirs(t *testing.T) {

	// Mock inputs
	command := GitStatus
	baseDirs := []string{invalidPath}

	// Call the function under test
	RunGitCommand(command, baseDirs)

	// Since RunGitCommand doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
}

func TestRunGitPull(t *testing.T) {

	// Mock inputs
	command := GitPull
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	// Call the function under test
	RunGitCommand(command, baseDirs)

	// Since RunGitCommand doesn't return anything, we can't make assertions about its return value.
	// We could potentially check for side effects (like changes to global state), but without more information, it's hard to say what to check.
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
