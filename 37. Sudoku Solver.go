package main

func next(i, j int) (ni, nj int) {
	nj = j + 1
	if nj == 9 {
		nj = 0
		ni = i + 1
	} else {
		ni = i
	}
	return
}

var entrys = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudokuxx(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	ni, nj := next(i, j)
	if board[i][j] != '.' {
		return solveSudokuxx(board, ni, nj)
	}
	for _, e := range entrys {
		if !inCel(board, i, j, e) && !inRow(board, i, e) && !inCol(board, j, e) {
			board[i][j] = e
			if solveSudokuxx(board, ni, nj) {
				return true
			}
		}
	}
	board[i][j] = '.'
	return false
}

func inCel(board [][]byte, i, j int, e byte) bool {
	ci, cj := i/3, j/3
	for x := range entrys {
		if board[3*ci+x/3][3*cj+x%3] == e {
			return true
		}
	}
	return false
}

func inCol(board [][]byte, j int, e byte) bool {
	for x := range entrys {
		if board[x][j] == e {
			return true
		}
	}
	return false
}

func inRow(board [][]byte, i int, e byte) bool {
	for x := range entrys {
		if board[i][x] == e {
			return true
		}
	}
	return false
}

func solveSudoku(board [][]byte) {
	solveSudokuxx(board, 0, 0)
}

func leetcode37() {
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
	solveSudoku(board)
	println(board)
}
