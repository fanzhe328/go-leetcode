package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

func totalNQueens(n int) (ans int) {
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}

	var solve func(row int, columns, diagonals1, diagonals2 map[int]bool)
	solve = func(row int, columns, diagonals1, diagonals2 map[int]bool) {
		if row == n {
			ans++
			return
		}

		for i := 0; i < n; i++ {
			if columns[i] {
				continue
			}
			diagonal1 := row + i
			if diagonals1[diagonal1] {
				continue
			}
			diagonal2 := row - i
			if diagonals2[diagonal2] {
				continue
			}

			columns[i] = true
			diagonals1[diagonal1], diagonals2[diagonal2] = true, true
			solve(row+1, columns, diagonals1, diagonals2)

			columns[i] = false
			diagonals1[diagonal1], diagonals2[diagonal2] = false, false
		}
	}

	solve(0, columns, diagonals1, diagonals2)
	return
}
