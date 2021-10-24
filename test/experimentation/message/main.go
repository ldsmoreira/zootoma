package main

import (
	"fmt"
)

type Metadata struct {
	meta []byte
}

type Message struct {
	body []byte
	size int
}

func (msg Message) PrintMessage() int {
	bytesPrinted, _ := fmt.Println(msg.body)
	return bytesPrinted
}

func main() {
	message := Message{[]byte("Here is a string...."), 10}
	message.PrintMessage()
}
