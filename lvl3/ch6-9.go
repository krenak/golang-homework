// Switch statement (Chapter 6)
// https://go.dev/ref/spec#Switch_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

func main() {
	// example a: simple cases
	x := 10
	switch { // default statement is eq to true
	case x < 5: // in case the first condition is true, the
		// switch statement stops once it meets this condition.
		fmt.Println("x is smaller than 5")
	case x > 5:
		fmt.Println("x is bigger than 5")
	case x == 5:
		fmt.Println("x is equal to 5")
	}
	// example b: cases with fallthrough and default action
	i := "Nobody"
	switch i { // it also accepts a var as check value.
	case "Bobby":
		fmt.Println("Bobby is on office.")
		fallthrough // once this case is true, fallthrough
		// allows the following condition to be also true.
	case "Eddy":
		fmt.Println("Eddy is on office as well.")
	case "Joe":
		fmt.Println("Joe is on office.")
	default: // default action in case there is no match
		// case before.
		fmt.Println("Office is empty.")
	}
	// example b: expression cases

	z := 1
	switch z { // it also accepts a var as check value.
	case 1, 2:
		fmt.Println("First squad is on.")
		fallthrough // once this case is true, fallthrough
		// allows the following condition to be also true.
	case 3, 4:
		fmt.Println("But not for long...")
	case 5, 6:
		fmt.Println("Second squad is on.")
	default: // default action in case there is no match
		// case before.
		fmt.Println("Squas is out.")
	}
}
