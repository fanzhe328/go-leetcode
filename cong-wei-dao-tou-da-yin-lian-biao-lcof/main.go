package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(reversePrint(s))
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	cur := head
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}
	res := make([]int, n)
	cur = head
	for i := n - 1; i >= 0; i-- {
		res[i] = cur.Val
		cur = cur.Next
	}
	return res
}
