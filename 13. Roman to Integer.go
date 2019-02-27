package main

func romanToInt(s string) int {
	var dp map[string]int = make(map[string]int)
	var result = 0
	dp["I"] = 1
	dp["IV"] = 4
	dp["V"] = 5
	dp["IX"] = 9
	dp["X"] = 10
	dp["XL"] = 40
	dp["L"] = 50
	dp["XC"] = 90
	dp["C"] = 100
	dp["CD"] = 400
	dp["D"] = 500
	dp["CM"] = 900
	dp["M"] = 1000


	var index = 0
	for {
		if index >= len(s) {
			break
		}
		if index + 1 < len(s) {
			value, OK := dp[s[index:index+2]]
			if OK {
				result += value
				index += 2
				continue
			}
		}
		value, ok := dp[s[index:index+1]]
		if ok {
			result += value
			index += 1
			continue
		}
	}
	return result
}

func leetcode13() {
	var s = "LVIII"
	println(romanToInt(s))
}