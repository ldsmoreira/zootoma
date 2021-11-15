package request

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
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
	Request *Request
}

// The getMainHeader method of the Parser struct is responsible to parse and
// validate the first line of the raw request
func (p Parser) GetMainHeader(mh []byte) (req *Request, err error) {

	var (
		isValidMethod bool
		isValidKey    bool
		isValidSize   bool
	)

	sl := bytes.Split(mh, protocol.MainHeaderSeparator)

	if len(sl) != protocol.MainHeaderCompQtt {
		return nil, errors.New("Main header line malformatted or corrupted")
	}

	method, key, size := sl[0], sl[1], sl[2]

	fmt.Println(string(method), string(key), string(size))

	isValidMethod = protocol.IsValidMethod(bytes.ToLower(method))
	isValidKey = protocol.IsValidKey(key)
	isValidSize = protocol.IsValidSize(size)

	fmt.Println(isValidMethod, isValidKey, isValidSize)

	if isValidMethod && isValidKey && isValidSize {
		p.Request.Method = string(method)
		p.Request.Key = string(key)
		p.Request.DataSize, _ = strconv.Atoi(string(size))

		return p.Request, nil
	} else {
		return nil, errors.New("Invalid main header options")
	}

}

func (p Parser) GetMetaHeader(mh []byte) (req *Request, err error) {
	key, value, valid := protocol.IsValidMetaHeader(mh)

	if valid {
		p.Request.Headers[string(key)] = value
		return p.Request, nil
	} else {
		return p.Request, errors.New("Invalid MetaHeader")
	}
}
