package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Car struct {
	name  string
	price int
}

// struct method can only decalre in same package
func (car Car) String() string {
	var builder strings.Builder
	builder.WriteString("name: ")
	builder.WriteString(car.name)
	builder.WriteString(" ,price: ")
	builder.WriteString(strconv.Itoa(car.price))

	return builder.String()
}

func structTest1() {
	car1 := Car{"Tesla", 100}
	fmt.Println("Override String():", car1)
}
