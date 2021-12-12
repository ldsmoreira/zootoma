package server

import (
	"net"
	"os"

	"github.com/moreira0102/zootoma/internal/server/request"
	"github.com/moreira0102/zootoma/internal/util/logging"
)

var listener_logger *logging.CustomLogger = logging.NewCustomLogger(logging.INFO)

//StartListen is the network layer entry point of the application
//It's start listening
func StartListen(host string, port string, conn_type string) {

	listener_logger.Info("Zootoma runnig on " + host + ":" + port)
	l, err := net.Listen(conn_type, host+":"+port)
	if err != nil {
		listener_logger.Error("Error listening:" + string(err.Error()))
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			listener_logger.Info("Error connecting:" + string(err.Error()))
			return
		}
		listener_logger.Info("Client connected.")

		listener_logger.Info("Client " + conn.RemoteAddr().String() + " connected.")

		session := request.NewSession(&conn)

		go session.Handle()
	}
}
