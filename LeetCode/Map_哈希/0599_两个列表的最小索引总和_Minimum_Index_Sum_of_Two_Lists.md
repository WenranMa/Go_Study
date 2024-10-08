# 599. 两个列表的最小索引总和 Minimum Index Sum of Two Lists

### Easy

Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.

### Example 1:

Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".

### Example 2:

Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).

Constraints:

1 <= list1.length, list2.length <= 1000
1 <= list1[i].length, list2[i].length <= 30
list1[i] and list2[i] consist of spaces ' ' and English letters.
All the stings of list1 are unique.
All the stings of list2 are unique.

### 解：

```go
// O(n + m) time, O(n) space.
// Map
func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	s := len(list1) + len(list2)
	res := []string{}
	for i, r := range list1 {
		m[r] = i
	}
	for i, r := range list2 {
		if v, ok := m[r]; ok {
			temp := i + v
			if s > temp {
				s = temp
				res = nil   // 这步很聪明，遇到小的，结果置零
				res = append(res, r)
			} else if s == temp {
				res = append(res, r)
			}
		}
	}
	return res
}
```