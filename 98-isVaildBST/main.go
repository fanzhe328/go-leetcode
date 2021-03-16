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
	t := &TreeNode{
		Val: 0,
		// Left: &TreeNode{
		// 	Val: 1,
		// },
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
	fmt.Println(isValidBST(t))
}
