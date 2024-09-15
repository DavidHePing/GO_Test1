package main

import "fmt"

// SliceQueue represents a queue data structure
type SliceQueue struct {
	items []int
}

// Enqueue adds an item to the queue
func (q *SliceQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes an item from the queue
func (q *SliceQueue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty!")
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func Fake_Queue_Slice_Test1() {
	q := &SliceQueue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue()) // Outputs: 1
	fmt.Println(q.Dequeue()) // Outputs: 2
	fmt.Println(q.Dequeue()) // Outputs: 3
}
