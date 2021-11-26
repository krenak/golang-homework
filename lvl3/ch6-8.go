// If/Else statement (Chapter 6)
// https://go.dev/ref/spec#If_statements
// https://go.dev/ref/spec#Operators

package main

import (
	"fmt"
)

func main() { // the corret structure is if {elseif{else{}}}.
	if x := 10; x > 100 {
		//		^NOT conditional
		fmt.Println("x is bigger than 100")
	} else if x == 10 {
		fmt.Println("x is equal to 100")
	} else {
		fmt.Println("x is smaller than 100")
	}
}
