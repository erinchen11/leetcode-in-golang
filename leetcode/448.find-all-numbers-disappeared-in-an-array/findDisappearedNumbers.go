package findallnumbersdisappearedinanarray

import "sort"

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	x := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > len(nums) {
			return []int{0}
		}

		if nums[i] != i + 1 {
			res[x] = i + 1
			x++
		}
	}

	return res

}
