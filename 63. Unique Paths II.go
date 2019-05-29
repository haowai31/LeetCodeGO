package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var dp = [][]int{}
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp = make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			dp[0][0] = 1
			continue
		}
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] != 1 {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				var tmp = dp[i-1][j] + dp[i][j-1]
				if tmp > dp[i][j] {
					dp[i][j] = tmp
				}
			} else {
				dp[i][j] = 0
			}

		}
	}

	return dp[m-1][n-1]
}

func leetcode63() {
	var obstrcleGrid = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	println(uniquePathsWithObstacles(obstrcleGrid))
}
