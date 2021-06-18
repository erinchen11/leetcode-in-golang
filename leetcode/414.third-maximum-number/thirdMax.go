package thirdmaximumnumber

import "math"

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, i := range nums {
		//先找到第一大
		if i > first {
			third = second
			second = first
			first = i
		} else if i < first && i > second { //再找第二大
			//找到第二大時 記得往前推third
			third = second
			second = i
		} else if i < second && i > third { //再找第三大
			third = i

		}

		//沒有第三大的數, 就是 third一直是0,沒有往前變成second

	}

	if third == math.MinInt64 {
		//那就回傳第一大
		return first
	}
	return third

}
