// Assignment 13 (Chapter 7 - episode 1)
// - Put on the screen: all numbers from 1 to 10000.
// - Solution: https://play.golang.org/p/MkdZiDW8SQ

package main

import(
	"fmt"
)

func main() {

	for x:=1; (x >= 1 && x <= 1000); x++ {
		fmt.Println(x)
	}
}
