package main

import (
	"fmt"
)

func main() {
	t := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(t))
}

func findRepeatNumber(nums []int) int {
	m := make([]int, 100001)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return nums[i]
		}
		m[nums[i]] = 1
	}
	return 0
}
