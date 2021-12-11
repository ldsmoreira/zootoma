package request

// The request module implements the Request structure and it's methods

type Request struct {
	MainHeader []byte
	MetaHeader [][]byte
	Data       []byte
}
