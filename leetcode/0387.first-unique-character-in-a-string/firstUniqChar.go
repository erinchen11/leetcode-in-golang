package firstuniquecharacterinastring

import "strings"

func firstUniqChar(s string) int {
	strings.ToLower(s)

	// result := make([]int, 26)

	// for i := 0; i < len(s); i++ {
	// 	result[s[i]-'a'] += 1
	// }

	// for i := 0; i < len(s); i++ {
	// 	if result[s[i]-'a'] == 1 {
	// 		return i
	// 	}

	// }

	// return -1

	sMap := make(map[rune]int)
	
	// 先將 string s 按出現次數存入 map
	for _, v := range s {
		sMap[v]++
	}
	
	for i, v := range s{
		if sMap[v] == 1{
			return i
		}
	}

	return -1


}
