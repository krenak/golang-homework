/* Data structure - Slice multi-dimensional (Chapter 8 - Episode 8)
https://pkg.go.dev/builtin@go1.17.5#append
https://youtu.be/dRNNC7VpztE

We've already seen all this here:
- Every slice has an underlying array.
- A slice is: a pointer/address to an array, plus len and cap (which is the len
to array).
- Example:
     - x := []int{...numbers}
     - y := append(x[:i], x[:i]...)
     - pkg/builtin/#append: "If it has sufficient capacity, the destination is
     resliced to accommodate the new elements. If it does not, the new
     underlying array will be allocated."
     - That is, y uses the same underlying array as x.
     - Which gives us an unexpected result.
- In other words, it's good to know in advance so that you don't have to learn
by force.
- Go Playground: https://play.golang.org/p/BBJLuIjU_i */

package main

import(
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append(slice1[:2], slice1[4:])
	fmt.Println(slice2)
}
