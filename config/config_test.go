package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHomeDirReturnsHomeEnv(t *testing.T) {
	homeDir := os.TempDir()
	_ = os.Setenv("HOME", homeDir)
	result := HomeDir()
	if result != homeDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestHomeDirHandlesEmptyHomeEnv(t *testing.T) {
	_ = os.Setenv("HOME", "")
	homeDir, _ := os.UserHomeDir()
	tmpDir := os.TempDir() // Alternative in cases where user home dir is not available (e.g. CICD)
	result := HomeDir()
	if result != homeDir && result != tmpDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestGitctlConfigDirReturnsHomeEnv(t *testing.T) {
	homeDir := os.TempDir()
	_ = os.Setenv("HOME", homeDir)
	configDir := filepath.Join(homeDir, ".config", "gitctl")
	result := GitctlConfigDir()
	if result != configDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func DisableTestGitctlConfigDirHandlesEmptyHomeEnv(t *testing.T) {
	_ = os.Setenv("HOME", "")
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".config", "gitctl")
	tmpDir := os.TempDir() // Alternative in cases where user home dir is not available (e.g. CICD)
	configTempDir := filepath.Join(tmpDir, ".config", "gitctl")
	result := GitctlConfigDir()
	if result != configDir && result != configTempDir {
		t.Errorf("expected %v, got %v", homeDir, result)
	}
}

func TestGitctlWorkingDirReturns(t *testing.T) {
	workingDir, _ := os.Getwd()
	result := GitctlWorkingDir()
	if result != workingDir {
		t.Errorf("expected %v, got %v", workingDir, result)
	}
}
