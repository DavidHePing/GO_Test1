package main

import "fmt"

func switchTest1() {
	switchTest1_1(0)
	switchTest1_1(2) //do next
	switchTest1_1(4)
	switchTest1_1(7) //do nothing
	switchTest1_1(8)
}

func switchTest1_1(num int) {
	switch num {
	case 0:
		fmt.Println("Is 0 !!!")
	case 2:
		fallthrough
	case 3:
		fmt.Println("num:", num, " fallthrough, Is 2~3")
	case 4, 5, 6:
		fmt.Println("Is  4~ 6")
	case 7:
	default:
		fmt.Println("Is over 7")
	}
}
