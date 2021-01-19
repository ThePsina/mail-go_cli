package repository

import (
	"cli/domain/entity"
	"net"
)

type AppLogic interface {
	Connect(host, port string) (net.Conn, error)
	CreatePackage(inf entity.ClientInformation) []byte
	Send(dst net.Conn, pkg []byte) error
	Receive(src net.Conn) (Response, error)
}
