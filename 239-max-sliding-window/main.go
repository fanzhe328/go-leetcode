package main

import "fmt"

func main() {
	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 1))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res, win := []int{}, []int{}

	appendWin := func(window []int, val int) []int {
		i := len(window) - 1
		for ; i >= 0; i-- {
			if window[i] >= val {
				break
			}
		}
		return append(window[:i+1], val)
	}

	i := 0
	for ; i < k; i++ {
		win = appendWin(win, nums[i])
	}
	res = append(res, win[0])

	for ; i < len(nums); i++ {
		if win[0] == nums[i-k] {
			win = win[1:]
		}
		win = appendWin(win, nums[i])
		res = append(res, win[0])
	}
	return res
}
