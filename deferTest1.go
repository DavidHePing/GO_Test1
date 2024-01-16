package main

import (
	"fmt"
	"sync"
)

func deferTest1() {
	deferTest1_1(1)
	deferTest1_1(2)
	deferTest1_1(3)
	deferTest1_1(4)
}

func deferTest1_1(num int) {
	defer fmt.Println("Is End!!!")   //do second
	defer fmt.Println("Nun is", num) //do first

	fmt.Println("deferTest1")

	if num == 1 {
		fmt.Println("do something with 1")
		return
	}

	if num == 2 {
		fmt.Println("do something with 2")

		return
	}

	if num == 3 {
		fmt.Println("do something with 3")

		return
	}
}

func deferTest2() {
	fmt.Println("deferTestForWrong")
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}

func deferTest3() {
	fmt.Println("deferTestForRight")
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		num := i
		go func() {
			defer wg.Done()
			fmt.Println(num)
		}()
	}

	wg.Wait()
}

func defer_panic_test1() {
	defer fmt.Println("Is End!!!")
	fmt.Println("defer_panic")

	panic("Is Panic!!!!!")
}

func defer_panic_test2() {
	fmt.Println("panic without defer")

	panic("Is Panic!!!!!")
	fmt.Println("Is End!!!")
}
