/* Data structure - Slice of slices (Chapter 8 - Episode 4)
 https://go.dev/ref/spec#Making_slices_maps_and_channels
 https://gobyexample.com/slices
 https://youtu.be/G0rxcnojV_U

x[:], x[a:], x[:b], x[a:b]
- "a" is included;
- "b" is not.
- Example: magnetic head of a hard disk (watch, tape).
     - Off-by-one error.
- Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
- It is by slicing that an item is deleted from a slice. In practice:
     - x := append(x[:i], x[:i]...)
     - Go Playground: https://play.golang.org/p/xK2HwCqvwd
- Exercise: try to access all items in a slice *without* using range.
- Solution: https://play.golang.org/p/aUC9qVCobH */

package main

import(
	"fmt"
)

func main() {
	flavours := []string{"pepperoni", "mozzarela", "fourcheeses",
			     "marguerita"}
	slice := flavours[0:2]
	fmt.Println(slice)
	slice2 := flavours[1:len(flavours)]
	fmt.Println(slice2)
	// extra exercise: access all slice items *without* use range statement
	// slice3 := flavours[0:len(flavours)] <-- i did *with* a range uwu
	slice3 := flavours[ : ] // one way of doing it
	fmt.Println(slice3)
	fmt.Println(flavours[0], flavours[1], flavours[2], flavours[3])
	/* another way of doing it... */
	for i := 0; i < len(flavours); i++ { // last idea of doing it
		fmt.Println(flavours[i])
	}
	// deleting an slice item
	flavours = append(flavours[:2],flavours[3:]...)
	/* it was sectioning the third element appending a new slice from 0 to 
	2, and finishing from 3 to the end. */
	fmt.Println(flavours)
}
