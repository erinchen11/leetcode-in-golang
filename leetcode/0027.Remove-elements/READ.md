#  Remove Element

### Question
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.



### Example1 :
nput: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2]
Explanation: Your function should return length = 2, with the first two elements of nums being 2.
It doesn't matter what you leave beyond the returned length. For example if you return 2 with nums = [2,2,3,3] or nums = [2,2,0,0], your answer will be accepted.

### Example 2 :
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3]
Explanation: Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4. Note that the order of those five elements can be arbitrary. It doesn't matter what values are set beyond the returned length.


### Constraints:
- 0 <= nums.length <= 100
- 0 <= nums[i] <= 50
### 思路
不能創建新的array
題目是說 給定一個array跟value
該value是要從array中刪除的
所求為 刪除後剩下的數有幾個

用two pointer, i去遍歷整個array, x去紀錄位置
看最後x記錄到的位置就是所求的剩下的數有幾個
i,x初始為0, 檢查nums[i]是否等於value
如果不是, 則把它放到nums[x],並且x++
最後回傳x

### Hint


#### Topics : Array
#### Difficulty : Easy



