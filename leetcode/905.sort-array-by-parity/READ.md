# Sort Array By Parity

### Question
Given an array nums of non-negative integers,
return an array consisting of all the even elements
of nums , followed by all the odd elements nums.

### Example1 :
Input: nums = [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

### Example 2 :

### Example 3 :




### Constraints:

### 思路
用 two pointer, i去遍歷array, x去紀錄要比較的位置
i,x起始位置都是0, 檢查 nums[i]是否為偶數
是偶數, 則交換x的位置,x++

### Hint


#### Topics : Array
#### Difficulty : Easy



