package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(climbStairs(i))
	}

}

// // 递归
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

// 迭代
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	first := 1
// 	second := 2
// 	third := 0
// 	for i := 3; i <= n; i++ {
// 		third = first + second
// 		first = second
// 		second = third
// 	}
// 	return second
// }

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	first := 1
	second := 2
	third := 4
	fourth := 0
	for i := 4; i <= n; i++ {
		fourth = first + second + third
		first = second
		second = third
		third = fourth
	}
	return third
}
