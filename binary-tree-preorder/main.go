package main

import (
	"fmt"

	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// t := &TreeNode{
	// 	Val: 1,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 			Left: &TreeNode{
	// 				Val: 4,
	// 			},
	// 		},
	// 	},
	// }
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(preorderTraversal(t))
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := list.New()

	stack.PushBack(root)
	for stack.Len() != 0 {
		item := stack.Back()
		cur := item.Value.(*TreeNode)
		res = append(res, cur.Val)
		stack.Remove(item)
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
	}
	return res
}
