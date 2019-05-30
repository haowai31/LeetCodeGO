package main

import "strconv"

func getPermutation(n int, k int) string {
	var x, y = make([]int, n), make([]int, n)
	var tmp = 1
	var result = ""
	y[0] = 1
	x[0] = 1
	for i := 1; i < n; i++ {
		tmp *= i
		x[i] = tmp
		y[i] = i + 1
	}
	var now = n
	for now > 0 {
		var tt = k / x[now-1]
		var yf = k % x[now-1]
		if tt > 0 && yf == 0 {
			result += strconv.Itoa(y[tt-1])
			y = append(y[:tt-1], y[tt:]...)
		}
		if tt >= 0 && yf != 0 {
			result += strconv.Itoa(y[tt])
			y = append(y[:tt], y[tt+1:]...)
		}
		if tt == 0 && yf == 0 {
			maxx := y[len(y)-1]
			result += strconv.Itoa(maxx)
			y = append(y[:len(y)-1])
		}

		k = yf
		now -= 1
	}
	return result
}

func leetcode60() {
	var n, k = 3, 1
	println(getPermutation(n, k))
}
