package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) (c []int) {
	

	//用最長的nums的長度
	record := make(map[int]bool)
	//把 nums1的值存入record
	for _, v := range nums1 {
		record[v] = true
	}

	//再去range nums2去比較 record
	// 如果找到一樣的值,就刪除record中的該值，並把該值加入到新slice
	for _, item := range nums2 {
		if _, ok := record[item]; ok{
			c = append(c, item)
			delete(record, item)
			
		}
	}
	return

}