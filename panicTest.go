package main

import "fmt"

//panic only for serious error, normal error would use Error
func panicTest1() {
	// panicTest1_1("")
	panicTest1_2("")
}

func panicTest1_1(st string) {
	if st == "" {
		panic("Panic!!!!")
	}
}

func panicTest1_2(st string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("keep going")
		}
	}()

	if st == "" {
		panic("Panic!!!!")
	}
}
