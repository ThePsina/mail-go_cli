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

func TestApp_RunFailConnect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logic := mock.NewMockAppLogic(ctrl)
	logic.EXPECT().Connect(host, port).Return(nil, errors.New("TestApp_RunFailConnect"))

	app := NewApp(logic)

	args := []string{"1", host, port, token, scope}
	err := app.Run(args)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestApp_RunFailSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logic := mock.NewMockAppLogic(ctrl)
	logic.EXPECT().Connect(host, port).Return(nil, nil)
	logic.EXPECT().
		CreatePackage(entity.ClientInformation{Token: token, Scope: scope}).
		Return([]byte{1, 0, 0, 1})
	logic.EXPECT().
		Send(nil, []byte{1, 0, 0, 1}).
		Return(errors.New("TestApp_RunFailCreatePackage"))

	app := NewApp(logic)

	args := []string{"1", host, port, token, scope}
	err := app.Run(args)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestApp_RunFailReceive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logic := mock.NewMockAppLogic(ctrl)
	logic.EXPECT().Connect(host, port).Return(nil, nil)
	logic.EXPECT().
		CreatePackage(entity.ClientInformation{Token: token, Scope: scope}).
		Return([]byte{1, 0, 0, 1})
	logic.EXPECT().
		Send(nil, []byte{1, 0, 0, 1}).
		Return(nil)
	logic.EXPECT().
		Receive(nil).
		Return(&entity.ResponseOk{}, errors.New("TestApp_RunFailReceive"))

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
	logic.EXPECT().Connect(host, port).Return(nil, nil)
	logic.EXPECT().
		CreatePackage(entity.ClientInformation{Token: token, Scope: scope}).
		Return([]byte{1, 0, 0, 1})
	logic.EXPECT().
		Send(nil, []byte{1, 0, 0, 1}).
		Return(nil)
	logic.EXPECT().
		Receive(nil).
		Return(&entity.ResponseOk{}, nil)

	app := NewApp(logic)

	args := []string{"1", host, port, token, scope}
	err := app.Run(args)

	if err != nil {
		t.Fatal("receive error")
	}
}
