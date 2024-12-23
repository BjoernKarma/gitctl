package gitrepo

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testDirPath         = "../test"
	microserviceDirPath = testDirPath + "/microservice"
	invalidPath         = "/path/does/not/exist"
)

func TestMain(m *testing.M) {
	dirPath := "../test/microservice/.git"
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("verbose", true)
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestFindGitRepos(t *testing.T) {

	// Call the function under test
	testDir, _ := filepath.Abs(testDirPath)
	repos, err := FindGitRepos(testDir)

	// Assert that there was no error and the result is as expected
	expectedPath, _ := filepath.Abs(microserviceDirPath)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(repos))
	assert.Equal(t, expectedPath, repos[0].path)
}

func TestFindGitReposInvalidPath(t *testing.T) {

	// Call the function under test
	repos, err := FindGitRepos(invalidPath)

	// Assert that there was an error and the result is nil
	assert.Error(t, err)
	assert.Nil(t, repos)
}

func TestGitRepoRunGitStatus(t *testing.T) {

	// Call the function under test
	testDir, _ := filepath.Abs(microserviceDirPath)
	gitRepo := GitRepo{path: testDir}

	output, err := gitRepo.RunGitCommand(GitStatus)

	// Assert that there was no error and the result is as expected
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestGitRepoEmptyRunGitStatus(t *testing.T) {

	// Call the function under test
	gitRepo := GitRepo{path: ""}

	output, err := gitRepo.RunGitCommand(GitStatus)

	// Assert that there was no error and the result is nil
	assert.NoError(t, err)
	assert.Nil(t, output)
}

func TestGitRepoRunGitPull(t *testing.T) {

	// Call the function under test
	testDir, _ := filepath.Abs(microserviceDirPath)
	gitRepo := GitRepo{path: testDir}

	output, err := gitRepo.RunGitCommand(GitPull)

	// Assert that there was an error and the result is nil (since the git repo is only simulated)
	// TODO - Fix this test
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestGitRepoRunGitCommand(t *testing.T) {

	// Call the function under test
	testDir, _ := filepath.Abs(microserviceDirPath)
	gitRepo := GitRepo{path: testDir}

	output, err := gitRepo.RunGitCommand(-1)

	// Assert that there was no error and the result is as expected
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
