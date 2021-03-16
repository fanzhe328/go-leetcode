package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(multiply("408", "0"))
}

func multiply(num1 string, num2 string) string {
	ans := make([]int, len(num1)+len(num2)+1)

	for i := len(num1) - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			res := x*int(num2[j]-'0') + ans[i+j]
			ans[i+j] = res % 10
			ans[i+j+1] += res / 10
		}
	}
	i := 0
	for ; i < len(ans)-1 && ans[i] == 0; i++ {
	}

	var res strings.Builder
	for ; i < len(ans); i++ {
		res.WriteString(strconv.Itoa(ans[i]))
	}
	return res.String()
}

// func multiply(num1 string, num2 string) string {
// 	mm := [][]int{}

// 	for i := len(num1) - 1; i >= 0; i-- {
// 		x := int(num1[i] - '0')
// 		line := []int{}
// 		ans, jinwei := 0, 0
// 		for j := len(num2) - 1; j >= 0; j-- {
// 			ans = x*int(num2[j]-'0') + jinwei
// 			line = append(line, ans%10)
// 			jinwei = ans / 10
// 		}
// 		for jinwei > 0 {
// 			line = append(line, jinwei%10)
// 			jinwei = jinwei / 10
// 		}
// 		mm = append(mm, line)
// 	}

// 	ans := mm[0]
// 	for i := 1; i < len(mm); i++ {
// 		ans = addSlice(ans, mm[i], i)
// 	}

// 	i := len(ans) - 1
// 	for ; i > 0 && ans[i] == 0; i-- {
// 	}

// 	ret := ""
// 	for ; i >= 0; i-- {
// 		ret += strconv.Itoa(ans[i])
// 	}
// 	return ret
// }

// func addSlice(num1, num2 []int, offset int) []int {
// 	ansSlice := []int{}

// 	for i := 0; i < offset; i++ {
// 		ansSlice = append(ansSlice, num1[i])
// 	}
// 	ans, add := 0, 0
// 	for i, j := offset, 0; i < len(num1) || j < len(num2) || add > 0; i, j = i+1, j+1 {
// 		var x, y int
// 		if i < len(num1) {
// 			x = num1[i]
// 		}
// 		if j < len(num2) {
// 			y = num2[j]
// 		}
// 		ans = x + y + add
// 		ansSlice = append(ansSlice, ans%10)
// 		add = ans / 10
// 	}
// 	return ansSlice
// }
