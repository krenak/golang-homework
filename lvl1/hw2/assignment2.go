// - Use var to declare three variables. They must have package-level scope. Do not assign values to these variables. Use the following identifiers and types for these variables:
// 		Identifier "x" must have type int
// 		Identifier "y" must have string type
// 		Identifier "z" must have bool type
// - In the main function:
// 		Demonstrate the values of each identifier
// 		The compiler assigned values to these variables. What are these values called?

package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
}

// https://play.golang.org/p/PmTebAQr0xu
