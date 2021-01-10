package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x/10 != 0) {
		return false
	}

	l, r := x, 0
	for l > r {
		r = 10*r + l%10
		l = l / 10
	}
	return l == r || r/10 == l
}
