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

// func main() { // My solution, which was not accurate, as
// required to convert unicode to ascii strings. But the first
// part was right.
// 	i := 0
// 	for i < 33 {
// 		i++
// 		for i >= 33 {
// 			for i > 120 {
// 				continue
// 			}
// 			fmt.Printf("%d\t %#x\t %#U\t\n", i, i, i)
// 			i++
// 		}
// 	}
// }

func main() { // Main solution provided by Ellen.
	// It diverges a bit from documentation (1), but I choose to
	// preserve it as it is in the video series.
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\n", string(i)) // (1) as here when is not
		// recommended to convert strings like that.
	}
}
