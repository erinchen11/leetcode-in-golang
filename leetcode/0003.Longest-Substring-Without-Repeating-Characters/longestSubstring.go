package leetcode

//Given a string s, find the length of the longest
//substring without repeating characters.

// solution 1

// func lengthOfLongestSubstring(s string) int {
// 	maxLen := 0

// 	for i := 0; i < len(s); i++ {
// 		//從第二個字元開始算，並且比較字串
// 		for j := i + 1; j <= len(s); j++ {
// 			subStr := s[i : j-1]
// 			if !strings.Contains(subStr, s[j-1:j]) {
// 				if maxLen < j-i {
// 					maxLen = j - i
// 				} else {
// 					break

// 				}
// 			}
// 		}
// 	}
// 	return maxLen

// }

func lengthOfLongestSubstring(s string) int {

	// location[s[i]] == j表示
	// s中的第i個字元,上次出現在S的j位置
	// 所以在 s[j+1:i]中沒有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出現

	location := [256]int{} // 假定輸入的字元是ASCII字元

	for i := range location {
		location[i] = -1 // 先設置所有的字元都沒有出現過
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		//說明s[i]已經在s[left:i+1]中重複了
		//並且s[i]上次出現的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen

}
