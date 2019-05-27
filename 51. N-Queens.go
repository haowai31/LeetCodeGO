package main

import (
	"strconv"
	"strings"
)

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	var res = []string{}
	solveNQueensHelper(&res, n, "")
	var result = [][]string{}
	result = make([][]string, len(res))

	for index := range res {
		var tmp = []string{}
		for i := 0; i < n; i++ {
			var once = ""
			for j := 0; j < n; j++ {
				if int(res[index][i]-48) == j {
					once += "Q"
				} else {
					once += "."
				}
			}
			tmp = append(tmp, once)
		}
		result[index] = tmp
	}

	return result
}

func solveNQueensHelper(res *[]string, n int, now string) {
	if len(now) == n {
		*res = append(*res, now)
		return
	}
	for i := 0; i < n; i++ {
		if strings.Contains(now, strconv.Itoa(i)) {
			continue
		}
		var flag = 0
		for j := range now {
			if judgequeenscex(len(now), i, j, int(now[j])-48) {
				flag = 1
				break
			}
		}
		if flag == 1 {
			continue
		}
		solveNQueensHelper(res, n, now+strconv.Itoa(i))
	}
}

func judgequeenscex(lx, ly, rx, ry int) bool {
	v := (ry - ly) / (rx - lx)
	if abs(v) == 1 {
		return true
	} else {
		return false
	}
}

func leetcode51() {
	var n = 4
	print(solveNQueens(n))
}
