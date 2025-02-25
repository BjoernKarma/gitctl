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

func createTestFilesAndFolders() {
	dirPath := filepath.Join(microserviceDirPath, ".git")
	// Set directory permissions to 0700 for secure access
	err := os.MkdirAll(dirPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	// Create a sample file in the directory
	filePath := filepath.Join(dirPath, "service.txt")
	err = os.WriteFile(filePath, []byte("sample content"), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanupTestFilesAndFolders() {
	// Remove the test directory and its contents
	err := os.RemoveAll(testDirPath)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	// Call the setup function to create files and folders
	createTestFilesAndFolders()

	viper.Set("verbose", true)
	// call flag.Parse() here if TestMain uses flags

	// Run the tests
	code := m.Run()

	// Call the cleanup function
	cleanupTestFilesAndFolders()

	// Exit with the code from m.Run()
	os.Exit(code)
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

	// Assert that there was an error and the result is the combined output (standard out/ standard error)
	assert.Error(t, err)
	assert.NotNil(t, output)
}

func TestGitRepoRunGitCommand(t *testing.T) {

	// Call the function under test
	testDir, _ := filepath.Abs(microserviceDirPath)
	gitRepo := GitRepo{path: testDir}

	output, err := gitRepo.RunGitCommand("hello")

	// Assert that there was no error and the result is as expected
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
