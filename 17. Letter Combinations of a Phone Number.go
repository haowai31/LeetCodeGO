package main

func findlettercombinations(dict map[string]string, digits string, result *[]string, value string) {
	if len(digits) == 0 {
		return
	}
	for _, vj := range dict[string(digits[0])] {
		tmp := value + string(vj)
		if len(digits) == 1 {
			*result = append(*result, tmp)
		}
		findlettercombinations(dict, digits[1:], result, tmp)
	}
	return
}

func letterCombinations(digits string) []string {
	var result = []string{}
	var dict map[string]string = make(map[string]string)
	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"

	findlettercombinations(dict, digits, &result, "")

	return result
}

func leetcode17() {
	//leetcode17
	var s = "23"
	println(letterCombinations(s))
}
