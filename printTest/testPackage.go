package printTest

import (
	"fmt"
	"strconv"
	"strings"
)

type TestObj struct {
	name string
	no   int
}

func (t TestObj) String() string {
	var builder strings.Builder
	builder.WriteString("name: ")
	builder.WriteString(t.name)
	builder.WriteString(" no: ")
	builder.WriteString(strconv.Itoa(t.no))

	return builder.String()
}

func TestPrint() {
	num1 := 123
	num2 := 456
	fmt.Println("TestPrint!!! Num1: ", num1, " Num2: ", num2)
	num1, num2 = num2, num1
	fmt.Println("Change 2 variable!!! Num1: ", num1, " Num2: ", num2)
	fmt.Println("10/2", 10/2, "11/2", 11/2)
	fmt.Println("11%10", 11%10, "1%10", 1%10)
}

func PrintTestLowercase() {
	fmt.Println("123")
}

// private method, can not be used in antoher package
func printTestLowercase() {
	fmt.Println("123")
}
