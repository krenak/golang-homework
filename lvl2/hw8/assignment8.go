// Assignment 8 (Chapter 4)
// - Write expressions using operators (==, !=, >=, >, <=, <), and assign their values to variables.
// - Demonstrate these values.

package main

import (
	"fmt"
)

func main() {
	// solution includes parenthesis for fast pattern recognize
	x := 10 > 10
	y := 10 < 10
	z := 10 >= 10
	w := 10 <= 10
	u := 10 != 10
	v := 10 == 10

	// fmt.Println(x) <- my original and not elegant answer...
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(w)
	// fmt.Println(u)
	// fmt.Println(v)
	// solution resusmes in one line and shows their types
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", x, y, z, w, u, v)
}
