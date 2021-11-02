package main

import (
	"zootoma/internals/maestro"
)

const (
	HOST      = "localhost"
	PORT      = "9000"
	CONN_TYPE = "tcp"
)

func main() {
	maestro.StartMaestroServer(HOST, PORT, CONN_TYPE)
}
