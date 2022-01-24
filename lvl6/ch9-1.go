/* Assignment 14 (Chapter 9 - episode 1)
- Using a composite literal:
      - Create an array that supports 5 values of type int
      - Assign values to your indexes
- Use range and show the values of the array.
- Using format printing, demonstrate the array type.
- Solution: https://play.golang.org/p/tpWDzzlO2l */

package main

import(
	"fmt"
)

func main() {
	array := [5]int{1,3,5,7,9}
	fmt.Println(array)
	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Printf("%T", array)
}


