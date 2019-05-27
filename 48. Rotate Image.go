package main

func rotate(matrix [][]int) {
	var n = len(matrix)
	if n <= 1 {
		return
	}
	for i := 0; i < n/2; i++ {
		var lens = n - 2*i
		for j := i; j < lens+i-1; j++ {
			var tmp = matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}

func leetcode48() {
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}
