# 1812. 判断国际象棋格子颜色 Determine Color of a Chessboard Square

### Easy

You are given coordinates, a string that represents the coordinates of a square of the chessboard. 

![board](/file/img/leetcode_1812.png)

Return true if the square is white, and false if the square is black.

The coordinate will always represent a valid chessboard square. The coordinate will always have the letter first, and the number second.

### Example 1:

Input: coordinates = "a1"
Output: false
Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.

### Example 2:

Input: coordinates = "h3"
Output: true
Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.
Example 3:

Input: coordinates = "c7"
Output: false

Constraints:

coordinates.length == 2
'a' <= coordinates[0] <= 'h'
'1' <= coordinates[1] <= '8'

### 解

```go
// '1' == 49
// 'a' == 97
// 两位加起来是偶数则为黑，奇数为白色
func squareIsWhite(coordinates string) bool {
	return (coordinates[0]+coordinates[1])%2 == 1
}

// 如果坐标都是奇数，或者都是偶数，则是黑色。
// 一奇一偶则是白色。
// 用亦或，看最后一位是不是相同。
func squareIsWhite(coordinates string) bool {
	x := coordinates[0] - 'a'
	y := coordinates[1] - '1'
	return (x^y)&1 == 1
}

```