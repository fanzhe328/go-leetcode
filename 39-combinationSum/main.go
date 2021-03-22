package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{1}, 2))
}

type Stack []int

func (s *Stack) Pop() (val int) {
	if len(*s) == 0 {
		return -1
	}
	val, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Copy() []int {
	new := make([]int, len(*s))
	copy(new, *s)
	return new
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	// fmt.Println(candidates)
	res := [][]int{}
	vals, idx := Stack{}, Stack{}
	curNum := 0

	for i := 0; i < len(candidates); {
		vals = append(vals, candidates[i])
		idx = append(idx, i)
		curNum += candidates[i]
		// fmt.Println(curNum, idx, vals)
		if curNum < target {
			continue
		}

		if curNum == target {
			res = append(res, vals.Copy())
		}

		curNum -= vals.Pop()
		curNum -= vals.Pop()
		idx.Pop()
		popIdx := idx.Pop()

		for !idx.IsEmpty() && popIdx+1 >= len(candidates) {
			curNum -= vals.Pop()
			popIdx = idx.Pop()
		}
		// fmt.Println(popIdx)
		if popIdx == -1 {
			i += 1
			continue
		}
		i = popIdx + 1
	}

	return res
}
