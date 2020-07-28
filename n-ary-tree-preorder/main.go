package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	t := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{Val: 5},
					&Node{Val: 6},
				},
			},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}
	fmt.Println(preorder(t))
}

// 迭代
func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	cur := root
	stack := list.New()
	stack.PushBack(cur)
	for stack.Len() != 0 {
		tmp := stack.Back()
		cur = tmp.Value.(*Node)
		res = append(res, cur.Val)
		stack.Remove(tmp)
		if len(cur.Children) == 0 {
			continue
		}
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack.PushBack(cur.Children[i])
		}
	}
	return res
}

// // 递归
// func preorder(root *Node) []int {
// 	res := []int{}
// 	cur := root

// 	helper(cur, &res)
// 	return res
// }

// func helper(cur *Node, res *[]int) {
// 	if cur == nil {
// 		return
// 	}
// 	*res = append(*res, cur.Val)
// 	for i := 0; i < len(cur.Children); i++ {
// 		helper(cur.Children[i], res)
// 	}
// }
