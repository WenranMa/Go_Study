# 925. 长按键入 Long Pressed Name

### Easy

Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

### Example 1:

Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

### Example 2:

Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it was not in the typed output.

Constraints:

1 <= name.length, typed.length <= 1000
name and typed consist of only lowercase English letters.

### 解：

双指针。

```go
/*
 Two pointers.
 test case: 
 1. alex, aaleexxxx, true.
 2. alex, aaleexbbb, false.
 3. alex, balex, false.
*/
func isLongPressedName(name string, typed string) bool {
	n := 0
	ln := len(name)
	lt := len(typed)
	for t := 0; t < lt; t += 1 {
		if n < ln && name[n] == typed[t] {
			n += 1
		} else if t == 0 || typed[t] != typed[t-1] {
			return false
		}
	}
	return n == ln
}
```