package main

func main() {

	funcs := map[string]func(){"leetcode1": leetcode1, "leetcode6": leetcode6, "leetcode7": leetcode7,
		"leetcode8": leetcode8, "leetcode9": leetcode9, "leetcode10": leetcode10, "leetcode11": leetcode11,
		"leetcode12": leetcode12, "leetcode13": leetcode13, "leetcode14": leetcode14, "leetcode15": leetcode15,
		"leetcode16": leetcode16, "leetcode17": leetcode17, "leetcode18": leetcode18, "leetcode19": leetcode19,
		"leetcode22": leetcode22, "leetcode23": leetcode23, "leetcode24": leetcode24, "leetcode25": leetcode25,
		"leetcode28": leetcode28, "leetcode29": leetcode29, "leetcode30": leetcode30, "leetcode31": leetcode31,
		"leetcode32": leetcode32, "leetcode33": leetcode33, "leetcode34": leetcode34, "leetcode35": leetcode35,
		"leetcode36": leetcode36, "leetcode37": leetcode37, "leetcode38": leetcode38, "leetcode39": leetcode39,
		"leetcode40": leetcode40, "leetcode41": leetcode41, "leetcode42": leetcode42, "leetcode43": leetcode43,
		"leetcode44": leetcode44, "leetcode45": leetcode45, "leetcode46": leetcode46, "leetcode47": leetcode47,
		"leetcode48": leetcode48, "leetcode49": leetcode49, "leetcode50": leetcode50, "leetcode51": leetcode51,
	}

	now := "leetcode51"
	funcs[now]()
}
