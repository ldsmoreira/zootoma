package main

import (
	"flag"

	"github.com/moreira0102/zootoma/internal/server/server"
)

const (
	HOST = "localhost"
)

func main() {

	PORT := flag.String("port", "9009", "The port to bind zootoma")
	CONN_TYPE := flag.String("conntype", "tcp", "The transport layer protocol used by zootoma")

	flag.Parse()

	server.StartListen(HOST, *PORT, *CONN_TYPE)
}
