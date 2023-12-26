package main

import "fmt"

func delegateTest2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddsArray := delegateTest1_1(nums, func(i int) bool { return i%2 != 0 })
	fmt.Println(oddsArray)
	evenArray := delegateTest1_1(nums, func(i int) bool { return i%2 == 0 })
	fmt.Println(evenArray)
}

func delegateTest2_1(nums []int, validMethod func(int) bool) []int {
	var result []int

	for _, num := range nums {
		if validMethod(num) {
			result = append(result, num)
		}
	}

	return result
}
