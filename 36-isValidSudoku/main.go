package main

import (
	"fmt"
)

func main() {
	a := 1 << 3
	fmt.Println(a)
}

func isValidSudoku(board [][]byte) bool {
	lens := len(board)
	row, col, box := make([]uint32, lens+1), make([]uint32, lens+1), make([]uint32, lens+1)

	for i := 0; i < lens; i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			x := uint32(1) << (board[i][j] - '0')
			if row[i]&x != 0 || col[j]&x != 0 || box[i/3*3+j/3]&x != 0 {
				return false
			}
			row[i] |= x
			col[j] |= x
			box[i/3*3+j/3] |= x
		}
	}
	return true
}
