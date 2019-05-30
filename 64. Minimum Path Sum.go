package main

func minPathSum(grid [][]int) int {
	var dp = [][]int{}
	var m = len(grid)
	if m == 0 {
		return 0
	}
	var n = len(grid[0])
	if n == 0 {
		return 0
	}
	for i := range grid {
		dp = append(dp, make([]int, n))
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			var tmp1 = grid[i][j] + dp[i-1][j]
			var tmp2 = grid[i][j] + dp[i][j-1]
			if tmp1 > tmp2 {
				dp[i][j] = tmp2
			} else {
				dp[i][j] = tmp1
			}
		}
	}
	return dp[m-1][n-1]
}

func leetcode64() {
	var grid = [][]int{
		{0, 1},
		{1, 0},
	}
	println(minPathSum(grid))
}
