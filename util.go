package matrixgo

import "fmt"

func MIntPrint(m MInt) {
	var x, y int
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			fmt.Print(m[i][j])
			fmt.Printf("\t")

		}
		fmt.Println()
	}
}
