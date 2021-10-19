package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'.','.','.','.','.','.','.','.'},{'.','B','.','.','W','.','.','.'},{'.','.','W','.','.','.','.','.'},{'.','.','.','W','B','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','B','W','.','.'},{'.','.','.','.','.','.','W','.'},{'.','.','.','.','.','.','.','B'}}
	rMove := 4
	cMove := 3
	color := byte('B')
	result := checkMove(board, rMove, cMove, color)
	fmt.Println(result)
}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	rMax := len(board)
	cMax := len(board[0])
	var op_color byte
	if color == 'B' {
		op_color = 'W'
	} else {
		op_color = 'B'
	}
	diffs := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	for di:=0;di < len(diffs);di++ {
		dr := diffs[di][0]
		dc := diffs[di][1]
		if rMove + dr * 2 < 0 || rMove + dr * 2 >= rMax {
			continue
		}
		if cMove + dc * 2 < 0 || cMove + dc * 2 >= cMax {
			continue
		}
		if board[rMove + dr][cMove + dc] != op_color {
			continue
		}
		for i:=2;;i++ {
			r := rMove + dr * i
			if r < 0 || r >= rMax {
				break
			}
			c := cMove + dc * i
			if c < 0 || c >= cMax {
				break
			}
			if board[r][c] == color {
				return true
			}
			if board[r][c] != op_color{
				break
			}
		}
	}

	return false
}