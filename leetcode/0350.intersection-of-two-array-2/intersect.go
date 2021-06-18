package intersectionoftwoarray2

func intersect(nums1, nums2 []int) []int {

	// if len(nums1) > len(nums2){
	//     nums1, nums2 = nums2, nums1
	// }

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// use map to store nums1
	record := make(map[int]int)
	result := []int{}
	for _, v := range nums1 {
		record[v]++
	}

	for _, v := range nums2 {
		if record[v] > 0 {
			result = append(result, v)
			// 刪除重複的
			record[v]--
		}
	}

	return result

}
