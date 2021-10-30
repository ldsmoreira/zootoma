package protocol

type Protocol struct {
	method    [4]byte
	path      [512]byte
	data_size [512]byte
	data      []byte
}
