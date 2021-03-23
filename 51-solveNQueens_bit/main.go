package main

import (
	"fmt"
	"math/bits"
)

func main() {
	boards := solveNQueens(4)
	for _, items := range boards {
		for _, item := range items {
			fmt.Println(item)
		}
		fmt.Println()
	}
}

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = make([][]string, 0)

	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}

	backtrace(queens, 0, n, 0, 0, 0)
	return solutions
}

func backtrace(queens []int, level, n int, col, dia1, dia2 int) {
	if level == n {
		board := genBoard(queens)
		solutions = append(solutions, board)
		return
	}
	avaliablePos := ((1 << n) - 1) & (^(col | dia1 | dia2))
	for avaliablePos != 0 {
		// x & (−x) 可以获得 x 的二进制表示中的最低位的 1 的位置
		// 即获得最低位可以放置皇后的位置
		pos := avaliablePos & -avaliablePos
		// x & (x−1) 可以将 x 的二进制表示中的最低位的 1 置成 0
		// 即下一个开始遍历下一个低位
		avaliablePos = avaliablePos & (avaliablePos - 1)

		// 获取pos值是在第几位，如 1000:3, 100:2
		column := bits.OnesCount(uint(pos - 1))
		queens[level] = column
		backtrace(queens, level+1, n, col|pos, (dia1|pos)>>1, (dia2|pos)<<1)
		queens[level] = -1
	}

}

func genBoard(queens []int) []string {
	result := make([]string, 0, len(queens))
	for i := 0; i < len(queens); i++ {
		s := make([]byte, len(queens))
		for j := 0; j < len(queens); j++ {
			s[j] = '.'
		}
		s[queens[i]] = 'Q'
		result = append(result, string(s))
	}
	return result
}
