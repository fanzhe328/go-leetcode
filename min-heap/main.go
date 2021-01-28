package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func main() {
	h := new(MinHeap)
	// h := &myHeap{7, 3, 5, 9}
	heap.Init(h)
	heap.Push(h, 7)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 9)
	// h.Push(7)
	// h.Push(3)
	// h.Push(5)
	// h.Push(9)
	fmt.Println(heap.Pop(h).(int))
	fmt.Println(heap.Pop(h).(int))
	fmt.Println(heap.Pop(h).(int))
	// fmt.Println(heap.Pop(h).(int))
	// 3 5 7
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
