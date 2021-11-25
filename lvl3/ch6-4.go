package main

import (
	"fmt"
)

func main() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println("tá chuveno")
		} else {
			fmt.Println("eita que parou de chuve")
			break
		}
	}
	fmt.Println("*toninho abre o guarda-chuva e sai andando*")
}
