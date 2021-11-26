// If statement (Chapter 6)
// https://go.dev/ref/spec#If_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

func main() {
	if x := 10; !(x > 100) {
		//		^NOT conditional
		fmt.Println("Hello")
	}
}
