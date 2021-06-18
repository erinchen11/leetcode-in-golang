package leetcode

import (
	"sort"
)

func sortedSquare(nums []int) []int {

	
	temp := make([]int, len(nums))

	for i := 0; i <len(nums); i++{
		temp[i] = nums[i] * nums[i]
	}
	sort.Ints(temp)
	return temp
	// error : index out of range [0] with length 0
	// solve : make([]int, len(nums))
	//因為 slice初始化是動態的，如果是用 var []int 
	// length是動態的會報錯
}
