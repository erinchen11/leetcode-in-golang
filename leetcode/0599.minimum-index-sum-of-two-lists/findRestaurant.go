package minimumindexsumoftwolists

func findRestaurant(list1, list2 []string) []string {
	// 創建1個map, 1個slice
	mapLike := map[string]int{}
	restaurants := []string{}

	for i, r := range list1 {
		mapLike[r] = i
	}

	for j, r := range list2{
		// 如果list2 和 mapLike所存的值一樣,則 mapLike的
		if _, exist := mapLike[r]; exist{
			//計算索引值之和S
			mapLike[r] += j
			if len(restaurants) == 0 || mapLike[r] == mapLike[restaurants[0]]{
				restaurants = append(restaurants, r)
			}else if mapLike[r] < mapLike[restaurants[0]] {
				restaurants = []string{r}
			}
		}
	}

	return restaurants

}
