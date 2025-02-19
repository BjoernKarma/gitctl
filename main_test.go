package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	os.Args = []string{"gitctl", "status", "--config", "gitctl.yaml"}
	main()
}
