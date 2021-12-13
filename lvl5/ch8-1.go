// Data structure - Arrays (Chapter 8 - Episode 1)
// https://go.dev/ref/spec#Array_types
// https://gobyexample.com/arrays
// https://youtu.be/i_3O4ooSVCM

package main

import (
	"fmt"
)

var x [5]int // even though they're arrays, the fact that one is 5
var y [6]int // and the other is 6 make them not compatible with each other

func main() { // arrays are not so useful, but knowing them are needed...
	x[0] = 1 // in 5 positions (0-4), "0" with integer "1" stored
	x[1] = 10 // as before, position "1" with integer "10" stored
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x)) // length of x array
	fmt.Println(len(y)) // length of y array
}
