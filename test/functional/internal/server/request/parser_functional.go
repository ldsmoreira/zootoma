package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"zootoma/internal/server/request"
)

func main() {
	file, err := os.Open("set_request_sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parser := request.Parser{Request: new(request.Request)}
	parser.Request.Headers = make(map[string][]byte)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

	}
	scanner.Scan()
	// fmt.Println(scanner.Bytes())
	parser.GetMainHeader(scanner.Bytes())
	// fmt.Println(parser.Request)
	scanner.Scan()
	scanner.Scan()
	parser.GetMetaHeader(scanner.Bytes())
	fmt.Println(parser.Request)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
