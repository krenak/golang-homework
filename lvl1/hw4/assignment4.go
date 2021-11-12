// - Create a type. The underlying type must be int.
// - Create a variable for this type, with the identifier "x", using the var keyword.
// - In the main function:
//      1. Show the value of the variable "x"
//      2. Demonstrate the type of variable "x"
//      3. Assign 42 to variable "x" using the operator "="
//      4. Show the value of the variable "x"
// - For the adventurous: https://golang.org/ref/spec#Types
// - Now we are almost ninjas lvl 1!
// - Solution: https://play.golang.org/p/snm4WuuYmG

package main

import (
	"fmt"
)

type ninja int

var x ninja

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

//
