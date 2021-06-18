# First Unique Character in a string

### Question
Given a string, return the first non-repeating character in it and return its index.If it does not exist, return -1.

### Example1 :
Input: s = "leetcode"
Output: 0

### Example 2 :
Input: s = "loveleetcode"
Output: 2

### Example 3 :
Input: s = "aabb"
Output: -1

### Constraints:
- 1 <= s.length <= 10^5
- s consists of only lowercase English letters

### 思路
先把 s轉成全小寫
用map建立每個字元和其出現次數的映射關係
然後按順序遍歷string，找到map中第一個出現次數為1的字元返回其index

### Hint


#### Topics : Hash Table
#### Difficulty : Easy



