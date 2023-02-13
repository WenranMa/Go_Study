# 1502. Can Make Arithmetic Progression From Sequence

### Easy

A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.

### Example 1:

Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

### Example 2:

Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic progression.

Constraints:

2 <= arr.length <= 1000
-10^6 <= arr[i] <= 10^6

```go
// O(n) time, O(n) space. Map
func canMakeArithmeticProgression(arr []int) bool {
	min := 1000001
	max := -1000001
	l := len(arr)
	m := make(map[int]int)
	for _, n := range arr {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
		m[n] += 1
	}
	if max == min {
		return true
	}
	if (max-min)%(l-1) != 0 {
		return false
	}
	step := (max - min) / (l - 1)
	for i := min; i <= max; i = i + step {
		if v, ok := m[i]; !ok || v != 1 {
			return false
		}
	}
	return true
}

// 改进，O(n) time, in place.
func canMakeArithmeticProgression(arr []int) bool {
	min := 1000001
	max := -1000001
	l := len(arr)
	for _, n := range arr {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
	}
	if max == min {
		return true
	}
	if (max-min)%(l-1) != 0 {
		return false
	}
	d := (max - min) / (l - 1)
	i := 0
	for i < l {
		if arr[i] == min+i*d {
			i++
		} else if (arr[i]-min)%d != 0 {  // 2 4 7这种
			return false
		} else {
			pos := (arr[i] - min) / d  //应该在的位置
			if arr[pos] == arr[i] {
				return false
			}
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	return true
}
```