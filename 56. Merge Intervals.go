package main

import "sort"

type IntSlice [][]int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int) {
	s[i][0], s[j][0] = s[j][0], s[i][0]
	s[i][1], s[j][1] = s[j][1], s[i][1]
}
func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }

func merge(intervals [][]int) [][]int {
	var result = [][]int{}
	sort.Sort(IntSlice(intervals))
	for i := range intervals {
		if i == 0 {
			result = append(result, intervals[i])

			continue
		}
		if result[len(result)-1][1] >= intervals[i][0] {
			if intervals[i][1] > result[len(result)-1][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func leetcode56() {
	var intervals = [][]int{{1, 4}, {4, 5}}
	println(merge(intervals))
}
