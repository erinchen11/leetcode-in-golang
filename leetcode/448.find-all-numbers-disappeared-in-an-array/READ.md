# Find all numbers disappeared in an array

### Question
Given an array nums of n integers where nums[i] is in the range [1, n], retuun an array of all the integers in the range [1, n] that do not appear in nums.

### Example1 :
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]


### Example 2 :
Input: nums = [1,1]
Output: [2]


### Constraints:

- n == nums.length
- 1 <= n <= 105
- 1 <= nums[i] <= n

### 思路
需要的變數：n 為陣列長度, res 為新的陣列
在nums裡面操作，
先遍歷一次標記已經出現過的數字，把它標記為負數



### Hint


#### Topics : Array
#### Difficulty : Easy



