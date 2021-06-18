# 

### Question
Given a string s, find the length of the longest substring without repeating characters.

### Example1 :

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

### Example 2 :
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

### Example 3 :
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

### Example 4 :
Input: s = ""
Output: 0

### Constraints:
0 <= s.length <= 5 * 10^4
s consists of English letters, digits, symbols and spaces.
### 思路

利用s[left:i+1]來表示 s[:i+1]中的包含s[i]的最長字串
location[s[i]]是字元s[i]在s[:i+1]倒數第二次出現的序列號
當left < location[s[i]]的時候，說明字元s[i]出現了兩次
需要設置left = location[s[i]] + 1，保證字元s[i]只出現一次



#### Topics : 
#### Difficulty : medium



