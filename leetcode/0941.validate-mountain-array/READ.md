# Valid Mountain Array

### Question
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

### Example1 :
Input: arr = [2,1]
Output: false

### Example 2 :
Input: arr = [3,5,5]
Output: false

### Example 3 :
Input: arr = [0,3,2,1]
Output: true


### Constraints:
- 1 <= arr.length <= 10^4
- 0 <= arr[i] <= 10^4
### 思路

用while迴圈
用i開始遍歷，從array的首嚴格遞增，當當數字小於右邊的數字時，i++
但是為了避免overflow, i只能遍歷到array倒數第二個數字
這樣可找到最後循環結束時,i指到的最大數字
再用j從最後一個數字遍歷到第二個數字，當數字小於左邊的數時, j--
這樣可找到最後循環結束時,j指到的最大數字
再比較i跟j最後指到的數字是否同一個

### Hint


#### Topics : Array
#### Difficulty : Easy



