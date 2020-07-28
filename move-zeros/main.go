package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 100, 3, 0, 0, 12, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var j int
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
