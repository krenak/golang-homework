// Assignment 14 (Chapter 7 - episode 2)
//- Put on screen: The unicode code point of all capital letters of the alphabet, three times each.
//- For example:
//     65
//         U+0041 'A'
//         U+0041 'A'
//         U+0041 'A'
//     66
//         U+0042 'B'
//         U+0042 'B'
//         U+0042 'B'
//     ...and so on.
//- Solution: https://play.golang.org/p/QlP6nVchq-


package main

import(
	"fmt"
)

func main() {
	// the range of ASCII's letters in all caps is from 65 to 90, then
	for i:=65; i <= 90; i++ { // (i >= 65 && i <= 90); i++ <-- my
		// overbloatted solution...
		fmt.Printf("%d\n %#U\n %#U\n %#U\n", i, i, i, i)
	}
}
