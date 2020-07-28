package main

import "fmt"

func main() {
	nums := []int{1, 1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	slow := 1
	for i := 1; i < lens; i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
