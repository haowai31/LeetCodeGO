package main

func isValidSudoku(board [][]byte) bool {
	var row, col, rex []map[byte]bool
	for i := 0; i < 9; i++ {
		row = append(row, make(map[byte]bool))
		col = append(col, make(map[byte]bool))
		rex = append(rex, make(map[byte]bool))
	}

	for i := range board {
		for j, num := range board[i] {
			if num == '.' {
				continue
			}
			if row[i][num] || col[j][num] || rex[(i/3)*3+j/3][num] {
				return false
			}
			row[i][num] = true
			col[j][num] = true
			rex[(i/3)*3+j/3][num] = true
		}
	}

	return true
}

func leetcode36() {
	var board [][]byte = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	println(isValidSudoku(board))
}
