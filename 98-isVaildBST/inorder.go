package main

// var prev *TreeNode

// // 注意：在leetcode在验证时会跑多个test case，如果使用全局变量，在每次执行完函数时，需要将全局变量重置为nil
// func isValidBST(root *TreeNode) bool {
// 	isValid := helper(root)
// 	prev = nil
// 	return isValid
// }

// // 递归方式的中序遍历，只保留前一个节点
// func helper(p *TreeNode) bool {
// 	if p == nil {
// 		return true
// 	}
// 	if !helper(p.Left) {
// 		return false
// 	}

// 	if prev != nil && p.Val <= prev.Val {
// 		return false
// 	}
// 	prev = p
// 	return helper(p.Right)
// }

// 递归法
// func isValidBST(root *TreeNode) bool {
// 	return helper(root, math.MinInt64, math.MaxInt64)
// }

// func helper(p *TreeNode, lower, higher int) bool {
// 	if p == nil {
// 		return true
// 	}
// 	if p.Val >= higher || p.Val <= lower {
// 		return false
// 	}
// 	return helper(p.Left, lower, p.Val) && helper(p.Right, p.Val, higher)
// }
