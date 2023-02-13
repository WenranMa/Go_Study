# 914. X of a Kind in a Deck of Cards

### Easy

In a deck of cards, each card has an integer written on it.

Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:

Each group has exactly X cards.
All the cards in each group have the same integer.
 
### Example 1:

Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].

### Example 2:

Input: deck = [1,1,1,2,2,2,3,3]
Output: false
Explanation: No possible partition.

Constraints:

1 <= deck.length <= 10^4
0 <= deck[i] < 10^4

```go
// Map + 最大公约数，O(n) time, O(n) space.
// 注意数组的最大公约数，用两个数的GCD去和下一个数继续计算。
func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	m := make(map[int]int)
	for _, c := range deck {
		m[c] += 1
	}
	d := m[deck[0]]
	for _, v := range m {
		d = gcd(d, v)
	}
	return d >= 2
}
func gcd(x, y int) int {
	temp := 0
	for {
		temp = x % y
		if temp > 0 {
			x = y
			y = temp
		} else {
			return y
		}
	}
}
``