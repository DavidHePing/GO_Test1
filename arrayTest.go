package main

import (
	"fmt"
)

func ArrayTestPrint() {
	list1 := []int{
		1, 2, 3, 4,
	}
	var list2 []int

	var matrix [2][2]bool

	list1 = append(list1, 5, 6, 7)
	list2 = append(list2, list1[:]...)
	copy(list2, list1)

	fmt.Println(list1)
	fmt.Println("list1: ", list1)
	fmt.Println("list2: ", list2)
	fmt.Println("matrix ", matrix)
}
