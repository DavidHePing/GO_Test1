package main

import (
	"container/list"
	"fmt"
)

func linkedList_test2() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ",")
	}
	fmt.Println()

	fmt.Println("Remove front")
	l.Remove(l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ",")
	}
	fmt.Println()
}
