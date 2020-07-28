package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isEqual('a', 'A'))
	fmt.Println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {

		if l == r {
			return true
		}
		for l < r && !isAlpanOrNum(s[l]) {
			l++
		}
		for l < r && !isAlpanOrNum(s[r]) {
			r--
		}
		if l < r && !isEqual(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlpanOrNum(b byte) bool {
	if b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z' ||
		b >= '0' && b <= '9' {
		return true
	}
	return false
}

func isEqual(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}
	return a == b
}
