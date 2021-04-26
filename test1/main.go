package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	fmt.Println(majorityElement([]int{3, 3, 3, 1, 1}))
}

func majorityElement(nums []int) int {
	val, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			val = nums[i]
			count = 1
		}

		if nums[i] == val {
			count++
		} else {
			count--
		}
	}
	return val
}
