package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l, r int, s string)
	dfs = func(l, r int, s string) {
		if l == 0 && r == 0 {
			res = append(res, s)
		}
		if l > 0 {
			dfs(l-1, r, s+"(")
		}
		if l < r {
			dfs(l, r-1, s+")")
		}
	}
	dfs(n, n, "")
	return res
}

// func dfs(l, r int, s string, res *[]string) {
// 	if l == 0 && r == 0 {
// 		*res = append(*res, s)
// 	}
// 	if l > 0 {
// 		dfs(l-1, r, s+"(", res)
// 	}
// 	if l < r {
// 		dfs(l, r-1, s+")", res)
// 	}
// }
