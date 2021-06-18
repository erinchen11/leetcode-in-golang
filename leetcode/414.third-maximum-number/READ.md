# Third Maximum Number

### Question
Given integer array nums, return the third maximum number in this array. If the third maximum does not exist, return the maximum number.



### Example1 :
Input: nums = [3,2,1]
Output: 1
Explanation: The third maximum is 1.

### Example 2 :
Input: nums = [1,2]
Output: 2
Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
### Example 3 :
Input: nums = [2,2,3,1]
Output: 1
Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.



### Constraints:

- 1 <= nums.length <= 104
- -2^31 <= nums[i] <= 2^31 - 1

### 思路
先找出第一大,再找出第二大, 第三大
需有三個變數儲存:
first, second, third
先遍歷過一次nums, 如果找到的值大於first,則該值 = first
同時first, second, third三個數往前推
即 second = first, third = second
需注意嚴格小於與大於

note: 因為constraints, 要用64位元的數去計算
### Hint


#### Topics : Array
#### Difficulty : Easy



