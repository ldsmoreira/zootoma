package maestro

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"zootoma/internals/memdata"
	dxp "zootoma/internals/protocol"
)

var memmap *memdata.PathMap = memdata.NewPathMap()

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StartMaestroServer(host string, port string, conn_type string) {

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

		go handleDXPConnection(conn)
	}
}

func handleDXPConnection(conn net.Conn) {

	method, path, data_size := make([]byte, 3), make([]byte, 30), make([]byte, 30)

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

	switch {
	case reflect.DeepEqual(method, []byte(dxp.SET)):
		dxp_obj.PutDataInMemory(*memmap)
	}

	fmt.Println(memmap)

	conn.Write(data_size)

}
