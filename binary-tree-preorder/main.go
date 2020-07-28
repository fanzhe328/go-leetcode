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
	fmt.Println(preorderTraversal(t))
}

// 栈
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := root
	queue := list.New()
	res := []int{}
	queue.PushBack(cur)
	for queue.Len() != 0 {
		tmp := queue.Back()
		cur = tmp.Value.(*TreeNode)
		queue.Remove(tmp)
		res = append(res, cur.Val)
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}
	}
	return res
}

// // 递归
// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	res := make([]int, 0)
// 	cur := root
// 	helper(cur, &res)
// 	return res
// }

// func helper(p *TreeNode, res *[]int) {
// 	if p == nil {
// 		return
// 	}
// 	*res = append(*res, p.Val)

// 	helper(p.Left, res)
// 	helper(p.Right, res)
// }

// // 栈
// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	stack := list.New()
// 	cur := root

// 	for cur != nil || stack.Len() != 0 {
// 		for cur != nil {
// 			stack.PushBack(cur)
// 			cur = cur.Left
// 		}
// 		tmp := stack.Back()
// 		cur = tmp.Value.(*TreeNode)
// 		res = append(res, cur.Val)
// 		stack.Remove(tmp)
// 		cur = cur.Right

// 	}
// 	return res
// }
