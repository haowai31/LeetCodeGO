package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var result = [][]int{}
	sort.Ints(candidates)
	combinationSumhelper(candidates, []int{}, target, 0, 0, &result)
	return result
}

func combinationSumhelper(candidates []int, subset []int, target, startIndex, sum int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		combinationSumhelper(candidates, append(subset, candidates[i]), target, i, sum+candidates[i], result)
	}
}

func leetcode39() {
	var candidates = []int{10, 1, 2, 7, 6, 1, 5}
	var target = 8
	var result = [][]int{}
	result = combinationSum(candidates, target)
	println(result)
}
