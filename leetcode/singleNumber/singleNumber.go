package leetcode

/*
Given a non-empty array of integers nums,
every element appears twice except for one.
Find that single one.

Follow up: Could you implement a solution with a linear runtime
complexity and without using extra memory?

給一個數列，找只出現一次的數字
*/

// func singleNumber(nums []int) int {

// 	newMap := make(map[int]int)
// 	for _, v := range nums {
// 		newMap[v]++
// 	}

// 	for k, v := range newMap {
// 		if v == 1{
// 			return k
// 		}
// 	}
// 	return -1

// }

//用 XOR的想法去寫

func singleNumber(nums []int) int {
	res := 0

	for _, v := range nums {
		res = res ^ v
	}
	return res
}
