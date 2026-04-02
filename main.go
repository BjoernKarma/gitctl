package main

import (
	"log"

	"github.com/bjoernkarma/gitctl/app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
