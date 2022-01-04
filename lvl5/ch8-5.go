/* Data structure - Append (Chapter 8 - Episode 5)
https://pkg.go.dev/builtin#append
https://youtu.be/MbvABKiAABA

Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → unfold, unroll
- Nome oficial: enumeration
- Go Playground: https://play.golang.org/p/RpkDCTumpT */

package main

import(
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{9, 10, 11, 12}
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1 = append(slice1, 5, 6, 7, 8)
	fmt.Println(slice1)

	slice2 = append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
