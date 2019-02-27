package main
func twoSum(nums []int, target int) []int {
	for i, v := range(nums){
		for j, v1 := range (nums) {
			if (i != j) && (v + v1 == target){
				result := []int{i, j}
				return result
			}
		}
	}
	return nil
}

func leetcode1() {
	//leetcode 1
	var nums = []int{2, 7, 13, 15}
	target := 9
	var result = twoSum(nums, target)
	println(result[0], result[1])
}