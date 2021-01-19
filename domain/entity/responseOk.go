package entity

type ResponseOk struct {
	ClientId   string
	ClientType int32
	Username   string
	ExpiresIn  int32
	UserId     int64
}

func (ok *ResponseOk) Print() {

}
