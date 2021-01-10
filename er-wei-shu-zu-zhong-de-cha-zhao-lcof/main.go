package main

import (
	"fmt"
)

func main() {
	t := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	// t := [][]int{{-1, 3}}
	fmt.Println(findNumberIn2DArray(t, 14))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix)
	x := 0
	y := len(matrix[0]) - 1
	for x < n && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		}
	}

	return false
}
