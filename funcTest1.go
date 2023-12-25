package main

import (
	"fmt"
	"strconv"
)

func funcTest1() {
	fmt.Println(funcTest1_1(1, "2"))
	fmt.Println(funcTest1_2(1))
	a, b, c := funcTest1_2(10)
	fmt.Println(a, b, c)
	fmt.Println(funcTest1_3(100))
	fmt.Println(funcTest1_4(100))
}

func funcTest1_1(num1 int, st1 string) string {
	return strconv.Itoa(num1) + st1
}

func funcTest1_2(num1 int) (int, int, int) {
	return num1 + 1, num1 + 2, num1 + 3
}

func funcTest1_3(num1 int) (a int, b int, c int) {
	return
}

func funcTest1_4(num1 int) (a int, b int, c int) {
	a = num1 + 10
	b = num1 + 20
	c = num1 + 30
	return
}
