package main

import (
	"fmt"

	pq "gopkg.in/dnaeon/go-priorityqueue.v1"
)

func PriorityQueue_Test2() {
	queue := pq.New[string, int64](pq.MinHeap)

	queue.Put("apple", 10)
	queue.Put("banana", 3)
	queue.Put("pear", 20)
	queue.Put("orange", 15)

	fmt.Println("Asc")
	for !queue.IsEmpty() {
		item := queue.Get()
		fmt.Println("value: ", item.Value, ", priority: ", item.Priority)
	}
}

func PriorityQueue_Test3() {
	queue := pq.New[string, int64](pq.MaxHeap)

	queue.Put("apple", 10)
	queue.Put("banana", 3)
	queue.Put("pear", 20)
	queue.Put("orange", 15)

	fmt.Println("Desc")
	for !queue.IsEmpty() {
		item := queue.Get()
		fmt.Println("value: ", item.Value, ", priority: ", item.Priority)
	}
}
