package leetcode

func TwoSum(nums []int, target int) []int {
	//建立一個 hash table,用來存放 nums的數值跟 index
	table := make(map[int]int, len(nums))

	//以for迴圈去輪詢nums，同時將數據用map存起來
	//翻轉 key和value,建立數字和位置之間的映射
	//並用差值target-value,去找hash table中有無存在的值，若有就返回該index

	for k, v := range nums {
		index, ok := table[target-v]
		if ok {
			return []int{index, k}
		}
		//如果不存在，則建立映射
		table[v] = k
	}
	return nil

}
