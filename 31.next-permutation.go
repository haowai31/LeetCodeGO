/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (30.01%)
 * Total Accepted:    232.3K
 * Total Submissions: 764.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */
package main

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	var index = -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}

	if index == -1 {
		sort.Ints(nums[:])
		return
	}

	var minn = math.MaxInt32
	var minindex = -1
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			if nums[i] < minn {
				minn = nums[i]
				minindex = i
			}
		}
	}
	tmp := nums[index]
	nums[index] = nums[minindex]
	nums[minindex] = tmp

	sort.Ints(nums[index+1:])
	return
}

func leetcode31() {
	var num = []int{1, 4, 3, 2}
	nextPermutation(num)
	println(num)
}
