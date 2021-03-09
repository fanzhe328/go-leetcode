package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	p := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	print(p)
	print(swapPairs(p))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHeadPre := &ListNode{0, head}
	pre, cur, head := newHeadPre, head, head.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = cur.Next
	}
	return newHeadPre.Next
}

func print(head *ListNode) {
	p := head
	for ; p.Next != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Printf("%d\n", p.Val)
}
