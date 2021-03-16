package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//fmt.Println(Merge([]int{}, []int{}))
	t := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		// Right: &TreeNode{
		// 	Val: 4,
		// 	Left: &TreeNode{
		// 		Val: 3,
		// 	},
		// 	Right: &TreeNode{
		// 		Val: 6,
		// 	},
		// },
	}
	fmt.Println(lowestCommonAncestor(t, &TreeNode{Val: 2}, &TreeNode{Val: 1}).Val)
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

// 递归方法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if right == nil {
		return left
	}
	if left == nil {
		return right
	}
	return root

}
