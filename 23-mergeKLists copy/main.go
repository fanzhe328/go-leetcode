package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println()
}

type MinHeap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := &MinHeap{}
	*h = make(MinHeap, 0, len(lists))
	heap.Init(h)
	for _, item := range lists {
		if item == nil {
			continue
		}
		heap.Push(h, item)
		// lists[i] = item.Next
	}
	if h.Len() == 0 {
		return nil
	}

	var head, p *ListNode
	for h.Len() != 0 {
		curr := heap.Pop(h).(*ListNode)
		if head == nil {
			head = curr
		} else {
			p.Next = curr
		}
		p = curr
		if curr.Next != nil {
			heap.Push(h, curr.Next)
		}
	}
	return head
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
