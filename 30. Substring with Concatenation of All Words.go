package main

func findSubstring(s string, words []string) []int {
	result := []int{}
	var wordlist = map[string]int{}
	if len(words) == 0 || len(s) == 0 {
		return result
	}
	for _, a := range words {
		_, ok := wordlist[a]
		if ok {
			wordlist[a] += 1
		} else {
			wordlist[a] = 1
		}
	}
	lens := len(words[0])
	for index := 0; index <= len(s)-lens*len(words); index++ {
		if compare30(index, lens, index+lens*len(words), s, wordlist) {
			result = append(result, index)
		}
		if index > len(s)-lens*len(words) {
			break
		}
	}

	return result
}

func compare30(index int, lens int, end int, s string, wordtmp map[string]int) bool {
	d := make(map[string]int)
	for i := index; i < end; i += lens {
		tmp := s[i : i+lens]
		d[tmp] = getoccurence(d, tmp) + 1
		if d[tmp] > getoccurence(wordtmp, tmp) {
			return false
		}
	}

	return true
}

func getoccurence(dict map[string]int, s string) int {
	if o, ok := dict[s]; ok {
		return o
	} else {
		return 0
	}
}

func leetcode30() {
	//leetcode30
	s := "a"
	var words = []string{"a", "a"}
	println(findSubstring(s, words))
}
