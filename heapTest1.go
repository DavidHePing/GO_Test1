package main

import (
	"container/heap"
	"fmt"
)

type IntHeap1 []int

func (h IntHeap1) Len() int {
	return len(h)
}

func (h IntHeap1) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap1) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap1) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeap1_Test1() {
	data := IntHeap1{2, 1, 5, 9, 34, 7, 91, 11}
	h1 := make(IntHeap1, len(data))
	copy(h1, data)

	heap.Init(&h1)

	for h1.Len() > 0 {
		fmt.Print(heap.Pop(&h1), ", ")
	}
	fmt.Println()

	h2 := make(IntHeap1, len(data))
	copy(h2, data)
	heap.Init(&h2)
	heap.Push(&h2, 14)
	heap.Push(&h2, 21)
	heap.Push(&h2, 32)
	heap.Push(&h2, 8)
	heap.Push(&h2, 10)

	for h2.Len() > 0 {
		fmt.Print(heap.Pop(&h2), ", ")
	}
	fmt.Println()
}
