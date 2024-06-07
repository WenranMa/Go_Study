# goroutine, channel

## 19. 关于channel，下面语法正确的是（）

- A. var ch chan int
- B. ch := make(chan int)
- C. <-ch
- D. ch<-

**答：A、B、C**

**解析：**

A、B都是申明channel；C读取channel；写channel是必须带上值，所以D错误。



## 86. 下面代码会触发异常吗？请说明。

```go
func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {
    case value := <-int_chan:
        fmt.Println(value)
    case value := <-string_chan:
        panic(value)
    }
}
```

**解析：**

`select` 会随机选择一个可用通道做收发操作，所以可能触发异常，也可能不会。

## 87. 关于channel的特性，下面说法正确的是？--- TBD

- A. 给一个 nil channel 发送数据，造成永远阻塞
- B. 从一个 nil channel 接收数据，造成永远阻塞
- C. 给一个已经关闭的 channel 发送数据，引起 panic
- D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

**答：A B C D**

## 96. 关于select机制，下面说法正确的是?

- A. select机制用来处理异步IO问题；
- B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
- C. golang在语言级别支持select关键字；
- D. select关键字的用法与switch语句非常类似，后面要带判断条件；

**答：A B C**


## 97. 下面的代码有什么问题？--- TBD

```go
func Stop(stop <-chan bool) {
    close(stop)
}
```

**答：有方向的 channel 不可以被关闭。**

## 121. 下面代码输出什么？

```go
func main() {
    var ch chan int
    select {
    case v, ok := <-ch:
        println(v, ok)
    default:
        println("default") 
    }
}
```

**答：default**

**解析：**

ch 为 nil，读写都会阻塞。


## 129. 关于 channel 下面描述正确的是？

- A. 向已关闭的通道发送数据会引发 panic；
- B. 从已关闭的缓冲通道接收数据，返回已缓冲数据或者零值；
- C. 无论接收还是接收，nil 通道都会阻塞；

**答：A B C**


## 138. 下面的代码输出什么？

```go
var o = fmt.Print

func main() {
    c := make(chan int, 1)
    for range [3]struct{}{} {
        select {
        default:
            o(1)
        case <-c:
            o(2)
            c = nil
        case c <- 1:
            o(3)
        }
    }
}
```

**答：321**

**解析：**

第一次循环，写操作已经准备好，执行 o(3)，输出 3；

第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；

第三次，由于 c 为 nil，走的是 default 分支，输出 1。


## 144. 下面的代码有什么问题？

```go
type data struct {
    sync.Mutex
}

func (d data) test(s string)  {
    d.Lock()
    defer d.Unlock()

    for i:=0;i<5 ;i++  {
        fmt.Println(s,i)
        time.Sleep(time.Second)
    }
}


func main() {

    var wg sync.WaitGroup
    wg.Add(2)
    var d data

    go func() {
        defer wg.Done()
        d.test("read")
    }()

    go func() {
        defer wg.Done()
        d.test("write")
    }()

    wg.Wait()
}
```

**答：锁失效**

**解析：**

将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。

修复代码：

```go
func (d *data) test(s string)  {     // 指针接收者
    d.Lock()
    defer d.Unlock()

    for i:=0;i<5 ;i++  {
        fmt.Println(s,i)
        time.Sleep(time.Second)
    }
}
```

或者可以通过嵌入 `*Mutex` 来避免复制的问题，但需要初始化。

```go
type data struct {
    *sync.Mutex     // *Mutex
}

func (d data) test(s string) {    // 值方法
    d.Lock()
    defer d.Unlock()

    for i := 0; i < 5; i++ {
        fmt.Println(s, i)
        time.Sleep(time.Second)
    }
}

func main() {

    var wg sync.WaitGroup
    wg.Add(2)

    d := data{new(sync.Mutex)}   // 初始化

    go func() {
        defer wg.Done()
        d.test("read")
    }()

    go func() {
        defer wg.Done()
        d.test("write")
    }()

    wg.Wait()
}
```

## 184. 下面的代码有什么问题？

```go
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        go func(wg sync.WaitGroup, i int) {
            wg.Add(1)
            fmt.Printf("i:%d\n", i)
            wg.Done()
        }(wg, i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```

