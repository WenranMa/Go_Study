# 130. 被围绕的区域

### 中等

给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：

连接：一个单元格与水平或垂直方向上相邻的单元格连接。
区域：连接所有 'O' 的单元格来形成一个区域。
围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
通过将输入矩阵 board 中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。

### 示例 1：

输入：board = [
        ["X","X","X","X"],
        ["X","O","O","X"],
        ["X","X","O","X"],
        ["X","O","X","X"]
    ]

输出：[
        ["X","X","X","X"],
        ["X","X","X","X"],
        ["X","X","X","X"],
        ["X","O","X","X"]
    ]

### 示例 2：

输入：board = [["X"]]

输出：[["X"]]

### 提示：

    m == board.length
    n == board[i].length
    1 <= m, n <= 200
    board[i][j] 为 'X' 或 'O'

### 解:

DFS, 找到边缘是 'O'的区域，这些都是不能翻转的，可以标记为一个别的字符，比如'N', 然后遍历矩阵，剩下的'O' 都是被包围的，翻转即可。再把标记为 'N' 的区域还原为 'O' 。

```go
func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	row, col := len(board), len(board[0])
	for i := 1; i < row-1; i++ {
		if board[i][0] == 'O' {
			mark(board, i, 1)
		}
		if board[i][col-1] == 'O' {
			mark(board, i, col-2)
		}
	}
	for j := 1; j < col-1; j++ {
		if board[0][j] == 'O' {
			mark(board, 1, j)
		}
		if board[row-1][j] == 'O' {
			mark(board, row-2, j)
		}
	}
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'N' {
				board[i][j] = 'O'
			}
		}
	}
}

func mark(board [][]byte, x, y int) {
	if board[x][y] != 'O' {
		return
	}
	board[x][y] = 'N'
	if x > 1 {
		mark(board, x-1, y)
	}
	if x < len(board)-2 {
		mark(board, x+1, y)
	}
	if y > 1 {
		mark(board, x, y-1)
	}
	if y < len(board[0])-2 {
		mark(board, x, y+1)
	}
}
```