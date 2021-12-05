// package main

// import "fmt"

// func update(age *int, text *string) {
// 	*age = *age + 5        // defrencing pointer address
// 	*text = *text + " mama" // defrencing pointer address
// }

// func main() {
// 	var age = 20
// 	var text = "John"
// 	fmt.Println("Before:", text, age)
// // usando o ponteiro da variavel
// 	for x := 1; x < 27; x++ {
// 		update(&age, &text)

// 		fmt.Println("After :", text, age)
// 	}
// }
package main

import (
	"fmt"
	"zootoma/internal/core/action"
	"zootoma/internal/core/memory/executor"
)

func main() {
	actn:= action.Action{Method: "get", DataSize: 9000, Data: nil, Key: "Toma"}
	resp:= executor.Execute(&actn)
	fmt.Println(resp)
}