// Assignment 17 (Chapter 7 - episode 6)
// - Create a program that demonstrates how the if statement works.
// - Solution: https://play.golang.org/p/OGPgTJZbiy
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

func main() {
	// simple example that enforces the boolean nature of if statement.
	i := 23
	if i >= 10 {
		fmt.Printf("%d", i)
	}
	// other interesting snippet I've found in Effective Go
	// if err := file.Chmod(0664); err != nil {
	//	log.Print(err)
	//	return err
	// }
}
