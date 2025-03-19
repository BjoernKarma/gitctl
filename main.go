package main

import (
	"github.com/bjoernkarma/gitctl/app/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
