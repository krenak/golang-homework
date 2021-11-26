// for statements (Chapter 6 - 4th episode)
// https://go.dev/ref/spec#For_statements
// https://go.dev/doc/effective_go#for

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 { // module which gives the rest of i, then, if the rest is 0, it means, in this case, that is a even number.
			continue // and implementing the continue statement to interrupt the section which it is included. In this example, it turns down all even values and allow to odd numbers to be printed.
		}
		fmt.Println(i)
	}
}
