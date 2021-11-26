// for statements (Chapter 6 - 4th episode)
// https://go.dev/ref/spec#For_statements
// https://go.dev/doc/effective_go#for

package main

import (
	"fmt"
)

func main() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println("tá chuveno")
		} else {
			fmt.Println("eita que parou de chuve")
			break // it breaks the loop
		}
	}
	fmt.Println("*toninho abre o guarda-chuva e sai andando*")
}
