/* Data structure - Slice multi-dimensional (Chapter 8 - Episode 7)
https://go.dev/doc/effective_go#allocation_make
https://youtu.be/o3yoYGWqrDE

Multi-dimensional slices are slices that contain slices.
- They are like spreadsheets.
- [][]type
- Go Playground: https://play.golang.org/p/vKyHiG1GtM
- Just for fun: https://play.golang.org/p/ZSU_8eJ9Yp */

package main

import(
	"fmt"
)

func main() {
	ss1 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(ss1) /* printing the whole array of slices, similar to excel
	rows */
	fmt.Println(ss1[1][1]) /* here was selected the second value from its
	second row */

	// literal example
	ss := [][][][]int{

		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{10, 20, 30, 40, 50},
				},
			},

		[][][]int{
			[][]int{
				[]int{2, 4, 6, 8, 10},
				},
			[][]int{
				[]int{3},
				},
			},

	}
	fmt.Println(ss)
	fmt.Println(ss[1][0][0][2]) /* here are the result of the
	multi-dimensional slices, nested one into another */
}
