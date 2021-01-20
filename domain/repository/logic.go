package repository

import (
	"cli/domain/entity"
)

type AppLogic interface {
	Send(connection entity.Connection, inf entity.ClientInformation) (Response, error)
}
