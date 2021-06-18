# Merge Sorted Array

### Question
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
 

### Example1 :


### Example 2 :

### Example 3 :




### Constraints:

### 思路

為了不大量移動元素，就要從兩個array長度之和的最後一個位置開始
依序取出兩個array中的大數
再從第一個array的尾巴開始從頭放
只要循環一次後就合併array了

### Hint


#### Topics : Array
#### Difficulty : Easy



