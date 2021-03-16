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
	fmt.Println(levelOrder(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		lens := len(queue)
		vals := []int{}
		for i := 0; i < lens; i++ {
			vals = append(vals, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[lens:]
		res = append(res, vals)
	}
	return res
}

// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	res := [][]int{}
// 	level, nextLevel := []*TreeNode{root}, []*TreeNode{}
// 	for {
// 		vals := []int{}
// 		for i := 0; i < len(level); i++ {
// 			vals = append(vals, level[i].Val)
// 			if level[i].Left != nil {
// 				nextLevel = append(nextLevel, level[i].Left)
// 			}
// 			if level[i].Right != nil {
// 				nextLevel = append(nextLevel, level[i].Right)
// 			}
// 		}
// 		res = append(res, vals)
// 		if len(nextLevel) == 0 {
// 			break
// 		}
// 		level = nextLevel
// 		nextLevel = []*TreeNode{}
// 	}
// 	return res
// }
