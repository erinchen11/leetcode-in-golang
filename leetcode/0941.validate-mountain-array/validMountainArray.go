package validatemountainarray

func ValidMountainArray(arr []int) bool {
	// 因為山形至少要3個元素
	if len(arr) < 3 {
		return false
	}
	//用兩個pointer開始找
	i, j := 0, len(arr)-1

	//目的是要找到i跟j指向同一個地方，如果是的話就回傳true
	for i != j {
		//從尾巴開始遍歷
		//左邊的數字比右邊大時，j左移一次
		if arr[j-1] > arr[j] {
			j--
			continue
		}
		//從首開始遍歷
		//右邊的數字比左邊大時，i右移一次
		if arr[i+1] > arr[i] {
			i++
			continue
		}

		return false
	}
	//如果都找不到
	if i == 0 || j == len(arr)-1 {
		return false
	}

	return true

}
