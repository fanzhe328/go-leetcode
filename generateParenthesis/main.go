package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	result := []string{}
	_generate(0, 0, n, "", &result)
	return result
}

func _generate(left, right, n int, s string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, s)
	}

	if left < n {
		_generate(left+1, right, n, s+"(", result)
	}
	if right < left {
		_generate(left, right+1, n, s+")", result)
	}
}
