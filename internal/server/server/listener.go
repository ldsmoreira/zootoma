package server

import (
	"fmt"
	"net"
	"os"
)

func StartListen(host string, port string, conn_type string) {

	fmt.Println("Starting " + conn_type + " server on " + host + ":" + port)
	l, err := net.Listen(conn_type, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")

		std := StdHandler{conn: conn}

		go std.handle()
	}
}
