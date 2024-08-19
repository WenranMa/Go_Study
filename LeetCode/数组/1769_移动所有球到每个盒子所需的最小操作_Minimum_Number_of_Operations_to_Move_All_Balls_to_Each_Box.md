# 1769. 移动所有球到每个盒子所需的最小操作 Minimum Number of Operations to Move All Balls to Each Box

### Medium

You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.

In one operation, you can move one ball from a box to an adjacent box. Box i is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may be more than one ball in some boxes.

Return an array answer of size n, where answer[i] is the minimum number of operations needed to move all the balls to the ith box.

Each answer[i] is calculated considering the initial state of the boxes.

### Example 1:

Input: boxes = "110"
Output: [1,1,3]
Explanation: The answer for each box is as follows:
1) First box: you will have to move one ball from the second box to the first box in one operation.
2) Second box: you will have to move one ball from the first box to the second box in one operation.
3) Third box: you will have to move one ball from the first box to the third box in two operations, and move one ball from the second box to the third box in one operation.

### Example 2:

Input: boxes = "001011"
Output: [11,8,5,4,3,4]

Constraints:

n == boxes.length
1 <= n <= 2000
boxes[i] is either '0' or '1'.

### 解：

```go
// O(n) time, two pass.
// L to R, ans[i] 是把从0到i-1的球移动到i所需的步数。
// R to L, ans[i] 是吧从l-1到i+1的球移动到i所需步数。
// 两者相加即可。

/*
比如  001011
第一次遍历，ans = [0,0,0,1,2,4]
第二次遍历，ans = [11,8,5,4,3,4]
*/

func minOperations(boxes string) []int {
	l := len(boxes)
	ans := make([]int, l)
	cnt := 0
	opt := 0
	for i := 0; i < l; i++ {
		ans[i] += opt
		if boxes[i] == 49 {
			cnt += 1
		}
		opt += cnt
	}
	cnt = 0
	opt = 0
	for i := l - 1; i >= 0; i-- {
		ans[i] += opt
		if boxes[i] == 49 {
			cnt += 1
		}
		opt += cnt
	}
	return ans
}
```

```go
// O(n^2) time.
func minOperations(boxes string) []int {
	ans := []int{}
	for i, _ := range boxes {
		s := 0
		for j, b := range boxes {
			if b == 49 && i != j { // 49 is '1'
				if i < j {
					s += j - i
				} else {
					s += i - j
				}
			}
		}
		ans = append(ans, s)
	}
	return ans
}
```
