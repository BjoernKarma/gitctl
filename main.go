package main

import (
	"github.com/bjoernkarma/gitctl/app/cmd"
)

func main() {
	_ = cmd.Execute() // Execute handles all error display and os.Exit internally
}
