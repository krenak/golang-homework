// Assignment 18 (Chapter 7 - episode 7)
// - Create a program that uses the switch statement, where the switch statement is a string variable with the identifier "esporteFavorito".
// - Solution: https://play.golang.org/p/4-iTPZwfEz
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
