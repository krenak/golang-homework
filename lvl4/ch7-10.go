// Assignment 21 (Chapter 7 - episode 10)
//- Write down (by hand) the result of the expressions:
//     - fmt.Println(true && true)
//     - fmt.Println(true && false)
//     - fmt.Println(true || true)
//     - fmt.Println(true || false)
//     - fmt.Println(!true)
//- Ninja Level 3! Congratulations!
// https://go.dev/ref/spec#Switch_statements
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

var esporteFavorito interface{}

func main() {
	esporteFavorito := "volleyball"

	switch esporteFavorito {
	case "soccer":
		// fmt.Println("How about Champions League this year?")
		fmt.Println(true && true)
		fmt.Println(true && false)
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(!true)
	case "volleyball":
		// fmt.Println("Does your volleyball team got any gold medal?")
		fmt.Println(true && true)
		fmt.Println(true && false)
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(!true)
	case "basketball":
		// fmt.Println("Dude, Celtics are over the roof this year, huh?")
		fmt.Println(true && true)
		fmt.Println(true && false)
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(!true)

	}
}
