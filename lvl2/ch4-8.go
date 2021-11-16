// IOTA
// https://golang.org/ref/spec#Iota

package main

import (
	"fmt"
)

const (
	a = iota
	_ = iota + 10 // dumping value using underscore amd repeting forward expression
	c             // using iota for all values only identifying the previous var
	x
	_
	z
)

func main() {
	fmt.Println(a, c, x, z)
}
