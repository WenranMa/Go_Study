# 1221. Split a String in Balanced Strings

### Easy

Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

Given a balanced string s, split it in the maximum amount of balanced strings.

Return the maximum amount of split balanced strings.

### Example 1:

Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

### Example 2:

Input: s = "RLLLLRRRLR"
Output: 3
Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

### Example 3:

Input: s = "LLLLRRRR"
Output: 1
Explanation: s can be split into "LLLLRRRR".

Constraints:

1 <= s.length <= 1000
s[i] is either 'L' or 'R'.
s is a balanced string.

```go
func balancedStringSplit(s string) int {
	res := 0
	pair := 0
	for _, c := range s {
		if c == 'R' {
			pair += 1
		} else if c == 'L' {
			pair -= 1
		}
		if pair == 0 {
			res += 1
		}
	}
	return res

	//     res:= 0
	//     pair:= 1
	//     l:= len(s)
	//     slow,fast:= 0,1
	//     for slow< l && fast< l{
	//         if s[fast] == s[slow] {
	//             pair += 1
	//             fast += 1
	//         } else {
	//             pair -= 1
	//             fast += 1
	//             if pair == 0 {
	//                 pair = 1
	//                 res += 1
	//                 slow = fast
	//                 fast += 1
	//             }
	//         }
	//     }
	//     return res
}
```