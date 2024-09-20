package main

import "container/list"

type LinkedListQueue struct {
	items list.List
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
