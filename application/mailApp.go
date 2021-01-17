package application

import (
	"cli/domain/entity"
	"net"
)

type MailApp struct {
}

func NewMailApp() *MailApp {
	return &MailApp{}
}

func (mailApp *MailApp) Connect(host, port string) (*net.Conn, error) {
	hostPort := net.JoinHostPort(host, port)
	conn, err := net.Dial("tcp", hostPort)
	return &conn, err
}

func (mailApp *MailApp) CreatePackage(inf entity.ClientInfReq) (string, error) {
	return "", nil
}

func (mailApp *MailApp) Send(dst *net.Conn, pkg string) (string, error) {
	return "", nil
}

func (mailApp *MailApp) ReadResponse(resp string) (entity.Response, error) {
	return entity.Response{}, nil
}
