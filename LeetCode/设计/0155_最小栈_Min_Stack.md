# 155. 最小栈

### 中等

### 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。

void push(int val) 将元素val推入堆栈。

void pop() 删除堆栈顶部的元素。

int top() 获取堆栈顶部的元素。

int getMin() 获取堆栈中的最小元素。

### 示例 1:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

### 提示：

-2^31 <= val <= 2^31 - 1
pop、top 和 getMin 操作总是在 非空栈 上调用
push, pop, top, and getMin最多被调用 3 * 10^4 次

### 解：

```go
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minStack) == 0 || val <= this.minStack[0] {
		this.minStack = append([]int{val}, this.minStack...)
	}
	this.stack = append([]int{val}, this.stack...)
}

func (this *MinStack) Pop() {
	if this.minStack[0] == this.stack[0] {
		this.minStack = this.minStack[1:len(this.minStack)]
	}
	this.stack = this.stack[1:len(this.stack)]
}

func (this *MinStack) Top() int {
	return this.stack[0]
}

func (this *MinStack) GetMin() int {
	return this.minStack[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```

用末尾的元素，不用头部，这样少些内存分配的过程。
```go
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	l := len(this.minStack)
	if l == 0 || val <= this.minStack[l-1] {
		this.minStack = append(this.minStack, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	lm, l := len(this.minStack), len(this.stack)
	if this.minStack[lm-1] == this.stack[l-1] {
		this.minStack = this.minStack[0 : lm-1]
	}
	this.stack = this.stack[0 : l-1]
}

func (this *MinStack) Top() int {
	l := len(this.stack)
	return this.stack[l-1]
}

func (this *MinStack) GetMin() int {
	l := len(this.minStack)
	return this.minStack[l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
```