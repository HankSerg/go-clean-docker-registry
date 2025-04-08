package main

import (
	"github.com/HankSerg/go-clean-docker-registry/cmd"
	"log"
	"os"
)

func main() {
	app := cmd.CreateApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
