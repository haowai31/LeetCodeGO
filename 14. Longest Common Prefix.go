package main

func longestCommonPrefix(strs []string) string {
	var result = ""
	var flag = 0
	if len(strs) == 0{
		return ""
	}

	for i :=0;i<len(strs[0]);i++ {
		for j:=1;j<len(strs);j++ {
			if len(strs[j]) <= i{
				flag = 1
				break
			}
			if strs[0][i] != strs[j][i] {
				flag = 1
				result = strs[0][:i]
				break
			}
		}
		if flag == 0 {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	if result == "" && flag == 0 {
		result = strs[0]
	}

	return result
}

func leetcode14() {
	//leetcode14
	var strs = []string{"aa", "a"}
	println(longestCommonPrefix(strs))
}