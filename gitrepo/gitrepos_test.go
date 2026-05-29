package gitrepo

import (
	"path/filepath"
	"strings"
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
	assert.NotNil(t, repos)
	assert.Len(t, repos, 1)
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

func TestRunGitCommandAggregatesErrorsFromInvalidAndValidBaseDirs(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{invalidPath, testDir}

	err := RunGitCommand(GitStatus, baseDirs)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "failed to find repositories"))
}

func TestRunGitCommandWithConcurrencyGreaterThanOneProcessesAllRepos(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)
	viper.Set(config.GitCtlConcurrency, 3)

	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(GitStatus, baseDirs)
	assert.NoError(t, err)
}

func TestRunWithWorkerPoolClampsNegativeConcurrencyToOne(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)
	viper.Set(config.GitCtlConcurrency, -1)

	testDir, _ := filepath.Abs(testDirPath)
	repos, err := findGitReposInBaseDirs([]string{testDir})
	assert.NoError(t, err)

	results := runWithWorkerPool(GitStatus, repos)
	assert.Len(t, results, len(repos))
}

func TestRunWithWorkerPoolClampsZeroConcurrencyToOne(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)
	viper.Set(config.GitCtlConcurrency, 0)

	testDir, _ := filepath.Abs(testDirPath)
	repos, err := findGitReposInBaseDirs([]string{testDir})
	assert.NoError(t, err)

	results := runWithWorkerPool(GitStatus, repos)
	assert.Len(t, results, len(repos))
}

func TestRunWithWorkerPoolPreservesDiscoveryOrder(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlConcurrency, 3)

	testDir, _ := filepath.Abs(microserviceDirPath)
	// Mix valid and invalid repos to confirm results are indexed by discovery order.
	repos := []GitRepo{
		{path: testDir},     // index 0: valid — no error expected
		{path: invalidPath}, // index 1: invalid — error expected
		{path: testDir},     // index 2: valid — no error expected
	}

	results := runWithWorkerPool(GitStatus, repos)

	assert.Len(t, results, 3)
	assert.NoError(t, results[0].err)
	assert.Error(t, results[1].err)
	assert.NoError(t, results[2].err)
}

func TestRunGitBranchCommand(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	command := GitBranch
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(command, baseDirs)
	assert.NoError(t, err)
}

func TestRunGitFetchCommand(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	viper.Set(config.GitCtlDryRun, true)

	command := GitFetch
	testDir, _ := filepath.Abs(testDirPath)
	baseDirs := []string{testDir}

	err := RunGitCommand(command, baseDirs)
	assert.NoError(t, err)
}
