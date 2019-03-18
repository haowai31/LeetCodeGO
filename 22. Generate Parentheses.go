package main

func getresult(left int, right int, result *[]string, now string) {
	if right == 0 {
		*result = append(*result, now)
		return
	}
	if left > 0 {
		getresult(left-1, right, result, now+"(")
		if left < right {
			getresult(left, right-1, result, now+")")
		}
	} else {
		getresult(left, right-1, result, now+")")
	}
	return
}

func generateParenthesis(n int) []string {
	result := []string{}
	getresult(n, n, &result, "")
	return result
}

func leetcode22() {
	//leetcode22
	n := 3
	generateParenthesis(n)
}
