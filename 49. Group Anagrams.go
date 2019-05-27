package main

func groupAnagrams(strs []string) [][]string {
	var result = [][]string{}
	var strsmap = map[int]bool{}
	if len(strs) == 1 {
		result = append(result)
	}
	for i := 0; i < len(strs); i++ {
		if strsmap[i] == true {
			continue
		}
		var tmp = []string{}
		var hs = []int{}
		hs = make([]int, 27)
		for xx := range strs[i] {
			var xxx = strs[i][xx] - 96
			hs[xxx] += 1
		}
		for j := i + 1; j < len(strs); j++ {
			if strsmap[j] == true {
				continue
			}
			if len(strs[j]) != len(strs[i]) {
				continue
			}
			hstmp := append([]int{}, hs...)
			var flag = 0
			for xx := range strs[j] {
				hstmp[strs[j][xx]-96] -= 1
				if hstmp[strs[j][xx]-96] < 0 {
					flag = 1
					break
				}
			}
			if flag == 0 {
				tmp = append(tmp, strs[j])
				strsmap[j] = true
			}
		}
		tmp = append(tmp, strs[i])
		strsmap[i] = true
		result = append(result, tmp)
	}
	return result
}

func leetcode49() {
	var strs = []string{"", "", ""}
	result := groupAnagrams(strs)
	println(result)
}
