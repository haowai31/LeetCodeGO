package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result = [][]int{}
	sort.Ints(candidates)
	combinationSumhelper2(candidates, []int{}, target, 0, 0, &result)
	return result
}

func combinationSumhelper2(candidates []int, subset []int, target, startIndex, sum int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, subset...))
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		if i != startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		combinationSumhelper2(candidates, append(subset, candidates[i]), target, i+1, sum+candidates[i], result)
	}
}

func leetcode40() {
	var candidates = []int{10, 1, 2, 7, 6, 1, 5}
	var target = 8
	var result = [][]int{}
	result = combinationSum2(candidates, target)
	println(result)
}
