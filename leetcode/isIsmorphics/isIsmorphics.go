package leetcode

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters
in s can be replaced to get t.

All occurrences of a character must be replaced with
another character while preserving the order of characters.
No two characters may map to the same character,
but a character may map to itself.

同構

思路：
建立兩個map, byteMap用來儲存字符的對應關係，另一個map用來儲存被映射過的字符，保證一個字符不會被映射兩次

1.遍歷s,t，判斷s[i]是否存在byteMap中，如果存在，則比較映射值和t[i]是否相等
如果相等則繼續下一次的遍歷，如果不相等，則兩字串不同構

2.如果s[i]不存在byteMap中，需要先判斷s[i], t[i]是否相等，如果相等，
把新增映射關係到map中，如果不相等，需要判断t[i]是否存在于tMap中，
如果存在，说明t[i]已经映射过一次，不能再次映射，返回false；如果不存在，
则把新增映射关系到Map中
如果所有s[i] -> t[i]都能保证一一映射，则返回true


*/

func isIsomorphics(s string, t string) bool {

	/*
			當字串長度為 N 時，我們至多需要遍歷整個陣列一次，時間複雜度為 O(N)。
		此方法建立了兩個陣列來儲存每個字元對應的數值，字元的ASCII碼數值範圍介於 [1,128]，
		與 N 大小無關，空間複雜度為 O(1)。
	*/

	var sArr, tArr [128]int
	var sB, tB byte

	for i := 0; i < len(s); i++ {
		sB, tB = s[i], t[i]
		if sArr[sB] != tArr[tB] {
			return false
		}
		sArr[sB], tArr[tB] = i+1, i+1
	}
	return true

}
