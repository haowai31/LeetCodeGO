package main

func main() {

	funcs := map[string]func(){"leetcode1": leetcode1, "leetcode6": leetcode6, "leetcode7": leetcode7,
		"leetcode8": leetcode8, "leetcode9": leetcode9, "leetcode10": leetcode10, "leetcode11": leetcode11,
		"leetcode12": leetcode12, "leetcode13": leetcode13, "leetcode14": leetcode14, "leetcode15": leetcode15,
		"leetcode16": leetcode16, "leetcode17": leetcode17,
	}

	now := "leetcode17"
	funcs[now]()
}
