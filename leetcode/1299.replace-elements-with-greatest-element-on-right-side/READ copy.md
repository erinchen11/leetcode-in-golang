# Replace Elements with Greatest Element on Right Side

### Question
Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.

### Example1 :
Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]
Explanation: 
- index 0 --> the greatest element to the right of index 0 is index 1 (18).
- index 1 --> the greatest element to the right of index 1 is index 4 (6).
- index 2 --> the greatest element to the right of index 2 is index 4 (6).
- index 3 --> the greatest element to the right of index 3 is index 4 (6).
- index 4 --> the greatest element to the right of index 4 is index 5 (1).
- index 5 --> there are no elements to the right of index 5, so we put -1.

### Example 2 :
Input: arr = [400]
Output: [-1]
Explanation: There are no elements to the right of index 0.


### Constraints:
- 1 <= arr.length <= 10^4
- 1 <= arr[i] <= 10^5
### 思路

陣列中的每個元素都替換成 它右邊以後中最大的元素，逐一替換
最後一個元素換成 -1,因為沒有右邊的元素了

思考點：
從首range還是從尾range？
需要幾個變數來儲存

因為是要換成右邊的最大值，可以直接從最右邊往左看
從首range要逐一比較剩餘的array,too cost,
選擇從尾range,需要兩個變數 x, y 存值

初始 ：x = -1, y = 0
需呼叫另個函數來做x,y的比較


### Hint


#### Topics : Array
#### Difficulty : Easy



