package removeduplicatesfromsortedarray

import "sort"

func RemoveDuplicates(nums []int) int {
	
	sort.Ints(nums)
	// 用兩個 pointer去跑array
	x, y := 0, 1

	// for y < len(nums){
	// 	if nums[x] == nums[y]{
	// 		y++
	// 	}else {
	// 		//如果 num[x]與nums[y]不一樣，則x右移一個
	// 		//並將當前的y值存到新的x位置
	// 		x++
	// 		nums[x] = nums[y]
	// 		// y再去下一個位置
	// 		y++
	// 	}
	// }
	
	for i := 0; i < len(nums) -1; i++{
		//如果 num[x] == num[y]，y向右移一個，y+1
		if nums[x] == nums[y]{
			y++
		}else {
			//如果 num[x]與nums[y]不一樣，則x右移一個
			//並將當前的y值存到新的x位置
			x++
			nums[x] = nums[y]
			// y再去下一個位置
			y++
		}
	}

	//因為x永遠指不到最後一個元素，所以要+1
	return x+1
}