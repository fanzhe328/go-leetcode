package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}

func dfs(p *TreeNode, val, target int) bool {
	if p == nil {
		return false
	}
	val += p.Val
	if p.Left == nil && p.Right == nil {
		if val == target {
			return true
		}
		return false
	}
	return dfs(p.Left, val, target) || dfs(p.Right, val, target)
}
