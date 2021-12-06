// Assignment 16 (Chapter 7 - episode 5)
//- Demonstrate the remainder by dividing by 4 of all numbers between 10 and 100
//- Solution: https://play.golang.org/p/zcEsXqnBr8
// https://github.com/ellenkorbes/aprendago/blob/master/OUTLINE.md#loops-break--continue

package main

import(
	"fmt"
)

func main() {
	// my solution was wrong, but I chose to keep it as a reminder.
	// for i := 10; i <= 100; i++ {
	//	if (i % 4 != 0 && i % 4 == 0) {
	//		continue
	//	}
	//	fmt.Println(i)
	// }
	// and her solution was way simpler
	for i :=10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
	// Also, I quite forgot or got it not right how to proceed with the
	// module thing... I think I got it now, finally.
}
