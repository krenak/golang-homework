// Data structure - Slices (Chapter 8 - Episode 2)
// https://go.dev/ref/spec#Composite_literals
// https://gobyexample.com/slices
// https://youtu.be/MMzTlWZ9Gjw

package main

import(
	"fmt"
)

func main() {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	slice = append(slice, 6) // using append we can input another
				 // index value to the slice
	fmt.Println(slice)

	fmt.Println(slice[3]) // slices are flexible once its value is
			      // mutable
	slice[3] = 878457
	fmt.Println(slice[3])

	slice = append(slice, 8) // using append we can input another
				 // slot to the slice, named "8"
	fmt.Println(slice)

}
