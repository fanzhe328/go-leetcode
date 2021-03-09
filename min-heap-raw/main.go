package main

import "fmt"

func main() {
	h := NewMinHeap()
	h.Push(1)
	h.Push(-1)
	h.Push(-11)
	h.Push(5)
	h.Push(9)
	h.Push(-21)
	fmt.Printf("%d\r\n", h.Pop())
	fmt.Printf("%d\r\n", h.Pop())
	fmt.Printf("%d\r\n", h.Pop())

	// fmt.Println(Merge([]int{1, 2, 2, 3}, []int{}))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

type MinHeap []int

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	*h = append(*h, 0) //
	return h
}

func (h *MinHeap) Push(val int) {
	*h = append(*h, val)
	h.siftUp()
}

func (h *MinHeap) Top() int {
	if len(*h) <= 1 {
		return -1
	}
	return (*h)[1]
}

func (h *MinHeap) siftUp() {
	idx := len(*h) - 1
	if idx == 1 {
		return
	}
	for i := idx / 2; i >= 1; i = idx / 2 {
		if (*h)[idx] < (*h)[i] {
			(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
		}
		idx = i
	}
}
func (h *MinHeap) siftDown() {
	lens := len(*h) - 1
	idx := 1
	for i := idx * 2; ; i = idx * 2 {
		if i > lens {
			break
		}
		if i+1 <= lens && (*h)[i] > (*h)[i+1] {
			i = i + 1
		}
		if (*h)[i] > (*h)[idx] {
			break
		}
		(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
	}
}

func (h *MinHeap) Pop() int {
	item := (*h)[1]
	idx := len(*h) - 1
	if idx != 1 {
		(*h)[1] = (*h)[idx]
		*h = (*h)[:idx]
		h.siftDown()
	}
	return item
}
