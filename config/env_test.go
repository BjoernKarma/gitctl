package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestIsQuietReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlQuiet, true)
	if !IsQuiet() {
		t.Errorf("expected true, got false")
	}
}

func TestIsQuietReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlQuiet, false)
	if IsQuiet() {
		t.Errorf("expected false, got true")
	}
}

func TestIsVerboseReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlVerbose, true)
	if !IsVerbose() {
		t.Errorf("expected true, got false")
	}
}

func TestIsVerboseReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlVerbose, false)
	if IsVerbose() {
		t.Errorf("expected false, got true")
	}
}

func TestIsDebugReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlDebug, true)
	if !IsDebug() {
		t.Errorf("expected true, got false")
	}
}

func TestIsDebugReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlDebug, false)
	if IsDebug() {
		t.Errorf("expected false, got true")
	}
}

func TestIsLocalReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlLocal, true)
	if !IsLocal() {
		t.Errorf("expected true, got false")
	}
}

func TestIsLocalReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlLocal, false)
	if IsLocal() {
		t.Errorf("expected false, got true")
	}
}

func TestIsDryRunReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlDryRun, true)
	if !IsDryRun() {
		t.Errorf("expected true, got false")
	}
}

func TestIsDryRunReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlDryRun, false)
	if IsDryRun() {
		t.Errorf("expected false, got true")
	}
}

func TestIsColoredReturnsTrueWhenEnabled(t *testing.T) {
	viper.Set(GitCtlColor, true)
	if !IsColored() {
		t.Errorf("expected true, got false")
	}
}

func TestIsColoredReturnsFalseWhenDisabled(t *testing.T) {
	viper.Set(GitCtlColor, false)
	if IsColored() {
		t.Errorf("expected false, got true")
	}
}

func TestGetConcurrencyReturnsCorrectValue(t *testing.T) {
	expected := "4"
	viper.Set(GitCtlConcurrency, expected)
	result := GetConcurrency()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetBaseDirsReturnsCorrectValueWhenLocal(t *testing.T) {
	viper.Set(GitCtlLocal, true)
	expected := []string{GitctlWorkingDir()}
	result := GetBaseDirs()
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetBaseDirsReturnsCorrectValueWhenNotLocal(t *testing.T) {
	viper.Set(GitCtlLocal, false)
	expected := []string{"dir1", "dir2"}
	viper.Set(GitCtlBaseDirs, expected)
	result := GetBaseDirs()
	if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
