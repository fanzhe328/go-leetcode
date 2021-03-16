package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	fmt.Println(majorityElement([]int{3, 3, 4}))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

// Boyer-Moore 投票算法，摩尔投票法
// https://leetcode-cn.com/problems/majority-element/solution/duo-shu-yuan-su-by-leetcode-solution/
// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElement(nums []int) int {
	cand, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			cand = nums[i]
			count++
			continue
		}
		if cand == nums[i] {
			count++
		} else {
			count--
		}
	}
	return cand
}
