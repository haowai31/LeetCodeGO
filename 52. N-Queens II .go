package main

import (
	"strconv"
	"strings"
)

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	var res = []string{}
	solveNQueensHelper(&res, n, "")

	return len(res)
}

func totalNQueensHelper(res *[]string, n int, now string) {
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
		totalNQueensHelper(res, n, now+strconv.Itoa(i))
	}
}

func leetcode52() {
	var n = 4
	print(totalNQueens(n))
}
