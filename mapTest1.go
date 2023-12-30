package main

import "fmt"

func mapTest1() {
	mapTest1_1()
	mapAddorUpdate()
}

func mapTest1_1() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"f": 4,
	}

	fmt.Println("map init!", m)
	value, isExist := m["a"]
	fmt.Println("map contains a", value, isExist)
	value, isExist = m["z"]
	fmt.Println("map contains z", value, isExist)
}

func mapAddorUpdate() {
	fmt.Println("--------")
	fmt.Println("map add or udpate")
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	fmt.Println(m)
	m["f"] = 5
	fmt.Println("add f:5", m)
	m["a"] = 100
	fmt.Println("update a:100", m)
}
