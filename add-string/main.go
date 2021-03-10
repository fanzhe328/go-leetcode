package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	fmt.Println(addStrings("999", "1000000000000"))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		res := x + y + add
		ans = strconv.Itoa(res%10) + ans
		add = res / 10
	}
	return ans
}
