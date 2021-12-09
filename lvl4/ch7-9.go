// Assignment 20 (Chapter 7 - episode 9)
// - Create a program that uses the switch statement, where the switch statement is a string variable with the identifier "esporteFavorito".
// - Solution: https://play.golang.org/p/4-iTPZwfEz
// https://go.dev/ref/spec#Switch_statements
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

var esporteFavorito interface{}
// var esporteFavorito1 interface{}

func main() {
	// simple switch implementation
//	esporteFavorito = "soccer"
//
//	switch esporteFavorito.(type) {
//	case string:
//		fmt.Println("How about Champions League this year?")
//	case int:
//		fmt.Println("Does your volleyball team got any gold medal?")
//	case float64:
//		fmt.Println("Dude, Celtics are over the roof this year, huh?")
//	}
	// funny thing is I did the correct answer before consult my notes...
	// in any case, here it is what I came up.
	switch esporteFavorito := "basketball"; {
	case esporteFavorito == "soccer":
		fmt.Println("How about Champions League this year?")
	case esporteFavorito == "volleyball":
		fmt.Println("Does your volleyball team got any gold medal?")
	case esporteFavorito == "basketball":
		fmt.Println("Dude, Celtics are over the roof this year, huh?")
	}
	// and here is her solution mixed with mine beforehand...
	esporteFavorito := "volleyball"

	switch esporteFavorito {
	case "soccer":
		fmt.Println("How about Champions League this year?")
	case "volleyball":
		fmt.Println("Does your volleyball team got any gold medal?")
	case "basketball":
		fmt.Println("Dude, Celtics are over the roof this year, huh?")
	}
}
