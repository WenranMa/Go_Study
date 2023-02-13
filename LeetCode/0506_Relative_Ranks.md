# 506. Relative Ranks

### Easy

You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on. The placement of each athlete determines their rank:

The 1st place athlete's rank is "Gold Medal".
The 2nd place athlete's rank is "Silver Medal".
The 3rd place athlete's rank is "Bronze Medal".
For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").
Return an array answer of size n where answer[i] is the rank of the ith athlete.

### Example 1:

Input: score = [5,4,3,2,1]
Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].

### Example 2:

Input: score = [10,3,8,9,4]
Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].

Constraints:

n == score.length
1 <= n <= 10^4
0 <= score[i] <= 10^6
All the values in score are unique.

```go
// O(n) time, O(n) space.
// 但因为分数过高，效果并不好，index数组过大。
func findRelativeRanks(score []int) []string {
	res := []string{}
	ind := make([]int, 1000000)
	m := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for _, sc := range score {
		ind[sc] = 1
	}
	for i := 1000000 - 1; i >= 1; i-- {
		ind[i-1] += ind[i]
	}
	for _, s := range score {
		if ind[s] <= 3 {
			res = append(res, m[ind[s]-1])
		} else {
			t := strconv.Itoa(ind[s])
			res = append(res, t)
		}
	}
	return res
}

// 优化，不把所有的score都当做index存下来，用map节省空间.
// 也不遍历整个1000000，而是遍历最大最小分值之间的数据.
// 但如果最大最小分值差距较大，遍历还是过多。
func findRelativeRanks(score []int) []string {
	res := make([]string, len(score))
	m := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	ind := make(map[int]int)
	max := score[0]
	min := score[0]
	for v, sc := range score {
		ind[sc] = v
		if max < sc {
			max = sc
		}
		if min > sc {
			min = sc
		}
	}
	rank := 0
	for i := max; i >= min; i-- {
		if _, ok := ind[i]; ok {
			if rank < 3 {
				res[ind[i]] = m[rank]
			} else {
				t := strconv.Itoa(rank + 1)
				res[ind[i]] = t
			}
			rank += 1
		}
	}
	return res
}

// 如果分值差距较大，可以用排序。
func findRelativeRanks(score []int) []string {
	l := len(score)
	res := make([]string, l)
	m := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	ind := make(map[int]int)
	for v, sc := range score {
		ind[sc] = v
	}
	sort.Ints(score)
	for i := l - 1; i >= 0; i-- {
		if l-i <= 3 {
			res[ind[score[i]]] = m[l-i-1]
		} else {
			t := strconv.Itoa(l - i)
			res[ind[score[i]]] = t
		}
	}
	return res
}
```