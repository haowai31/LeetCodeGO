package main

func permute(nums []int) [][]int {
	var result = [][]int{}
	permuteHelper(nums, []int{}, &result)
	return result
}

func permuteHelper(nums []int, now []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, now)
		return
	}
	for i, v := range nums {
		tmp := append([]int{}, nums...)
		tmpv := append([]int{}, now...)
		permuteHelper(append(tmp[:i], tmp[i+1:]...), append(tmpv, v), result)
	}
}

func leetcode46() {
	var nums = []int{1, 2, 3}
	println(len(permute(nums)))
}
