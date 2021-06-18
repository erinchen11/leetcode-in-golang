package leetcode

func Reverse(x int) int {

	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10

	}
	// <<31 左移31次 相當於 乘2乘31次
	if tmp > (1<<31-1) || tmp < -(1<<31) {
		return 0
	}
	return tmp

}
