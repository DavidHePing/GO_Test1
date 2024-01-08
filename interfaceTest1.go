package main

import "fmt"

type calculator interface {
	Add(a int, b int) int
	Minus(a int, b int) int
}

type newCalculator struct {
}

func (calculator *newCalculator) Add(a int, b int) int {
	return a + b
}

func (calculator *newCalculator) Minus(a int, b int) int {
	return a - b
}

func useCalculator(c calculator) {
	fmt.Println("calc Add (100, 1): ", c.Add(100, 1))
	fmt.Println("calc Minus (100, 1): ", c.Minus(100, 1))
}

func interfaceTest1() {
	c := &newCalculator{}
	useCalculator(c)
}
