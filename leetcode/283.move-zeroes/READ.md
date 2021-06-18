# Move Zeroes

### Question
Given an integer array nums, move all 0's to thhe end
of it while maintaining the relative order of the
non-zero elements.

### Example1 :
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

### Example 2 :
Input: nums = [0]
Ouptput: [0]


### Constraints:
- 1 <= nums.length <= 10^4
- -2^31 <= nums[i] <= 2^31-1

### 思路
用two pointer, i去遍歷整個array, x 去紀錄當前要交換的值
用交換的方式去把0移到右邊
0,1,0,3,12
i起始在位置0, x起始位置在0,用nums[i]檢查是否為0
不是0,則nums[i]和nums[x]交換, x+1

### Hint


#### Topics : Array
#### Difficulty : Easy



