package main

import (
	"container/list"
	"fmt"
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
	fmt.Println(postorderTraversal(t))
}

func postorderTraversal(root *TreeNode) []int {
	type Item struct {
		Flag int
		Node *TreeNode
	}

	res := []int{}
	stack := list.New()
	cur := root
	stack.PushBack(Item{0, cur})
	for stack.Len() != 0 {
		tmp := stack.Back()
		item := tmp.Value.(Item)
		if item.Node == nil {
			stack.Remove(tmp)
			continue
		}
		if item.Flag == 0 {
			stack.Remove(tmp)
			stack.PushBack(Item{1, item.Node})
			stack.PushBack(Item{0, item.Node.Right})
			stack.PushBack(Item{0, item.Node.Left})
			continue
		}
		stack.Remove(tmp)
		res = append(res, item.Node.Val)
	}
	return res
}
