package main

import "fmt"

func sliceTest2() {
	li := []int{1}

	fmt.Println(append(li, 3, 4)) //append would decide if create a new array or not
	fmt.Println(li)

	li = append(li, 2)
	fmt.Println(li)
}
