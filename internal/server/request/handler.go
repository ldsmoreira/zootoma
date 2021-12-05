package request

import (
	"bufio"
	"fmt"
	"net"
	action "zootoma/internal/core/action"
	"zootoma/internal/core/memory/executor"
	"zootoma/internal/server/protocol"
	"zootoma/internal/util/logging"
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
	fmt.Println("Data stored in Action Structure:")

	action:= *h.Parser.Action
	balb:= executor.Execute(&action)
	fmt.Println(balb)
	// response, err = executor.Execute(action)

	fmt.Println(h.Parser.Action)
	(*h.conn).Write([]byte("Relaxou"))
	logger := logging.NewCustomLogger(logging.INFO)
	logger.Info("Handle exited with success")

}
