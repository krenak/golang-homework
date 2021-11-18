// bitwise and IOTA
// https://golang.org/ref/spec#Iota
// https://play.golang.org/p/7MOnbhx4R4
// https://splice.com/blog/iota-elegant-constants-golang/
// https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

package main

import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	x := 1      // moving bits backward
	y := x << 1 // moving bits forward
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	// ^ previous session
	// ellen's session
	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
