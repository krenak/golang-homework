// Assignment 16 (Chapter 7 - episode 4)
//- Create a loop using the syntax: for {}
//- Use it to demonstrate the years since you were born.
//- Solution: https://play.golang.org/p/dIbfdiuw0ms

package main

import(
	"fmt"
)

func main() {
	// now it seems the correct time to implement the __break__ argument.
	// let's see...
	for i:=1989; i <= 2021; i++ {
		for i == 1989 {
			fmt.Println("The year I was born is", i,
				     "and since have pass these ")
			i++
			break
		}
		fmt.Printf("%d\n", i)
	}
	// it wasn't... but it was pretty close. She mentioned that the main
	// goal was to use break argument.
	anoemqueeunasci := 1988
	anoatequaleuquerocontar := 2088
	for {
		if anoemqueeunasci > anoatequaleuquerocontar {
			break
		}
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
