package main

func lengthOfLastWord(s string) int {
	var index = 0
	var result = ""
	var last = ""
	var flag = 0
	for index = range s {
		if s[index] != ' ' {
			flag = 1
			break
		}
	}
	if flag != 1 || len(s) == 0 || s == " " {
		return 0
	}
	for i := index; i < len(s); i++ {
		if s[i] != ' ' {
			result += string(s[i])
		} else {
			if result != "" {
				last = result
				result = ""
			}
		}
	}
	if result != "" {
		return len(result)
	} else {
		return len(last)
	}
}

func leetcode58() {
	var s = "a"
	println(lengthOfLastWord(s))
}
