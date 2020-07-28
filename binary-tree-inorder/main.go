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
	t := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(t))
}

// 栈
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := list.New()
	cur := root

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		tmp := stack.Back()
		cur = tmp.Value.(*TreeNode)
		res = append(res, cur.Val)
		stack.Remove(tmp)
		cur = cur.Right

	}
	return res
}

// 递归法
// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	helper(root, &res)
// 	return res
// }

// func helper(p *TreeNode, res *[]int) {
// 	if p == nil {
// 		return
// 	}
// 	helper(p.Left, res)
// 	*res = append(*res, p.Val)
// 	helper(p.Right, res)
// }
