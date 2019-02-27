package main

func isMatchDP(s string, p string) bool {
	var dp [][]bool
	dp = make([][]bool, len(s)+1)
	for i:=0; i<=len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i :=0; i<=len(s);i++ {
		for j:=1;j<=len(p);j++{
			if (j > 1 && p[j-1] == "*"[0]) {
				dp[i][j] = dp[i][j-2] || (i > 0 &&(s[i-1] == p[j-2] || p[j-2]=="."[0]) && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == "."[0])
			}
		}
	}
	return dp[len(s)][len(p)]
}

