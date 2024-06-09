# Go并发
## Goroutines
每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。

示例：Spinner动画，B_08_Spinner

示例：并发的Clock服务，参见B_09_Clock

示例：并发的Echo服务，参见B_10_Echo

## Channel
goroutine是Go语言程序的并发体，channel则是它们之间的通信机制。一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有可发送数据的类型。创建一个channel: `ch := make(chan int) // ch has type 'chan int'`。

和map类似，channel也一个对应make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者何被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较。`如果两个channel引用的是相同的对象，那么比较的结果为真`。一个channel也可以和nil进行比较。

一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都是用`<‐`运算符。一个不使用接收结果的接收操作也是合法的。
```go
ch <- x     // a send statement
x = <-ch    // a receive expression in an assignment statement 
<-ch        // a receive statement; result is discared
```
make可以指定第二个整形参数，对应channel的容量。如果channel的容量大于零，那么该channel就是带缓存的channel。
```go
ch = make(chan int)     // unbuffered channel
ch = make(chan int, 0)  // unbuffered channel
ch = make(chan int, 3)  // buffered channel with capacity 3
```
一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channel成功传输之后，两个goroutine可以继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channel上执行发送操作。

基于无缓存Channel的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为`同步Channel`。

Channel也可以用于将多个goroutine链接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channel就是所谓的管道(pipeline)。

Channel还支持close操作，用于关闭channel，使用内置的close函数就可以关闭一个channel: `close(ch)`。close后发送操作都将导致panic异常。`对一个已经被close过的channel执行接收操作依然可以接受到之前已经成功发送的数据；后续的接收操作将不再阻塞，它们会立即返回一个零值(一直可以读取零值)。`

没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式:它多接收一个结果，多接收的第二个结果是一个布尔值ok，true表示成功从channel接收到值，false表示channels已经被关闭并且里面没有值可接收。`x, ok := <‐naturals`

Go语言的range循环可直接在channel上面迭代 (使用range循环是上面ok判断的简洁语法)。它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。

例子：
```go
package main

import (
    "fmt"
    "time"
)

func counter(out chan<- int) { //类型chan<‐ int表示一个只发送int的channel，只能发送不能接收。
    for x := 0; x < 5; x++ {
        out <- x
        time.Sleep(1 * time.Second)
    }
    close(out)
}
func squarer(out chan<- int, in <-chan int) { //类型<‐chan int表示一个只接收int的channel，只能接收不能发送。
    for v := range in {
        out <- v * v
    }
    close(out)
}
func printer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
    }
}
func main() {
    naturals := make(chan int)
    squares := make(chan int)
    go counter(naturals)
    go squarer(squares, naturals)
    printer(squares)
}
// 每隔一秒打印0，1，4，9，16
```

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

可以用内置的cap函数获取channel内部缓存的容量: `fmt.Println(cap(ch))`。
内置的len函数返回channel内部缓存队列中有效元素的个数: `fmt.Println(len(ch))`。

## 并发循环 WaitGroup
多个goroutine同时工作是，为了知道最后一个goroutine什么时候结束(最后一个结束并不一定是最后一个开始)，我们需要一个递增的计数器，在每一个goroutine启动时加一，在goroutine退出时减一。这个计数器需要在多个goroutine操作时做到安全并且提供提供在其减为零之前一直等待的一种方法。这种计数类型被称为`sync.WaitGroup`，下面的代码就用到了这种方法：
```go
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doSomething(&wg, "Task "+fmt.Sprint(i))
	}
	wg.Wait()
	fmt.Println("All tasks completed.")
}

func doSomething(wg *sync.WaitGroup, taskName string) {
	defer wg.Done() // 完成任务后通知WaitGroup
	fmt.Printf("Starting task: %s\n", taskName)
	time.Sleep(1 * time.Second) // 模拟耗时操作
	fmt.Printf("Finished task: %s\n", taskName)
}
/*
Starting task: Task 4
Starting task: Task 1
Starting task: Task 2
Starting task: Task 3
Starting task: Task 0
Finished task: Task 0
Finished task: Task 4
Finished task: Task 1
Finished task: Task 2
Finished task: Task 3
All tasks completed.
*/
```
注意`Add`和`Done`方法的不对称。`Add是为计数器加一，必须在worker goroutine开始之前调用，而不是在goroutine中`；否则的话我们没办法确定Add是在"closer" goroutine调用Wait之前被调用。并且Add还有一个参数，但Done却没有任何参数；其实它和Add(-1)是等价的。我们使用defer来确保计数器即使是在出错的情况下依然能够正确地被减掉。上面的程序代码结构是当我们使用并发循环，但又不知道迭代次数时很通常而且很地道的写法。

