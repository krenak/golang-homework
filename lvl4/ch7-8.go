// Assignment 19 (Chapter 7 - episode 8)
// - Create a program that uses the switch statement, with no switch statement specified.
// - Solution: https://play.golang.org/p/TyIGp4Hi8B
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
