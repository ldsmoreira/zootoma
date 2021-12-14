package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	dxp "github.com/moreira0102/zootoma/test/experimentation/protocol"
)

const (
	connHost = "localhost"
	connPort = "9000"
	connType = "tcp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleDXPConnection(c)
	}
}

func handleDXPConnection(conn net.Conn) {

	method, path, data_size := make([]byte, 3), make([]byte, 30), make([]byte, 30)
	// var data []byte

	_, err := conn.Read(method)
	check(err)
	_, err = conn.Read(path)
	check(err)
	_, err = conn.Read(data_size)
	check(err)

	dxp_obj := dxp.NewDxpProtocolMapping(method, path, data_size)
	data := make([]byte, dxp_obj.Data_size)

	_, err = conn.Read(data)
	dxp_obj.Data = data

	fmt.Println(dxp_obj)

	conn.Write(data_size)

}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes(0)

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
