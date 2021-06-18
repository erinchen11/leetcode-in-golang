#  Height Checker

### Question
A school is trying to take an annual photo of all the students. the students are asked to stand in a single file line in non-decreasing order by height.
Let this ordering be represented by the integer array
expected where expected[i] is the expected height of the i th student in line.

You artr given an integer array height representing the current order thate the students are standing in.
Each heights[i] is the height of the i th student in line (0-indexed).

Return the number of indices where heights[i] != expected[i]



### Example1 :
Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation: 
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.

### Example 2 :
Input: heights = [5,1,2,3,4]
Output: 5
Explanation:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.

### Example 3 :
Input: heights = [1,2,3,4,5]
Output: 0
Explanation:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.



### Constraints:

- 1 <= heights.length <= 100
- 1 <= heights[i] <= 100

### 題意
給定一個未排序的原始array heights,
將heights以遞增方式排序後的array為 expected
兩個array互相比較
相同位置的值不相等時,另外記錄起來, 看總共紀錄了幾個



### 思路
先將heights遞增排序為expected

用 two pointer方法
i去遍歷array heights和expected
x去紀錄當兩個相同位置的值不同出現之次數


### Constraints
- 1 <= heights.length <= 100
- 1 <= heights[i] <= 100


#### Topics : Array
#### Difficulty : Easy



