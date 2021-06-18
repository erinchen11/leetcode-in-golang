# Maximum Subarray

### Question
Given an integer array nums, find the contiguous subarray 
(containing at least one number) 
which has the largest sum and return its sum.

### Example1 :
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6


### Example 2 :
Input: nums = [1]
Output: 1

### Example 3 :

Input: nums = [5,4,-1,7,8]
Output: 23


### Constraints:

- 1 <= nums.length <= 3 * 10^4
- -10^5 <= nums[i] <= 10^5
### 思路

使用DP來解

最大總和(res)與目前為止包含當前元素的最大總和(curr)

如果小於0則 curr = nums[i],從nums[i]開始 捨棄原本的curr值
如果curr大於0，則繼續累加


#### Topics : Dynamic programming
#### Difficulty : easy



