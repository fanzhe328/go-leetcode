package main

import (
	"fmt"
)

func main() {
	s := "Test1ng-Leet=code-Q!"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(isAlpha(s[i]))
	// }
	fmt.Println(reverseOnlyLetters(s))
}

func reverseOnlyLetters(S string) string {
	head := 0
	tail := len(S) - 1
	res := make([]byte, len(S))

	for head <= tail {
		if head == tail {
			res[head] = S[head]
			break
		}
		for head <= tail && !isAlpha(S[head]) {
			res[head] = S[head]
			head++
		}
		for head < tail && !isAlpha(S[tail]) {
			// fmt.Println(S[tail])
			res[tail] = S[tail]
			tail--
		}
		if head < tail {
			res[head], res[tail] = S[tail], S[head]
			head++
			tail--
		}

	}
	return string(res)
}

func isAlpha(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}
