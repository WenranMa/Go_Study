# 222. 完全二叉树的节点个数

### 简单

给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2^h 个节点。

### 示例 1：
![complete](/file/img/complete.jpg)

输入：root = [1,2,3,4,5,6]

输出：6

### 示例 2：

输入：root = []

输出：0

### 示例 3：

输入：root = [1]

输出：1
 
### 提示：

树中节点的数目范围是[0, 5 * 10^4]

0 <= Node.val <= 5 * 10^4

题目数据保证输入的树是 完全二叉树

### 进阶：
遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

### 解：
方法1：遍历 O(n) time.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```

方法2：

规定根节点位于第0层，完全二叉树的最大层数为h。根据完全二叉树的特性可知，完全二叉树的最左边的节点一定位于最底层，因此从根节点出发，每次访问左子节点，直到遇到叶子节点，该叶子节点即为完全二叉树的最左边的节点，经过的路径长度即为最大层数h。

当 `0 <= i < h`时，第i层包含`2^i`个节点，最底层包含的节点数最少为`1`，最多为`2^h`。

当最底层包含`1`个节点时，完全二叉树的节点个数是`2^h`
 
当最底层包含`2^h`个节点时，完全二叉树的节点个数是`2^(h+1)-1`

因此对于最大层数为h的完全二叉树，节点个数一定在 `[2^h,2^(h+1)−1]`的范围内，可以在该范围内通过二分查找的方式得到完全二叉树的节点个数。

具体做法是，根据节点个数范围的上下界得到当前需要判断的节点个数 k，如果第k个节点存在，则节点个数一定大于或等于k，如果第k个节点不存在，则节点个数一定小于k，由此可以将查找的范围缩小一半，直到得到节点个数。

如何判断第k个节点是否存在呢？如果第k个节点位于第h层，则k的二进制表示包含`h+1`位，其中最高位是`1`，其余各位从高到低表示从根节点到第k个节点的路径，`0`表示移动到`左子节点`，`1`表示移动到`右子节点`。通过位运算得到第k个节点对应的路径，判断该路径对应的节点是否存在，即可判断第k个节点是否存在。

![complete_tree](/file/img/complete_tree_with_bits.png)

```go
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    level := 0
    for node := root; node.Left != nil; node = node.Left {
        level++
    }
    return sort.Search(1<<(level+1), func(k int) bool {
        if k <= 1<<level {
            return false
        }
        bits := 1 << (level - 1)
        node := root
        for node != nil && bits > 0 {
            if bits&k == 0 {
                node = node.Left
            } else {
                node = node.Right
            }
            bits >>= 1
        }
        return node == nil
    }) - 1
}
```