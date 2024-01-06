package structTest2

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
func (car *Car) String() string {
	var builder strings.Builder
	builder.WriteString("{")
	builder.WriteString("name: ")
	builder.WriteString(car.name)
	builder.WriteString(" ,price: ")
	builder.WriteString(strconv.Itoa(car.price))
	builder.WriteString("}")

	return builder.String()
}

type Buyer struct {
	car1 *Car
	car2 *Car
	name string
	age  int
}

func (buyer *Buyer) String() string {
	var builder strings.Builder
	builder.WriteString("car1: ")
	builder.WriteString(buyer.car1.String())
	builder.WriteString(" ,car2: ")
	builder.WriteString(buyer.car2.String())
	builder.WriteString(" ,name: ")
	builder.WriteString(buyer.name)
	builder.WriteString(" ,age: ")
	builder.WriteString(strconv.Itoa(buyer.age))

	return builder.String()
}

func StructTestRefType() {
	car1 := &Car{name: "Tesla", price: 100}
	fmt.Println("Override String():", car1)
	car1.name = "Toyota"
	fmt.Println("assign struct value", car1)
	buyer1 := &Buyer{
		car1: &Car{
			name:  "Toyota",
			price: 50,
		},
		car2: &Car{
			name:  "Tesla",
			price: 100,
		},
		name: "David",
		age:  18,
	}

	fmt.Println(buyer1)

}
