package main

import (
	"fmt"
	"sync"
)

func syncMapTest1() {
	m := sync.Map{}
	m.Store("a", 1)
	m.Store(1, "b")

	fmt.Println(m.Load("a"))
	fmt.Println(m.Load(1))
	fmt.Println(m.Load("1"))
}
