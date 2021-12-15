// Data structure - Range (Chapter 8 - Episode 3)
// https://go.dev/ref/spec#Composite_literals
// https://gobyexample.com/range
// https://youtu.be/l6O8HWaoy_w

package main

import(
	"fmt"
)

func main() {
	slice := []string{"banana", "apple", "watermelon"}
//	for índice, valor := range slice {
//		fmt.Println("No índice", índice,"temos o valor:", valor)
//	}

	slice = append(slice, "coxinha")
	for índice, valor := range slice { // range iterates over elements in 
					   // a variety of data structures
		fmt.Println("No índice", índice,"temos o valor:", valor)
	}
	// another example provided by Ellen's
	slice1 := []int{20,21,22,23,1,13}
	total := 0
	for _, valor1 := range slice1 { // the before índice value is discharged
		total += valor1		// and here total is summed with valor1
	}
	fmt.Println("O valor total é:", total)
}
