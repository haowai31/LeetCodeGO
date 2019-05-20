package main

func solveSudokuxx(board [][]byte) bool {
	for i := range board {
		for j, num := range board[i] {
			if num == '.' {
				for index := 1; index <= 9; index++ {
					board[i][j] = byte(index) + '0'
					if isValade_Sudoku(board, i, j) && solveSudokuxx(board) {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
	}
	return true
}

func isValade_Sudoku(xx [][]byte, i int, j int) bool {
	var tmp map[byte]bool = make(map[byte]bool)
	for ix := range xx[i] {
		if tmp[xx[ix][j]] && xx[ix][j] != '.' {
			return false
		}
		if xx[ix][j] != '.' {
			tmp[xx[ix][j]] = true
		}
	}
	tmp = make(map[byte]bool)
	for jx := 0; jx < 9; jx++ {
		if tmp[xx[i][jx]] && xx[i][jx] != '.' {
			return false
		}
		if xx[i][jx] != '.' {
			tmp[xx[i][jx]] = true
		}
	}
	tmp = make(map[byte]bool)
	for ix := i / 3; ix < i/3+3; ix++ {
		for jx := j / 3 * 3; jx < j/3*3+3; jx++ {
			if tmp[xx[ix][jx]] && xx[ix][jx] != '.' {
				return false
			}
			if xx[ix][jx] != '.' {
				tmp[xx[ix][jx]] = true
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solveSudokuxx(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

		}
	}
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
