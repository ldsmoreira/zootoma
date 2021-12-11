package request

import (
	"bufio"
	"net"

	action "github.com/moreira0102/zootoma/internal/core/action"
	"github.com/moreira0102/zootoma/internal/core/memory/executor"
	"github.com/moreira0102/zootoma/internal/server/protocol"
)

type Handler struct {
	conn   *net.Conn
	Parser Parser
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewHandler(conn *net.Conn) (handler Handler) {

	conn = conn

	parser := Parser{
		Request: new(Request),
		Action:  new(action.Action),
	}

	reader := bufio.NewReader(*conn)
	writer := bufio.NewWriter(*conn)

	handler = Handler{
		conn:   conn,
		Parser: parser,
		Reader: reader,
		Writer: writer}

	handler.Parser.Action.Headers = make(map[string][]byte)

	return handler
}

func (h Handler) Handle() {

	var buffer []byte
	var offset = 0

	for offset < 2 {
		buffer, _ = h.Reader.ReadBytes(protocol.STATEMENTS_DELIMITER)
		if buffer[0] != protocol.STATEMENTS_DELIMITER {
			h.Parser.BuildAction(&buffer, offset)
		} else {
			offset++
		}
	}

	buffer = make([]byte, h.Parser.Action.DataSize)
	h.Reader.Read(buffer)
	h.Parser.BuildAction(&buffer, protocol.DATA_BLOCK_POSITION)

	action := *h.Parser.Action
	actionresp := executor.Execute(&action)
	if actionresp.Data != nil {
		(*h.conn).Write([]byte(*actionresp.Data))
	} else {
		(*h.conn).Write([]byte(actionresp.Message))
	}
}
