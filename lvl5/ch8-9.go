/* Data structure - Maps: intro (Chapter 8 - Episode 9)
https://go.dev/ref/spec#Map_types
https://youtu.be/clobcQqAgos

- Uses key:value format.
- E.g. name and phone
- Excellent performance for lookups.
- map[key]value{ key: value }
- Access: m[key]
- Key without value returns zero. This can bring problems.
- To check: comma ok idiom.
     - v, ok := m[key]
     - ok it's a boolean, true/false
- In practice: if v, ok := m[key]; OK { }
- To add an item: m[v] = value
- Maps *no order.*
- Go Playground: https://play.golang.org/p/JXDdJan8Ev'' */

package main

import(
	"fmt"
)

func main() {
	friends := map[string]int{
		"alfredo": 5555123,
		"joana": 99995526,
	}
	fmt.Println(friends)
	fmt.Println(friends["joana"], "\n")
	// below, an if statement using map function to print values
	if x, ok := friends["jack"]; !ok {
		fmt.Println("does not exist")
	} else {
		fmt.Println(x)
	}
}

