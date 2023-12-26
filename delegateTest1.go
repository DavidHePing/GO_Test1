package main

import "fmt"

type isValidNum func(int) bool

func delegateTest1() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(delegateTest1_1(nums, isOdd))
	fmt.Println(delegateTest1_1(nums, isEven))
}

func delegateTest1_1(nums []int, validMethod isValidNum) []int {
	var result []int

	for _, num := range nums {
		if validMethod(num) {
			result = append(result, num)
		}
	}

	return result
}

func isOdd(num int) bool {
	return num%2 != 0
}

func isEven(num int) bool {
	return num%2 == 0
}
