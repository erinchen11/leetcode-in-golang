package maxconsecutiveones

//連續的1不會有0出現
//可以透過兩個0來計算有多少1

func FindMaxConsecutiveOnes(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxCount, currCount := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currCount++
		} else {
			currCount = 0
		}
		if currCount > maxCount {
			maxCount = currCount
		}
	}
	return maxCount

}
