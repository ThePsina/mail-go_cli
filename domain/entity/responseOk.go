package entity

type ResponseOk struct {
	StrLenClientId int32
	// slice of chars

	ClientType int32

	StrLenUsername int32
	// slice of chars

	ExpiresIn      int32
	UserId         int64
}

func (ok *ResponseOk) Print() {

}
