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
	fmt.Println()
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1

	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}
