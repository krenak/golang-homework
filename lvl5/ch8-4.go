// Data structure - Slice of slices (Chapter 8 - Episode 4)
// https://go.dev/ref/spec#Making_slices_maps_and_channels
// https://gobyexample.com/slices
// https://youtu.be/G0rxcnojV_U

//x[:], x[a:], x[:b], x[a:b]
//- "a" is included;
//- "b" is not.
//- Example: magnetic head of a hard disk (watch, tape).
//     - Off-by-one error.
//- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
//- It is by slicing that an item is deleted from a slice. In practice:
//     - x := append(x[:i], x[:i]...)
//     - Go Playground: https://play.golang.org/p/xK2HwCqvwd
//- Exercise: try to access all items in a slice *without* using range.
//- Solution: https://play.golang.org/p/aUC9qVCobH

package main

import(
	"fmt"
)

func main() {
	flavours := []string{"pepperoni", "mozzarela", "fourcheeses", "marguerita"}
	slice := flavours[0:2]
	fmt.Println(slice)
	slice2 := flavours[1:len(flavours)]
	fmt.Println(slice2)
	// extra exercise: access all slice items *without* use range statement.
	slice3 := flavours[0:len(flavours)]
	fmt.Println(slice3)
}
