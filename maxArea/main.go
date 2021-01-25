package main

import "fmt"

func main() {
	// fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{1, 2, 1, 2, 10, 1, 200, 400, 1, 2, 1, 2}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	minW := 0
	max := 0
	for l < r {
		if height[l] < height[r] {
			minW = height[l]
			l++
		} else {
			minW = height[r]
			r--
		}

		area := (r - l + 1) * minW
		if max < area {
			max = area
		}
	}
	return max
}
