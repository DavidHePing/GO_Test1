package main

import (
	"fmt"
)

func SliceTestPrint() {
	list1 := []int{
		1, 2, 3, 4,
	}
	fmt.Println("list1: ", list1)

	list2 := make([]int, 4)
	fmt.Println("list2: ", list2)
	list2[3] = 12
	fmt.Println("list2: ", list2)

	var matrix [2][2]bool
	fmt.Println("matrix ", matrix)

	list1 = append(list1, 5, 6, 7)
	fmt.Println("list1 append: ", list1)
	list2 = append(list2, list1[:]...)
	fmt.Println("list2 append list1: ", list2)
	//copy element from src.slice
	copy(list2, list1)
	fmt.Println("list2 copy list1: ", list2)

}
