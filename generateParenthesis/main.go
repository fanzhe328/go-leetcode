package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := []string{}
	dfs(n, n, "", &res)
	return res
}

func dfs(l, r int, s string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, s)
	}
	if l > 0 {
		dfs(l-1, r, s+"(", res)
	}
	if l < r {
		dfs(l, r-1, s+")", res)
	}
}
