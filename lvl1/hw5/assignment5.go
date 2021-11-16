// - Using the solution from the previous exercise:
//      1. In the package-level scope, using the var keyword, create a variable with the identifier "y". The type of this variable must be the underlying type of the type you created in the previous exercise.
//      2. In the main function:
//          1. This should already be done:
//              1. Show the value of the variable "x"
//              2. Demonstrate the type of variable "x"
//              3. Assign 42 to variable "x" using the operator "="
//              4. Show the value of the variable "x"
//          2. Now do too:
//              1. Use cast to type the value of the variable "x" into its underlying type and, using the "=" operator, assign the value of "x" to "y"
//              2. Demonstrate the value of "y"
//              3. Demonstrate the type of "y"
// - Solution: https://play.golang.org/p/uq81T_fasP

package main

import (
	"fmt"
)

type ninja int

var x ninja
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

// https://play.golang.org/p/dNV1AKCL5W7
