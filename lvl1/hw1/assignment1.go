// - Using the short declaration operator, assign these values to variables with identifiers "x", "y", and "z".
//      1. 42
//      2. "James Bond"
//      3. true
// - Now demonstrate the values in these variables, with:
//      1. A single print statement
//      2. Multiple print statements
// - Solution: https://play.golang.org/p/yYXnWXIQNa

package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Print("The actor of ", y, "'s movies have an age average of ", x, " which is ", z, ".")
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// https://play.golang.org/p/dglPr_cqwTb
