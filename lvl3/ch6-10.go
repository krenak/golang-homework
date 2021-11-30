// Flow control: Switch and doc (Chapter 6-10)
// https://go.dev/ref/spec#Switch_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

var x interface{}
var y interface{}

func main() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("it's an int")
	case bool:
		fmt.Println("it's a bool")
	case string:
		fmt.Println("it's a string")
	case float64:
		fmt.Println("it's a float")
	}

	switch y := 2; { // it can accept an initialization as in
	// for and if statements
	case y == 2:
		fmt.Println("it's an int")
	case y <= 2:
		fmt.Println("it's a bool")
	case y == 3:
		fmt.Println("it's a string")
	case y == 4:
		fmt.Println("it's a float")
	}
}
