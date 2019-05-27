package main

func isMatch44(s string, p string) bool {
	var dp = [][]bool{}
	dp = make([][]bool, len(p)+1)
	for i := 0; i <= len(p); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		dp[i][0] = (p[i-1] == '*') && dp[i-1][0]
		for j := 1; j <= len(s); j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				if p[i-1] == '?' || p[i-1] == s[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[len(p)][len(s)]
}

func leetcode44() {
	var s = "acdcb"
	var p = "*a*b"
	println(isMatch44(s, p))
}
