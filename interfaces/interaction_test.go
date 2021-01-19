package interfaces

import (
	"cli/domain/entity"
	"cli/infrastructure/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

const (
	host  = "host"
	port  = "port"
	token = "token"
	scope = "scope"
)

func TestApp_RunFailArgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := NewApp(mock.NewMockAppLogic(ctrl))

	args := []string{"1", host, port, token}
	err := app.Run(args)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestApp_RunFailSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logic := mock.NewMockAppLogic(ctrl)
	logic.EXPECT().
		Send(entity.Connection{Host: host, Port: port}, entity.ClientInformation{Token: token, Scope: scope}).
		Return(nil, errors.New("TestApp_RunFailCreatePackage"))

	app := NewApp(logic)

	args := []string{"1", host, port, token, scope}
	err := app.Run(args)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestApp_RunSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logic := mock.NewMockAppLogic(ctrl)
	logic.EXPECT().
		Send(entity.Connection{Host: host, Port: port}, entity.ClientInformation{Token: token, Scope: scope}).
		Return(&entity.ResponseOk{}, nil)

	app := NewApp(logic)

	args := []string{"1", host, port, token, scope}
	err := app.Run(args)

	if err != nil {
		t.Fatal("expected no error")
	}
}
