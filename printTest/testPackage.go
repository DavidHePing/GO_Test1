package printTest

import (
	"fmt"
)

func TestPrint() {
	num1 := 123
	num2 := 456
	fmt.Print("Num1: ", num1, " Num2: ", num2)
}

func PrintTestLowercase() {
	fmt.Println("123")
}

// private method, can not be used in antoher package
func printTestLowercase() {
	fmt.Println("123")
}
