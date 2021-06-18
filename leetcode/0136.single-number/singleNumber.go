package singlenumber

func singleNumber(nums []int) int {
	
	// use map to record 紀錄出現的次數
	record := make(map[int]int)

	for i := 0; i < len(nums); i++{
		record[nums[i]]++
	}

	for k, v := range record {
		// 表示只有出現一次
		if v == 1 {
			return k
		}

	}
	return -1


}