### 示例：并发Web爬虫

## Select
火箭发射的例子，两个channel，一个负责计时，一个负责终止。现在每一次计数循环的迭代都需要等待两个channel中的其中一个返回事件了：ticker channel正常计数，或者异常时返回的abort事件。我们无法做到从每一个channel中接收信息，如果我们这么做的话，如果第一个channel中没有事件发过来那么程序就会立刻被阻塞，这样我们就无法收到第二个channel中发过来的事件。这时候我们需要多路复用(multiplex)这些操作了，为了能够多路复用，我们使用了select语句。

```go
select { 
    case <-ch1:
    // ...
    case x := <-ch2: 
    // ...use x...
    case ch3 <- y: 
    // ...
    default: // ...
}
```
select语句的一般形式：会有几个case和最后的default。每一个case代表一个通信操作(在某个channel上进行发送或者接收)并且会包含一些语句组成的一个语句块。

select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等待下去。如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。

如果 select 控制结构中包含 default 语句，那么这个 select 语句在执行时会遇到以下两种情况：
- 当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
- 当不存在可以收发的 Channel 时，执行 default 中的语句；

例1：
```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    fmt.Println("Commencing countdown. Press return to abort.")
    //tick := time.Tick(1 * time.Second)
    ticker := time.NewTicker(1 * time.Second)

    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1)) // read a single byte
        abort <- struct{}{}
    }()

    for s := 10; s >= 0; s-- {
        select {
        //case <-tick:
        case <-ticker.C:
            fmt.Println("T minis: ", s)
        case <-abort:
            fmt.Println("Launch aborted!")
            return
        }
    }
    ticker.Stop()
    launch()
}

func launch() {
    fmt.Println("Rocket launched!")
}
```

例2：ch这个channel的buffer大小是1，所以会交替的为空或为满，所以只有一个case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8。注意必须是buffer大小为1的channel，如果是没有buffer，这个例子会死锁。
```go
ch := make(chan int, 1)
for i := 0; i < 10; i++ {
    select {
    case x := <-ch:
        fmt.Println(x) // "0" "2" "4" "6" "8"
    case ch <- i:
    }
}
```
channel的零值是nil。对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。

### 示例：并发的字典遍历

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

### 超时控制，goroutine退出

为了能够达到退出多个goroutine的目的，我们需要靠谱的策略，来通过一个channel把消息广播出去，这样goroutine们能够看到这条事件消息，并且在事件完成之后，可以知道这件事已经发生过了。

可以使用context:

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 创建带有超时的上下文
	defer cancel()
	for i := 0; i < 5; i++ {
		go worker(ctx, i)
	}
	<-ctx.Done() // 程序会自动在超时或取消时结束
	time.Sleep(1 * time.Second)
	fmt.Println("Gracefully shutting down due to timeout or cancellation...")
}

