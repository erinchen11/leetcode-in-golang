package leetcode

/*
Calculate the sum of two integers a and b,
but you are not allowed to use the operator + and -.

思路：
用位元做計算

*/

func getSum(a, b int) int {

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		//進位值
		temp := (a & b)
		//相加
		a = a ^ b
		//進位
		b = temp << 1
	}
	return a
}
