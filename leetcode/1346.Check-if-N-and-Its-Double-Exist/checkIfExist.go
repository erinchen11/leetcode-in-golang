package checkifnanditsdoubleexist

func CheckIfExist(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}

	set := make(map[int]int)

	//把 array存在 set裡面
	for i,_ := range nums{
		set[nums[i]] = i
	}

	//再遍歷nums和set裡面的數字比較
	// set裡面的值*2
	for i,_:= range nums{
		if j, ok := set[nums[i]*2]; ok && j != i {
			return true
		}
	}
	return false
	



}
