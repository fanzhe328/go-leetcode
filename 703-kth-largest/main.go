package main

import (
	"fmt"
)

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Printf("%d\n", obj.Head())
	fmt.Printf("%d\r\n", obj.Add(3))
	fmt.Printf("%d\r\n", obj.Len())
	fmt.Printf("%d\r\n", obj.Add(5))
	// fmt.Println()
}

type MinHeap []int

type KthLargest struct {
	h *MinHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	instance := KthLargest{
		h: NewMinHeap(),
		k: k,
	}
	for _, item := range nums {
		if instance.h.Len() == instance.k && item <= instance.h.Top() {
			continue
		}
		instance.h.Push(item)
		if instance.h.Len() > k {
			instance.h.Pop()
		}
	}
	return instance
}

func (this *KthLargest) Head() int {
	return this.h.Top()
}

func (this *KthLargest) Len() int {
	return this.h.Len()
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k || val > this.h.Top() {
		this.h.Push(val)
	}

	if this.h.Len() > this.k {
		this.h.Pop()
	}
	if this.h.Len() != this.k {
		return -1
	}
	return this.h.Top()
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	*h = append(*h, 0)
	return h
}

func (h *MinHeap) Len() int {
	return len(*h) - 1
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
	for i := idx * 2; i <= lens; i = idx * 2 {
		if i+1 <= lens && (*h)[i] > (*h)[i+1] {
			i = i + 1
		}
		if (*h)[i] > (*h)[idx] {
			break
		}
		(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
		idx = i
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
