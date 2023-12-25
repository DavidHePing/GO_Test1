package main

import "fmt"

func deferTest1() {
	deferTest1_1(1)
	deferTest1_1(2)
	deferTest1_1(3)
	deferTest1_1(4)
}

func deferTest1_1(num int) {
	defer fmt.Println("Is End!!!")   //do second
	defer fmt.Println("Nun is", num) //do first

	if num == 1 {
		fmt.Println("do something with 1")
		return
	}

	if num == 2 {
		fmt.Println("do something with 2")

		return
	}

	if num == 3 {
		fmt.Println("do something with 3")

		return
	}
}
