// numeric types
// https://golang.org/ref/spec#Numeric_types

package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e" // byte allocation with 1 byte
	b := "é" // byte allocation with 2 byte
	c := "香" // byte allocation with 3 byte
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a) // slice showing "a" byte allocation
	e := []byte(b) // slice showing "b" byte allocation
	f := []byte(c) // slice showing "c" byte allocation
	fmt.Printf("%v, %v, %v\n", d, e, f)

	g := 10   // integer
	h := 10.1 // float point
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)

	fmt.Println(runtime.GOOS)   // identify Operational System
	fmt.Println(runtime.GOARCH) // identify architecture
}
