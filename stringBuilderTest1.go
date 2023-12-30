package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringBuilderTest1() {
	slice1 := []string{}

	for i := 1; i <= 10; i++ {
		slice1 = append(slice1, strconv.Itoa(i))
	}

	fmt.Println(slice1)

	var builder strings.Builder

	for _, st := range slice1 {
		builder.WriteString(st)
	}

	fmt.Println(builder.String())
}
