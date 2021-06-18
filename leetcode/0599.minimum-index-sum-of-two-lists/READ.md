# Minimum index sum of two lists

### Question
Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restauranted by strings.

You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.


### Example1 :
Input: 
list1 = ["Shogun","Tapioca Express","Burger King","KFC"],
list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".

### Example 2 :
Input: 
list1 = ["Shogun","Tapioca Express","Burger King","KFC"], 
list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).

### Example 3 :
Input: 
list1 = ["Shogun","Tapioca Express","Burger King","KFC"], 
list2 = ["KFC","Burger King","Tapioca Express","Shogun"]
Output: ["KFC","Burger King","Tapioca Express","Shogun"]

### Example 4 :
Input: 
list1 = ["Shogun","Tapioca Express","Burger King","KFC"], 
list2 = ["KNN","KFC","Burger King","Tapioca Express","Shogun"]
Output: ["KFC","Burger King","Tapioca Express","Shogun"]

### Constraints:

- 1 <= list1.length, list2.length <= 1000
- 1 <= list1[i].length, list2[i].length <= 30
- list1[i] and list2[i] consist of spaces ' ' and English letters.
- All the stings of list1 are unique.
- All the stings of list2 are unique.

### 思路
需要一個hash table與 restaurant list
將其中一個list的字串存入 hash table，使得每個字串都對應到其索引值
然後再遍歷另一個list,每碰到一個字串就去hash table裡找有沒有一樣的字串
如果有，則將hash table的索引值加上現在這個字串的索引值得到索引值之和 S
如果S小於目前的索引值之和最小值,則更新restaurant list的內容
將現在的字串放進restaurant list並更新最小值
如果S與最小值一樣大,則將此字串加入到restaurant list之中

### Hint


#### Topics : Hashtable
#### Difficulty : Easy



