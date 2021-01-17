package repository

import (
	"cli/domain/entity"
	"net"
)

type AppLogic interface {
	Connect(host, port string) (*net.Conn, error)
	CreatePackage(inf entity.ClientInfReq) (string, error)
	Send(dst *net.Conn, pkg string) (string, error)
	ReadResponse(resp string) (entity.Response, error)
}
