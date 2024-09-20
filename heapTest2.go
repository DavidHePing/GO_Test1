package main

import (
	"container/heap"
	"fmt"
)

type IntHeap2Desc []int

func (h IntHeap2Desc) Len() int {
	return len(h)
}

func (h IntHeap2Desc) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap2Desc) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap2Desc) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap2Desc) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeap2Desc_Test1() {
	data := IntHeap2Desc{2, 1, 5, 9, 34, 7, 91, 11}
	h1 := make(IntHeap2Desc, len(data))
	copy(h1, data)

	heap.Init(&h1)

	fmt.Println("Desc")
	for h1.Len() > 0 {
		fmt.Print(heap.Pop(&h1), ", ")
	}
	fmt.Println()

	h2 := make(IntHeap2Desc, len(data))
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
