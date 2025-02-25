package config

import (
	"log"
	"os"
	"path/filepath"
)

// GitctlWorkingDir returns the current gitctl working directory path
func GitctlWorkingDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error trying to find working directory due to %s", err)
	}
	return workingDir
}

// GitctlConfigDir returns the gitctl config directory path
func GitctlConfigDir() string {
	return filepath.Join(HomeDir(), ".config", "gitctl")
}

// HomeDir finds the users home directory
func HomeDir() string {
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
		if err != nil || !info.IsDir() {
			log.Fatalf("The path %s is not a valid directory", home)
		}
	}

	return home
}
