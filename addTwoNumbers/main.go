package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	print(addTwoNumbers(newList([]int{2, 4, 3}), newList([]int{5, 6, 4})))
	print(addTwoNumbers(newList([]int{9, 9, 9, 9, 9, 9, 9}), newList([]int{9, 9, 9, 9})))
	// print(newList([]int{2, 4, 3}))
	// print(newList([]int{5, 6, 4}))
}

func newList(nums []int) *ListNode {
	res := &ListNode{}
	pre, p := res, res
	for _, num := range nums {
		p.Val = num
		p.Next = &ListNode{}
		pre = p
		p = p.Next
	}
	pre.Next = nil
	return res
}

func print(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p := res
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{sum % 10, nil}
		p = p.Next
		sum /= 10
	}
	if sum != 0 {
		p.Next = &ListNode{sum, nil}
	}
	return res.Next
}
