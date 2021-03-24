package leetcode

/*
Given an integer array nums, find the contiguous subarray
(containing at least one number)
which has the largest sum and return its sum

range nums一次，記錄當前的最大值，當最大值小於0時，重新開始
*/

func maximumSubarray_error(nums []int) int {
	current := nums[0]
	sum := nums[0]

	for _, v := range nums {
		if sum >= 0 {
			sum += v
		} else {
			sum = v
		}
		if current < sum {
			current = sum
		}
	}
	return current

}

func maximumSubarray(nums []int) int {
	current := nums[0]
	sum := nums[0]
	
	for i := 1; i < len(nums); i++ {
		if sum >= 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if current < sum {
			current = sum
		}

	}
	return current
}
