/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.33%)
 * Total Accepted:    480.9K
 * Total Submissions: 2.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
package main

import "sort"

func threeSum(nums []int) [][]int {
	var result = [][]int{}

	sort.Ints(nums)
	for i := 0;i<len(nums)-2;i++ {
		if i >0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right{
			target := nums[i] + nums[left] + nums[right]
			if target == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				result = append(result, tmp)
				left += 1
				for nums[left] == nums[left-1] {
					left += 1
					if left >= len(nums) {
						break
					}
				}
				right -= 1
			} else {
				if target > 0 {
					right -= 1
				} else {
					left += 1
					for nums[left] == nums[left-1] {
						left += 1
						if left >= len(nums) {
							break
						}
					}
				}
			}
		}
	}
	return result
}

func leetcode15() {
	//leetcode15
	var nums = []int{-1, 0, 1, 2, -1, -4}
	println(threeSum(nums))
}
