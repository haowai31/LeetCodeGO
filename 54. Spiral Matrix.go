package main

func spiralOrder(matrix [][]int) []int {
	var result = []int{}
	var mapp = []map[int]bool{}
	var count = 0
	var n = len(matrix)
	if n == 0 {
		return result
	}
	var m = len(matrix[0])
	for i := 0; i < n; i++ {
		mapp = append(mapp, map[int]bool{})
	}
	var step = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	var ix, iy = 0, 0
	var v = 0
	result = append(result, matrix[0][0])
	mapp[0][0] = true
	for {
		for {
			tmpx := ix + step[v][0]
			tmpy := iy + step[v][1]
			if tmpx < 0 || tmpx >= n || tmpy < 0 || tmpy >= m {
				break
			}
			if mapp[tmpx][tmpy] == true {
				break
			}
			mapp[tmpx][tmpy] = true
			result = append(result, matrix[tmpx][tmpy])
			ix = tmpx
			iy = tmpy
			count = 0
		}
		v = (v + 1) % 4
		count += 1
		if count >= 4 {
			break
		}
	}
	return result
}

func leetcode54() {
	var matrix = [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 8},
		{7, 8, 9, 12},
	}
	println(spiralOrder(matrix))
}
