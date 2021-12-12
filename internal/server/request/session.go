package request

import (
	"bufio"
	"encoding/json"
	"net"

	"github.com/moreira0102/zootoma/internal/core/memory/executor"
	"github.com/moreira0102/zootoma/internal/server/protocol"
	"github.com/moreira0102/zootoma/internal/util/logging"
)

var session_logger *logging.CustomLogger = logging.NewCustomLogger(logging.INFO)

type Session struct {
	conn    *net.Conn
	Handler Handler
	Reader  *bufio.Reader
	Writer  *bufio.Writer
}

func NewSession(conn *net.Conn) (session Session) {
	conn = conn

	reader := bufio.NewReader(*conn)
	writer := bufio.NewWriter(*conn)

	session = Session{
		conn:   conn,
		Reader: reader,
		Writer: writer,
	}

	return session
}

func (s Session) handle() (conn_status bool) {

	var buffer []byte
	var err error
	var offset = 0

	s.Handler = NewHandler()
	for offset < 2 {
		buffer, err = s.Reader.ReadBytes(protocol.STATEMENTS_DELIMITER)
		if err != nil {
			session_logger.Error(err.Error())
			return false
		}
		if buffer[0] != protocol.STATEMENTS_DELIMITER {
			s.Handler.Parser.BuildAction(&buffer, offset)
		} else {
			offset++
		}
	}

	buffer = make([]byte, s.Handler.Parser.Action.DataSize)
	s.Reader.Read(buffer)
	s.Handler.Parser.BuildAction(&buffer, protocol.DATA_BLOCK_POSITION)

	action := *s.Handler.Parser.Action
	actionresp := executor.Execute(&action)
	if actionresp.Data != nil {
		resp, _ := json.Marshal(actionresp)
		(*s.conn).Write(resp)
	} else {
		resp, _ := json.Marshal(actionresp)
		(*s.conn).Write(resp)
	}

	return true

}

func (s Session) Handle() {

	conn_status := true
	for conn_status {
		conn_status = s.handle()
	}

	session_logger.Info("Client " + (*s.conn).RemoteAddr().String() + " disconnected")
	(*s.conn).Close()
}
