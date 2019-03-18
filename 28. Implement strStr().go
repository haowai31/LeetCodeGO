package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			flag := 0
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					flag = 1
					break
				}
			}
			if flag == 0 {
				return i
			}
		}
	}

	return -1
}

func leetcode28() {
	//leetcode28
	haystack := "h"
	needle := "l"
	println(strStr(haystack, needle))
}
