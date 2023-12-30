package main

import (
	"fmt"
	"strconv"
)

func strconvTest1() {
	itoa()
	IntToBinary()
	BoolToString()
}

func itoa() {
	num1 := 123
	num2 := 456

	result := strconv.Itoa(num1) + strconv.Itoa(num2)
	fmt.Println("itoa!", result)
}

func IntToBinary() {
	number := 42
	st2 := strconv.FormatInt(int64(number), 2)
	fmt.Println("IntToBinary! source num :", number, " after convert: ", st2)
}

func BoolToString() {
	isTrue := true
	isFalse := false

	stTrue := strconv.FormatBool(isTrue)
	stFalse := strconv.FormatBool(isFalse)

	fmt.Println("Bool To String!", stTrue, stFalse)
}
