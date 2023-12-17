package main

import (
	"Go_Test1/arrayTest"
	"fmt"
)

func main() {

	list1 := []int{
		1, 2, 3, 4,
	}
	var list2 []int

	var matrix [2][2]bool

	list1 = append(list1, 5, 6, 7)
	list2 = append(list2, list1[:]...)
	copy(list2, list1)

	fmt.Printf("list1 %v\n", list1)
	fmt.Printf("list2 %v\n", list2)
	fmt.Printf("matrix %v\n", matrix)

	arrayTest.PrintArray()
	arrayTest.PrintTestLowercase()

}
