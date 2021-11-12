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
