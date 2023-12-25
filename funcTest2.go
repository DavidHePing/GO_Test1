package main

import "fmt"

func funcTest2() {
	funcTest2_1(1, 2, 3, 4, 5, 6, 7)
}

func funcTest2_1(nums ...int) {
	fmt.Println(nums)

	for _, num := range nums {
		fmt.Print(10+num, ", ")
	}

	fmt.Println("")
}
