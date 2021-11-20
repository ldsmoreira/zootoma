package request

import (
	"bytes"
	"errors"
	"strconv"
	action "zootoma/internal/core"
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
	Action  *action.Action
}

func (p Parser) SetRawMainHeader(mainHeader []byte) {
	p.Request.MainHeader = mainHeader
}

func (p Parser) SetRawMetaHeader(metaHeader []byte) {
	p.Request.MetaHeader = append(p.Request.MetaHeader, metaHeader)
}

func (p Parser) SetActionData(data []byte) {
	p.Action.Data = data
}

// The getMainHeader method of the Parser struct is responsible to parse and
// validate the first line of the raw request
func (p Parser) ParseMainHeader() (err error) {

	var (
		isValidMethod bool
		isValidKey    bool
		isValidSize   bool
	)

	if p.Request.MainHeader == nil {
		return errors.New("Main header is null")
	}

	sl := bytes.Split(p.Request.MainHeader, protocol.MAIN_HEADER_SEPARATOR)

	if len(sl) != protocol.MAIN_HEADER_ITEMS {
		return errors.New("Main header line malformatted or corrupted")
	}

	method, key := sl[0], sl[1]
	size, _ := strconv.Atoi(string(sl[2][:len(sl[2])-1]))

	isValidMethod = protocol.IsValidMethod(bytes.ToLower(method))
	isValidKey = protocol.IsValidKey(key)
	isValidSize = protocol.IsValidSize(size)

	if isValidMethod && isValidKey && isValidSize {
		p.Action.Method = string(method)
		p.Action.Key = string(key)
		p.Action.DataSize = size
		return nil
	} else {
		return errors.New("Invalid main header options")
	}
}

func (p Parser) ParseMetaHeader() (err error) {
	err = nil
	for _, value := range p.Request.MetaHeader {
		key, value, valid := protocol.IsValidMetaHeader(value)
		if valid {
			p.Action.Headers[string(key)] = value[:len(value)-1]
		} else {
			err = errors.New("Invalid MetaHeader")
		}
	}
	return err
}

func (p Parser) BuildAction(buffer *[]byte, block_position int) (err error) {

	switch block_position {

	case protocol.MAIN_HEADER_POSITION:

		p.SetRawMainHeader(*buffer)
		err = p.ParseMainHeader()
		return err

	case protocol.META_HEADER_BLOCK_POSITION:

		p.SetRawMetaHeader(*buffer)
		err = p.ParseMetaHeader()
		return err

	case protocol.DATA_BLOCK_POSITION:

		p.SetActionData(*buffer)
		return nil

	default:
		return errors.New("Wrong format for request !!")
	}
}
