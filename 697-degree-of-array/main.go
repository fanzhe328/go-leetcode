package main

import "fmt"

func main() {
	fmt.Printf("%d\r\n", findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}

type Node struct {
	Num, Degree      int
	StartIdx, EndIdx int
}

func NewNode(num, startIdx int) *Node {
	return &Node{
		Num:      num,
		Degree:   1,
		StartIdx: startIdx,
		EndIdx:   startIdx,
	}
}

func findShortestSubArray(nums []int) int {
	cache := map[int]*Node{}
	for i, num := range nums {
		node, ok := cache[num]
		if !ok {
			node = NewNode(num, i)
			cache[num] = node
		} else {
			node.Degree++
			node.EndIdx = i
		}
	}

	maxDegree, minLens := 0, 0
	for _, node := range cache {
		if maxDegree == 0 || node.Degree > maxDegree {
			maxDegree = node.Degree
			minLens = node.EndIdx - node.StartIdx + 1
		}
		if node.Degree == maxDegree {
			lens := node.EndIdx - node.StartIdx + 1
			if lens < minLens {
				minLens = lens
			}
		}
	}
	return minLens
}
