// strings
// https://go.dev/blog/strings

package main

import (
	"fmt"
)

func main() {
	s := "hello, playground"
	// sb := []byte(s) // slice of bytes
	for _, v := range s { // range gives character by character, not bytes
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}
	fmt.Println("")               // space in between
	for i := 0; i < len(s); i++ { // this for loop gives bytes by bytes
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
	// fmt.Printf("%v\n%T\n", sb, sb)
}
