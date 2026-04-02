package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GitctlWorkingDir returns the current gitctl working directory path
func GitctlWorkingDir() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error trying to find working directory: %w", err)
	}
	return workingDir, nil
}

// GitctlConfigDir returns the gitctl config directory path
func GitctlConfigDir() (string, error) {
	homeDir, err := HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", "gitctl"), nil
}

// HomeDir finds the users home directory
func HomeDir() (string, error) {
	// Find home directory.
	var home string
	home = os.Getenv("HOME")
	if home == "" {
		var err error
		home, err = os.UserHomeDir()
		if err != nil {
			home = os.TempDir()
			log.Printf("Error trying to find users home directory due to %s", err)
			log.Println("Using temporary directory as home directory instead")
		}
	} else {
		info, err := os.Stat(home)
		if err != nil {
			return "", fmt.Errorf("failed to stat home directory %s: %w", home, err)
		}
		if !info.IsDir() {
			return "", fmt.Errorf("the path %s is not a valid directory", home)
		}
	}

	return home, nil
}
