// Numeric systems
// https://pkg.go.dev/fmt

package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Printf("%d\t %b <- binary, %#x <- hex\n", x, x, x)
}
