package main

func longestValidParentheses(s string) int {
	var result = 0

	var dp []int = make([]int, len(s)+1)
	if len(dp) == 0 {
		return 0
	}
	dp[0] = 0
	for i := 1; i <= len(s); i++ {
		j := i - dp[i-1] - 2
		if j < 0 || s[i-1] == '(' || s[j] == ')' {
			dp[i] = 0
			continue
		} else {
			dp[i] = dp[i-1] + 2 + dp[j]
			if result < dp[i] {
				result = dp[i]
			}
		}

	}

	return result
}

func leetcode32() {
	var s = ")()())"
	println(longestValidParentheses(s))
}
