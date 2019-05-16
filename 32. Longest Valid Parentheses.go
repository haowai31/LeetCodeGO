package main

func longestValidParentheses(s string) int {
	var result = 0
	var stack []string
	var stackindex []int
	var index = 0
	for {
		if index >= len(s) {
			break
		}
		if s[index] == "(" {
			stack.
		}
	}



	return result
}

func leetcode32() {
	var s = ")()())"
	println(longestValidParentheses(s))
}
