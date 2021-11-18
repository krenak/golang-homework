// Assignment 12 (Chapter 4)
// - Using iota, create 4 constants whose values are the next 4 years.
// - Demonstrate these values.
// - Solution: https://play.golang.org/p/zRBEOoRvo4

package main

import (
	"fmt"
)

const (
	// x = 2022 + iota <- my original solution giving up one year, unfortunately
	_ = 2021 + iota
	x
	y
	z
	w
)

func main() {
	// fmt.Printf("%v\t %v\t %v\t %v\t\n", x, y, z, w) <- my original solution with too much info
	fmt.Println(x, y, z, w)
}
