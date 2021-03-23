package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
}

func totalNQueens(n int) (ans int) {
	var solve func(row int, col, dia1, dia2 int)
	solve = func(row int, col, dia1, dia2 int) {
		if row == n {
			ans++
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
			solve(row+1, col|pos, (dia1|pos)>>1, (dia2|pos)<<1)
		}
	}

	solve(0, 0, 0, 0)
	return
}
