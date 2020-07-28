package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		// Right: &TreeNode{
		// 	Val: 15,
		// 	Left: &TreeNode{
		// 		Val: 6,
		// 	},
		// 	Right: &TreeNode{
		// 		Val: 20,
		// 	},
		// },
	}
	fmt.Println(isValidBST(t))
}
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(p *TreeNode, lower, higher int) bool {
	if p == nil {
		return true
	}
	if p.Val >= higher || p.Val <= lower {
		return false
	}
	return helper(p.Left, lower, p.Val) && helper(p.Right, p.Val, higher)
}
