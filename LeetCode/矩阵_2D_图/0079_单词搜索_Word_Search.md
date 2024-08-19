# 79. 单词搜索

### 中等

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

### 示例 1：
![w1](/file/img/word1.jpg)

    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    输出：true

### 示例 2：
![w1](/file/img/word2.jpg)

    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
    输出：true

### 示例 3：
![w1](/file/img/word3.jpg)

    输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
    输出：false
 
### 提示：
- m == board.length
- n = board[i].length
- 1 <= m, n <= 6
- 1 <= word.length <= 15
- board 和 word 仅由大小写英文字母组成

### 进阶：
你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

### 解：
DFS.

```go
func exist(board [][]byte, word string) bool {
	visit := make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if dfs(board, visit, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visit [][]bool, x, y int, word string, index int) bool {
	if visit[x][y] || word[index] != board[x][y] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	visit[x][y] = true
	if x > 0 && dfs(board, visit, x-1, y, word, index+1) {
		return true
	}
	if x < len(board)-1 && dfs(board, visit, x+1, y, word, index+1) {
		return true
	}
	if y > 0 && dfs(board, visit, x, y-1, word, index+1) {
		return true
	}
	if y < len(board[0])-1 && dfs(board, visit, x, y+1, word, index+1) {
		return true
	}
	visit[x][y] = false
	return false
}

```