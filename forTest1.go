package main

import "fmt"

func ForTest1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
}

func ForTestArray() {
	ary1 := [10]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for i := 0; i < len(ary1); i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println("")

	for num := range ary1 {
		fmt.Print(num, ", ")
	}
	fmt.Println("")

	map1 := map[int]int{
		1: 123,
		2: 456,
	}

	for key, value := range map1 {
		fmt.Println("key: ", key, "value: ", value)
	}
}
