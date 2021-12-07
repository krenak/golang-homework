// Assignment 18 (Chapter 7 - episode 7)
// - Using the previous solution, add the else if and else options.
// - Solution: https://play.golang.org/p/_cO7E-yL0o
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

func main() {
	// simple switch implementation
	i := 100
	switch {
	case i >= 300:
		fmt.Println("It's f*cking hot out here!!")
	case i < 300 && i != 100:
		fmt.Println("Hmm, it's chilling now...")
	case i == 100:
		fmt.Println("Also, it is getting fricking cold.")
	}
}
