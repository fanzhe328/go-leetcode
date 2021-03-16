package main

import "math"

// 递归法
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
