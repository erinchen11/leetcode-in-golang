package movezeroes

func MoveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return []int{0}
	}

	x := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[x] = nums[x], nums[i]
			x++
		}
	}
	return nums

}
