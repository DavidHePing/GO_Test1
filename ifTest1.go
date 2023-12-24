package main

import "fmt"

func ifTest1() {
	test1(11)
	test1(10)
	test1(9)
}

func test1(num int) {
	if num > 10 {
		fmt.Println("greate than 10")
	} else if num == 10 {
		fmt.Println("equal 10")
	} else {
		fmt.Println("less than 10")
	}
}

func ifTest2() {
	test2(2, 2)
	test2(1, 2)
	test2(0, 2)
}

func test2(num1 int, num2 int) {
	if a := test21(num1, num2); a > 3 {
		fmt.Println("a: ", a, "is over 3")
	} else if a == 3 {
		fmt.Println("a: ", a, "is over 3")
	} else {
		fmt.Println("a: ", a, "is over 3")
	}
}

func test21(a int, b int) int {
	return a + b
}
