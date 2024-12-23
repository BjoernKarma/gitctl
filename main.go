package main

import (
	"ethical-developer/cli/gitctl/app/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
