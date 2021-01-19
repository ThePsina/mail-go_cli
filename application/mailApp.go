package application

import (
	"bytes"
	"cli/domain/entity"
	"cli/domain/repository"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
)

const (
	svcId        int32 = 2
	bodyBeginPos       = 16
	svcMsg       int32 = 1
)

type MailApp struct {
}

func NewMailApp() *MailApp {
	return &MailApp{}
}

func (mailApp *MailApp) Connect(host, port string) (net.Conn, error) {
	return net.Dial("tcp", net.JoinHostPort(host, port))
}

/*
	<packet> ::= <request> | <response>
	<request> ::= <header><svc_request_body>
*/
func (mailApp *MailApp) CreatePackage(inf entity.ClientInformation) []byte {
	header := mailApp.createHeader(inf)
	svcRequestBody := mailApp.createSvcRequestBody(inf)

	return append(header, svcRequestBody...)
}

/*
	<header> ::= <svc_id><body_length>

	<svc_id> ::= <int32> - идентификатор CUBE сервиса
	<body_length> ::= <int32> - длина тела запроса
*/
func (mailApp *MailApp) createHeader(inf entity.ClientInformation) []byte {
	svcIdBinary := make([]byte, 4) // int32 - 4 byte
	binary.LittleEndian.PutUint32(svcIdBinary, uint32(svcId))

	bodyLength := int32(len(inf.Token)+len(inf.Scope)) + svcMsg
	bodyLengthBinary := make([]byte, 4) // int32 - 4 byte
	binary.LittleEndian.PutUint32(bodyLengthBinary, uint32(bodyLength))

	return append(svcIdBinary, bodyLengthBinary...)
}

/*
	<svc_request_body> ::= <svc_msg><token><scope>

	<svc_msg> ::= <int32> - номер сообщения для проверки access token и scope, равен 0x00000001
	<token> ::= <string> - проверяемый токен
	<scope> ::= <string> - проверяемый scope
*/
func (mailApp *MailApp) createSvcRequestBody(inf entity.ClientInformation) []byte {
	svcMsgBinary := make([]byte, 4) // int32 - 4 byte
	binary.LittleEndian.PutUint32(svcMsgBinary, uint32(svcMsg))

	token := mailApp.createString(inf.Token)
	scope := mailApp.createString(inf.Scope)

	svcRequestBody := append(svcMsgBinary, token...)
	svcRequestBody = append(svcRequestBody, scope...)
	return svcRequestBody
}

/*
	<string> ::= <str_len><str>

	<str_len> ::= <int32> - длина строки, больше 0
	<str> ::= <int8>+ - строка
	<int8> ::= целочисленное число со знаком в бинарном виде, размер 1 байт
*/
func (mailApp *MailApp) createString(str string) []byte {
	strLen := make([]byte, 4) // int32 - 4 byte
	binary.LittleEndian.PutUint32(strLen, uint32(len(str)))

	strBytes := []byte(str)

	return append(strLen, strBytes...) // len bytes + str bytes
}

func (mailApp *MailApp) Send(dst net.Conn, pkg []byte) error {
	_, err := dst.Write(pkg)
	return err
}

func (mailApp *MailApp) Receive(src net.Conn) (repository.Response, error) {
	tmp := make([]byte, 256)
	_, err := src.Read(tmp)
	if err != nil {
		return nil, err
	}

	return mailApp.parseResponse(tmp)
}

func (mailApp *MailApp) parseResponse(rawResponse []byte) (repository.Response, error) {
	reader := bytes.NewReader(rawResponse)
	respInf := entity.ResponseInformation{}
	if err := binary.Read(reader, binary.LittleEndian, respInf); err != nil {
		return nil, err
	}

	body := rawResponse[bodyBeginPos:]
	if respInf.ReturnCode == 0 {  // no error
		response := &entity.ResponseOk{}
		err := mailApp.fillResponse(response, body)
		return response, err
	}
	response := &entity.ResponseErr{}
	err := mailApp.fillResponse(response, body)
	return response, err
}

func (mailApp *MailApp) fillResponse(inf interface{}, data []byte) error {
	reader := bytes.NewReader(data)

	val := reflect.ValueOf(inf)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("input is not pointer")
	}

	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

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
