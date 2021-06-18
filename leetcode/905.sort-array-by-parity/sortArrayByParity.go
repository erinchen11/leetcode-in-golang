package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	if len(nums) == 0 {
		return []int{0}
	}

	x := 0 

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			return []int{0}
		}
		if (nums[i]%2 == 0) {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
		
	}
	return nums

}
