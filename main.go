package main

import (
	"cli/application"
	"cli/interfaces"
	"log"
	"os"
)

func main() {
	mailer := application.NewMockedApp()
	app := interfaces.NewApp(mailer)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
