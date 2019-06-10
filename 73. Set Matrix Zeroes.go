package main

func setZeroes(matrix [][]int) {
	var xzero = map[int]int{}
	var yzero = map[int]int{}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				xzero[i] = 1
				yzero[j] = 1
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if xzero[i] == 1 || yzero[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
	return
}

func leetcode73() {
	var matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}}
	setZeroes(matrix)
}
