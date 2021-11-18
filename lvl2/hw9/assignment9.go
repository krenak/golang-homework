// Assignment 9 (Chapter 4)
// - Create typed and untyped constants.
// - Demonstrate your values.
// Solution: https://play.golang.org/p/eWnKI59ual

package main

import (
	"fmt"
)

const (
	a      = iota
	x int  = 10
	y bool = true
	z      = 13 > 19
)

func main() {
	fmt.Printf("%v, %v, %v, %v\n", a, x, y, z)
}
