package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func float_test1() {
	num := 10.0 / 3.0
	fmt.Println(num)

	num2, _ := decimal.NewFromString("10")
	num3, _ := decimal.NewFromString("3")
	fmt.Println(num2.Div(num3))
}
