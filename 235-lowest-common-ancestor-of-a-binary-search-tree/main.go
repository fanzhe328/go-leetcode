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

// 一次遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := root
	for {
		if ans.Val < p.Val && ans.Val < q.Val {
			ans = ans.Right
		} else if ans.Val > p.Val && ans.Val > q.Val {
			ans = ans.Left
		} else {
			return ans
		}
	}
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	if root.Val > p.Val && root.Val > q.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	if root.Val < p.Val && root.Val < q.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }

// 递归方法
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if p.Val > q.Val {
// 		return lca(root, q, p)
// 	}
// 	return lca(root, p, q)
// }

// func lca(root, l, h *TreeNode) *TreeNode {
// 	if root == nil || (root.Val >= l.Val && root.Val <= h.Val) {
// 		return root
// 	}

// 	left := lca(root.Left, l, h)
// 	right := lca(root.Right, l, h)

// 	if right == nil {
// 		return left
// 	}
//  if left == nil {
// 		return right
// 	}
//  return root
// }
