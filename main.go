package main

import (
	"log"
	"os"

	"github.com/larse514/go-cli-boilerplate/config"
)

func main() {
	app := config.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
