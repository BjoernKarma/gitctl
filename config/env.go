package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// GitCtlQuiet is the environment variable for quiet mode
	GitCtlQuiet = "verbosity.quiet"
	// GitCtlVerbose is the environment variable for verbose mode
	GitCtlVerbose = "verbosity.verbose"
	// GitCtlDebug is the environment variable for debug mode
	GitCtlDebug = "verbosity.debug"
	// GitCtlLocal is the environment variable for local mode
	GitCtlLocal = "run_mode.local"
	// GitCtlDryRun is the environment variable for dry run mode
	GitCtlDryRun = "run_mode.dry_run"
	// GitCtlColor is the environment variable for color output
	GitCtlColor = "output.color"
	// GitCtlConcurrency is the environment variable for concurrency level
	GitCtlConcurrency = "run_mode.concurrency"
	// GitCtlBaseDirs is the environment variable for base directories
	GitCtlBaseDirs = "base_dirs"
)

// IsQuiet checks if the quiet mode is enabled
func IsQuiet() bool {
	return viper.GetBool(GitCtlQuiet)
}

// IsVerbose checks if the verbose mode is enabled
func IsVerbose() bool {
	return viper.GetBool(GitCtlVerbose)
}

// IsDebug checks if the debug mode is enabled
func IsDebug() bool {
	return viper.GetBool(GitCtlDebug)
}

// IsLocal checks if the local mode is enabled
func IsLocal() bool {
	return viper.GetBool(GitCtlLocal)
}

// IsDryRun checks if the dry-run mode is enabled
func IsDryRun() bool {
	return viper.GetBool(GitCtlDryRun)
}

// IsColored checks if the color output is enabled
func IsColored() bool {
	return viper.GetBool(GitCtlColor)
}

// GetConcurrency returns the concurrency level as a string
func GetConcurrency() string {
	return viper.GetString(GitCtlConcurrency)
}

// GetBaseDirs returns the base directories as a slice of strings
func GetBaseDirs() []string {
	var baseDirs []string
	if IsLocal() {
		baseDirs = []string{GitctlWorkingDir()}
	} else {
		baseDirs = viper.GetStringSlice(GitCtlBaseDirs)
	}

	var validPaths []string
	for _, dir := range baseDirs {
		absPath, err := filepath.Abs(dir)
		if err != nil {
			continue // Skip invalid paths
		}
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			continue // Skip non-existent paths
		}
		validPaths = append(validPaths, absPath)
	}
	return validPaths
}
