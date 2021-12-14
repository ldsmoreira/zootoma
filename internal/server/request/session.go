package request

import (
	"bufio"
	"encoding/binary"
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
			session_logger.Info("Client " + (*s.conn).RemoteAddr().String() + " disconnected")
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
	_, resp := executor.Execute(&action)

	content_length := len(resp)

	prefix := make([]byte, 8)
	binary.LittleEndian.PutUint64(prefix, uint64(content_length))
	resp = append(prefix, resp...)
	(*s.conn).Write(resp)

	return true

}

func (s Session) Handle() {

	for conn_status := true; conn_status; {
		conn_status = s.handle()
	}

	(*s.conn).Close()
}
