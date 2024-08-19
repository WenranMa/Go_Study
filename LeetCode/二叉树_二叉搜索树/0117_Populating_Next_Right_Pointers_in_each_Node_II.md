# 117. 填充每个节点的下一个右侧节点指针 II

### 中等

给定一个二叉树：
```go
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。

### 示例 1：
![next](/file/img/next_pointer_node_2.png)

    输入：root = [1,2,3,4,5,null,7]
    输出：[1,#,2,3,#,4,5,7,#]
    解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next 指针连接），'#' 表示每层的末尾。

### 示例 2：

    输入：root = []
    输出：[]
 

### 提示：
- 树中的节点数在范围 [0, 6000] 内
- -100 <= Node.val <= 100

### 进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。

### 解：


```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			p := root.Next
			for p != nil && p.Left == nil && p.Right == nil {
				p = p.Next
			}
			if p != nil {
				if p.Left != nil {
					root.Left.Next = p.Left
				} else if p.Right != nil {
					root.Left.Next = p.Right
				}
			}
		}
	}
	if root.Right != nil {
		p := root.Next
		for p != nil && p.Left == nil && p.Right == nil {
			p = p.Next
		}
		if p != nil {
			if p.Left != nil {
				root.Right.Next = p.Left
			} else if p.Right != nil {
				root.Right.Next = p.Right
			}
		}
	}
	connect(root.Right)
	connect(root.Left)
	return root
}
```