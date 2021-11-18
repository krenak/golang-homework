// Assignment 7 (chapter 4)
// Write a program that displays a number in decimal, binary, and hexadecimal.

package main

import (
	"fmt"
)

func main() {
	x := 356
	y := 124
	z := 1313

	fmt.Printf("%d <- decimal\n", x)
	fmt.Printf("%b <- binary\n", y)
	fmt.Printf("%#x <- hex\n", z) // include # to hex explicitation
}
