package main

func generateMatrix(n int) [][]int {
	var result = make([][]int, n)
	if n == 0 {
		return result
	}
	var mapp = []map[int]bool{}
	for i := range result {
		result[i] = make([]int, n)
		mapp = append(mapp, map[int]bool{})
	}
	var step = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	var tor = 0
	var ix = 0
	var iy = 0
	mapp[ix][iy] = true
	var count = 1
	result[ix][iy] = count
	var flag = 0
	for {
		for {
			tmpx := ix + step[tor][0]
			tmpy := iy + step[tor][1]
			if tmpx < 0 || tmpx >= n || tmpy < 0 || tmpy >= n {
				break
			}
			if mapp[tmpx][tmpy] {
				break
			}
			ix = tmpx
			iy = tmpy
			mapp[ix][iy] = true
			count += 1
			result[ix][iy] = count
			flag = 0
		}
		tor = (tor + 1) % 4
		flag += 1
		if flag >= 4 {
			break
		}
	}
	return result
}

func leetcode59() {
	var n = 0
	println(generateMatrix(n))
}
