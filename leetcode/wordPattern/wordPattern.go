package leetcode

import "strings"

/*

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is
a bijection between a letter in pattern and a non-empty word in s.

    1 <= pattern.length <= 300
    pattern contains only lower-case English letters.
    1 <= s.length <= 3000
    s contains only lower-case English letters and spaces ' '.
    s does not contain any leading or trailing spaces.
    All the words in s are separated by a single space.

*/

/*
思路：每個pattern中的字元，只能對應到s中的一個字串，是一對一的關係
*/

func wordPattern(pattern string, s string) bool {
	//先把s字串以空格為間隔做字串取出
	strList := strings.Split(s, " ")
	patternByte := []byte(pattern)
	//字元數和字串個數不合時，就回傳false
	if pattern == "" || len(patternByte) != len(strList) {
		return false
	}
	/* 用2個map去實現 pattern和s的一對一關係
	記錄雙向關係：a 對應了dog,如果b再對應dog是錯誤的
	所以需要s的dog去查詢它是否已經和pattern中某個字元匹配過了
	1個map記錄 pattern和s的匹配關係
	另一個map記錄s和pattern匹配關係
	*/

	pMap := make(map[byte]string)
	sMap := make(map[string]byte)

	for index, b := range patternByte {
		////輪詢patternByte，如果
		if _, ok := pMap[b]; !ok {
			if _, ok := sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true

	
}
