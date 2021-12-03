// Assignment 15 (Chapter 7 - episode 3)
// Create a loop using the syntax: for condition {}
// - Use it to demonstrate the years since you were born.
// - Solution: https://play.golang.org/p/qnFjiDJzLor

package main

import(
	"fmt"
)

func main() {
	// the range of ASCII's letters in all caps is from 65 to 90, then
	for i:=65; i <= 90; i++ { // (i >= 65 && i <= 90); i++ <-- my
		// overbloatted solution...
		fmt.Printf("%d\n %#U\n %#U\n %#U\n", i, i, i, i)
	}
}
