package main

import (
	"container/list"
	"fmt"
)

type LinkedListQueue struct {
	items list.List
}

func (l *LinkedListQueue) Len() int {
	return l.items.Len()
}

func (l *LinkedListQueue) EnQueue(num int) {
	l.items.PushBack(num)
}

func (l *LinkedListQueue) Dequeue() int {
	if l.items.Len() == 0 {
		return -1
	}

	tmp := l.items.Front().Value
	l.items.Remove(l.items.Front())
	return tmp.(int)
}

func Fake_Queue_LinkedList_Test2() {
	var l LinkedListQueue

	l.EnQueue(1)
	l.EnQueue(2)
	l.EnQueue(3)
	l.EnQueue(4)

	fmt.Println("Len:", l.Len())

	for l.Len() > 0 {
		fmt.Print(l.Dequeue(), ", ")
	}

	fmt.Println("")
	fmt.Println("Len:", l.Len())
}
