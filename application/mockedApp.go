package application

import (
	"bytes"
	"cli/domain/entity"
	"cli/domain/repository"
	"encoding/binary"
	"fmt"
	"reflect"
)

type MockedApp struct {
}

func NewMockedApp() *MockedApp {
	return &MockedApp{}
}

/*
	<packet> ::= <request> | <response>
	<request> ::= <header><svc_request_body>
*/
func (mockedApp *MockedApp) createMockedResponse(inf entity.ClientInformation) []byte {
	if inf.Token == "abracadabra" && inf.Scope == "test" {
		return []byte{
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
	}

	if inf.Token == "abracadabra" && inf.Scope == "xxx" {
		return []byte{
				/* header */
			2, 0, 0, 0,  /* svc_id */
			23, 0, 0, 0, /* body_length */
			0, 0, 0, 0,  /* request_id */

				/* body */
			1, 0, 0, 0,  /* return code */

			/* message */
			15, 0, 0, 0,
			116, 111, 107, 101, 110, 32, 110, 111, 116, 32, 102, 111, 117, 110, 100,
		}
	}
	return []byte{
		/* header */
		2, 0, 0, 0,  /* svc_id */
		16, 0, 0, 0, /* body_length */
		0, 0, 0, 0,  /* request_id */

		/* body */
		2, 0, 0, 0,  /* return code */

		/* message */
		8, 0, 0, 0,
		100, 98, 32, 101, 114, 114, 111, 114,
	}
}

func (mockedApp *MockedApp) Send(connection entity.Connection, inf entity.ClientInformation) (repository.Response, error) {
	return mockedApp.parseResponse(mockedApp.createMockedResponse(inf))
}

func (mockedApp *MockedApp) parseResponse(rawResponse []byte) (repository.Response, error) {
	reader := bytes.NewReader(rawResponse)
	respInf := entity.ResponseInformation{}
	if err := binary.Read(reader, binary.LittleEndian, &respInf); err != nil {
		return nil, err
	}

	body := rawResponse[bodyBeginPos:]
	if respInf.ReturnCode == 0 { // no error
		response := &entity.ResponseOk{}
		err := mockedApp.fillResponse(response, body)
		return response, err
	}
	response := &entity.ResponseErr{}
	err := mockedApp.fillResponse(response, body)
	response.ReturnCode = respInf.ReturnCode
	return response, err
}

func (mockedApp *MockedApp) fillResponse(inf interface{}, data []byte) error {
	reader := bytes.NewReader(data)

	val := reflect.ValueOf(inf)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("input is not pointer")
	}

	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int32:
			var value uint32
			if err := binary.Read(reader, binary.LittleEndian, &value); err != nil {
				return err
			}
			valueField.SetInt(int64(value))
		case reflect.Int64:
			var value uint64
			if err := binary.Read(reader, binary.LittleEndian, &value); err != nil {
				return err
			}
			valueField.SetInt(int64(value))
		case reflect.String:
			var lenRaw uint32
			if err := binary.Read(reader, binary.LittleEndian, &lenRaw); err != nil {
				return err
			}

			dataRaw := make([]byte, lenRaw)
			if err := binary.Read(reader, binary.LittleEndian, &dataRaw); err != nil {
				return err
			}

			valueField.SetString(string(dataRaw))
		default:
			return fmt.Errorf("bad type: %v for field %v", typeField.Type.Kind(), typeField.Name)
		}
	}

	return nil
}
