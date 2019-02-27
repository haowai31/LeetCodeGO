package main
func intToRoman(num int) string {
	var dp map[int]string = make(map[int]string)
	var result = ""
	dp[1] = "I"
	dp[4] = "IV"
	dp[5] = "V"
	dp[9] = "IX"
	dp[10] = "X"
	dp[40] = "XL"
	dp[50] = "L"
	dp[90] = "XC"
	dp[100] = "C"
	dp[400] = "CD"
	dp[500] = "D"
	dp[900] = "CM"
	dp[1000] = "M"



	var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var index = 0
	for {
		if num == 0{
			break
		}
		var tmp = keys[index]
		for {
			if num < tmp {
				break
			}
			result += dp[tmp]
			num -= tmp
		}
		index += 1
	}
	return result
}

func leetcode12() {
	//leetcode12
	var num = 58
	println(intToRoman(num))
}