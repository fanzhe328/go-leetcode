package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// func hasCycle(head *ListNode) bool {
// 	if head == nil || head.Next == nil || head.Next.Next == nil {
// 		return false
// 	}
// 	slow, fast := head.Next, head.Next.Next
// 	for slow != fast {
// 		if fast.Next == nil || fast.Next.Next == nil {
// 			return false
// 		}
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 	}
// 	return true
// }
