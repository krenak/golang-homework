// overflow
// https://pkg.go.dev/builtin#uint16

package main

import "fmt"

func main() {
	var i uint16
	i = 65535 // limit number for uint16
	fmt.Println(i)
	i++
	fmt.Println(i) // if 65536, it would be technical overflow for uint16 limit
	i++
	fmt.Println(i) // if 65537, it would be technical overflow for uint16 limit
}
