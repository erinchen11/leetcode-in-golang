# Contains duplicate 2

### Question
Given an integer array nums and integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

### Example1 :
Input: nums = [1,2,3,1], k = 3
Output: true


### Example 2 :
Input: nums = [1,0,1,1], k = 1
Output: true

### Example 3 :
Input: nums = [1,2,3,1,2,3], k = 2
Output: false




### Constraints:

### 思路
用map保存出現過的元素，找到後去判斷是否小於K

### Hint


#### Topics : Hash Table
#### Difficulty : Easy



