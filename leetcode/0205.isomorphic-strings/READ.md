# Isomorphic strings

### Question
Given two strings s and t, determine if they art isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No Two characters may map to the same character, but a character may map to itself.

### Example1 :
Input: s = "egg", t = "add"
Output: true

### Example 2 :
Input: s = "foo", t = "bar"
Output: false

### Example 3 :
Input: s = "paper", t = "title"
Output: true


### Constraints:
- 1 <= s.length <= 5 * 104
- t.length == s.length
- s and t consist of any valid ascii character.

### 思路
跟290很類似
同構字串，原本字串中的每個字元可以由另外一個字元替代
相同的字元一定要被同一個字元替代，且一個字元不能被多個字元替代，即不能出現一對多映射
是一對一映射，根據此 需要用兩個 HashMap分別用來紀錄原本字串與目標字串中 字元的出現情況

由於 ASCII碼只有256個字元, 可以用byte
如果可以映射則輸出true, 反之則輸出false

創建兩個 map, map[byte]byte{}, 用來建例雙向映射關係
將原本的輸入字串轉成 byte型別, 

再遍歷原本的字串, 遍歷的同時, 檢查兩個map的映射關係
如果沒有找到對應關係，則回傳false
如果有則該map的字元換成目標字元

example2 : s = "foo", t = "bar"

以s為基準建立一個map去映射t
1. s = 'foo'
   sMap = map[byte]byte{} //空的map

2. s[0] = 'f'
    //在sMap中沒有找到'f',就把該值與映射值存入
   sMap = {'f':'b'} // s[0]去映射t[0]
   
   
3. s[1] = 'o'
   s2t = {'f':'b', 'o':'a'} // s[1]去映射t[1]

4. s[2] = 'o'
   sMap = {'f':'b', 'o':'a'} // s[2]去映射t[2]
   但是 t[2]是 'r', 不是'a'
   所以return false

以t為基準建立一個map去映射s, 才能達到雙向映射
1. t = 'bar'
   t2s = map[byte]byte{}

2. t[0] = 'b'
   t2s = {'b':'f'}

3.  t[1] = 'a'
    t2s = {'b':'f', 'a':'o'}

4. t[2] = 'r'
    t2s = {'b':'f', 'a':'o', 'r':'o'}
    return true ,但實際上是return false

### Hint


#### Topics : Array
#### Difficulty : Easy



