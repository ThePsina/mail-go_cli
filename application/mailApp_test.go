package application

import (
	"cli/domain/entity"
	"fmt"
	"reflect"
	"testing"
)

const (
	token = "token"
	scope = "scope"
)

func TestMailApp_createHeader(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: token, Scope: scope}
	header := app.createHeader(inf)

	fmt.Println(header)
	if len(header) != 8 {
		t.Fatalf("wrong bytes number\nexpected: 8\ngot: %d", len(header))
	}
}

func TestMailApp_createString(t *testing.T) {
	app := NewMailApp()

	strBytes := app.createString(token)

	expected := []byte{
		5, 0, 0, 0,
		116, 111, 107, 101, 110,
	}

	if !reflect.DeepEqual(strBytes, expected) {
		t.Fatal("not deep equal")
	}
}

func TestMailApp_createSvcRequestBody(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: token, Scope: scope}
	body := app.createSvcRequestBody(inf)

	if len(body) != 22 {
		t.Fatalf("wrong bytes number\nexpected: 22\ngot: %d", len(body))
	}
}

func TestMailApp_CreatePackage(t *testing.T) {
	app := NewMailApp()

	inf := entity.ClientInformation{Token: token, Scope: scope}
	header := app.createPackage(inf)

	if len(header) != 30 {
		t.Fatalf("wrong bytes number\nexpected: 30\ngot: %d", len(header))
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
		t.Fatalf("got error: %s", err.Error())
	}
}

func TestMailApp_parseResponseOk(t *testing.T) {
	app := NewMailApp()
	resp := []byte{
		/* header */
		2, 0, 0, 0,  /* svc_id */
		44, 0, 0, 0, /* body_length */
		0, 0, 0, 0,  /* request_id */

		/* body */
		0, 0, 0, 0,  /* return code */

		/* client_id */
		9, 0, 0, 0,
		99, 108, 105, 101, 110, 116, 95, 105, 100,

		4, 0, 0, 0,  /* client_type*/

		/* username */
		7, 0, 0, 0,
		77, 105, 99, 104, 97, 101, 108,

		132, 3, 0, 0,  			  /* expires_in */
		14, 0, 0, 0, 0, 0, 0, 0,  /* user_id */
	}
	/*
	client_id: client_id
	client_type: 4
	expires_in: 900
	user_id: 14
	username: Michael
	*/

	_, err := app.parseResponse(resp)
	if err != nil {
		t.Fatalf("got error: %s", err.Error())
	}
}
