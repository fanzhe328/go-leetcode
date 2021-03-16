package main

import (
	"fmt"
)

func main() {
	//fmt.Println(Merge([]int{}, []int{}))

	fmt.Println(isAnagram("abc", "cbad"))
	//fmt.Println(Merge([]int{1, 2, 2, 3, 3}, []int{1, 2, 2, 3}))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1 [26]int
	for _, v := range s {
		c1[v-'a']++
	}

	for _, v := range t {
		if c1[v-'a'] <= 0 {
			return false
		}
		c1[v-'a']--
	}
	return true
}

// func isAnagram(s string, t string) bool {
// 	var c1, c2 [26]int
// 	for _, v := range s {
// 		c1[v-'a']++
// 	}

// 	for _, v := range t {
// 		c2[v-'a']++
// 	}
// 	return c1 == c2
// }

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	cnt := map[rune]int{}
// 	for _, v := range s {
// 		cnt[v]++
// 	}

// 	for _, v := range t {
// 		if cnt[v] <= 0 {
// 			return false
// 		}
// 		cnt[v]--
// 	}
// 	return true
// }
