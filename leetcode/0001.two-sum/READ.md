# Two Sum

### Question
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

### Example1 :
Input: nums = [2,7,11,15], target = 9

Output: [0,1]

Output: Because nums[0] + nums[1] == 9.

### Example 2 :
Input: nums = [3,2,4], target = 6

Output: [1,2]

### Example 3 :
Input: nums = [3,3], target = 6

Output: [0,1]

### Constraints:
- 2 <= nums.length <= 103
- 109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.

### 思路

a + b = target 想成 a = target - b

建立一個hash table，在一個for迴圈中利用差值所得到的數值，用來查詢hash table中是否存在，有的話返回該index
#### Topics : Array, Hash table
#### Difficulty : Easy



