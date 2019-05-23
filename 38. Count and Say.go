package main

import "strconv"

func countAndSay(n int) string {
	var result = ""
	result = "1"
	for i := 1; i < n; i++ {
		result = getnext(result)
	}

	return result
}

func getnext(x string) string {
	var result = ""
	var index = 0
	var now = ""
	var count = 0
	for index < len(x) {
		if now == "" {
			now = string(x[index])
			count = 1
			index += 1
			continue
		} else {
			if now == string(x[index]) {
				count += 1
				index += 1
				continue
			} else {
				result += strconv.Itoa(count) + now
				now = string(x[index])
				count = 1
				index += 1
				continue
			}
		}
	}
	result = result + strconv.Itoa(count)
	result = result + now
	return result
}

func leetcode38() {
	var n = 6
	println(countAndSay(n))
}
