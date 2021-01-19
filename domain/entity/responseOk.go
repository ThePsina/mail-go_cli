package entity

import "fmt"

type ResponseOk struct {
	ClientId   string
	ClientType int32
	Username   string
	ExpiresIn  int32
	UserId     int64
}

func (ok *ResponseOk) Print() {
	fmt.Printf(
		"client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n",
		ok.ClientId,
		ok.ClientType,
		ok.ExpiresIn,
		ok.UserId,
		ok.Username,
	)
}
