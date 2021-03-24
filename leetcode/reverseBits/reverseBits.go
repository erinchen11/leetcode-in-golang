package leetcode

/*
Reverse bits of a given 32 bits unsigned integer.
思路 ： input num 從高位（第31位）開始翻轉
res從低位開始翻轉，如果該位是1，則res+1
該位是0，則res+0
把32位全部range過

num & 1 === 1, 说明n的最后一位是1
num & 1 === 0, 说明n的最后一位是0
*/

func reverseBits(num uint32) uint32 {

	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res <<= 1
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res

}
