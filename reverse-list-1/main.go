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

	p := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	print(p)
	print(reverseList(p))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

// 非递归
// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	curr := head

// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	return prev
// }

// 非递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func print(head *ListNode) {
	p := head
	for ; p.Next != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Printf("%d\n", p.Val)
}
