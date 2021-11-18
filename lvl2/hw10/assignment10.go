// Assignment 10 (Chapter 4)
// - Create a program that:
//      - Assign an int value to a variable
//      - Show this value in decimal, binary and hexadecimal
//      - Shift the bits of this variable 1 to the left, and assign the result to another variable
//      - Demonstrate this other variable in decimal, binary and hexadecimal
// - Solution: https://play.golang.org/p/IiwgT0v3Mp

package main

import (
	"fmt"
)

var x int

func main() {
	x := 55
	fmt.Printf("%d, %b, %#x\n", x, x, x)
	y := x >> 1
	fmt.Printf("%d, %b, %#x\n", y, y, y)
}
