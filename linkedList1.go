package main

import (
	"container/list"
	"fmt"
)

func linkedList_test1() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	fmt.Println("Front", l.Front().Value, ",Back", l.Back().Value)

	// Iterate through the list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ",")
	}
	fmt.Println()
}
