package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 1))
}

type Stack []int

func (s *Stack) PushBackDesc(v int) {
	for len(*s) != 0 {
		item := s.Back()
		if v <= item {
			break
		}
		s.Pop()
	}
	*s = append(*s, v)
}

func (s *Stack) Back() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Head() int {
	return (*s)[0]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) RemoveLessThan(value int) {
	if (*s)[0] == value {
		*s = (*s)[1:len(*s)]
	}
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) == 0 {
// 		return []int{}
// 	}
// 	res := []int{}
// 	s := new(Stack)
// 	for i := 0; i < k; i++ {
// 		s.PushBackDesc(nums[i])
// 	}
// 	res = append(res, s.Head())
// 	for i := k; i < len(nums); i++ {
// 		s.PushBackDesc(nums[i])
// 		s.RemoveLessThan(nums[i-k])
// 		res = append(res, s.Head())
// 	}
// 	return res
// }

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
// 辅助数组，保存滑动窗口内的单调递减队列
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res, win := []int{}, []int{}
	winTailIdx := func(window []int, key int) int {
		j := len(win) - 1
		for ; j >= 0; j-- {
			if key <= win[j] {
				break
			}
		}
		return j + 1
	}
	for i := 0; i < k; i++ {
		win = append(win[:winTailIdx(win, nums[i])], nums[i])
	}
	res = append(res, win[0])
	for i := k; i < len(nums); i++ {
		win = append(win[0:winTailIdx(win, nums[i])], nums[i])
		if win[0] == nums[i-k] {
			win = win[1:]
		}
		res = append(res, win[0])
	}
	return res
}
