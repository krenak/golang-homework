// Assignment 18 (Chapter 7 - episode 7)
// - Using the previous solution, add the else if and else options.
// - Solution: https://play.golang.org/p/_cO7E-yL0o
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

var z int

func main() {
	// if else implementation with a var intrude...
	i := z
	if i < 0 {
		fmt.Println("This number is not a valid and rational number.")
	} else if i >= 0 {
		fmt.Println("This number is valid.") // interesting result.
	} else {
		fmt.Println("This is not a number.")
	}
}
