package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHomeDirReturnsHomeEnv(t *testing.T) {
	homeDir := os.TempDir()
	t.Setenv("HOME", homeDir)
	result, err := HomeDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != homeDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestHomeDirHandlesEmptyHomeEnv(t *testing.T) {
	t.Setenv("HOME", "")
	homeDir, _ := os.UserHomeDir()
	tmpDir := os.TempDir() // Alternative in cases where user home dir is not available (e.g. CICD)
	result, err := HomeDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != homeDir && result != tmpDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestHomeDirReturnsErrorForInvalidHomeEnv(t *testing.T) {
	notADir := filepath.Join(t.TempDir(), "not-a-dir")
	err := os.WriteFile(notADir, []byte("x"), 0600)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	t.Setenv("HOME", notADir)

	_, homeErr := HomeDir()
	if homeErr == nil {
		t.Fatal("expected error when HOME is not a directory")
	}
}

func TestGitctlConfigDirReturnsHomeEnv(t *testing.T) {
	homeDir := os.TempDir()
	t.Setenv("HOME", homeDir)
	configDir := filepath.Join(homeDir, ".config", "gitctl")
	result, err := GitctlConfigDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != configDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func DisableTestGitctlConfigDirHandlesEmptyHomeEnv(t *testing.T) {
	t.Setenv("HOME", "")
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".config", "gitctl")
	tmpDir := os.TempDir() // Alternative in cases where user home dir is not available (e.g. CICD)
	configTempDir := filepath.Join(tmpDir, ".config", "gitctl")
	result, err := GitctlConfigDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != configDir && result != configTempDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestGitctlWorkingDirReturns(t *testing.T) {
	workingDir, _ := os.Getwd()
	result, err := GitctlWorkingDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != workingDir {
		t.Errorf("expected %v, got %v", workingDir, result)
	}
}
