package main

func uniquePaths(m int, n int) int {
	var dp = [][]int{}
	dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var tmp = dp[i-1][j] + dp[i][j-1]
			if tmp > dp[i][j] {
				dp[i][j] = tmp
			}
		}
	}

	return dp[m-1][n-1]
}

func leetcode62() {
	var m, n = 3, 2
	println(uniquePaths(m, n))
}
