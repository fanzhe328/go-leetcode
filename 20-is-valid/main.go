package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	fmt.Println(isValid("{()}[]"))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

func isValid(s string) bool {
	res := make([]byte, 0, len(s))
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	for i := 0; i <= len(s)-1; i++ {
		if _, ok := m[s[i]]; !ok {
			res = append(res, s[i])
		} else if len(res) <= 0 || res[len(res)-1] != m[s[i]] {
			return false
		} else {
			res = res[:len(res)-1]
		}
	}
	if len(res) == 0 {
		return true
	}
	return false
}

// func isValid(s string) bool {
// 	res := make([]byte, 0, len(s))
// 	for i := 0; i <= len(s)-1; i++ {
// 		switch s[i] {
// 		case '(', '{', '[':
// 			res = append(res, s[i])
// 		case ')':
// 			if len(res) <= 0 || res[len(res)-1] != '(' {
// 				return false
// 			}
// 			res = res[:len(res)-1]
// 		case '}':
// 			if len(res) <= 0 || res[len(res)-1] != '{' {
// 				return false
// 			}
// 			res = res[:len(res)-1]
// 		case ']':
// 			if len(res) <= 0 || res[len(res)-1] != '[' {
// 				return false
// 			}
// 			res = res[:len(res)-1]
// 		default:
// 			continue
// 		}
// 	}
// 	if len(res) == 0 {
// 		return true
// 	}
// 	return false
// }
