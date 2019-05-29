package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var result = [][]int{}
	var index = 0
	var flag = 0
	for index = range intervals {
		if newInterval[0] < intervals[index][0] {
			flag = 1
			break
		}
	}
	if flag == 0 {
		intervals = append(intervals, newInterval)
	} else {
		tmp := append([][]int{}, intervals[index:]...)
		intervals = append(intervals[0:index], newInterval)
		intervals = append(intervals, tmp...)
	}

	for i := range intervals {
		if i == 0 {
			result = append(result, intervals[i])
		}
		if intervals[i][0] <= result[len(result)-1][1] {
			if result[len(result)-1][1] < intervals[i][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func leetcode57() {
	var intervals = [][]int{
		{1, 5},
	}
	var newInterval = []int{0, 3}

	println(insert(intervals, newInterval))
}
