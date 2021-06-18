package containsduplicate2

func containsNearbyDuplicate(nums []int, k int) bool {
	// if length of nums < 2 , return false
	if len(nums) < 2{
		return false
	}
	// use map to store 
	record := make(map[int]int)

	for i, v := range nums {
		// 從map中找到跟nums一樣的數字,就用nums和map當前的index相減判斷是否小於等於k
		if n, ok := record[v]; ok && (i - n ) <= k{
			return true
		}
		//沒有找到就紀錄數字
		record[v] = i
	}

	return false



}