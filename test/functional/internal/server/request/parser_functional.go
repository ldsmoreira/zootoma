package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	action "zootoma/internal/core/action"
	"zootoma/internal/server/protocol"
	"zootoma/internal/server/request"
)

func main() {
	file, err := os.Open("set_request_sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parser := request.Parser{Request: new(request.Request), Action: new(action.Action)}
	parser.Action.Headers = make(map[string][]byte)
	reader := bufio.NewReader(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	counter := 0
	for counter < 3 {
		var b []byte
		if counter < 2 {
			b, _ = reader.ReadBytes(protocol.STATEMENTS_DELIMITER)
		} else {
			b = make([]byte, parser.Action.DataSize)
			_, _ = reader.Read(b)
			b = b[:len(b)-1]
		}
		parser.BuildAction(&b, counter)
		if len(b) == 0 {
			counter++
		}
	}
	fmt.Println(parser.Action.Data)
	fmt.Printf("Pointer value request=>%p:\t%p\n", parser.Request, &parser.Request)
	fmt.Printf("Pointer value request=>%p:\t%p\n", parser.Action, &parser.Action)
}
