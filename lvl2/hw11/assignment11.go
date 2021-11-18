// Assignment 11 (Chapter 4)
// - Create a variable of type string using a raw string literal.
// - Demonstrate it.
// - Solution: https://play.golang.org/p/RkpqPpRWuo

package main

import (
	"fmt"
)

var x string

func main() {
	x := `this
					is
								fine`
	fmt.Println(x)
}
