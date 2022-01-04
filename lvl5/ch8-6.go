/* Data structure - Make (Chapter 8 - Episode 6)
https://go.dev/doc/effective_go#allocation_make
https://youtu.be/IMO5_ancK9w

Slices are made up of arrays.
- They are dynamic, can change size.
- Whenever this happens, a new array is created and the data is copied.
- It's convenient, but it comes at a computational cost.
- To optimize things, we can use make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the
limits of the underlying array; just assign it to a slice of itself. The
capacity of a slice, accessible by the built-in function cap, reports the
maximum length the slice may assume."
- len(x), cap(x)
- x[n] where n is greater than len is out of range. Use append.
- Append greater than cap modifies the underlying array.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is
resliced to accommodate the new elements. If it does not, the new underlying
array will be allocated."
- Effective Go.
- Go Playground: https://play.golang.org/p/e8GWzyEEL8 */

package main

import(
	"fmt"
)

func main() {
	slice1 := make([]int, 5, 10) // slice definitions
	slice1[0], slice1[1], slice1[2], slice1[3] = 1, 2, 3, 4
	// asssigning values to the slices' slots
	slice1 = append(slice1, 10) // expanding slice capacity using append
	fmt.Println(slice1, len(slice1), cap(slice1))
	/* Printing slice1 configs of values, lenght and capacity. Once its
	capacity is reached, the make function discards the previous array,
	creates a new one and allocate all values into there. Hence, its
	expanded once again. */
}

