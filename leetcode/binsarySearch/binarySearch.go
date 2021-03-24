package leetcode

/*
Given a sorted (in ascending order) integer array nums of n elements
and a target value, write a function to search target in nums.
If target exists, then return its index, otherwise return -1.
*/

// func binarySearch(nums []int, target int) int {

// 	index := -1
// 	for k, v := range nums {
// 		if v == target {
// 			index = k
// 		}
// 	}
// 	return index
// }

func binarySearch(nums []int, target int) int {
	left, pivot := 0, 0
	right := len(nums) - 1

	if nums == nil || len(nums) == 0 {
		return -1
	}

	for left <= right {
		pivot = (left + right) / 2

		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1

}
