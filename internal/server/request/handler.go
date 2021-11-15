package request

import "net"

type Handler struct {
	conn *net.Conn
}
