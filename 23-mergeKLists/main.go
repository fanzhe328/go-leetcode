package main

func main() {
	// h := NewMinHeap()
	// h.Push(&ListNode{Val: 2})
	// h.Push(&ListNode{Val: 3})
	// h.Push(&ListNode{Val: 1})
	// fmt.Println(h.Pop().Val, h.Len())
	// fmt.Println(h.Pop().Val, h.Len())
	// fmt.Println(h.Pop().Val, h.Len())
	// fmt.Println(h.Pop().Val, h.Len())
	// fmt.Println(h.Pop().Val)
	// fmt.Println(h.Pop().Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	heap := NewMinHeap()

	for _, item := range lists {
		if item == nil {
			continue
		}
		heap.Push(item)
	}
	if heap.Len() == 0 {
		return nil
	}

	var head, p *ListNode
	cur := heap.Pop()
	head, p = cur, cur
	if cur.Next != nil {
		heap.Push(cur.Next)
	}
	for heap.Len() > 0 {
		cur := heap.Pop()
		p.Next = cur
		p = cur
		if cur.Next != nil {
			heap.Push(cur.Next)
		}
	}
	return head
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	*h = append(*h, &ListNode{})
	return h
}

func (h *MinHeap) Len() int {
	return len(*h) - 1
}

func (h *MinHeap) Push(val *ListNode) {
	*h = append(*h, val)
	h.siftUp()
}

func (h *MinHeap) Top() *ListNode {
	if len(*h) <= 1 {
		return nil
	}
	return (*h)[1]
}

func (h *MinHeap) siftUp() {
	idx := len(*h) - 1
	if idx == 1 {
		return
	}
	for i := idx / 2; i >= 1; i = idx / 2 {
		if (*h)[idx].Val < (*h)[i].Val {
			(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
		}
		idx = i
	}
}
func (h *MinHeap) siftDown() {
	lens := len(*h) - 1
	idx := 1
	for i := idx * 2; i <= lens; i = idx * 2 {
		if i+1 <= lens && (*h)[i].Val > (*h)[i+1].Val {
			i = i + 1
		}
		if (*h)[i].Val > (*h)[idx].Val {
			break
		}
		(*h)[idx], (*h)[i] = (*h)[i], (*h)[idx]
		idx = i
	}
}

func (h *MinHeap) Pop() *ListNode {
	if len(*h) <= 1 {
		return nil
	}
	item := (*h)[1]
	idx := len(*h) - 1
	if idx != 1 {
		(*h)[1] = (*h)[idx]
		*h = (*h)[:idx]
		h.siftDown()
	} else {
		*h = (*h)[:idx]
	}
	return item
}
