# 为什么 Go 标准库中有些函数只有签名，没有函数体？

如果你看过 Go 语言标准库，应该有见到过，有一些函数只有签名，没有函数体。你有没有感觉到很奇怪？这到底是怎么回事？我们自己可以这么做吗？

首先，函数肯定得有实现，没有函数体，一定是在其他某个地方。Go 中一般有两种形式。

1. 函数签名使用Go,然后通过该包中的汇编文件来实现它。

比如，在标准库 sync/atomic 包中的函数基本只有函数签名。比如：`atomic.StoreInt32`
```go
// StoreInt32 atomically stores val into *addr.
func StoreInt32(addr *int32, val int32)
```
它的函数实现在哪呢？其实只要稍微留意一下发现该目录下有一个文件：asm.s，它提供了具体的实现，即通过汇编来实现：
```
TEXT ·StoreInt32(SB),NOSPLIT,$0
    JMP    runtime∕internal∕atomic·Store(SB)
```
具体的实现，在 runtime∕internal 文件夹中，有兴趣你可以打开 asm_amd64.s 看看。

很明显，这种方式一方面会是效率的考虑，另一方面，有一些代码只能汇编实现。

2. 通过`//go:linkname`指令来实现
比如，在标准库 time 包中的 Sleep 函数：
```go
// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d Duration)
```
它的实现在哪里呢？在 time 包中并没有找到相应的汇编文件。

按照 Go 源码的风格，这时候一般需要去 runtime 包中找。我们会找到 time.go，其中有一个函数：
```go
// timeSleep puts the current goroutine to sleep for at least ns nanoseconds.
//go:linkname timeSleep time.Sleep
func timeSleep(ns int64) {
    ...
}
```
这就是我们要找的 time.Sleep 的实现。

对于 `//go:linkname`,它的格式是：`//go:linkname 函数名 包名.函数名`

因此我们在遇到函数没有实现，但汇编又不存在时，可以通过尝试搜索：go:linkname xxx xx.xxx 的形式来找。

这里面要提示一点，使用 //go:linkname，必须导入 unsafe 包，所以，有时候会见到：import _ "unsafe" 这样的代码。
