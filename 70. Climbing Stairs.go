package main

func climbStairs(n int) int {
	var dp = []int{}
	dp = append(dp, 0)
	dp = append(dp, 1)
	dp = append(dp, 2)
	if n <= 2 {
		return dp[n]
	}
	var count = 2
	for {
		count += 1
		dp = append(dp, dp[count-1]+dp[count-2])
		if count >= n {
			break
		}
	}
	return dp[n]
}

func leetcode70() {
	var n = 5
	println(climbStairs(n))
}
