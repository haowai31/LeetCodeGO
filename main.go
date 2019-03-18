package main

func main() {

	funcs := map[string]func(){"leetcode1": leetcode1, "leetcode6": leetcode6, "leetcode7": leetcode7,
		"leetcode8": leetcode8, "leetcode9": leetcode9, "leetcode10": leetcode10, "leetcode11": leetcode11,
		"leetcode12": leetcode12, "leetcode13": leetcode13, "leetcode14": leetcode14, "leetcode15": leetcode15,
		"leetcode16": leetcode16, "leetcode17": leetcode17, "leetcode18": leetcode18, "leetcode19": leetcode19,
		"leetcode22": leetcode22, "leetcode23": leetcode23, "leetcode24": leetcode24, "leetcode25": leetcode25,
		"leetcode28": leetcode28,
	}

	now := "leetcode28"
	funcs[now]()
}
