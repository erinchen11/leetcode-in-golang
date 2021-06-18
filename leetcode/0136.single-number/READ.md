# Single Number

### Question
Given a non-empty array of intergers nums, every element appears twice except for one. Find that single on.
### Example1 :
Input: nums = [2,2,1]
Output: 1

### Example 2 :
Input: nums = [4,1,2,1,2]
Output: 4
### Example 3 :
Input: nums = [1]
Output: 1

### Constraints:

- 1 <= nums.length <= 3 * 104
- -3 * 104 <= nums[i] <= 3 * 104
- Each element in the array appears twice except for one element which appears only once.


### 思路
map[int]int
用map去紀錄nums中每個數出現的次數
用nums[i]當map的key,每次紀錄時
map[nums[i]]++
如果同一個數出現兩次以上，則map的key值會大於2
如果一個數只出現一次, 則map的key值會等於1

### Hint


#### Topics : Array
#### Difficulty : Easy



