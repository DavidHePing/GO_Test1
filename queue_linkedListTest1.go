package main

import (
	"container/list"
	"fmt"
)

func Fake_Queue_LinkedList_Test1() {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	for l.Len() > 0 {
		fmt.Print(l.Front().Value, ", ")
		l.Remove(l.Front())
	}

}
