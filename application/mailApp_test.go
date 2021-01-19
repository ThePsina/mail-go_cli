package application

import (
	"cli/domain/entity"
	"fmt"
	"reflect"
	"testing"
)

func TestMailApp_createHeader(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: "token", Scope: "scope"}
	header := app.createHeader(inf)

	fmt.Println(header)
	if len(header) != 8 {
		t.Fatal("wrong bytes number")
	}
}

func TestMailApp_createString(t *testing.T) {
	app := NewMailApp()

	strBytes := app.createString("token")

	expected := []byte{
		5, 0, 0, 0,
		116, 111, 107, 101, 110,
	}

	if !reflect.DeepEqual(strBytes, expected) {
		t.Fatal("wrong answer")
	}
}

func TestMailApp_createSvcRequestBody(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: "token", Scope: "scope"}
	body := app.createSvcRequestBody(inf)

	if len(body) != 22 {
		t.Fatal("wrong bytes number")
	}
}

func TestMailApp_CreatePackage(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: "token", Scope: "scope"}
	header := app.CreatePackage(inf)

	if len(header) != 30 {
		t.Fatal("wrong bytes number")
	}
}

func TestMailApp_parseResponseErrorMsg(t *testing.T) {
	app := NewMailApp()
	resp := []byte{
		/* header */
		2, 0, 0, 0,  /* svc_id */
		23, 0, 0, 0, /* body_length */
		11, 0, 0, 0, /* request_id */
		1, 0, 0, 0,  /* return code */
		/* body */
		15, 0, 0, 0,
		116, 111, 107, 101, 110, 32, 110, 111, 116, 32, 102, 111, 117, 110, 100,
	}
	/*
	error: CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND
	message: token not found
	*/

	_, err := app.parseResponse(resp)
	if err != nil {
		t.Fatal("something went wrong")
	}
}

func TestMailApp_parseResponseOk(t *testing.T) {
	app := NewMailApp()
	resp := []byte{
		/* header */
		2, 0, 0, 0, /* svc_id */
		23, 0, 0, 0, /* body_length */
		11, 0, 0, 0, /* request_id */
		0, 0, 0, 0, /* return code */
		/* body */
		/* client_id */
		14, 0, 0, 0,
		116, 101, 115, 116, 95, 99, 108, 105, 101, 110, 116, 95, 105, 100,

		2, 2, 2, 2, /* client_type */

		/* username */
		7, 0, 0, 0,
		77, 105, 99, 104, 97, 101, 108,

		23, 1, 1, 12, /* expires_in */
		23, 1, 1, 12, 23, 1, 1, 12, /* user_id */
	}

	p, err := app.parseResponse(resp)
	if err != nil {
		t.Fatal("something went wrong")
	}

	p.Print()
}
