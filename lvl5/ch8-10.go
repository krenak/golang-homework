/* Data structure - Maps: range and deleting (Chapter 8 - Episode 10)
https://go.dev/ref/spec#Map_types
https://www.youtube.com/watch?v=7a6I-GnqtSE

- Range: for k, v := range map { }
- To reiterate: maps *has no order* and a range will use a random order.
- Go Playground: https://play.golang.org/p/6zEMfIP-AE
- delete(map, key)
- Deleting a non-existent key does not return errors!
- Go Playground: https://play.golang.org/p/0uuIicU3Zz */

package main

import(
	"fmt"
)

func main() {
	x := map[int]string{
		123: "nice",
		98: "less nice",
		983: "this one is ok",
		18: "age to party",
	}
	fmt.Println(x, "\n")
	/* below, an example of for statement using map function to verify its
	range */
	for key, value := range x {
		fmt.Println(key, value)
	}
	// and to delete
	delete(x, 123)
	fmt.Println(x)
}
