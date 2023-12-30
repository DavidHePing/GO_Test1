package main

import (
	"fmt"
	"strings"
)

func stringTest1() {
	stringDefault()
	stringConcatFmt()
	stringJoin()
}

func stringDefault() {
	var st1 string
	if st1 == "" {
		fmt.Println("default string == \"\"!!! ", st1 == "")
	}
}

func stringConcatFmt() {
	st1 := "1"
	st2 := "2"
	st3 := "3"

	result := fmt.Sprint("stringConcatFmt: First: ", st1, " Second: ", st2, " Third: ", st3)
	fmt.Println(result)
}

func stringJoin() {
	slice1 := []string{"1", "2", "3"}
	result := strings.Join(slice1, ", ")
	fmt.Println("stringJoin: ", result)
}
