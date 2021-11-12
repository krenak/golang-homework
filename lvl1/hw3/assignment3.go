// - Using the solution from the previous exercise:
// 		In package-level scope, assign the following values to the variables:
// 		for "x" assign 42
// 		for "y" assign "James Bond"
// 		for "z" set true
// - In the main function:
// 		Use fmt.Sprintf to assign all these values to a single variable. Make this string type assignment to a variable named "s" using the short declaration operator.
// 		Demonstrate the variable "s".

package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v, %v, %v", x, y, z)
	fmt.Println(s)
}

// https://play.golang.org/p/MWZsNe1MD3Z
