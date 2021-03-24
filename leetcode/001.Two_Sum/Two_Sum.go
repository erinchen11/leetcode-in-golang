package leetcode

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.
*/
func TwoSum(nums []int, target int) []int {
	// m := make(map[int]int)
	// for k, v := range nums {
	// 	c := target - v
	// 	if i, ok := m[c]; ok {
	// 		return []int{k, i}
	// 	}
	// 	m[v] = k
	// }
	// return nil
	results := []int{}
	valueMap := make(map[int]int)
	for idx, val := range nums {
		value, ok := valueMap[target-nums[idx]]
		if ok && value != idx {
			if idx < value {
				return []int{idx, value}
			} else {
				return []int{value, idx}
			}
		} else {
			valueMap[val] = idx
		}
	}
	return results
}
