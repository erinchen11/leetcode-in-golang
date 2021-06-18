package heightchecker

import (
	"sort"
)

func heightChecker(height []int) int {
	if len(height) < 0 || len(height) > 100 {
		return 0
	}

	//複製一個array
	expected := make([]int, len(height))
	copy(expected, height)
	//再把他遞增排序
	sort.Ints(expected)
	//紀錄不同值出現的次數
	x := 0

	for i := 0; i < len(height); i++ {
		if height[i] < 0 || height[i] > 100 {
			return 0
		}

		if expected[i] != height[i] {
			x++
		}
	}

	return x

}
