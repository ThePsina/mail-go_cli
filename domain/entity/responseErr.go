package entity

import "fmt"

const numOfErrorsCode = 6

type ResponseErr struct {
	ErrorString string
	ReturnCode  int32 `unpack:"-"`
}

var errsCode = make(map[int32]string, numOfErrorsCode)
var errsMsg = make(map[string]string, numOfErrorsCode)
func init() {
	if len(errsCode) != 0 || len(errsMsg) != 0 {
		return
	}
	errsCode[1] = "CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND"
	errsCode[2] = "CUBE_OAUTH2_ERR_DB_ERROR"
	errsCode[3] = "CUBE_OAUTH2_ERR_UNKNOWN_MSG"
	errsCode[4] = "CUBE_OAUTH2_ERR_BAD_PACKET"
	errsCode[5] = "CUBE_OAUTH2_ERR_BAD_CLIENT"
	errsCode[6] = "CUBE_OAUTH2_ERR_BAD_SCOPE"
}

func (err *ResponseErr) Print() {
	fmt.Printf("error: %s\nmessage: %s\n", errsCode[err.ReturnCode], err.ErrorString)
}
