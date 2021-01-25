package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{1, &ListNode{4, &ListNode{2, nil}}}
	fmt.Println(reversePrint(l))
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	p, n := head, 0
	for p != nil {
		n++
		p = p.Next
	}
	res := make([]int, n)
	p = head
	for p != nil {
		n = n - 1
		res[n] = p.Val
		p = p.Next
	}
	return res
}
