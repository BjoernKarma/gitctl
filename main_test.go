package main

import (
	"os"
	"testing"
)

func TestRun(_ *testing.T) {
	os.Args = []string{"gitctl", "status", "--config", "gitctl.yaml"}
	main()
}
