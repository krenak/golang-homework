// Data structure - Range (Chapter 8 - Episode 3)
// https://go.dev/ref/spec#Composite_literals
// https://gobyexample.com/ranges
// https://youtu.be/l6O8HWaoy_w

package main

import(
	"fmt"
)

func main() {
	slice := []string{"banana", "apple", "watermelon"}
	for índice, valor := range slice {
		fmt.Println("No índice", índice,"temos o valor:", valor)
	}
}
