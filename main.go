package main

import (
	"cli/interfaces"
	"log"
	"os"
)

func main() {
	app := interfaces.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