func worker(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done(): // 监听上下文的Done通道
			fmt.Printf("Worker %d received cancellation signal. Exiting...\n", num)
			return
		default:
			fmt.Printf("Worker %d is running...\n", num)
			time.Sleep(1 * time.Second)
		}
	}
}
```

健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行。withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已，他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。

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

### withCancel
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

### 示例：聊天服务 -- Todo

---

## 基于共享变量的并发

### 竞争条件
在一个线性(只有一个goroutine的)的程序中，程序的执行顺序只由程序的逻辑来决定。在有两个或更多goroutine的程序中，每一个goroutine内的语句也是按照既定的顺序去执行的，但是一般情况下没法知道分别位于两个goroutine的事件x和y的执行顺序，x是在y之前还是之后还是同时发生是没法判断的。这也说明x和y这两个事件是并发的。

一个函数在线性、在并发的情况下，这个函数可以正确地工作，那么这个函数是并发安全的，并发安全的函数不需要额外的同步工作。可以把这个概念概括为一个特定类型的一些方法和操作函数，如果这个类型是并发安全的话，那么所有它的访问方法和操作就都是并发安全的。

在一个程序中有非并发安全的类型的情况下，我们依然可以使这个程序并发安全。并发安全的类型是例外，而不是规则，所以只有当文档中明确地说明了其是并发安全的情况下，你才可以并发地去访问它。

数据竞争：只要有两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。有三种方式可以避免数据竞争：第一种方法是不要去写变量。第二种避免多个goroutine访问变量。第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问，这种方式被称为“互斥”。

### sync.Mutex互斥锁
可以用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量。
```go
var (
    sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
    balance int
)
func Deposit(amount int) {
    sema <- struct{}{} // acquire token
    balance = balance + amount
    <-sema // release token
}
func Balance() int {
    sema <- struct{}{} // acquire token
    b := balance
    <-sema // release token
    return b
}
```
sync包里的Mutex类型，它的Lock方法能够获取到token(这里叫锁)，并且Unlock方法会释放这个token。

Mutex实现了Locker接口
```go
type Locker interface {
    Lock()
    Unlock()
}
```
也就是互斥锁 Mutex 提供两个方法 Lock 和 Unlock
```go
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
```
下面的写法与上面的功能一样：
```go
import "sync"
var (
    mu      sync.Mutex // guards balance
    balance int
)
func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock() // defer 更好
}
func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()  // defer 更好
    return b
}
```
每次一个goroutine访问bank变量时(这里只有balance余额变量)，它都会调用mutex的Lock方法来获取一个互斥锁。如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态。mutex会保护共享变量。

更好的方法是用defer来调用Unlock。无法对一个已经锁上的mutex来再次上锁­­这会导致程序死锁，没法继续执行下去。

Mutex可能处于两种操作模式下：正常模式和饥饿模式

1. 正常模式（Normal Mode）

- 默认模式：当创建一个新的 sync.Mutex 时，默认处于正常模式。
- FIFO队列：等待获取锁的goroutines会形成一个先进先出（FIFO）的等待队列。
- 被唤醒的goroutine不会直接持有锁，会和新进来的锁进行竞争，新请求进来的锁会更容易抢占到锁，因为正在CPU上运行，因此刚唤醒的goroutine可能会竞争失败，回到队列头部。
- 潜在饥饿问题：由于自旋和锁的释放并不保证FIFO队列中的下一个goroutine能立即获得锁，因此在高竞争环境下，某些goroutines可能长时间无法获得锁，导致“饥饿”现象。

2. 饥饿模式（Starvation Mode）
- 触发条件：当一个goroutine等待锁的时间超过一定阈值（通常是1毫秒），Mutex会被切换到饥饿模式。
- 直接传递：在饥饿模式下，解锁操作会直接将锁传递给等待队列中的第一个goroutine，而不是让其重新竞争，从而避免了自旋和可能的饥饿问题。
- 队列保证：后续到达的goroutines即使发现锁是未锁定状态，也不会尝试获取锁，而是直接加入到队列的末尾等待。
- 模式切换：一旦等待队列中的某个goroutine获得了锁并完成其临界区操作后，如果此时没有其他goroutine在自旋（如果当前goroutine已经是队列的最后一个或者当前goroutine等待时间小于1毫秒），Mutex会从饥饿模式切换回正常模式。

正常模式下，性能更好，但饥饿模式解决取锁公平问题，性能较差。

易错场景

- Lock/Unlock没有成对出现（加锁后必须有解锁操作），如果Lock之后，没有Unlock会出现死锁的情况，或者是因为 Unlock 一个未Lock的 Mutex 而导致 panic
- 复制已经使用过的Mutex，因为复制了已经使用了的Mutex，导致锁无法使用，程序处于死锁的状态
- 重入锁，Mutex是不可重入锁，如果一个线程成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态
- 死锁，两个或两个以上的goroutine争夺共享资源，互相等待对方的锁释放

### 内存同步
因为赋值和打印指向不同的变量，编译器可能会断定两条语句的顺序不会影响执行结果，并且会交换两个语句的执行顺序。可能的话，将变量限定在goroutine内部;如果是多个goroutine都需要访问的变量，使用互斥条件来访问。

### sync.RWMutex读写锁
允许多个只读操作并行执行，但写操作会完全互斥。这种锁叫作“多读单写”锁(multiple readers, single writer lock)，Go语言提供的这样的锁是sync.RWMutex。
```go
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock defer 
    mu.RUnlock()
    return balance
}
```
RWMutex 是一个 reader/writer 互斥锁。RWMutex 在某一时刻只能由任意数量的 reader goroutine 持有，或者是只被单个的 writer goroutine 持有，适用于读多写少的场景。

- Lock/Unlock：写操作时调用的方法
- RLock/RUnlock：读操作时调用的方法
- RLocker：这个方法的作用是为读操作返回一个 Locker 接口的对象。它的 Lock 方法会调用 RWMutex 的 RLock 方法，它的 Unlock 方法会调用 RWMutex 的 RUnlock 方法。

使用示例：
```go
func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				fmt.Println("worker: ", i, " count: ", counter.Count()) // 计数器读操作
				time.Sleep(1 * time.Second)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
