# Reverse Integer

### Question
Given a signed 32-bit integer x, return x with its digits reversed. 
If reversing x causes the value to go outside the signed 32-bit integer
range [-2^31, 2^31 - 1], then return 0.



### Example1 :

Input: x = 123
Output: 321
Explanation: 342 + 465 = 807.

### Example 2 :
Input: x = -123
Output: -321

### Example 3 :
Input: x = 120
Output: 21

### Example 4 :
Input: x = 0
Output: 0


### Constraints:

- -2^31 <= x <= 2^31 - 1


### 思路

反轉時要注意數字範圍必須要在 [-2^31 <= x <= 2^31 - 1]
超過就回傳0

#### Topics : 
#### Difficulty : medium



