// Logical condition operators (Chapter 6-11)
// https://go.dev/ref/spec#Switch_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

func main() {

	switch x.(type) {
	case int:
		fmt.Println("it's an int")
	}
}
