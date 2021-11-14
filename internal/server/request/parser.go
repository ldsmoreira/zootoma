package request

import (
	"bytes"
	"net"
	"zootoma/internal/server/protocol"
)

// The parser module implements the function that receives a connection object
// and return a Request object if it is valid or return an error if it isn't

// Example of a valid request:

// set /home/lucas/data.txt 30000
//
// status::ok
// host::com.toma
//
// (every data that fits 30000 bytes)

type Parser struct {
	conn *net.Conn
}

func isValidMainHeader(mh []byte) bool {

	var (
		isValidMethod bool
		isValidKey    bool
		isValidSize   bool
	)

	sl := bytes.Split(mh, []byte(" "))

	if len(sl) != protocol.MainHeaderCompQtt {
		return false
	}

	method, key, size := sl[0], sl[1], sl[2]

	isValidMethod = protocol.IsValidMethod(bytes.ToLower(method))
	isValidKey = protocol.IsValidKey(key)
	isValidSize = protocol.IsValidSize(size)

	return (isValidMethod && isValidKey && isValidSize)

}
