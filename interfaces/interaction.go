package interfaces

import (
	"cli/domain/entity"
	"errors"
)

const (
	numOfArgs = 5
	hostPos   = 1
	portPos   = 2
	tokenPos  = 3
	scopePos  = 4
)

func (app *App) Run(args []string) error {
	if err := app.checkArgs(args); err != nil {
		return err
	}

	conn, err := app.mailer.Connect(args[hostPos], args[portPos])
	if err != nil {
		return err
	}

	err = app.mailer.Send(conn, app.mailer.CreatePackage(entity.ClientInformation{Token: args[tokenPos], Scope: args[scopePos]}))
	if err != nil {
		return err
	}

	resp, err := app.mailer.Receive(conn)
	if err != nil {
		return err
	}

	resp.Print()
	return nil
}

func (app *App) checkArgs(args []string) error {
	if len(args) != numOfArgs {
		return errors.New("wrong number of arguments")
	}

	return nil
}
