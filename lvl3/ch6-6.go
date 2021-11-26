// for statements (Chapter 6 - 6th episode)
// https://go.dev/ref/spec#For_statements
// https://go.dev/doc/effective_go#for
// - Surprise challenge!
// - FORMAT PRINTING:
//      - Decimal% D
//      - hexadecimal% # x
//      - Unicode% # u
//      - Tab \ t
//      - New Line \ N
// - Loop the numbers 33 to 122, and use Format
// Printing to demonstrate them as text / string.
// - Solution: https://play.golang.org/p/rem2whyzzz

package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 33 {
		i++
		for i >= 33 {
			for i > 120 {
				continue
			}
			fmt.Printf("%d\t %#x\t %#U\t\n", i, i, i)
			i++
		}
	}
}
