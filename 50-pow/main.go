package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))
	x := 34.00515
	n := -3
	fmt.Println(myPow(x, n))
	// fmt.Println(math.Pow(x, float64(n)))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

// 递归
// func myPow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1.0
// 	}
// 	if n < 0 {
// 		return myPow(1/x, -1*n)
// 	}
// 	if n&1 != 0 {
// 		return myPow(x*x, n>>1) * x
// 	}
// 	return myPow(x*x, n>>1)
// }

// 循环
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n != 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
