// Logical condition operators (Chapter 6-11)
// Extra assignment
// https://go.dev/ref/spec#Switch_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

func main() {
	// borrowed from extra assigment in this chapter
	// here are presented all logical conditions inserted in
	// Golang's if statement.
	// case 1 - true && true
	x := 3
	if (x == 2 && x == 3) {
		fmt.Println("x case #1 pass")
	}

	// case 2 - true && false
	y := 2
	if (y == 2 && y == 3) {
		fmt.Println("y case #2 pass")
	}

	// case 3 - true || true
	w := 3
	if (w == 2 || w == 3) {
		fmt.Println("w case #3 pass")
	}

	// case 4 - true || false
	z := 3
	if (z == 2 || z == 3) {
		fmt.Println("z case #4 pass")
	}

	// case 5 - !true
	u := 9
	if !(u == 0) {
		fmt.Println("u case #5 pass")
	}

}
