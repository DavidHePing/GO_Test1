package main

import "fmt"

type Car struct {
	name  string
	price int
}

func (car *Car) String() string {
	return fmt.Sprint("name: ", car.name, " price: ", car.price)
}

func (car *Car) ShowCarPrice() int {
	return car.price
}

type Tesla struct {
	Car
}

func inheritTest1() {
	tesla1 := &Tesla{Car{
		name:  "Tesla",
		price: 1000,
	}}

	fmt.Println("inherit String:", tesla1)
	fmt.Println("ShowCarPrice:", tesla1.ShowCarPrice())
}