```

底层结构
```go
type RWMutex struct {
    w           Mutex   // 互斥锁解决多个writer的竞争
    writerSem   uint32  // writer信号量
    readerSem   uint32  // reader信号量
    readerCount int32   // reader的数量（以及是否有 writer 竞争锁）
    readerWait  int32   // writer等待完成的reader的数量
}

const rwmutexMaxReaders = 1 << 30
```
实现原理:

一个 writer goroutine 获得了内部的互斥锁，就会反转 readerCount 字段，把它从原来的正整数 readerCount(>=0) 修改为负数（readerCount - rwmutexMaxReaders），让这个字段保持两个含义（既保存了 reader 的数量，又表示当前有 writer）。也就是说当readerCount为负数的时候表示当前writer goroutine持有写锁中，reader goroutine会进行阻塞。

当一个 writer 释放锁的时候，它会再次反转 readerCount 字段。可以肯定的是，因为当前锁由 writer 持有，所以，readerCount 字段是反转过的，并且减去了 rwmutexMaxReaders 这个常数，变成了负数。所以，这里的反转方法就是给它增加 rwmutexMaxReaders 这个常数值。

易错场景

- 复制已经使用的读写锁，会把它的状态也给复制过来，原来的锁在释放的时候，并不会修改你复制出来的这个读写锁，这就会导致复制出来的读写锁的状态不对，可能永远无法释放锁
- 重入导致死锁，因为读写锁内部基于互斥锁实现对 writer 的并发访问，而互斥锁本身是有重入问题的，所以，writer 重入调用 Lock 的时候，就会出现死锁的现象, 在 reader 的读操作时调用 writer 的写操作（调用 Lock 方法），那么，这个 reader 和 writer 就会形成互相依赖的死锁状态。
当一个 writer 请求锁的时候，如果已经有一些活跃的 reader，它会等待这些活跃的 reader 完成，才有可能获取到锁，但是，如果之后活跃的 reader 再依赖新的 reader 的话，这些新的 reader 就会等待 writer 释放锁之后才能继续执行，这就形成了一个环形依赖： writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader 依赖 writer
- 释放未加锁的 RWMutex，和互斥锁一样，Lock 和 Unlock 的调用总是成对出现的，RLock 和 RUnlock 的调用也必须成对出现。Lock 和 RLock 多余的调用会导致锁没有被释放，可能会出现死锁，而 Unlock 和 RUnlock 多余的调用会导致 panic

### sync.Once初始化
sync包提供了一个专门的方案来解决一次性初始化的问题:sync.Once。一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了;互斥量用来保护boolean变量和客户端数据结构。
```go
var loadIconsOnce sync.Once
var icons map[string]image.Image
// Concurrency‐safe.
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
//每一次对Do(loadIcons)的调用都会锁定mutex，并会检查boolean变量。在第一次调用时，变量的值是false，Do会调用loadIcons并会将boolean设置为true。
```

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

#### 源码分析
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

### 示例：并发的非阻塞缓存 -- todo

### Goroutines和线程
每一个OS线程都有一个固定大小的内存块(一般会是2MB)来做栈，这个栈会用来存储当前正在被调用或挂起(指在调用其它函数时)的函数的内部变量。2MB的栈对于一个小小的goroutine来说是很大的内存浪费。修改固定的大小可以提升空间的利用率允许创建更多的线程，并且可以允许更深的递归调用，不过这两者是没法同时兼备的。

相反，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。一个goroutine的栈，和操作系统线程一样，会保存其活跃或挂起的函数调用的本地变量，但是和OS线程不太一样的是一个goroutine的栈大小并不是固定的;栈的大小会根据需要动态地伸缩。而goroutine的栈的最大值有1GB，比传统的固定大小的线程栈要大得多。

OS线程会被操作系统内核调度。每几毫秒，一个硬件计时器会中断处理器，这会调用一个叫作scheduler的内核函数。这个函数会挂起当前执行的线程并保存内存中它的寄存器内容，检查线程列表并决定下一次哪个线程可以被运行，并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。因为操作系统线程是被内核所调度，所以从一个线程向另一个“移动”需要完整的上下文切换，也就是说，保存一个用户线程的状态到内存，恢复另一个线程的到寄存器，然后更新调度器的数据结构。这几步操作很慢，因为其局部性很差需要几次内存访问，并且会增加运行的cpu周期。

Go的运行时包含自己的调度器，这个调度器使用了一些技术手段，比如m:n调度，因为其会在n个操作系统线程上多工(调度)m个goroutine。Go调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的goroutine(按程序独立)。
和操作系统的线程调度不同的是，Go调度器并不是用一个硬件定时器而是被Go语言"建筑"本身进行调度的。这种调度方式不需要进入内核的上下文，所以重新调度一个goroutine比调度一个线程代价要低得多。

Go的调度器使用了一个叫做 `GOMAXPROCS` 的变量来决定会有多少个操作系统的线程同时执行Go的代码。其默认的值是运行机器上的CPU的核心数，所以在一个有8个核心的机器上时，调度器一次会在8个OS线程上去调度GO代码。(GOMAXPROCS是m:n调度中的n)。在休眠中的或者在通信中被阻塞的goroutine是不需要一个对应的线程来做调度的。
可以用GOMAXPROCS的环境变量来显式地控制这个参数，或者也可以在运行时用runtime.GOMAXPROCS函数来修改它。
`GOMAXPROCS=1 go run main.go`

在大多数支持多线程的操作系统和程序语言中，当前的线程都有一个独特的身份(id)，并且这个身份信息可以以一个普通值的形式被被很容易地获取到。
goroutine没有可以被程序员获取到的身份(id)的概念。

## sync.Cond
Cond 通常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。
```go
// 基本方法
func NeWCond(l Locker) *Cond 
func (c *Cond) Broadcast() 
func (c *Cond) Signal() 
func (c *Cond) Wait()
// Singal(): 唤醒一个等待此 Cond 的 goroutine
// Broadcast(): 唤醒所有等待此 Cond 的 goroutine
// Wait(): 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒，使用该方法是需要搭配满足条件
```

使用示例：
```go
func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员#%d 已准备就绪\n", i)
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
```

实现原理
```go
type Cond struct {
    noCopy noCopy

    // 当观察或者修改等待条件的时候需要加锁
    L Locker

    // 等待队列
    notify  notifyList
    checker copyChecker
}

