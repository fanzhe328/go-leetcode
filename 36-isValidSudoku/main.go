package main

import (
	"fmt"
)

func main() {
	a := 7
	fmt.Println(a & 8)
}

func isValidSudoku(board [][]byte) bool {
	lens := len(board)
	box := make([][]int, lens+1)
	for i := 0; i < lens+1; i++ {
		box[i] = make([]int, lens+1)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			x := board[i][j]
			if box[i][x-'0']&1 != 0 || box[j][x-'0']&2 != 0 || box[i/3*3+j/3][x-'0']&4 != 0 {
				return false
			}
			box[i][x-'0'] |= 1
			box[j][x-'0'] |= 2
			box[i/3*3+j/3][x-'0'] |= 4
		}
	}
	return true
}