**解析：**

知识点：WaitGroup 的使用。存在两个问题：

- 在协程中使用 wg.Add()；
- 使用了 sync.WaitGroup 副本；

修复代码：

```go
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            fmt.Printf("i:%d\n", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```

或者：

```go
func main() {

    wg := &sync.WaitGroup{}

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg *sync.WaitGroup,i int) {
            fmt.Printf("i:%d\n", i)
            wg.Done()
        }(wg,i)
    }

    wg.Wait()

    fmt.Println("exit")
}
```


## 171. 下面代码有什么问题，请说明？

```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i:=0;i<10 ;i++  {
            fmt.Println(i)
        }
    }()

    for {}
}
```

**答：** 

以上代码在go1.14版本之前(不含1.14版本): for {} 独占 CPU 资源导致其他 Goroutine 饿死，

在go1.14版本之后(包含go1.14): 会打印0123456789, 并且主程会进入死循环。

这是因为1.14版本(含1.14版本)之后goroutine抢占式调度设计改为基于信号的抢占式调度，当调度器监控发现某个goroutine执行时间过长且有别的goroutine在等待时，会把执行时间过长的goroutine暂停，转而调度等待的goroutine. 所以for循环的goroutine得以执行.

**解析：**

可以通过阻塞的方式避免 CPU 占用，修复代码：

```go
func main() {
    runtime.GOMAXPROCS(1)

    go func() {
        for i:=0;i<10 ;i++  {
            fmt.Println(i)
        }
        os.Exit(0)
    }()

    select {}
}
```

### 两个goroutine 交替打印 1-20
```go
package main

import (
	"fmt"
	"time"
)
func main() {
	ch0 := make(chan int)
	ch1 := make(chan int)
	go func() {
		for i := 1; i <= 20; i += 2 {
			<-ch0
			fmt.Println(i)
			ch1 <- 0
		}
	}()
	go func() {
		for i := 2; i <= 20; i += 2 {
			<-ch1
			fmt.Println(i)
			ch0 <- 0
		}
	}()
	ch0 <- 0	
	time.Sleep(10*time.Second)
}



// 可以用一个管道
func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 20; i += 1 {
			ch <- 0
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 20; i += 1 {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
```

