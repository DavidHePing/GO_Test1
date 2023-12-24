package main

import "fmt"

func ArrayTestPrint() {
	ary1 := [4]int{}
	fmt.Println("ary1: ", ary1)
	ary1[3] = 12
	fmt.Println("ary1: ", ary1)
	matrix := [4][4]int{}
	fmt.Println("matrix", matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = i*10 + j
		}
	}

	fmt.Println("matrix", matrix)
}
