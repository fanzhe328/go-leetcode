package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{1, &ListNode{4, &ListNode{2, nil}}}
	// l := &ListNode{1, nil}
	printList(l)
	printList(reverseList(l))
}

func reverseList(head *ListNode) *ListNode {
	p := head
	var (
		pre  *ListNode
		next *ListNode
	)
	for p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}
