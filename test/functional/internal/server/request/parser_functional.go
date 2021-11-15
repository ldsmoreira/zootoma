package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"zootoma/internal/core"
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
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	counter := 0
	for scanner.Scan() {
		parser.BuildAction(scanner.Bytes(), counter)
		fmt.Println(scanner.Bytes())
		if len(scanner.Bytes()) == 0 {
			counter++
		}
	}
	fmt.Println(parser)
	fmt.Printf("Pointer value request=>%p:\t%p\n", parser.Request, &parser.Request)
	fmt.Printf("Pointer value request=>%p:\t%p\n", parser.Action, &parser.Action)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
