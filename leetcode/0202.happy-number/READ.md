# Happy Number

### Question
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process :

* Starting with any positive integer, replace the number by the sum of the squares of its figits.
* Repeat the process until the number equal 1 (wherer it will stay),
or it loops endlessly in a cycle which does not include 1.
* Those numbers for which this process ends in 1 are happy.



### Example1 :
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

### Example 2 :
Input: n = 2
Output: false

### Constraints:
1 <= n <= 2^31 - 1
### 思路

建立一個map,用來存放每次計算出的新數字
並且判斷該值是否為1, 是1返回true
如果不是1則判斷該值在map中是否存在
如果存在則返回false
如果不存在則將新數字放入map

### Hint


#### Topics : Array
#### Difficulty : Easy