## 176. 下面的代码有什么问题？

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        fmt.Println("1")
        wg.Done()
        wg.Add(1)
    }()
    wg.Wait()
}
```

**解析：**

`panic: sync: WaitGroup is reused before previous Wait has returned`

协程里面，使用 wg.Add(1) 但是没有 wg.Done()，导致 panic()。





### sync.Once如果中途panic，后续能否再执行一次？
不会


## 133. 关于 channel 下面描述正确的是？

- A. close() 可以用于只接收通道；
- B. 单向通道可以转换为双向通道；
- C. 不能在单向通道上做逆向操作（例如：只发送通道用于接收）；

**答：C**


## 95. 下面代码输出什么？

```go
func main() {
    ch := make(chan int, 100)
    // A
    go func() {              
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
    // B
    go func() {
        for {
            a, ok := <-ch
            if !ok {
                fmt.Println("close")
                return
            }
            fmt.Println("a: ", a)
        }
    }()
    close(ch)
    fmt.Println("ok")
    time.Sleep(time.Second * 10)
}
```

**答：程序抛异常**

**解析：**

先定义下，第一个协程为 A 协程，第二个协程为 B 协程；当 A 协程还没起时，主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic，`panic: send on closed channel`。

## 93. 关于无缓冲和有冲突的channel，下面说法正确的是？--- TBD

- A. 无缓冲的channel是默认的缓冲为1的channel；
- B. 无缓冲的channel和有缓冲的channel都是同步的；
- C. 无缓冲的channel和有缓冲的channel都是非同步的；
- D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；

**答：D**


## 75. 关于协程，下面说法正确是（）

- A. 协程和线程都可以实现程序的并发执行；
- B. 线程比协程更轻量级；
- C. 协程不存在死锁问题；
- D. 通过 channel 来进行协程间的通信；

**答：A D**


### map可以被多个goroutine操作吗？
可以，线程不安全，怎么办，枷锁，其他办法，sync.Map 线程安全。




## Context

context可以用来在goroutine之间传递上下文信息，相同的context可以传递给运行在不同goroutine中的函数，上下文对于多个goroutine同时使用是安全的，context包定义了上下文类型，可以使用background、TODO创建一个上下文，在函数调用链之间传播context，也可以使用WithDeadline、WithTimeout、WithCancel 或 WithValue 创建的修改副本替换它，听起来有点绕，其实总结起就是一句话：context的作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。

### context的使用

context包主要提供了两种方式创建context:

- context.Backgroud()
- context.TODO()

这两个函数其实只是互为别名，没有差别，官方给的定义是：

context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来。
context.TODO 应该只在不确定应该使用哪种上下文时使用；
所以在大多数情况下，我们都使用context.Background作为起始的上下文向下传递。

上面的两种方式是创建根context，不具备任何功能，具体实践还是要依靠context包提供的With系列函数来进行派生：
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

#### 超时控制

通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，所以一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行。withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已，他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。

withTimeout、WithDeadline不同在于WithTimeout将持续时间作为参数输入而不是时间对象，这两个方法使用哪个都是一样的，看业务场景和个人习惯了，因为本质withTimout内部也是调用的WithDeadline。

```go
//达到超时时间终止接下来的执行
func main()  {
    HttpHandler()
}

func NewContextWithTimeout() (context.Context,context.CancelFunc) {
    return context.WithTimeout(context.Background(), 3 * time.Second)
}

func HttpHandler()  {
    ctx, cancel := NewContextWithTimeout()
    defer cancel()
    deal(ctx)
}

func deal(ctx context.Context)  {
    for i:=0; i< 10; i++ {
        time.Sleep(1*time.Second)
        select {
        case <- ctx.Done():
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Printf("deal time is %d\n", i)
        }
    }
}
// 输出结果：
// 
// deal time is 0
// deal time is 1
// deal time is 2 ??
// context deadline exceeded
```

```go
// 没有达到超时时间终止接下来的执行
func main()  {
    HttpHandler1()
}

func NewContextWithTimeout1() (context.Context,context.CancelFunc) {
    return context.WithTimeout(context.Background(), 3 * time.Second)
}

func HttpHandler1()  {
    ctx, cancel := NewContextWithTimeout1()
    defer cancel()
    deal1(ctx, cancel)
}

func deal1(ctx context.Context, cancel context.CancelFunc)  {
    for i:=0; i< 10; i++ {
        time.Sleep(1*time.Second)
        select {
        case <- ctx.Done():
            fmt.Println(ctx.Err())
            return
        default:
            fmt.Printf("deal time is %d\n", i)
            cancel()
        }
    }
}
// 输出结果：
// 
// deal time is 0
// context canceled
```
使用起来还是比较容易的，既可以超时自动取消，又可以手动控制取消。

#### withCancel
日常业务开发中我们往往为了完成一个复杂的需求会开多个gouroutine去做一些事情，这就导致我们会在一次请求中开了多个goroutine确无法控制他们，这时我们就可以使用withCancel来衍生一个context传递到不同的goroutine中，当我想让这些goroutine停止运行，就可以调用cancel来进行取消。

来看一个例子：
```go
func main()  {
    ctx,cancel := context.WithCancel(context.Background())
    go Speak(ctx)
    time.Sleep(10*time.Second)
    cancel()
    time.Sleep(1*time.Second)
}

func Speak(ctx context.Context)  {
    for range time.Tick(time.Second){
        select {
        case <- ctx.Done():
            fmt.Println("我要闭嘴了")
            return
        default:
            fmt.Println("balabalabalabala")
        }
    }
}
// 运行结果：
// 
// balabalabalabala
// ....省略
// balabalabalabala
// 我要闭嘴了
```
我们使用withCancel创建一个基于Background的ctx，然后启动一个讲话程序，每隔1s说一话，main函数在10s后执行cancel，那么speak检测到取消信号就会退出。

context.Context 是 Go 语言在 1.7 版本中引入标准库的接口1，该接口定义了四个需要实现的方法，其中包括：

- Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
- Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
- Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
	- 如果 context.Context 被取消，会返回 Canceled 错误；
	- 如果 context.Context 超时，会返回 DeadlineExceeded 错误；
- Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

context reference:
https://segmentfault.com/a/1190000040917752



## sync.Once

Once 可以用来执行且仅仅执行一次动作，常常用于单例对象的初始化场景。Once 常常用来初始化单例资源，或者并发访问只需初始化一次的共享资源，或者在测试的时候初始化一次测试资源。

sync.Once 只暴露了一个方法 Do，你可以多次调用 Do 方法，但是只有第一次调用 Do 方法时 f 参数才会执行，这里的 f 是一个无参数无返回值的函数。

demo:
```go
import (
	"fmt"
	"sync"
)
func main() {
	var o sync.Once
	func1:= func() {	
		fmt.Println("only once")	
	}
	done:= make(chan bool)
	for i:= 0; i< 10; i++ {
		go func() {
			o.Do(func1)
			done <- true
		}()
	}
	for i:= 0; i< 10; i++ {
		<- done
	}
}
```

只输出一次 “only once”.

### 源码分析
接下来分析 sync.Do 究竟是如何实现的，它存储在包sync下 once.go 文件中，源代码如下:

```go
// sync/once.go

type Once struct {
	done uint32 // 初始值为0表示还未执行过，1表示已经执行过
	m    Mutex
}

func (o *Once) Do(f func()) {
	// 判断done是否为0，若为0，表示未执行过，调用doSlow()方法初始化
	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

// 加载资源
func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	// 采用双重检测机制 加锁判断done是否为零
	if o.done == 0 {
		// 执行完f()函数后，将done值设置为1
		defer atomic.StoreUint32(&o.done, 1)
		// 执行传入的f()函数
		f()
	}
}
```

为了防止多个goroutine调用 doSlow() 初始化资源时，造成资源多次初始化，因此采用 Mutex 锁机制来保证有且仅初始化一次
Do

调用 Do 函数时，首先判断done值是否为0，若为1，表示传入的匿名函数 f() 已执行过，无需再次执行；若为0，表示传入的匿名函数 f() 还未执行过，则调用 doSlow() 函数进行初始化。

在 doSlow() 函数中，若并发的goroutine进入该函数中，为了保证仅有一个goroutine执行 f() 匿名函数。为此，需要加互斥锁保证只有一个goroutine进行初始化，同时采用了双检查的机制(double-checking)，再次判断 o.done 是否为 0，如果为 0，则是第一次执行，执行完毕后，就将 o.done 设置为 1，然后释放锁。

即使此时有多个 goroutine 同时进入了 doSlow 方法，因为双检查的机制，后续的 goroutine 会看到 o.done 的值为 1，也不会再次执行 f。

这样既保证了并发的 goroutine 会等待 f 完成，而且还不会多次执行 f。

举例：
```go
func main() {
	panicDo()
	nestedDo2()
}

func panicDo() {
	once := &sync.Once{}
	defer func() {
		if err := recover(); err != nil {
			once.Do(func() {
				fmt.Println("panic happened")
			})
		}
	}()
	once.Do(func() {
		panic("panic")
	})
}

func nestedDo() {
	once := &sync.Once{}
	once.Do(func(){
		once.Do(func(){
			fmt.Println("nested do")
		})
	})
}

func nestedDo2() {
	once1 := &sync.Once{}
	once2 := &sync.Once{}
	once1.Do(func(){
		once2.Do(func(){
			fmt.Println("nested do")
		})
	})
}
```
1. sync.Once()方法中传入的函数发生了panic，重复传入还会执行吗？

执行panicDo方法,不会打印任何东西. sync.Once.Do 方法中传入的函数只会被执行一次,哪怕函数中发生了 panic；

2. sync.Once()方法传入的函数中再次调用sync.Once()方法会有什么问题吗？

会发生死锁! 执行nestedDo方法,会报 fatal error: all goroutines are asleep - deadlock! 根据源码实现,可知在第二个do方法会一直等doshow()中锁的释放导致发生了死锁;

3. 执行nestedDo2,会输出什么?

会打印出 ‘nested do’. once1，once2是两个对象,互不影响. 所以sync.Once是使方法只执行一次对象的实现。


