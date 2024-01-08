package main

import (
	"Go_Test1/weekday"
	"fmt"
)

func enumTest1() {
	fmt.Println(weekday.Sunday.GetDesciption(), weekday.Sunday)
	fmt.Println(weekday.Monday.GetDesciption(), weekday.Monday)
	fmt.Println(weekday.Thursday.GetDesciption(), weekday.Thursday)
	fmt.Println(weekday.Saturday.GetDesciption(), weekday.Saturday)
}
