package main

import (
	"fmt"
)

func main() {
	// x := uint32(1) << ('4' - '0' - 1)
	// fmt.Println(x)
	// x := byte(2)+'0'
	// a := uint32(12)
	// x := uint32(2)
	// a |= x
	// fmt.Printf("%b\n", a)
	// a &= (^x)
	// fmt.Printf("%b\n", a)
	a := [1]int{0}
	a[0] = 1
	test(a)
	fmt.Println(a[0])
}

func test(a [1]int) {
	a[0] += 1
}

func solveSudoku(board [][]byte) {
	var backtrace func(n, i, j int, row, col, box [10]uint32) bool
	backtrace = func(n, i, j int, row, col, box [10]uint32) bool {
		if j == n {
			return backtrace(n, i+1, 0, row, col, box)
		}
		if i == n {
			return true
		}
		if board[i][j] != '.' {
			return backtrace(n, i, j+1, row, col, box)
		}
		for k := 1; k <= 9; k++ {
			x := uint32(1) << (k-1)
			if row[i]&x != 0 || col[j]&x != 0 || box[i/3*3+j/3]&x != 0 {
				continue
			}
			row[i] |= x
			col[j] |= x
			box[i/3*3+j/3] |= x
			board[i][j] = byte(k) + '0'
			if backtrace(n, i, j+1, row, col, box) {
				return true
			}
			board[i][j] = '.'
			row[i] &= ^x
			col[j] &= ^x
			box[i/3*3+j/3] &= ^x
		}
		return false
	}
    row, col, box := [10]uint32{},[10]uint32{}, [10]uint32{}
	for i := 0; i<9; i++ {
		for j :=0; j<9; j++ {
			if board[i][j] != '.' {
				x := uint32(1) << (board[i][j]-'0'-1)
				row[i] |= x
				col[j] |= x
				box[i/3*3+j/3] |= x
			}
		}
	}
    backtrace(9, 0, 0, row, col, box)
}