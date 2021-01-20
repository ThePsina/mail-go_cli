package interfaces

import (
	"cli/domain/repository"
)

type App struct {
	mailer repository.AppLogic
}

func NewApp(handler repository.AppLogic) *App {
	return &App{mailer: handler}
}