func NewCond(l Locker) *Cond {
    return &Cond{L: l}
}

func (c *Cond) Wait() {
    c.checker.check()
    // 增加到等待队列中
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    // 阻塞休眠直到被唤醒
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}

func (c *Cond) Signal() {
    c.checker.check()
    runtime_notifyListNotifyOne(&c.notify)
}

func (c *Cond) Broadcast() {
    c.checker.check()
    runtime_notifyListNotifyAll(&c.notify）
}
```
在上述的实现源码中，Signal和Broadcast调用了底层的通知方法；重点在Wait方法中，把调用者加入到等待队列时会释放锁，在被唤醒之后还会请求锁。在阻塞休眠期间，调用者是不持有锁的，这样能让其他 goroutine 有机会检查或者更新等待变量，因此在使用Wait方法的时候必须持有锁。

易错场景

- 调用Wait方法没有加锁
- 没有检查等待条件是否满足

--- 
## 练习：

下面语法正确的是（）
```go
// A. var ch chan int
// B. ch := make(chan int)
// C. <-ch
// D. ch<-
/*
答：A、B、C
A、B都是申明channel；C读取channel；写channel是必须带上值，所以D错误。
*/
```

关于channel的特性，下面说法正确的是？ 答：A B C D

- A. 给一个 nil channel 发送数据，造成永远阻塞
- B. 从一个 nil channel 接收数据，造成永远阻塞
- C. 给一个已经关闭的 channel 发送数据，引起 panic
- D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

关于select机制，下面说法正确的是? 答：A B C

- A. select机制用来处理异步IO问题；
- B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
- C. golang在语言级别支持select关键字；
- D. select关键字的用法与switch语句非常类似，后面要带判断条件；

关于 channel 下面描述正确的是？ 答：C

- A. close() 可以用于只接收通道；
- B. 单向通道可以转换为双向通道；
- C. 不能在单向通道上做逆向操作（例如：只发送通道用于接收）；

关于无缓冲和有冲突的channel，下面说法正确的是？ 答：D

- A. 无缓冲的channel是默认的缓冲为1的channel；
- B. 无缓冲的channel和有缓冲的channel都是同步的；
- C. 无缓冲的channel和有缓冲的channel都是非同步的；
- D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；

关于协程，下面说法正确是？ 答：A D

- A. 协程和线程都可以实现程序的并发执行；
- B. 线程比协程更轻量级；
- C. 协程不存在死锁问题；
- D. 通过 channel 来进行协程间的通信；

下面代码会触发异常吗？请说明。
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
// `select` 会随机选择一个可用通道做收发操作，所以可能触发异常，也可能不会。
```

下面的代码有什么问题？
```go
func Stop(stop <-chan bool) {
    close(stop)
}
```
`有方向的 channel 不可以被关闭。`

下面代码输出什么？
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
//default
//ch 为 nil，读写都会阻塞。
```

下面的代码输出什么？
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
/* 输出： 321
第一次循环，写操作已经准备好，执行 o(3)，输出 3；
第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
第三次，由于 c 为 nil，走的是 default 分支，输出 1。
*/
```

下面的代码有什么问题？
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
/*
锁失效
将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
*/

//修复代码：
func (d *data) test(s string)  {  // 指针接收者
    d.Lock()
    defer d.Unlock()
    for i:=0;i<5 ;i++  {
        fmt.Println(s,i)
        time.Sleep(time.Second)
    }
}
//或者可以通过嵌入 `*Mutex` 来避免复制的问题，但需要初始化。

type data struct {
    *sync.Mutex     // *Mutex
}

func (d data) test(s string) {  // 值方法
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

下面的代码有什么问题？
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
// 在协程中使用 wg.Add()；
// 使用了 sync.WaitGroup 副本；

// 修复代码：
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

// 或者：
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

下面代码有什么问题，请说明？
```go
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
	}
}
/*
以上代码在go1.14版本之前(不含1.14版本): for {} 独占 CPU 资源导致其他 Goroutine 饿死，

在go1.14版本之后(包含go1.14): 会打印0123456789, 并且主程会进入死循环。

这是因为1.14版本(含1.14版本)之后goroutine抢占式调度设计改为基于信号的抢占式调度，当调度器监控发现某个goroutine执行时间过长且有别的goroutine在等待时，会把执行时间过长的goroutine暂停，转而调度等待的goroutine. 所以for循环的goroutine得以执行.

可以通过阻塞的方式避免 CPU 占用，修复代码：
*/
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

两个goroutine 交替打印 1-20
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

下面的代码有什么问题？
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
// `panic: sync: WaitGroup is reused before previous Wait has returned`
// 协程里面，使用 wg.Add(1) 但是没有 wg.Done()，导致 panic()。
```

下面代码输出什么？
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
//当 A 协程还没起时，主协程已经将 channel 关闭了，当 A 协程往关闭的 channel 发送数据时会 panic，`panic: send on closed channel`。
```

map可以被多个goroutine操作吗？ 可以，线程不安全，怎么办，枷锁，其他办法，sync.Map 线程安全。





07 Channel

底层结构
qcount 已经接收但还未被取走的元素个数 内置函数len获取到

datasiz 循环队列的大小 暂时认为是cap容量的值

elemtype和elemsize 声明chan时到元素类型和大小 固定

buf 指向缓冲区的指针 无缓冲通道中 buf的值为nil

sendx 处理发送进来数据的指针在buf中的位置 接收到数据 指针会加上elemsize，移向下一个位置

recvx 处理接收请求（发送出去）的指针在buf中的位置

recvq 如果没有数据可读而阻塞， 会加入到recvq队列中

sendq 向一个满了的buf 发送数据而阻塞，会加入到sendq队列中

实现原理
向channel写数据的流程：

有缓冲区：
优先查看recvq是否为空，如果不为空，优先唤醒recvq的中goroutine，并写入数据；如果队列为空，则写入缓冲区，如果缓冲区已满则写入sendq队列；

无缓冲区：
直接写入sendq队列

向channel读数据的流程：

有缓冲区：优先查看缓冲区，如果缓冲区有数据并且未满，直接从缓冲区取出数据；
如果缓冲区已满并且sendq队列不为空，优先读取缓冲区头部的数据，并将队列的G的数据写入缓冲区尾部；

无缓冲区：将当前goroutine加入recvq队列，等到写goroutine的唤醒

易错点
channel未初始化，写入或者读取都会阻塞
往close的channel写入数据会发生panic
close未初始化channel会发生panic
close已经close过的channel会发生panic



08 SingleFlight
基本概念
SingleFlight 是 Go 开发组提供的一个扩展并发原语。它的作用是，在处理多个 goroutine 同时调用同一个函数的时候，只让一个 goroutine 去调用这个函数，等到这个 goroutine 返回结果的时候，再把结果返回给这几个同时调用的 goroutine，这样可以减少并发调用的数量。

与sync.Once的区别
sync.Once 不是只在并发的时候保证只有一个 goroutine 执行函数 f，而是会保证永远只执行一次，而 SingleFlight 是每次调用都重新执行，并且在多个请求同时调用的时候只有一个执行。
sync.Once 主要是用在单次初始化场景中，而 SingleFlight 主要用在合并并发请求的场景中
应用场景
使用 SingleFlight 时，可以通过合并请求的方式降低对下游服务的并发压力，从而提高系统的性能，常常用于缓存系统中

基本方法
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)
提供一个 key，对于同一个 key，在同一时间只有一个在执行，同一个 key 并发的请求会等待。第一个执行的请求返回的结果，就是它的返回结果。函数 fn 是一个无参的函数，返回一个结果或者 error，而 Do 方法会返回函数执行的结果或者是 error，shared 会指示 v 是否返回给多个请求

func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result
类似 Do 方法，只不过是返回一个 chan，等 fn 函数执行完，产生了结果以后，就能从这个 chan 中接收这个结果

func (g *Group) Forget(key string)
告诉 Group 忘记这个 key。这样一来，之后这个 key 请求会执行 f，而不是等待前一个未完成的 fn 函数的结果

实现方法
SingleFlight 定义一个辅助对象 call，这个 call 就代表正在执行 fn 函数的请求或者是已经执行完的请求

在Do方法中，传入key与执行函数，加锁，查询是否存在key，如果存在，等待第一个请求完成并返回。如果不存在，创建一个call，将这个call加入到key map中，释放锁，执行doCall函数，执行完实际函数后，删除key。

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
  g.mu.Lock()
  if g.m == nil {
    g.m = make(map[string]*call)
  }
  if c, ok := g.m[key]; ok {//如果已经存在相同的key
    c.dups++
    g.mu.Unlock()
    c.wg.Wait() //等待这个key的第一个请求完成
    return c.val, c.err, true //使用第一个key的请求结果
  }
  c := new(call) // 第一个请求，创建一个call
  c.wg.Add(1)
  g.m[key] = c //加入到key map中
  g.mu.Unlock()


  g.doCall(c, key, fn) // 调用方法
  return c.val, c.err, c.dups > 0
}

func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {
  c.val, c.err = fn()
  c.wg.Done()


  g.mu.Lock()
  if !c.forgotten { // 已调用完，删除这个key
    delete(g.m, key)
  }
  for _, ch := range c.chans {
    ch <- Result{c.val, c.err, c.dups > 0}
  }
  g.mu.Unlock()
}