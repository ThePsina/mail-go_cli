package interfaces

import "cli/application"

type App struct {
	mailer *application.MailApp
}

func NewApp() *App {
	return &App{mailer: application.NewMailApp()}
}
