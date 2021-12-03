// Assignment 15 (Chapter 7 - episode 3)
// Create a loop using the syntax: for condition {}
// - Use it to demonstrate the years since you were born.
// - Solution: https://play.golang.org/p/qnFjiDJzLor

package main

import(
	"fmt"
)

func main() {
	// my solution went a bit ahead of what was asked...
	for i:=1989; i <= 2021; i++ {
		for i == 1989 {
			fmt.Println("The year I was born is", i,
				     "and since have pass these ")
			i++
			break
		}
		fmt.Printf("%d\n", i)
	}
	// and her solution was
	anoemqueeunasci := 1988
	anoatequaleuquerocontar := 2088

	for anoemqueeunasci <= anoatequaleuquerocontar {
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
