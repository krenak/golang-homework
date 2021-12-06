// Assignment 17 (Chapter 7 - episode 6)
// - Create a program that demonstrates how the if statement works.
// - Solution: https://play.golang.org/p/OGPgTJZbiy
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

func main() {
	// my solution was wrong, but I chose to keep it as a reminder.
	for i := 10; i <= 100; i++ {
		if (i % 4 != 0 && i % 4 == 0) {
			continue
		}
		fmt.Println(i)
	}
}
