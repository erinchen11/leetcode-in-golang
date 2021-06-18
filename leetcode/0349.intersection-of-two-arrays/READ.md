# Intersection of two arrays

### Question
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

### Example1 :
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]


### Example 2 :
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.



### Constraints:

- 1 <= nums1.length, nums2.length <= 1000
- 0 <= nums1[i], nums2[i] <= 1000

### 思路
map[int]bool
先把最長的nums都存入map
再去range另一個nums,同時比較map
如果一樣就添加到新slice


### Hint


#### Topics : Array
#### Difficulty : Easy



