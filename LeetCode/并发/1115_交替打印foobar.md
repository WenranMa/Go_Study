# 1115. 交替打印 FooBar

### 中等

给你一个类：
```java
class FooBar {
    public void foo() {
        for (int i = 0; i < n; i++) {
            print("foo");
        }
    }

    public void bar() {
        for (int i = 0; i < n; i++) {
            print("bar");
        }
    }
}
```
两个不同的线程将会共用一个 FooBar 实例：

线程 A 将会调用 foo() 方法，而
线程 B 将会调用 bar() 方法
请设计修改程序，以确保 "foobar" 被输出 n 次。

### 示例 1：

    输入：n = 1
    输出："foobar"
    解释：这里有两个线程被异步启动。其中一个调用 foo() 方法, 另一个调用 bar() 方法，"foobar" 将被输出一次。

### 示例 2：

    输入：n = 2
    输出："foobarfoobar"
    解释："foobar" 将被输出两次。


### 解：

两个channel分别当两个开关

```go
type FooBar struct {
	n int
	f chan int
	b chan int
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n: n,
		f: make(chan int, 1),
		b: make(chan int, 1),
	}
	fb.f <- 1
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.f
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.b <- 1
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.b
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.f <- 1
	}
}
```