# Go Timer

在 Go 语言中，time 包提供了 Timer 类型，用于创建一次性定时器。

## 举例
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，5秒后发送当前时间到通道
	timer := time.After(5 * time.Second)
	// 等待定时器触发，然后从通道接收时间并打印
	<-timer
	fmt.Println("定时器触发！")
}
```

**使用带回调的定时器**

如果你想在定时器触发时执行一个特定的函数，可以使用以下方式：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，5秒后触发
	t := time.NewTimer(5 * time.Second)
	// 在另一个 goroutine 中等待定时器触发并执行回调
	go func() {
		<-t.C
		fmt.Println("定时器触发！")
	}()
	// 阻塞主 goroutine，直到定时器触发
	time.Sleep(6 * time.Second)
}
```

**取消定时器**

如果需要在定时器触发前取消它，可以使用 Stop 方法：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(5 * time.Second)
	go func() {
		select {
		case <-t.C:
			fmt.Println("定时器触发！")
		}
	}()
	// 如果你想在定时器触发前取消它
	t.Stop()
	fmt.Println("定时器已取消或触发。")
}
// t.Stop()返回true, 则表示定时器已经触发，否则表示定时器还未触发  
```

## 源码 --- TBD 

```go
type timer struct {
    tb *timersBucket
    i  int

    when   int64
    period int64
    f      func(interface{}, uintptr)
    arg    interface{}
    seq    uintptr
}
```

- `timer` 就是 Golang 定时器内部表示,每一个`timer` 其实都存在堆中

- `tb` 就是用于存储当前定时器的桶

- `i` 是当前定时器在堆中的索引,可以通过这两个变量找到当前定时器在堆中的位置

- `when` 表示当前定时器(Timer) 被唤醒的时间

- `period` 表示两次被唤醒的间隔,每当定时器被唤醒时都会调用`f(args,now)` 函数并传入`args` 和当前时间作为参数

- 这里的`timer`作为一个私有结构体其实只是定时器的运行时表示,`time` 包对外暴露的定时器是如下结构

  ```go
  type Timer struct {
    C <-chan Time
    r runtimeTimer
  }
  Copy
  ```

- `Timer` 定时器必须通过`NewTimer` 或者 `AfterFunc` 函数进行创建,其中的`runtimeTimer` 其实就是上面的`timer`结构体, 当定时器失效时,失效的时间就会被发送给当前定时器持有的Channel`C`, 订阅管道中消息的Goroutine就会接收到当前定时器失效的时间

### 创建

- ```go
  time
  ```

   

  包对外提供了两种创建定时器的方法

  - `NewTimer` 接口创建用于通知触发时间的Channel,调用 `startTimer` 方法并返回一个创建指向`Timer`结构体的指针

  - `AfterFunc` 也提供了相似的结构,与上面不同的是,它只会在定时器到期时调用传入的方法

  - ```go
    startTimer
    ```

    是创建定时器的入口,所有的定时器的创建和重启基本上都需要这个函数

    - 调用`addTimer`函数,首先通过`assignBucket`方法为当前定时器选择一个`timersBucket` 桶,根据当前的Goroutine所在处理器P的id选择一个合适的桶,随后调用`addTimerLocked`方法将当前定时器加入桶中
    - `addtimerLocked` 会先将最新加入的定时器加到队列的末尾,随后调用`siftipTimer`将当前定时器与四叉树(或者四叉堆)中的父节点进行比较,保证父节点的到期时间一定小于子节点

### 触发

- 定时器的触发都是由`timerproc`中的一个双层`for`循环控制的,外层的`for`循环主要负责对当前的`Goroutine` 进行控制,它不仅会负责锁的获取和释放,还会在合适的时机触发当前Goroutine的休眠
- 内部循环
  - 如果桶中不包含任何定时器就会直接返回并陷入休眠等待定时器加入当前桶
  - 如果四叉树最上面的定时器还没有到期会通过`notetsleepg`方法陷入休眠等待最近定时器的到期
  - 如果四叉树最上面的定时器已经到期
    - 当定时器 `preiod > 0` 就会设置下一次会触发定时器的时间并将当前定时器向下移动到对应位置
    - 当定时器`preios <= 0` 就会将当前定时器从四叉树中移除
  - 在每次循环的最后都会从定时器中取出定时器中的函数,参数 和序列号并调用函数触发该计数器
- 使用`NewTimer`创建的定时器,传入的函数时`sendTime`,它会将当前时间发送到定时器持有的Channel中,而使用`AfterFunc` 创建的定时器,在内层循环中调用的函数就会是调用方法传入的函数了

### 休眠

- `timeSleep` 会创建一个新的`timer`结构体,在初始化的过程中我们会传入当前Goroutine 应该被唤醒的时间以及唤醒时需要调用的函数`goroutineReady`,随后会调用 `goparklock` 将当前GOroutine陷入休眠状态,当定时器到期时也会调用 `goroutineReady` 方法唤醒当前的Goroutine
- `time.Sleep` 方法其实只是创建了一个会在到期时唤醒当前Goroutine的定时器并通过`goparkunlock`将当前的协程陷入休眠状态等待定时器触发的唤醒

## Ticker

- 除了只用于一次的定时器(Timer)之外,Go语言的`time`包中还提供了用于多次通知的`Ticker`计时器,计时器中包含了一个用于接受通知的Channel 和一个定时器,这个两个字段组成了用于连续多次触发事件的计时器
- 想要在Go中创建一个计时器只有两种方法,一种是使用`NewTicker`方法显示的创建`Ticker`计时器指针,另一种可以直接通过`Tick`方法获取一个会定期发送消息的Channel
- 每一个`NewTicker`方法开启的计时器都需要在不需要使用时调用`Stop`进行关闭,如果不显示调用`Stop`方法,创建的计时器就没有办法被垃圾回收,而通过`Tick`创建的计时器由于只对外提供了Channel,所以是一定没有办法关闭的,我们一定要谨慎使用这一接口创建计时器

## 性能分析

- 定时器在内部使用四叉树的方式进行实现和存储,高并发的场景下会有比较明显的性能问题

