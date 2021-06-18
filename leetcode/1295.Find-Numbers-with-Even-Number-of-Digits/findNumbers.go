package leetcode

import "strconv"

func FindNumbers(nums []int) int {

	res := 0

	for _, v := range nums {
		if calculateDigits(v) %2 == 0 {
			res++
		}

	}
	return res

}


func calculateDigits(num int) int {
	result := strconv.Itoa(num)
	return len(result)
}