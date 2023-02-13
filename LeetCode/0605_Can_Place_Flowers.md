# 605. Can Place Flowers

### Easy

You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.

### Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

### Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false

Constraints:

1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length

```go
/* 细节很多。
1.只有一个数的情况
2.只有两个数的情况
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	c := 0
	l := len(flowerbed)
	i := 0
	for i < l {
		if flowerbed[i] == 1 {
			i += 2
		} else {
			if i == 0 {
				if i+1 < l && flowerbed[i+1] == 0 {
					c += 1
				} else if i+1 == l {
					c += 1
					break
				}
				i += 2
			} else if i == l-1 {
				if i-1 >= 0 && flowerbed[i-1] == 0 {
					c += 1
				}
				break
			} else {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					c += 1
					i += 2
				} else {
					i += 1
				}
			}
		}
	}
	return c >= n
}
```