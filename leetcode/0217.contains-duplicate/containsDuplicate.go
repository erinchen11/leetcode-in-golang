package containsduplicate

func containsDuplicate(nums []int) bool {
	// use map to record bool
	record := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++{
		//從 record去找
		if _, found := record[nums[i]]; found{
			return true
		}
		record[nums[i]] = true
	}

	return false
}
