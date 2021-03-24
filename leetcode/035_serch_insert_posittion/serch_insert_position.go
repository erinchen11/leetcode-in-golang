package leetcode

func searchInsert(nums []int, target int) int {

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] >= target {
	// 		return i
	// 	} else {
	// 		if i == len(nums)-1 {
	// 			return i + 1
	// 		}
	// 	}
	// }
	// return len(nums) - 1

	//思路：二元搜尋法
	// low, up := 0, len(nums)-1
	// mid := (low + up)/2

	// if target < nums[0]{
	// 	return 0
	// }
	// if target > nums[up]{
	// 	return up + 1
	// }

	// for low <= up{
	// 	if nums[mid] == target{
	// 		return mid
	// 	}
	// 	if nums[mid] < target{
	// 		low = mid + 1
	// 	}else {
	// 		up = mid
	// 	}
	// }

	// // low = up
	// return low
	low, up := 0, len(nums)-1
	// mid := (low + up) / 2
	if target > nums[up] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}
	for low < up {
		mid := (low + up) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			up = mid
		}
	}
	return low
}
