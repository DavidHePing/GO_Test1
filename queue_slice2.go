package main

import "fmt"

func Fake_Queue_Slice_Test2() {
	var q []int
	var dequeueItem int
	item := 1

	q = append(q, item)
	fmt.Println("enqueue", item, "queue", q)
	item = 2
	q = append(q, item)
	fmt.Println("enqueue", item, "queue", q)
	item = 3
	q = append(q, item)
	fmt.Println("enqueue", item, "queue", q)

	dequeueItem = q[0]
	q = q[1:]
	fmt.Println("dequeue", dequeueItem, "queue", q)
	dequeueItem = q[0]
	q = q[1:]
	fmt.Println("dequeue", dequeueItem, "queue", q)
	dequeueItem = q[0]
	q = q[1:]
	fmt.Println("dequeue", dequeueItem, "queue", q)
}
