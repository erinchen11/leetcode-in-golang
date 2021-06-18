# Duplicate Zeros 

### Question
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

### Example1 :
Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

### Example 2 :
Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]



### Constraints:
- 1 <= arr.length <= 10000
- 0 <= arr[i] <= 9
### 思路
以右移Array的方向去想，要想一下有沒有地方可以複寫，有需要右移的東西嗎
遍歷整個 Array,
如果是0且後面有地方複寫，

index+2

### Hint


#### Topics : Array
#### Difficulty : Easy



