# Go并发

## 基本概念
### 串行、并发与并行
- 串行：顺序执行。
- 并发：同一时间段内执行多个任务。
- 并行：同一时刻执行多个任务。

### 进程、线程和协程
- 进程（process）：是应用程序的启动实例，程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
- 线程（thread）：操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位。多个线程之间可以共享进程的资源并通过共享内存等线程间的通信方式来通信。
- 协程（coroutine）：为轻量级线程，与线程相比，协程不受操作系统的调度，协程的调度器由用户应用程序提供，协程调度器按照调度策略把协程调度到线程中运行。

早期单进程时代不需要调度器，一切的软件都是跑在操作系统上，真正用来干活(计算)的是CPU。早期的操作系统每个程序就是一个进程，一个程序运行完，才能进行下一个进程，就是“单进程时代”，一切的程序只能串行发生。

单进程一旦进程阻塞，CPU的时间就浪费了，后来操作系统就具有了最早的并发能力：多进程并发，当一个进程阻塞的时候，切换到另外等待执行的进程，这样就能尽量把CPU利用起来，CPU就不浪费了。而且调度cpu的算法可以保证在运行的进程都可以被分配到cpu的运行时间片。这样从宏观来看，似乎多个进程是在同时被运行。

现代操作系统调度线程都是抢占式的，我们不能依赖用户代码主动让出CPU，或者因为IO、锁等待而让出，这样会造成调度的不公平。基于经典的时间片算法，当线程的时间片用完之后，会被时钟中断给打断，调度器会将当前线程的执行上下文进行保存，然后恢复下一个线程的上下文，分配新的时间片令其开始执行。这种抢占对于线程本身是无感知的，系统底层支持，不需要开发人员特殊处理。

但新的问题就又出现了，进程拥有太多的资源，进程的创建、切换、销毁，都会占用很长的时间，CPU虽然利用起来了，但如果进程过多，CPU有很大的一部分都被用来进行进程调度了。（linux对待进程，线程类似）

并且虽然多进程、多线程已经提高了系统的并发能力，但高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存(进程虚拟内存会占用4GB(32位操作系统), 而线程也要大约4MB)。

其实一个线程分为“内核态“线程和”用户态“线程。一个“用户态线程”必须要绑定一个“内核态线程”，但是CPU并不知道有“用户态线程”的存在，它只知道它运行的是一个“内核态线程”。

再去细化去分类一下，内核线程依然叫“线程(thread)”，用户线程叫“协程(co-routine)”。既然一个协程(co-routine)可以绑定一个线程(thread)，那么能不能多个协程(co-routine)绑定一个或者多个线程(thread)上呢。

CPU负责线程，而用户自己开发调度器，调度携程。​之后，我们就看到了有3中协程和线程的映射关系：

1. `N:1关系`：N个协程绑定1个线程，优点就是协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。但也有很大的缺点，1个进程的所有协程都绑定在1个线程上，用不了硬件的多核加速能力，一旦某协程阻塞，造成线程阻塞，本进程的其他协程都无法执行了，根本就没有并发的能力了。？？？？？？？？？？
2. `1:1 关系`：1个协程绑定1个线程，这种最容易实现。协程的调度都由CPU完成了，不存在N:1缺点，协程的创建、删除和切换的代价都由CPU完成，有点略显昂贵了。这不就是普通多线程吗。
3. `M:N关系`：M个协程绑定N个线程，是N:1和1:1类型的结合，克服了以上2种模型的缺点，但实现起来最为复杂。压力全在协程调度。

### 并发模型

业界将如何实现并发编程总结归纳为各式各样的并发模型，常见的并发模型有以下几种：

- 线程&锁模型
- Actor模型
- CSP模型
- Fork&Join模型

Go语言中的并发程序主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现，当然也支持使用传统的多线程共享内存的并发方式。

## goroutine

Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。

### go关键字

Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上`go`关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。

```go
go f()  // 创建一个新的 goroutine 运行函数f
```

匿名函数也支持使用`go`关键字创建 goroutine 去执行。

```go
go func(){
  // ...
}()
```

一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。

```go
func hello() {
	fmt.Println("hello")
}

func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
}
// 输出：main goroutine done!
```

这一次的执行结果没有打印 `hello`。这是为什么呢？

其实在 Go 程序启动时，Go 程序就会为 main 函数创建一个默认的 goroutine 。在上面的代码中我们在 main 函数中使用 go 关键字创建了另外一个 goroutine 去执行 hello 函数，而此时 main goroutine 还在继续往下执行，我们的程序中此时存在两个并发执行的 goroutine。当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，所有由 main goroutine 创建的 goroutine 也会一同退出。也就是说我们的 main 函数退出太快，另外一个 goroutine 中的函数还未执行完程序就退出了，导致未打印出“hello”。

Go 语言中通过`sync`包为我们提供了一些常用的并发原语，我们会在后面的小节单独介绍`sync`包中的内容。在这一小节，我们会先介绍一下 sync 包中的`WaitGroup`。当你并不关心并发操作的结果或者有其它方式收集并发操作的结果时，`WaitGroup`是实现等待一组并发操作完成的好方法。

下面的示例代码中我们在 main goroutine 中使用`sync.WaitGroup`来等待 hello goroutine 完成后再退出。

```go
package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done() // 告知当前goroutine完成
}

func main() {
	wg.Add(1) // 登记1个goroutine
	go hello()
	fmt.Println("你好")
	wg.Wait() // 阻塞等待登记的goroutine完成
}
// 你好
// hello
// hello goroutine 执行完毕后程序直接退出。
```

### 启动多个goroutine

在 Go 语言中实现并发就是这样简单，我们还可以启动多个 goroutine 。让我们再来看一个新的代码示例。这里同样使用了`sync.WaitGroup`来实现 goroutine 的同步。

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
```

多次执行上面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。

### 动态栈

操作系统的线程一般都有固定的栈内存（通常为2MB）， 这个栈会用来存储当前正在被调用或挂起(指在调用其它函数时)的函数的内部变量，2MB的栈对于一个小小的goroutine来说是很大的内存浪费而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的，一个goroutine的栈，和操作系统线程一样，会保存其活跃或挂起的函数调用的本地变量，区别于操作系统线程由系统内核进行调度，goroutine 是由Go运行时（runtime）负责调度。goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。

### goroutine调度

OS线程会被操作系统内核调度。每几毫秒，一个硬件计时器会中断处理器，这会调用一个叫作scheduler的内核函数。这个函数会挂起当前执行的线程并保存内存中它的寄存器内容，检查线程列表并决定下一次哪个线程可以被运行，并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。因为操作系统线程是被内核所调度，所以从一个线程向另一个“移动”需要完整的上下文切换，也就是说，保存一个用户线程的状态到内存，恢复另一个线程的到寄存器，然后更新调度器的数据结构。这几步操作很慢，因为其局部性很差需要几次内存访问，并且会增加运行的cpu周期。

goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行（比如m:n调度）。完全是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身 goroutine 的超轻量级，以上种种特性保证了 goroutine 调度方面的性能。

在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 `GPM` 调度模型。

#### 被废弃的goroutine调度器
Go目前使用的调度器是2012年重新设计的，因为之前的调度器性能存在问题，所以使用4年就被废弃了。（用G来表示Goroutine，用M来表示线程）

只有一个全局队列，​M想要执行、放回G都必须访问全局G队列，并且M有多个，即多线程访问同一资源需要加锁进行保证互斥/同步，所以全局G队列是有互斥锁进行保护的。有几个缺点：

- 创建、销毁、调度G都需要每个M获取锁，这就形成了激烈的锁竞争。
- M转移G会造成延迟和额外的系统负载。比如当G中包含创建新协程的时候，M创建了G’，为了继续执行G，需要把G’交给M’执行，也造成了很差的局部性，因为G’和G是相关的，最好放在M上执行，而不是其他M’。
- 系统调用(CPU在M之间的切换)导致频繁的线程阻塞和取消阻塞操作增加了系统开销。

### Goroutine调度器的GMP模型

![gpm](/file/img/go_gpm.png) 

面对之前调度器的问题，Go设计了新的调度器。在新调度器中，除了M(thread)和G(goroutine)，又引进了P(Processor)。
Processor，它包含了运行goroutine的资源，如果线程想运行goroutine，必须先获取P，P中还包含了可运行的G队列。

线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上。

1. 全局队列（Global Queue）：存放等待运行的G。
2. P的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建G’时，G’优先加入到P的本地队列，如果队列满了，则会把本地队列中一半的G移动到全局队列。
3. P列表：所有的P都在程序启动时创建，并保存在数组中，最多有GOMAXPROCS(可配置)个。
4. M：线程想运行任务就得获取P，从P的本地队列获取G，P队列为空时，M也会尝试从全局队列拿一批G放到P的本地队列，或从其他P的本地队列偷一半放到自己P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断重复下去。

Goroutine调度器和OS调度器是通过M结合起来的，每个M都代表了1个内核线程，OS调度器负责把内核线程分配到CPU的核上执行。

**P和M何时会被创建**
1. P何时创建：在确定了P的最大数量n后，运行时系统会根据这个数量创建n个P。
2. M何时创建：没有足够的M来关联P并运行其中的可运行的G。比如所有的M此时都阻塞住了，而P中还有很多就绪任务，就会去寻找空闲的M，而没有空闲的，就会去创建新的M。

**G,P,M 的个数问题**：

1. G 的个数理论上是无限制的，但是受内存限制，
2. P 的数量一般建议是逻辑 CPU 数量的 2 倍，由启动时环境变量`$GOMAXPROCS`或者是由runtime的方法`GOMAXPROCS()`决定。这意味着在程序执行的任意时刻都只有`$GOMAXPROCS`个goroutine在同时运行。**P决定了并行数，注意不是并发数**
3. M 的数量
	- go语言本身的限制：go程序启动时，会设置M的最大数量，默认10000.但是内核很难支持这么多的线程数，所以这个限制可以忽略。
	- runtime/debug中的SetMaxThreads函数，设置M的最大数量
	- 一个M阻塞了，会创建新的M。
	- 空闲的M会回收或者休眠。

M与P的数量没有绝对关系，一个M阻塞，P就会去创建或者切换另一个M，所以，即使P的默认数量是1，也有可能会创建很多个M出来。

**为什么要有 P**
- 每个 P 有自己的本地队列，大幅度的减轻了对全局队列的直接依赖，所带来的效果就是锁竞争的减少。而 GM 模型的性能开销大头就是锁竞争。
- 每个 P 相对的平衡上，在 GMP 模型中也实现了 Work Stealing （工作量窃取机制）算法，如果 P 的本地队列为空，则会从全局队列或其他 P 的本地队列中窃取可运行的 G 来运行，减少空转，提高了资源利用率。

那为什么不直接在 M 实现本地队列、Work Stealing 算法？
- 一般来讲，M 的数量都会多于 P。像在 Go 中，M 的数量默认是 10000，P 的默认数量的 CPU 核数。另外由于 M 的属性，也就是如果存在系统阻塞调用，阻塞了 M，又不够用的情况下，M 会不断增加。
- M 不断增加的话，如果本地队列挂载在 M 上，那就意味着本地队列也会随之增加。这显然是不合理的，因为本地队列的管理会变得复杂，且 Work Stealing 性能会大幅度下降。
- M 被系统调用阻塞后，我们是期望把他既有未执行的任务分配给其他继续运行的，而不是一阻塞就导致全部停止。

因此使用 M 是不合理的，那么引入新的组件 P，把本地队列关联到 P 上，就能很好的解决这个问题。

### 调度器的设计策略

**复用线程**：避免频繁的创建、销毁线程，而是对线程的复用。
1. work stealing（工作量窃取）机制，当本线程无可运行的G时，会优先从全局队列里进行窃取，之后会从其它的P队列里窃取一半的G，放入到本地P队列里，而不是销毁线程。
2. hand off（移交）机制，当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。

**利用并行**：GOMAXPROCS设置P的数量，最多有GOMAXPROCS个线程分布在多个CPU上同时运行。GOMAXPROCS也限制了并发的程度，比如GOMAXPROCS = 核数/2，则最多利用了一半的CPU核进行并行。

**抢占**：在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU 10ms，防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方。

**全局G队列**：在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G。

**抢占式调度是如何抢占的？**

就像操作系统要负责线程的调度一样，Go的runtime要负责goroutine的调度。基于**时间片**的抢占式调度有个明显的优点，能够避免CPU资源持续被少数线程占用，从而使其他线程长时间处于饥饿状态。goroutine的调度器也用到了时间片算法，但是和操作系统的线程调度还是有些区别的，因为整个Go程序都是运行在用户态的，所以不能像操作系统那样利用时钟中断来打断运行中的goroutine。也得益于完全在用户态实现，goroutine的调度切换更加轻量。

**go func()流程**

![go func](/file/img/go_gmp_go_func_workflow.jpeg)

1. 我们通过 go func()来创建一个goroutine；
2. 有两个存储G的队列，一个是局部调度器P的本地队列、一个是全局G队列。新创建的G会先保存在P的本地队列中，如果P的本地队列已经满了就会保存在全局的队列中；
3. G只能运行在M中，一个M必须持有一个P，M与P是1：1的关系。M会从P的本地队列弹出一个可执行状态的G来执行，如果P的本地队列为空，就会想其他的MP组合偷取一个可执行的G来执行；
4. 一个M调度G执行的过程是一个循环机制 （调度，执行，销毁G，返回）；
5. 当M执行某一个G时候如果发生了syscall或则其余阻塞操作，M会阻塞，如果当前有一些G在执行，runtime会把这个线程M从P中摘除(detach)，然后再创建一个新的操作系统的线程(如果有空闲的线程可用就复用空闲线程)来服务于这个P；
6. 当M系统调用结束时候，这个G会尝试获取一个空闲的P执行，并放入到这个P的本地队列。如果获取不到P，那么这个线程M变成休眠状态，加入到空闲线程中，然后这个G会被放入全局队列中。

### 4、调度器的生命周期

![img](/file/img/go_gmp_sheduler_life.png)

特殊的M0和G0
- M0是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G(一般为main的gouroutine)， 在之后M0就和其他的M一样了。
- G0是每次启动一个M都会第一个创建的gourtine，G0仅用于负责调度的G，G0不指向任何可执行的函数, 每个M都会有一个自己的G0。在调度或系统调用时会使用G0的栈空间, 全局变量的G0是M0的G0。

我们来跟踪一段代码，它会经历如上图所示的过程：

```go
package main 
import "fmt" 
func main() {
    fmt.Println("Hello world") 
}
```

1. runtime创建最初的线程m0和goroutine g0，并把2者关联。
2. 调度器初始化：初始化m0、栈、垃圾回收，以及创建和初始化由GOMAXPROCS个P构成的P列表。
3. 示例代码中的main函数是main.main，runtime中也有1个main函数——runtime.main，代码经过编译后，runtime.main会调用main.main，程序启动时会为runtime.main创建goroutine，称它为main goroutine吧，然后把main goroutine加入到P的本地队列。
4. 启动m0，m0已经绑定了P，会从P的本地队列获取G，获取到main goroutine。
5. G拥有栈，M根据G中的栈信息和调度信息设置运行环境
6. M运行G
7. G退出，再次回到M获取可运行的G，这样重复下去，直到main.main退出，runtime.main执行Defer和Panic处理，或调用runtime.exit退出程序。

调度器的生命周期几乎占满了一个Go程序的一生，runtime.main的goroutine执行之前都是为调度器做准备工作，runtime.main的goroutine运行，才是调度器的真正开始，直到runtime.main结束而结束。

reference:
- https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmeduvk27bo0

### GOMAXPROCS

Go运行时的调度器使用`GOMAXPROCS`参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。Go语言中可以通过`runtime.GOMAXPROCS`函数设置当前程序并发时占用的 CPU逻辑核心数。（Go1.5版本之前，默认使用的是单核心执行。Go1.5 版本之后，默认使用全部的CPU 逻辑核心数。）

(GOMAXPROCS是m:n调度中的n)。在休眠中的或者在通信中被阻塞的goroutine是不需要一个对应的线程来做调度的。可以用GOMAXPROCS的环境变量来显式地控制这个参数，或者也可以在运行时用runtime.GOMAXPROCS函数来修改它。
`GOMAXPROCS=1 go run main.go`

在大多数支持多线程的操作系统和程序语言中，当前的线程都有一个独特的身份(id)，并且这个身份信息可以以一个普通值的形式被被很容易地获取到。goroutine没有可以被程序员获取到的身份(id)的概念。

### 练习题

1. 请写出下面程序的执行结果。

```go
func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
// 可能打印不出任何输出，因为main goroutine 结束后，go 函数中的 goroutine 也会结束。
// 可以这样打印 0 1 2 3 4，但顺序不定
func main() {
	var wg3 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func(i int) {
			defer wg3.Done()
			fmt.Println(i)
		}(i)
	}
	wg3.Wait()
}
```

## channel

单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言采用的并发模型是`CSP（Communicating Sequential Processes）`，提倡**通过通信共享内存**而不是**通过共享内存而实现通信**。

如果说 goroutine 是Go程序并发的执行体，`channel`就是它们之间的连接。`channel`是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

### channel类型

`channel`是 Go 语言中一种特有的类型。声明通道类型变量的格式如下：

```go
var 变量名称 chan 元素类型
```
- chan：是关键字
- 元素类型：是指通道中传递元素的类型

举几个例子：

```go
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
```

### 初始化channel

未初始化的通道类型变量其默认零值是`nil`。

```go
var ch chan int
fmt.Println(ch) // <nil>
```

声明的通道类型变量需要使用内置的`make`函数初始化之后才能使用。具体格式如下：

```go
make(chan 元素类型, [缓冲大小])
```
- channel的缓冲大小是可选的。

举几个例子：

```go
ch4 := make(chan int)
ch5 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道
fmt.Println(ch5) //	0xc000100000
```

当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者何被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较。`如果两个channel引用的是相同的对象，那么比较的结果为真`。一个channel也可以和nil进行比较。

### channel操作

通道共有发送（send）、接收(receive)和关闭（close）三种操作。而发送和接收操作都使用`<-`符号。

```go
ch := make(chan int)
ch <- 10 // 把10发送到ch中
```

```go
x := <- ch // 从ch中接收值并赋值给变量x
<-ch       // 从ch中接收值，忽略结果
```

我们通过调用内置的`close`函数来关闭通道。

```go
close(ch)
```

**注意：**一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

关闭后的通道有以下特点：

1. 对一个关闭的通道再发送值就会导致 panic。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致 panic。

### 无缓冲的通道

无缓冲的通道又称为阻塞的通道。我们来看一下如下代码片段。

```go
func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
/*
上面这段代码能够通过编译，但是执行的时候会出现以下错误：

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../main.go:8 +0x54
*/
```

`deadlock`表示我们程序中的 goroutine 都被挂起导致程序死锁了。为什么会出现`deadlock`错误呢？

因为我们使用`ch := make(chan int)`创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。

上面的代码会阻塞在`ch <- 10`这一行代码形成死锁，那如何解决这个问题呢？

其中一种可行的方法是创建一个 goroutine 去接收值，例如：

```go
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
// 注意，这里虽然不会阻塞，但输出有两种可能：
// 1. 接收成功，10 发送成功
// 2. 只打印发送成功
```

首先无缓冲通道`ch`上的发送操作会阻塞，直到另一个 goroutine 在该通道上执行接收操作，这时数字10才能发送成功，两个 goroutine 将继续执行。相反，如果接收操作先执行，接收方所在的 goroutine 将阻塞，直到 main goroutine 中向该通道发送数字10。

使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为`同步通道`。

### 有缓冲的通道

还有另外一种解决上面死锁问题的方法，那就是使用有缓冲区的通道。我们可以在使用 make 函数初始化通道时，可以为其指定通道的容量，例如：

```go
func main() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}
```

只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。

我们可以使用内置的`len`函数获取通道内元素的数量，使用`cap`函数获取通道的容量，虽然我们很少会这么做。

### 多返回值模式

当向通道中发送完数据时，我们可以通过`close`函数来关闭通道。当一个通道被关闭后，再往该通道发送值会引发`panic`，从该通道取值的操作会先取完通道中的值。通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。那我们如何判断一个通道是否被关闭了呢？

对一个通道执行接收操作时支持使用如下多返回值模式。

```go
value, ok := <- ch
```
- value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
- ok：通道ch关闭时返回 false，否则返回 true。

下面代码片段中的`f2`函数会循环从通道`ch`中接收所有值，直到通道被关闭后退出。

```go
func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}
// v:1 ok:true
// v:2 ok:true
// 通道已关闭
```

### for range接收值

通常我们会选择使用`for range`循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。上面那个示例我们使用`for range`改写后会很简洁。

```go
func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
```

**注意：**目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过`len(ch)`操作来判断通道是否被关闭。

### 单向通道

在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，比如限制通道在某个函数中只能执行发送或只能执行接收操作。想象一下，我们现在有`Producer`和`Consumer`两个函数，其中`Producer`函数会返回一个通道，并且会持续将符合条件的数据发送至该通道，并在发送完成后将该通道关闭。而`Consumer`函数的任务是从通道中接收值进行计算，这两个函数之间通过`Producer`函数返回的通道进行通信。完整的示例代码如下。

```go
package main

import (
	"fmt"
)

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch // 通道工厂？
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
}
```

从上面的示例代码中可以看出正常情况下`Consumer`函数中只会对通道进行接收操作，但是这不代表不可以在`Consumer`函数中对通道进行发送操作。作为`Producer`函数的提供者，我们在返回通道的时候可能只希望调用方拿到返回的通道后只能对其进行接收操作。但是我们没有办法阻止在`Consumer`函数中对通道进行发送操作。

Go语言中提供了**单向通道**来处理这种需要限制通道只能进行某种操作的情况。

```go
<- chan int // 只接收通道，只能接收不能发送
chan <- int // 只发送通道，只能发送不能接收
```

其中，箭头`<-`和关键字`chan`的相对位置表明了当前通道允许的操作，这种限制将在编译阶段进行检测。另外对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。

我们使用单向通道将上面的示例代码进行如下改造。

```go
// Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25
}
```

这一次，`Producer`函数返回的是一个只接收通道，这就从代码层面限制了该函数返回的通道只能进行接收操作，保证了数据安全。很多读者看到这个示例可能会觉着这样的限制是多余的，但是试想一下如果`Producer`函数可以在其他地方被其他人调用，你该如何限制他人不对该通道执行发送操作呢？并且返回限制操作的单向通道也会让代码语义更清晰、更易读。

在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。

```go
var ch3 = make(chan int, 1)
ch3 <- 10
close(ch3)
Consumer2(ch3) // 函数传参时将ch3转为单向通道

var ch4 = make(chan int, 1)
ch4 <- 10
var ch5 <-chan int // 声明一个只接收通道ch5
ch5 = ch4          // 变量赋值时将ch4转为单向通道
<-ch5
```

另一个例子：
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

### 总结

下面的表格中总结了对不同状态下的通道执行相应操作的结果。

![img](/file/img/go_channel_state.png)

图里差2个：
**注意：**对已经关闭的通道再执行 close 也会引发 panic。往关闭的管道写数据会 panic。

### channel控制并发数量 （gotoutine数量）
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, sem chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- true // 请求访问
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
	<-sem // 释放访问
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan bool, 3) // 信号量，同时只允许3个协程访问资源
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go doWork(i, sem, &wg)
	}
	wg.Wait()
}
//- 创建了一个带缓冲的通道 sem 作为信号量。它的容量是 3，这意味着同时只允许 3 个协程访问资源。
//- 此代码展示了如何使用 Go 的通道来实现信号量模式，从而限制并发访问某个资源。
```

### 通道误用示例
#### 示例1

```go
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

将上述代码编译执行后，匿名函数所在的 goroutine 并不会按照预期在通道被关闭后退出。因为`task := <- ch`的接收操作在通道被关闭后会一直接收到零值，而不会退出。此处的接收操作应该使用`task, ok := <- ch`，通过判断布尔值`ok`为假时退出；或者使用select 来处理通道。也可以用for range读取。

```go
func demo012() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for true {
				task, ok := <-ch
				if !ok {
					break
				}
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func demo011() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for task := range ch {
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

#### 示例2
```go
func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}
```

上述代码片段可能导致 goroutine 泄露（goroutine 并未按预期退出并销毁）。由于 select 命中了超时逻辑，导致通道没有消费者（无接收操作），而其定义的通道为无缓冲通道，因此 goroutine 中的`ch <- "job result"`操作会一直阻塞，最终导致 goroutine 泄露。

## Channel 原理

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	timer    *timer // timer feeding this chan
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```

对于有缓冲的channel存储数据，借助的是循环数组的结构

- `qcount` 已经接收但还未被取走的元素个数 内置函数len获取到
- `datasiz` 循环队列的大小 暂时认为是cap容量的值
- `buf` 指向底层循环数组的指针, 无缓冲通道中 buf的值为nil
- `elemsize` 声明chan时到元素类型和大小 固定, 能够收发元素的大小
- `closed` channel是否关闭的标志
- `elemtype` channel中的元素类型

有缓冲channel内的缓冲数组会被作为一个“环型”来使用。当下标超过数组容量后会回到第一个位置，所以需要有两个字段记录当前读和写的下标位置
- `sendx` 下一次发送数据的下标位置,处理发送进来数据的指针在buf中的位置 接收到数据 指针会加上elemsize，移向下一个位置
- `recvx` 下一次读取数据的下标位置,处理接收请求（发送出去）的指针在buf中的位置

当循环数组中没有数据时，收到了接收请求，那么接收数据的变量地址将会写入读等待队列，当循环数组中数据已满时，收到了发送请求，那么发送数据的变量地址将写入写等待队列
- `recvq` 读等待队列, 如果没有数据可读而阻塞， 会加入到recvq队列中
- `sendq` 写等待队列, 向一个满了的buf 发送数据而阻塞，会加入到sendq队列中

- `lock` 互斥锁，保证读写channel时不存在并发竞争问题

### 发送流程：

1. 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程；
2. 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
3. 如果缓冲区中没有空余位置，或没有缓冲区，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；

简单流程图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661788117541-f82a3d7e-8b22-46cd-9bd9-dde26f0d290c.png)

### 接收流程：

1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；
2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；
3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；
4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；

简单流程图如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661788153163-c386fedf-84b2-42ed-9965-d5d80743650c.png)


### 总结hchan结构体的主要组成部分有四个：

![img](https://cdn.nlark.com/yuque/0/2022/webp/22219483/1661787750459-2608e3a8-f5f9-4d1c-a97f-314d4d83fecf.webp)

- 用来保存goroutine之间传递数据的循环链表。=====> buf。
- 用来记录此循环链表当前发送或接收数据的下标值。=====> sendx和recvx。
- 用于保存向该chan发送和从改chan接收数据的goroutine的队列。=====> sendq 和 recvq
- 保证channel写入和读取数据时线程安全的锁。 =====> lock

### 关闭channel

关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic。

**使用场景：** 消息传递、消息过滤，信号广播，事件订阅与广播，请求、响应转发，任务分发，结果汇总，并发控制，限流，同步与异步

## sync.WaitGroup
- `WaitGroup` 是Go语言`sync` 包中比较常见的同步机制,它可以用于等待一系列的Goroutine的返回。
- 通过`WaitGroup` 我们可以在多个Goroutine之间非常轻松的同步信息,原本顺序执行的代码也可以在多个Goroutine中并发执行,加快了程序处理的速度。

多个goroutine同时工作是，为了知道最后一个goroutine什么时候结束(最后一个结束并不一定是最后一个开始)，我们需要一个递增的计数器，在每一个goroutine启动时加一，在goroutine退出时减一。这个计数器需要在多个goroutine操作时做到安全并且提供提供在其减为零之前一直等待的一种方法。这种计数类型被称为`sync.WaitGroup`。

有以下几个方法：

|                方法名                 |        功能         |
| :----------------------------------: | :-----------------: |
| func (wg * WaitGroup) Add(delta int) |    计数器+delta     |
|        (wg *WaitGroup) Done()        |      计数器-1       |
|        (wg *WaitGroup) Wait()        | 阻塞直到计数器变为0  |

下面的代码就用到了这种方法：
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
注意`Add`和`Done`方法的不对称。`Add是为计数器加一，必须在worker goroutine开始之前调用，而不是在goroutine中`；否则的话我们没办法确定Add是在"closer" goroutine调用Wait之前被调用。并且Add还有一个参数，但Done却没有任何参数；其实它和Add(-1)是等价的，`Done` 只是调用了`wg.Add(-1)`。我们使用defer来确保计数器即使是在出错的情况下依然能够正确地被减掉。上面的程序代码结构是当我们使用并发循环，但又不知道迭代次数时很通常而且很地道的写法。

多生产者channel关闭示例：
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	workN := 5 // 生产者数
	var wg sync.WaitGroup
	wg.Add(workN)
	for i := 0; i < workN; i++ {
		go func(i int) {
			n := i * i
			ch <- n
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
```

### 源码

```go
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}
```

在 Go 语言中，`noCopy` 是一种用于标记某个结构体不应该被值拷贝的机制。Go 的工具链中有一个工具叫做 govet，它可以检测出可能违反 noCopy 约定的情况。当一个结构体嵌入了 noCopy 字段后，如果尝试对该结构体执行值拷贝，govet 或其他静态分析工具会发出警告，因为复制包含 noCopy 的结构体会导致不可预料的行为。noCopy 并不会阻止 Go 编译器或运行时执行拷贝操作，它只是一个约定，依赖于开发者遵循这一约定并利用工具进行检查。

- `noCopy` 的主要作用就是保证`WaitGroup` 不会被开发者通过再赋值的方式进行拷贝,进而导致一些诡异的行为。
- `copyLock` 包就是一个用于检测类似错误的分析器,它的原理就是在 编译期间 检查被拷贝的变量中是否包含的`noCopy` 或者`sync` 关键字
- `state` 是一个64位的无符号整数,其高32位表示计数器,低32位表示等待的Goroutine个数

- `Add`方法的主要作用就是更新`WaitGroup`中持有的计数器`counter`,64位状态的高32位,虽然`Add`方法传入的参数可以为负数,但是一个`WaitGroup`的计数器只能是非负数,当调用`Add`方法导致计数器归零并且还有等待的Goroutine时,就会通过`runtime_Semrelease` 唤醒处于等待状态的所有Goroutine。
- Add方法在等待goroutine为0，如果传入负值，则panic. `sync: negative WaitGroup counter`.

- 另一个`WaitGroup`的方法`Wait`就会在当前计数器中保存的数据大于0 时修改等待Goroutine 的个数`waiter`并调用 `runtime_Semacquire`陷入睡眠状态。
- 陷入睡眠的`Goroutine` 就会等待`Add`方法在计数器为 0 时唤醒。

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

select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等待下去 （可用于阻塞 main 函数，防止退出）。如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。

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
//第一次循环时 i = 0，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
//第二次 for 循环时，i = 1，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值0并赋值给变量 x ，所以会在终端打印出 0；
//后续的 for 循环以此类推会依次打印出2、4、6、8。
```
channel的零值是nil。对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。

## Context
Golang context是Golang应用开发常用的并发控制技术，它与WaitGroup最大的不同点是context对于派生goroutine有更强的控制力，它可以控制多级的goroutine。

context翻译成中文是”上下文”，即它可以控制一组呈树状结构的goroutine，每个goroutine拥有相同的上下文。

典型的使用场景如下图所示：

![goroutine](/file/img/goroutine_tree.png)

上图中由于goroutine派生出子goroutine，而子goroutine又继续派生新的goroutine，这种情况下使用WaitGroup就不太容易，因为子goroutine个数不容易确定。而使用context就可以很容易实现。context的作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。

context包主要提供了两种方式创建context:

- context.Backgroud()
- context.TODO()

这两个函数其实只是互为别名，没有差别，官方给的定义是：

context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来。
context.TODO 应该只在不确定应该使用哪种上下文时使用；
所以在大多数情况下，我们都使用context.Background作为起始的上下文向下传递。

### Context实现原理

context实际上只定义了接口，凡是实现该接口的类都可称为是一种context，官方包中实现了几个常用的context，分别可用于不同的场景。

源码包中`src/context/context.go:Context` 定义了该接口：

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
基础的context接口只定义了4个方法，下面分别简要说明一下：

- Deadline()

该方法返回一个deadline （context.Context 被取消的时间，也就是完成工作的截止日期）和标识是否已设置deadline的bool值，如果没有设置deadline，则ok == false，此时deadline为一个初始值的time.Time值

- Done()

该方法返回一个channel，需要在select-case语句中使用，如`case <-context.Done():`。

当context关闭后，Done()返回一个被关闭的管道，关闭的管道仍然是可读的，据此goroutine可以收到关闭请求；当context还未关闭时，Done()返回nil。

- Err()

该方法描述context关闭的原因。关闭原因由context实现控制，不需要用户设置。比如Deadline context，关闭原因可能是因为deadline，也可能提前被主动关闭，那么关闭原因就会不同:

- 因deadline关闭：“context deadline exceeded”；
- 因主动关闭： “context canceled”。

当context关闭后，Err()返回context的关闭原因；当context还未关闭时，Err()返回nil；

- Value()

有一种context，它不是用于控制呈树状分布的goroutine，而是用于在树状分布的goroutine间传递信息。Value()方法就是用于此种类型的context，该方法根据key值查询map中的value。具体使用后面示例说明。

### 空context
context包中定义了一个空的context， 名为emptyCtx，用于context的根节点，空的context只是简单的实现了Context，本身不包含任何值，仅用于其他context的父节点。

emptyCtx类型定义如下代码所示：

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
}

func (*emptyCtx) Done() <-chan struct{} {
    return nil
}

func (*emptyCtx) Err() error {
    return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
    return nil
}
```

context包中定义了一个公用的emptCtx全局变量，名为background，可以使用context.Background()获取它，实现代码如下所示：

```go
var background = new(emptyCtx)
func Background() Context {
    return background
}
```

context包提供了`4个方法`创建不同类型的context，使用这四个方法时如果没有父context，都需要传入backgroud，即backgroud作为其父节点：

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

context包中实现Context接口的struct，除了emptyCtx外，还有cancelCtx、timerCtx和valueCtx三种struct，正是基于这三种context实例，实现了上述4种类型的context。

context包中各context类型之间的关系，如下图所示：

![go_context](/file/img/go_context.png)

struct cancelCtx、timerCtx、valueCtx都继承于Context，下面分别介绍这三个struct。

### cancelCtx

源码包中`src/context/context.go:cancelCtx` 定义了该类型context：

```go
type cancelCtx struct {
    Context

    mu       sync.Mutex            // protects following fields
    done     chan struct{}         // created lazily, closed by first cancel call
    children map[canceler]struct{} // set to nil by the first cancel call
    err      error                 // set to non-nil by the first cancel call
}
```

children中记录了由此context派生的所有child，此context被cancel时会把其中的所有child都cancel掉。

cancelCtx与deadline和value无关，所以只需要实现Done()和Err()外露接口即可。

#### Done()接口实现

按照Context定义，Done()接口只需要返回一个channel即可，对于cancelCtx来说只需要返回成员变量done即可。

这里直接看下源码，非常简单：

```go
func (c *cancelCtx) Done() <-chan struct{} {
    c.mu.Lock()
    if c.done == nil {
        c.done = make(chan struct{})
    }
    d := c.done
    c.mu.Unlock()
    return d
}
```

由于cancelCtx没有指定初始化函数，所以cancelCtx.done可能还未分配，所以需要考虑初始化。
cancelCtx.done会在context被cancel时关闭，所以cancelCtx.done的值一般经历如下三个阶段：
`nil –> chan struct{} –> closed chan`。

#### Err()接口实现

按照Context定义，Err()只需要返回一个error告知context被关闭的原因。对于cancelCtx来说只需要返回成员变量err即可。

还是直接看下源码：

```go
func (c *cancelCtx) Err() error {
    c.mu.Lock()
    err := c.err
    c.mu.Unlock()
    return err
}
```

cancelCtx.err默认是nil，在context被cancel时指定一个error变量： `var Canceled = errors.New("context canceled")`。

#### cancel()实现

cancel()内部方法是理解cancelCtx的最关键的方法，其作用是关闭自己和其后代，其后代存储在cancelCtx.children的map中，其中key值即后代对象，value值并没有意义，这里使用map只是为了方便查询而已。

cancel方法实现伪代码如下所示：

```go
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
    c.mu.Lock()

    c.err = err                          //设置一个error，说明关闭原因
    close(c.done)                     //将channel关闭，以此通知派生的context

    for child := range c.children {   //遍历所有children，逐个调用cancel方法
        child.cancel(false, err)
    }
    c.children = nil
    c.mu.Unlock()

    if removeFromParent {            //正常情况下，需要将自己从parent删除
        removeChild(c.Context, c)
    }
}
```

实际上，WithCancel()返回的第二个用于cancel context的方法正是此cancel()。

#### WithCancel()方法实现

WithCancel()方法作了三件事：

- 初始化一个cancelCtx实例
- 将cancelCtx实例添加到其父节点的children中(如果父节点也可以被cancel的话)
- 返回cancelCtx实例和cancel()方法

其实现源码如下所示：

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    c := newCancelCtx(parent)
    propagateCancel(parent, &c)   //将自身添加到父节点
    return &c, func() { c.cancel(true, Canceled) }
}
```

这里将自身添加到父节点的过程有必要简单说明一下：

1. 如果父节点也支持cancel，也就是说其父节点肯定有children成员，那么把新context添加到children里即可；
2. 如果父节点不支持cancel，就继续向上查询，直到找到一个支持cancel的节点，把新context添加到children里；
3. 如果所有的父节点均不支持cancel，则启动一个协程等待父节点结束，然后再把当前context结束。

#### 典型使用案例

一个典型的使用cancel context的例子如下所示：

```go
package main

import (
    "fmt"
    "time"
    "context"
)

func HandelRequest(ctx context.Context) {
    go WriteRedis(ctx)
    go WriteDatabase(ctx)
    for {
        select {
        case <-ctx.Done():
            fmt.Println("HandelRequest Done.")
            return
        default:
            fmt.Println("HandelRequest running")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteRedis(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("WriteRedis Done.")
            return
        default:
            fmt.Println("WriteRedis running")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteDatabase(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("WriteDatabase Done.")
            return
        default:
            fmt.Println("WriteDatabase running")
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go HandelRequest(ctx)

    time.Sleep(5 * time.Second)
    fmt.Println("It's time to stop all sub goroutines!")
    cancel()

    //Just for test whether sub goroutines exit or not
    time.Sleep(5 * time.Second)
}
```

上面代码中协程HandelRequest()用于处理某个请求，其又会创建两个协程：WriteRedis()、WriteDatabase()，main协程创建context，并把context在各子协程间传递，main协程在适当的时机可以cancel掉所有子协程。

程序输出如下所示：

```
HandelRequest running
WriteDatabase running
WriteRedis running
HandelRequest running
WriteDatabase running
WriteRedis running
HandelRequest running
WriteDatabase running
WriteRedis running
It's time to stop all sub goroutines!
WriteDatabase Done.
HandelRequest Done.
WriteRedis Done.
```

另一个例子：
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

### timerCtx

源码包中`src/context/context.go:timerCtx` 定义了该类型context：

```go
type timerCtx struct {
    cancelCtx
    timer *time.Timer // Under cancelCtx.mu.

    deadline time.Time
}
```

timerCtx在cancelCtx基础上增加了deadline用于标示自动cancel的最终时间，而timer就是一个触发自动cancel的定时器。

由此，衍生出WithDeadline()和WithTimeout()。实现上这两种类型实现原理一样，只不过使用语境不一样：

- deadline: 指定最后期限，比如context将2018.10.20 00:00:00之时自动结束
- timeout: 指定最长存活时间，比如context将在30s后结束。

对于接口来说，timerCtx在cancelCtx基础上还需要实现Deadline()和cancel()方法，其中cancel()方法是重写的。

#### Deadline()接口实现

Deadline()方法仅仅是返回timerCtx.deadline而矣。而timerCtx.deadline是WithDeadline()或WithTimeout()方法设置的。

#### cancel()实现

cancel()方法基本继承cancelCtx，只需要额外把timer关闭。

timerCtx被关闭后，timerCtx.cancelCtx.err将会存储关闭原因：

- 如果deadline到来之前手动关闭，则关闭原因与cancelCtx显示一致；
- 如果deadline到来时自动关闭，则原因为：”context deadline exceeded”

#### WithDeadline()方法实现

WithDeadline()方法实现步骤如下：

- 初始化一个timerCtx实例
- 将timerCtx实例添加到其父节点的children中(如果父节点也可以被cancel的话)
- 启动定时器，定时器到期后会自动cancel本context
- 返回timerCtx实例和cancel()方法

也就是说，timerCtx类型的context不仅支持手动cancel，也会在定时器到来后自动cancel。

#### WithTimeout()方法实现

WithTimeout()实际调用了WithDeadline，二者实现原理一致。

看代码会非常清晰：

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}
```

#### 典型使用案例

下面例子中使用WithTimeout()获得一个context并在其子协程中传递：

```go
package main

import (
    "fmt"
    "time"
    "context"
)

func HandelRequest(ctx context.Context) {
    go WriteRedis(ctx)
    go WriteDatabase(ctx)
    for {
        select {
        case <-ctx.Done():
            fmt.Println("HandelRequest Done.")
            return
        default:
            fmt.Println("HandelRequest running")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteRedis(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("WriteRedis Done.")
            return
        default:
            fmt.Println("WriteRedis running")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteDatabase(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("WriteDatabase Done.")
            return
        default:
            fmt.Println("WriteDatabase running")
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    go HandelRequest(ctx)

    time.Sleep(10 * time.Second)
}
```

主协程中创建一个5s超时的context，并将其传递给子协程，5s自动关闭context。程序输出如下：

```
HandelRequest running
WriteRedis running
WriteDatabase running
HandelRequest running
WriteRedis running
WriteDatabase running
HandelRequest running
WriteRedis running
WriteDatabase running
HandelRequest Done.
WriteDatabase Done.
WriteRedis Done.
```

另一个例子：
```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 创建带有超时的上下文
	defer cancel()
	for i := 0; i < 5; i++ {
		go worker(ctx, i)
	}
	<-ctx.Done() // 程序会自动在超时或取消时结束 (阻塞)
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

```go
//达到超时时间终止接下来的执行
func main()  {
    HttpHandler()
}

func NewContextWithTimeout() (context.Context,context.CancelFunc) {
    return context.WithTimeout(context.Background(), 3 * time.Second)
}

func HttpHandler() {
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

### valueCtx

源码包中`src/context/context.go:valueCtx` 定义了该类型context：

```go
type valueCtx struct {
    Context
    key, val interface{}
}
```

valueCtx只是在Context基础上增加了一个key-value对，用于在各级协程间传递一些数据。

由于valueCtx既不需要cancel，也不需要deadline，那么只需要实现Value()接口即可。

#### Value（）接口实现

由valueCtx数据结构定义可见，valueCtx.key和valueCtx.val分别代表其key和value值。 实现也很简单：

```go
func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}
```

这里有个细节需要关注一下，即当前context查找不到key时，会向父节点查找，如果查询不到则最终返回interface{}。也就是说，可以通过子context查询到父的value值。

#### WithValue() 方法实现

WithValue()实现也是非常的简单, 伪代码如下：

```go
func WithValue(parent Context, key, val interface{}) Context {
    if key == nil {
        panic("nil key")
    }
    return &valueCtx{parent, key, val}
}
```

#### 典型使用案例

下面示例程序展示valueCtx的用法：

```go
package main

import (
    "fmt"
    "time"
    "context"
)

func HandelRequest(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("HandelRequest Done.")
            return
        default:
            fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx := context.WithValue(context.Background(), "parameter", "1")
    go HandelRequest(ctx)

    time.Sleep(10 * time.Second)
}
```

上例main()中通过WithValue()方法获得一个context，需要指定一个父context、key和value。然后通将该context传递给子协程HandelRequest，子协程可以读取到context的key-value。

注意：本例中子协程无法自动结束，因为context是不支持cancle的，也就是说<-ctx.Done()永远无法返回。如果需要返回，需要在创建context时指定一个可以cancel的context作为父节点，使用父节点的cancel()在适当的时机结束整个context。

### 总结

- Context仅仅是一个接口定义，根据实现的不同，可以衍生出不同的context类型；
- cancelCtx实现了Context接口，通过WithCancel()创建cancelCtx实例；
- timerCtx实现了Context接口，通过WithDeadline()和WithTimeout()创建timerCtx实例；
- valueCtx实现了Context接口，通过WithValue()创建valueCtx实例；
- 三种context实例可互为父节点，从而可以组合成不同的应用形式；

## 并发安全

### 竞争条件
在一个线性(只有一个goroutine的)的程序中，程序的执行顺序只由程序的逻辑来决定。在有两个或更多goroutine的程序中，每一个goroutine内的语句也是按照既定的顺序去执行的，但是一般情况下没法知道分别位于两个goroutine的事件x和y的执行顺序，x是在y之前还是之后还是同时发生是没法判断的。这也说明x和y这两个事件是并发的。

一个函数在线性、在并发的情况下，这个函数可以正确地工作，那么这个函数是并发安全的，并发安全的函数不需要额外的同步工作。可以把这个概念概括为一个特定类型的一些方法和操作函数，如果这个类型是并发安全的话，那么所有它的访问方法和操作就都是并发安全的。

数据竞争：只要有两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。

我们用下面的代码演示一个数据竞争的示例。

```go
package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup // 等待组
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```
每次执行都会输出诸如9537、5865、6527等不同的结果。

在上面的示例代码片中，我们开启了两个 goroutine 分别执行 add 函数，这两个 goroutine 在访问和修改全局的`x`变量时就会存在数据竞争，某个 goroutine 中对全局变量`x`的修改可能会覆盖掉另一个 goroutine 中的操作，所以导致最后的结果与预期不符。

有一种方法，利用channel的阻塞特性，用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量。
```go
package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  chan struct{}
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		m <- struct{}{}
		x = x + 1
		<-m
	}
	wg.Done()
}

func main() {
	m = make(chan struct{}, 1)
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// 另一个例子：
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

### sync.Mutex互斥锁
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。Go 语言中使用`sync`包中提供的`Mutex`类型来实现互斥锁。

`sync.Mutex`提供了两个方法供我们使用。

|          方法名          |    功能    |
| ---------------------- | -------- |
|  func (m *Mutex) Lock()  | 获取互斥锁 |
| func (m *Mutex) Unlock() | 释放互斥锁 |

sync包里的Mutex类型，它的Lock方法能够获取到token(这里叫锁)，并且Unlock方法会释放这个token。
我们在下面的示例代码中使用互斥锁限制每次只有一个 goroutine 才能修改全局变量`x`，从而修复上面代码中的问题。

```go
package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup // 等待组
	m sync.Mutex // 互斥锁
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
//每一次都会得到预期中的结果——10000。
```

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

下面的写法与之前的例子功能一样：
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

每次一个goroutine调用mutex的Lock方法来获取一个互斥锁。如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态。mutex会保护共享变量。

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

#### 源码
```go
type Mutex struct {
	state int32
	sema  uint32
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota

	// 这段注释讲了正常，饥饿模式，看原文
	starvationThresholdNs = 1e6
)
```
- 由`state` 和`sema` 组成, `state` 表示 当前互斥锁的状态, 而`sema` 真正用于控制锁状态的信号量, 这两个加起来只占8字节空间的结构体就表示了Go语言中的互斥锁。
- 互斥锁的状态是用`int32`来表示的,但是锁的状态并不是互斥的,它的最低三位分别表示`mutexLocked`,`mutexWoken`,`mutexStarving`,剩下的位置都用来表示当前有多少个Goroutine等待互斥锁被释放。
- 互斥锁在被创建出来时,所有的状态位的默认值都是`0`,当互斥锁被锁定时,`mutexLocked` 就会被置成`1`,当互斥锁被在正常模式下被唤醒`mutexWoken`就会被置成`1`,`mutexStarving`用于表示当前的互斥锁进入了饥饿状态,最后的几位是在当前互斥锁上等待的Goroutine个数。

#### 加锁

- 互斥锁`Mutex`的加锁是靠`Lock` 方法完成的,当锁的状态是 `0` 时直接将`mutexLocked` 位置成 `1`。
- 当`Lock` 方法被调用时`Mutex`的状态不是 `0` 时就会进入 `lockSlow` 方法尝试通过自旋或者其他方法等待锁的释放并获取互斥锁。
- 自旋其实是在多线程同步的过程中使用的一种机制, 一个goroutine在尝试获取锁失败时继续占用CPU资源，并持续检查锁的状态，直到锁变为可用，当锁可用时，自旋中的goroutine可以直接获取锁并继续执行，而不需要经历调度过程。在多核的CPU上,自旋的优点是避免了Goroutine的切换,如果使用恰当会对性能带来非常大的增益。

- 互斥锁中,只有在普通模式下才可能进入自旋,除了模式的限制之外 `runtime_canSpin` 方法中会判断当前方法是否可以进入自旋,进入自旋的条件非常苛刻:
	- 运行在多CPU的机器上
	- 当前Goroutine 为了获取该锁进入自旋的次数小于四次
	- 当前机器上至少存在一个正在运行的处理器`P`并且处理的运行队列是空的

- 如果互斥锁处于`mutexLocked` 并且在普通模式下工作,就会进入自旋,执行30次`PAUSE`指令消耗CPU时间等待锁的释放

- `runtime_SemacquireMutex` 方法主要作用就是使用互斥锁中的信号量保证资源不会被两个Goroutine获取,从这里我们就能看出`Mutex`其实就是对更底层的信号量进行封装,对外提供更加易用的API,`runtime_SemacquireMutex` 会在方法中不断调用 `goparkunlock`将当前 Goroutine陷入休眠等待信号量可以被获取.
- 一旦当前Goroutine 可以获取信号量,就证明互斥锁已经被解锁,该方法就会立刻返回,`Lock`方法的剩余代码也会继续执行下去了,当前互斥锁处于饥饿模式时,如果该Goroutine是队列中最后的一个Goroutine 或者等待锁的时间小于 `starvationThresho1dNs(1ms)` 当前Goroutine 就会直接获得互斥锁并且从饥饿模式中退出并获得锁.

#### 解锁

- 互斥锁的解锁过程相比之下就非常简单,`Unlock` 方法会直接使用`atomic` 包提供的`AddInt32`减一,如果返回的新状态不等于 `0` 就会进入 `unlockSlow` 方法
- `unlockSlow` 方法首先会对锁的状态进行校验,如果当前互斥锁已经被解锁过了就会直接抛出异常`sync: unlock of unlocked mutex` 中止当前程序,在正常情况下会根据当前互斥锁的状态是正常模式还是饥饿模式进入不同的分支!
- 如果当前互斥锁的状态是饥饿模式就会直接调用`runtime_Semrelease` 方法直接将当前锁交给下一个正在正在尝试获取锁的等待者,等待者会在被唤醒之后设置`mutexLocked`状态,由于此时仍然处于`mutexStarving`,所以新的Goroutine也无法获得锁。
- 在正常模式下,如果当前互斥锁不存在等待者或者最低三位表示的状态都是0, 那么当前方法就不需要唤醒其他Goroutine可以直接返回,当有 Goroutine 正在处于等待状态时,还是会通过`runtime_Semrelease` 唤醒对应的 Goroutine并移交锁的所有权

### sync.RWMutex读写锁
互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用`sync`包中的`RWMutex`类型。

`sync.RWMutex`提供了以下5个方法。

|               方法名                |              功能              |
| :---------------------------------: | :----------------------------: |
|      func (rw *RWMutex) Lock()      |            获取写锁            |
|     func (rw *RWMutex) Unlock()     |            释放写锁            |
|     func (rw *RWMutex) RLock()      |            获取读锁            |
|    func (rw *RWMutex) RUnlock()     |            释放读锁            |
| func (rw *RWMutex) RLocker() Locker | 返回一个实现Locker接口的读写锁 |

读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。

RLocker：这个方法的作用是为读操作返回一个 Locker 接口的对象。它的 Lock 方法会调用 RWMutex 的 RLock 方法，它的 Unlock 方法会调用 RWMutex 的 RUnlock 方法。

下面我们使用代码构造一个读多写少的场景，然后分别使用互斥锁和读写锁查看它们的性能差异。

```go
var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)
}

func main() {
	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 100) // x:10 cost:1.466500951s

	// 使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 100) // x:20 cost:117.207592ms
}
```
从最终的执行结果可以看出，使用读写互斥锁在读多写少的场景下能够极大地提高程序的性能。不过需要注意的是如果一个程序中的读操作和写操作数量级差别不大，那么读写互斥锁的优势就发挥不出来。

#### 源码

读写互斥锁在Go语音中的实现是 `RWMutex`,其中不仅包含一个互斥锁,还持有两个信号量,分别用于写等待读和读等待写。
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
- `readerCount` 存储了当前正在执行的读操作的数量,最后的`readerWait`表示当写操作被阻塞时等待完成读操作个数

#### 读写锁

- 当资源的使用者想要获取读写锁时,就需要通过`Lock` 方法了,在`Lock` 方法中首先调用了读写互斥锁持有的`Mutex`的`Lock`方法保证其他获取读写锁的Goroutine 进入等待状态,随后的`atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders)` 其实是为了堵塞后续的读操作，（此时readerCount为负数，一个 writer goroutine 获得了内部的互斥锁，就会反转 readerCount 字段，把它从原来的正整数 readerCount(>=0) 修改为负数，让这个字段保持两个含义（既保存了 reader 的数量，又表示当前有 writer）。也就是说当readerCount为负数的时候表示当前writer goroutine持有写锁中，reader goroutine会进行阻塞。）。
- 如果当时仍然有其他Goroutine持有读锁,该Goroutine就会调用`runtime_SemacquireMutex`进入休眠状态,等待读锁释放时触发`writerSem`信号量将当前协程唤醒。
- 对资源的读写操作完成之后就会将通过`atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)` 变回正数，并通过for循环出发所有由于获取读锁而陷入等待的 Goroutine （当一个 writer 释放锁的时候，它会再次反转 readerCount 字段。这里的反转方法就是给它增加 rwmutexMaxReaders 这个常数值。）
- 最后,`RWMutex` 会释放持有的互斥锁让其他的协程能够重新获取读写锁

#### 读锁

- 读锁的加锁非常简单,通过`atomic.AddInt32`方法为`readerCount`加一,如果该方法返回了负数说明当前有Goroutine获得了写锁,当前Goroutine就会调用`runtime_SemacquireMutex`陷入休眠等待唤醒。
- 如果没有写操作获取当前互斥锁,当前方法就会在`readerCount`加一后返回,当Goroutine想要释放读锁时会调用`RUnlock`方法,该方法会减少正在读资源的`readerCount`,当前方法如果遇到了返回值小于零的情况,说明有一个正在进行的写操作,在这时就应该通过`rUnlockSlow`方法减少当前写操作等待的读操作数`readerWait`并在所有都被释放之后出发写操作的信号量`writerSem`,`writerSem`被触发之后,尝试获取读写锁的进程就会被唤醒并获得锁。

#### 总结

- `readerSem` - 读写锁释放时，通知由于获取读锁等待的Goroutine
- `writerSem` - 读锁释放时，通知由于获取读写锁等待的Goroutine
- `w`互斥锁 - 保证写操作之间的互斥
- `readerCount` - 统计当前进行读操作的协程数,触发写锁时会将其减少`rwmutexMaxReaders`阻塞后续的读操作
- `readerWait` - 当前读写锁等待的进行读操作的协程数,在出发`Lock`之后的每次`RUnlock`都会将其减一,当它归零时该Goroutine就会获得读写锁
- 当读写锁被释放`Unlock`时，首先通知所有的读操作,然后才会释放持有的互斥锁,这样能够保证读操作不会被连续的写操作`饿死`。

易错场景

- 复制已经使用的读写锁，会把它的状态也给复制过来，原来的锁在释放的时候，并不会修改你复制出来的这个读写锁，这就会导致复制出来的读写锁的状态不对，可能永远无法释放锁
- 重入导致死锁，因为读写锁内部基于互斥锁实现对 writer 的并发访问，而互斥锁本身是有重入问题的，所以，writer 重入调用 Lock 的时候，就会出现死锁的现象, 在 reader 的读操作时调用 writer 的写操作（调用 Lock 方法），那么，这个 reader 和 writer 就会形成互相依赖的死锁状态。
- 释放未加锁的 RWMutex，和互斥锁一样，Lock 和 Unlock 的调用总是成对出现的，RLock 和 RUnlock 的调用也必须成对出现。Lock 和 RLock 多余的调用会导致锁没有被释放，可能会出现死锁，而 Unlock 和 RUnlock 多余的调用会导致 panic

### sync.Once初始化
sync包提供了一个专门的方案来解决一次性初始化的问题:sync.Once。一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了;互斥量用来保护boolean变量和客户端数据结构。

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
//只输出一次 only once.
```

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

### sync.Map
Go 语言中内置的 map 不是并发安全的，请看下面这段示例代码。

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

将上面的代码编译后执行，会报出`fatal error: concurrent map writes`错误。我们不能在多个 goroutine 中并发对内置的 map 进行读写操作，否则会存在数据竞争问题。

像这种场景下就需要为 map 加锁来保证并发的安全性了，Go语言的`sync`包中提供了一个开箱即用的并发安全版 map——`sync.Map`。开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。同时`sync.Map`内置了诸如`Store`、`Load`、`LoadOrStore`、`Delete`、`Range`等操作方法。

|                            方法名                            		 |              功能               |
| :---------------------------------------------------------------: | :-----------------------------: |
| func (m *Map) Store(key, value interface{})          				|   存储key-value数据        |
| func (m *Map) Load(key interface{}) (value interface{}, ok bool)  |   查询key对应的value        |
| func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) |    查询或存储key对应的value     |
| func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool) |          查询并删除key          |
| func (m *Map) Delete(key interface{})             |             删除key             |
| func (m *Map) Range(f func(key, value interface{}) bool)   | 对map中的每个key-value依次调用f |

下面的代码示例演示了并发读写`sync.Map`。

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的map
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

## sync.Cond
Cond 通常应用于等待某个条件的一组 goroutine，等条件变为 true 的时候，其中一个 goroutine 或者所有的 goroutine 都会被唤醒执行。

Go语言在标准库中提供的`Cond` 其实是一个条件变量,通过`Cond`我们可以让一系列的 Goroutine 都在触发某个事件或者条件时才被唤醒,每一个`Cond`结构体都包含一个互斥锁`L`。

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
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := sync.NewCond(&sync.Mutex{})
	// 共享资源初始化完成的标志
	sharedResourceReady := false
	// 添加5个 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 加锁
			c.L.Lock()
			// 如果资源还没有准备好，等待
			c.Wait()
			// 资源准备好了，开始工作
			fmt.Printf("Worker %d: Resource is ready, starting work...\n", id)
			c.L.Unlock()
		}(i)
	}

	// 模拟资源初始化需要的时间
	time.Sleep(2 * time.Second)

	// 当资源准备好后，唤醒所有等待的 goroutines
	//c.L.Lock()
	sharedResourceReady = true
	c.Broadcast()
	//c.L.Unlock()

	// 等待所有 goroutines 完成
	fmt.Println(sharedResourceReady)
	wg.Wait()
	fmt.Println("All workers have finished.")
}
```

### 源码
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
    runtime_notifyListNotifyAll(&c.notify)
}
```
- `Cond` 结构体中包含`noCopy` 和 `copyChecker` 两个字段,前者用于保证`Cond` 不会再编译期间拷贝,后者保证在运行期间发生拷贝会直接`panic`,持有的另一个锁`L`其实是一个接口`Locker`,任意实现`Lock`和`Unlock`方法的结构体都可以作为`NewCond`方法的参数
- 结构体中的最后的变量`notifyList` 其实也就是为了实现`Cond`同步机制,该结构体其实就是一个Goroutine的链表
- `Cond` 对外暴露的`Wait` 方法会将当前Goroutine陷入休眠状态,它会先调用`runtime_notifyListAdd`将等待计数器 +1,然后解锁并调用 `runtime_notifyListWait` 等待其他Goroutine的唤醒
- `notifyListWait` 方法的主要作用就是获取当前的Goroutine并将它追加到 `notifyList`链表的最末端
- 除了将当前Goroutine追加到链表的最末端之外,我们还会调用`goparkunlock`陷入睡眠状态,该函数也是在Go语音切换Goroutine 时经常会使用的方法,它会直接让当前处理器的使用权并等待调度器的唤醒
- `Cond`对外提供的`Signal` 和 `Broadcast` 方法就是用来唤醒调用`Wait`陷入休眠的Goroutine,前者会唤醒队列最前面的Goroutine,后者会唤醒队列中全部的Goroutine
- `notifyListNotifyAll`方法会从链表中取出全部的Goroutine并为他们依次调用 `readyWithTime`, 该方法会通过`goready` 将目标的Goroutine 唤醒
- 虽然它会唤醒全部的Goroutine,但是这里唤醒的顺序其实还是按照加入队列的先后顺序,先加入的会先被`goready`唤醒,后加入的Goroutine可能就需要等待调度器的调度
- `notifyListNotifyOne` 函数就只会从 `sudog` 构成的链表中满足 `sudog.ticket == l.notify`的Goroutine并通过`readyWithTime`唤醒
- 在一般情况下我们会选择在不满足特定条件时调用`Wait`陷入休眠,当某些Goroutine检测到当前满足了唤醒的条件,就可以选择使用`Signal`通过一个或者 `Broadcast`通知全部的Goroutine当前条件已经满足,可以继续完成工作了

### 总结

- `Wait`方法在调用之前一定要使用`L.Lock` 持有该资源,否则会发生`panic`导致程序崩溃。把调用者加入到等待队列时会释放锁，在被唤醒之后还会请求锁。在阻塞休眠期间，调用者是不持有锁的，这样能让其他 goroutine 有机会检查或者更新等待变量，因此在使用Wait方法的时候必须持有锁。
- `Signal`方法唤醒Goroutine都是队列最前面 等待最久的Goroutine
- `Broadcast`虽是广播通知等待的Goroutine,但是真正被唤醒时也是按照一定顺序的

## 原子操作

针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。Go语言中原子操作由内置的标准库`sync/atomic`提供。

### atomic包

五种类型操作：读取操作，写入操作，修改操作，交换操作，比较并交换操作

| 方法                                                         |      解释      |
| :----------------------------------------------------------- | :------------: |
| func LoadInt32(addr *int32) (val int32) func LoadInt64(addr *int64) (val int64) func LoadUint32(addr *uint32) (val uint32) func LoadUint64(addr *uint64) (val uint64) func LoadUintptr(addr *uintptr) (val uintptr) func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) |    读取操作    |
| func StoreInt32(addr *int32, val int32) func StoreInt64(addr *int64, val int64) func StoreUint32(addr *uint32, val uint32) func StoreUint64(addr *uint64, val uint64) func StoreUintptr(addr *uintptr, val uintptr) func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) |    写入操作    |
| func AddInt32(addr *int32, delta int32) (new int32) func AddInt64(addr *int64, delta int64) (new int64) func AddUint32(addr *uint32, delta uint32) (new uint32) func AddUint64(addr *uint64, delta uint64) (new uint64) func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) |    修改操作    |
| func SwapInt32(addr *int32, new int32) (old int32) func SwapInt64(addr *int64, new int64) (old int64) func SwapUint32(addr *uint32, new uint32) (old uint32) func SwapUint64(addr *uint64, new uint64) (old uint64) func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) |    交换操作    |
| func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool) func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) | 比较并交换操作 |

### 示例

我们填写一个示例来比较下互斥锁和原子操作的性能。

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

// CommonCounter  普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

// MutexCounter 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// AtomicCounter 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全				//	0 4.165322ms
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全		//	10000 3.837352ms
	test(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率高	//	10000 3.276938ms
	test(&c3)
}
```

`atomic`包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者 sync 包的函数/类型实现同步更好。

### Compare and Swap
不同的类型有不同的CompareAndSwap操作，如 `atomic.CompareAndSwapInt64` 等。这是一个原子操作，用于在多线程环境中更新共享变量。这个操作允许线程检查一个变量的当前值是否与期望值相等，如果相等，则将变量的值设置为新值；如果不相等，则不做任何改变并返回当前值。整个检查和更新的过程是原子的，意味着它不能被其他线程中断。

```go
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
```
- addr 是指向 int64 变量的指针。
- old 是期望的旧值。
- new 是要设置的新值。
- 返回值 swapped 是一个布尔值，表示操作是否成功（即 addr 指向的值是否确实为 old）。

CompareAndSwap 的主要作用包括：
- 保证原子性：确保读取、比较和修改操作作为一个整体执行，防止其他线程在这些操作之间插入额外的操作。
- 实现无锁编程：可以用来实现某些数据结构和算法，比如无锁队列、栈或其他同步结构，从而避免了使用锁带来的性能开销。
- 避免死锁和优先级反转：由于没有锁的竞争，可以减少死锁的可能性以及因线程调度引起的优先级反转问题。
- 在并发编程中，CompareAndSwap 是构建更复杂同步原语的基础，如自旋锁、信号量和原子引用等。

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

多个 goroutine 对同一个 map 写会 panic，异常是否可以用 defer 捕获？答：可以捕获异常，但是只能捕获一次

下面代码有什么问题？
```go
package main
 
import (
    "fmt"
    "io"
    "net/http"
	"time"
)

func httpget(ch chan int){
    resp, err := http.Get("http://localhost:8000/rest/api/user")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
	ch <- 1
}

func main() {
    start := time.Now()
    // 注意设置缓冲区大小要和开启协程的个人相等
    chs := make([]chan int, 2000)
    for i := 0; i < 2000; i++ {
        chs[i] = make(chan int)
        go httpget(chs[i])
    }
    for _, ch := range chs {
        <- ch
    }
    end := time.Now()
    consume := end.Sub(start).Seconds()
    fmt.Println("程序执行耗时(s)：", consume)
}
// 有泄露风险：
// httpGet中，如果err!= nil, 提前return，导致 chs[i] <- 1 不执行。main函数中 <- ch 会阻塞。
```

下面代码有什么问题?
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func httpget(lock *sync.Mutex) {
	lock.Lock()
	// defer lock.Unlock()
	counter++
	resp, err := http.Get("http://localhost:8000/rest/api/user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	lock.Unlock() // 应该注释
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 800; i++ {
		go httpget(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 800 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
// httpget 如果err!=nil 直接返回，没有释放锁，导致死锁。
```

Go 如何实现原子操作？

答：原子操作就是不可中断的操作，外界是看不到原子操作的中间状态，要么看到原子操作已经完成，要么看到原子操作已经结束。在某个值的原子操作执行的过程中，CPU 绝对不会再去执行其他针对该值的操作，那么其他操作也是原子操作。Go 语言的标准库代码包 sync/atomic 提供了原子的读取（Load 为前缀的函数）或写入（Store 为前缀的函数）某个值。

原子操作与互斥锁的区别？
- 互斥锁是一种数据结构，用来让一个线程执行程序的关键部分，完成互斥的多个操作。
- 原子操作是针对某个值的单个互斥操作。

channel 是否线程安全？锁用在什么地方？ 答：
- Golang的Channel,发送一个数据到Channel 和 从Channel接收一个数据 都是 原子性的。
- 而且Go的设计思想就是:不要通过共享内存来通信，而是通过通信来共享内存，前者就是传统的加锁，后者就是Channel。
- 设计Channel的主要目的就是在多任务间传递数据的，这当然是安全的


1. 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
    1. 开启一个 goroutine 循环生成int64类型的随机数，发送到`jobChan`
    2. 开启24个 goroutine 从`jobChan`中取出随机数计算各位数的和，将结果发送到`resultChan`
    3. 主 goroutine 从`resultChan`取出结果并打印到终端输出


## 一些代码

### 1. context相关
```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result := make(chan int, 0)
	asyncDoStuffWithTimeout(ctx, result)
	fmt.Printf("result get: %v", <-result)
}

func asyncDoStuffWithTimeout(ctx context.Context, result chan int) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx is done, %v\n", ctx.Err())
			result <- 0
			return
		case <-time.After(2 * time.Second):
			fmt.Println("set result")
			result <- 10
		}
	}()
}
// ctx is done, context deadline exceeded
// restult get: 0
```

### 2. 10 个goroutine顺序打印 1-10
```go
package main

import (
	"fmt"
	"sync"
)

//整体思路
//简单讲解下，由于goroutine的执行顺序是没有保证的，所以当需要顺序打印时，需要借助其他的一些全局变量控制打印顺序，我此处是通过一个tag，通过tag这个变量控制数据进入chan的顺序，进而使得输出可以按序输出
//需要注意的一个小点
//通过协程进行写入时，不能直接用变量i，需要通过参数把i传递进去，因为i对于这个协程来说是全局的，而i这个全局变量又是在自增的，所以不能直接用i

//使用channel按顺序输出1-10
func orderchan() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	chs := make(chan int)
	//通过一个全局变量控制进channel的顺序
	tag := 1
	for i := 1; i <= 10; i++ {
		go func(value int) {
			//死循环，保证按顺序进chan
			for {
				if tag == value {
					chs <- value
					break
				}
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(<-chs)
		wg.Done()
		tag++
	}
	wg.Wait()
}

func main() {
	orderchan()
}
```

### 3. 写一个死锁
```go
// 死锁
package main

import (
	"fmt"
	"sync"
)

func f1(in chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in:===", <-in)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	out := make(chan int)
	//out := make(chan int, 1) // solution 2
	//go f1(out, &wg) // solution 1
	out <- 2
	go f1(out, &wg) // solution 2
	wg.Wait()
}

// 这段代码会造成死锁。问题的根源是在 out <- 2 处。在此处，主 goroutine 尝试向 out 通道发送一个整数值，但是没有其他的 goroutine 在接收这个值，导致主 goroutine 被阻塞。
//这是一个常见的 Go 并发模型中的问题。当一个 goroutine 尝试向一个无缓冲的通道发送数据时，它会阻塞，直到另一个 goroutine 从该通道读取数据。但在这段代码中，在发送数据之前没有启动任何接收数据的 goroutine，所以会出现死锁。
// 有两种解决方案：
// 向通道发送数据之前，启动一个 goroutine 来接收数据。
// 将通道的容量设置为 1，这样就可以发送一个数据，然后再阻塞。
```

### 4. 只消费一次
```go
// 只消费一次。
package main

import (
	"fmt"
	"time"
)

var cnt = 0

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
	time.Sleep(time.Second)
	fmt.Println(cnt) // prints 1
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i // the channel will be block due to lack of consumer
		cnt++   // this code will only execute once
	}
}
```

### 5. 用一个channel done 阻塞main
```go
package main

import "fmt"

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch) // 关闭通道，这很重要，否则 main 函数中的 for 循环会永远阻塞等待
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go tel(ch)
	go cus(ch, done)
	<-done
}

func cus(ch chan int, done chan bool) {
	go func() {
		for i := range ch {
			fmt.Printf("The counter is at %d\n", i)
		}
		done <- true
	}()
}

// The counter is at 0
// The counter is at 1
// ...
// The counter is at 13
// The counter is at 14
```

使用select 语句阻塞
```go
package main

import "fmt"

func tel2(ch chan int, done chan struct{}) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
	done <- struct{}{}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go tel2(ch, done)
	// 使用 select 语句来从 ch 和 done 中接收数据
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Printf("The counter is at %d\n", i)
			}
		case <-done:
			fmt.Println("Finished receiving!")
			return
		}
	}
}
```

### 6. select 随机打印1 0
```go
func main() {
	c := make(chan int)
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
// 随机打印1 0
```

### 7. fibonacci
fibonacci的一些写法：
```go
// 1. select
func goFibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	goFibonacciSelect(c, quit)
}

// 2
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// 3
```

### 8. 优先级选择
```go
func priority_select(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
			select {
			case val1 := <-ch1:
				fmt.Println(val1)
			default:
				fmt.Println(val2)
			}
		}
	}
}
```

### 9. 实现sync.WaitGroup的三个功能：Add、Done、Wait

- 定义一个count用来实现Add, Done的加一减一功能。
- `cwg.done`是一个`chan struct{}`类型的通道，当调用`close(cwg.done)`时，会向该通道发送一个零值的结构体，表示通道已经关闭。Wait()方法就是调用`<-cwg.done`这一行代码，它会阻塞等待，直到收到通道关闭的信号。

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CustomWaitGroup struct {
	count int32
	done  chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		count: 0,
		done:  make(chan struct{}),
	}
}

func (cwg *CustomWaitGroup) Add(delta int) {
	atomic.AddInt32(&cwg.count, int32(delta))
}

func (cwg *CustomWaitGroup) Done() {
	if atomic.AddInt32(&cwg.count, -1) == 0 {
		close(cwg.done)
	}
}

func (cwg *CustomWaitGroup) Wait() {
	<-cwg.done
}

func main() {
	cwg := NewCustomWaitGroup()

	for i := 0; i < 3; i++ {
		cwg.Add(1)
		go func(i int) {
			defer cwg.Done()
			fmt.Printf("Task %d started\n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Task %d completed\n", i)
		}(i)
	}

	cwg.Wait()
	fmt.Println("All tasks completed")
}
```

### 10. 一个简单的携程池
```go
// 主函数
package main

import (
	"fmt"
	"time"
)

/* 有关Task任务相关定义及操作 */
//定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error //一个无参的函数类型
}

// 通过NewTask来创建一个Task
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

// 执行Task任务的方法
func (t *Task) Execute() {
	t.f() //调用任务所绑定的函数
}

/* 有关协程池的定义及操作 */
//定义池类型
type Pool struct {
	EntryChannel chan *Task //对外接收Task的入口
	worker_num   int        //协程池最大worker数量,限定Goroutine的个数
	JobsChannel  chan *Task //协程池内部的任务就绪队列
}

// 创建一个协程池
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		worker_num:   cap,
		JobsChannel:  make(chan *Task),
	}
	return &p
}

// 协程池创建一个worker并且开始工作
func (p *Pool) worker(work_ID int) {
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel {
		//如果拿到任务,则执行task任务
		//time.Sleep(1 * time.Second)
		task.Execute()
		fmt.Println("worker ID ", work_ID, " 执行完毕任务")
	}
}

// 让协程池Pool开始工作
func (p *Pool) Run() {
	//1,首先根据协程池的worker数量限定,开启固定数量的Worker,
	//  每一个Worker用一个Goroutine承载
	for i := 0; i < p.worker_num; i++ {
		fmt.Println("开启固定数量的Worker:", i)
		go p.worker(i)
	}

	//2, 从EntryChannel协程池入口取外界传递过来的任务
	//   并且将任务送进JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}

	//3, 执行完毕需要关闭JobsChannel
	close(p.JobsChannel)
	fmt.Println("执行完毕需要关闭JobsChannel")

	// 需要延时一下
}

// 主函数
func main() {
	//创建一个Task
	t := NewTask(func() error {
		fmt.Println("创建一个Task:", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for i := 0; i < 10; i++ {
			p.EntryChannel <- t
		}
		close(p.EntryChannel)
		fmt.Println("执行完毕需要关闭EntryChannel")
	}()

	//启动协程池p
	p.Run()
}
```
这不对，还是得用waitgroup

```go
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	var wg sync.WaitGroup

	// 发送数据到c1
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1) // 关闭c1通道
	}()

	// 创建工作协程
	for i := 0; i < 3; i++ {
		wg.Add(1) // 增加WaitGroup计数器
		go func(i int) {
			defer wg.Done() // 完成时减小计数器
			for in := range c2 {
				fmt.Println("Worker:", i, "接收到外界传递过来的任务:", in)
			}
		}(i)
	}

	// 从c1读取数据并转发到c2
	for n := range c1 {
		c2 <- n
	}
	close(c2) // 关闭c2通道

	// 等待所有工作协程完成
	wg.Wait()
}
```

## 4种Golang并发操作中常见的死锁情形

什么是死锁，在Go的协程里面死锁通常就是永久阻塞了，你拿着我的东西，要我先给你然后再给我，我拿着你的东西又让你先给我，不然就不给你。我俩都这么想，这事就解决不了了。

### 第一种情形：无缓存能力的管道，自己写完自己读
```go
func main() {
    ch := make(chan int, 0)

    ch <- 666
    x := <- ch
    fmt.Println(x)
}
```
一个没有缓存能力的管道，然后往里面写666，然后就去管道里面读。这样肯定会出现问题啊！一个无缓存能力的管道，没有人读，你也写不了，没有人写，你也读不了，这正是一种死锁！

fatal error: all goroutines are asleep - deadlock!
解决办法很简单，开辟两条协程，一条协程写，一条协程读。

### 第二种情形：协程来晚了
```go
func main() {
    ch := make(chan int,0)
    ch <- 666
    go func() {
        <- ch
    }()
}
```
协程开辟在将数字写入到管道之后，因为没有人读，管道就不能写，然后写入管道的操作就一直阻塞。

### 第三种情形：管道读写时，相互要求对方先读/写
如果相互要求对方先读/写，自己再读/写，就会造成死锁。
```go
func main() {
    chHusband := make(chan int,0)
    chWife := make(chan int,0)

    go func() {
        select {
        case <- chHusband:
            chWife<-888
        }
    }()

    select {
        case <- chWife:
            chHusband <- 888
    }
}
```

### 第四种情形：读写锁相互阻塞，形成隐形死锁
```go
func main() {
    var rmw09 sync.RWMutex
    ch := make(chan int,0)

    go func() {
        rmw09.Lock()
        ch <- 123
        rmw09.Unlock()
    }()

    go func() {
        rmw09.RLock()
        x := <- ch
        fmt.Println("读到",x)
        rmw09.RUnlock()
    }()

    for {
        runtime.GC()
    }
}
```
这两条协程，如果第一条协程先抢到了只写锁，另一条协程就不能抢只读锁了，那么另外一条协程没有读，所以第一条协程就写不进。

如果第二条协程先抢到了只读锁，另一条协程就不能抢只写锁了，那么因为另外一条协程没有写，所以第二条协程就读不到。

## TBD

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


`go tool trace trace.out` 可以看调度信息

`GODEBUG=schedtrace=1000 ./trace` 可以debug信息

```bash
➜  gotest GODEBUG=schedtrace=1000 go run main.go
SCHED 0ms: gomaxprocs=4 idleprocs=2 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 [1 0 0 0]
SCHED 1ms: gomaxprocs=4 idleprocs=2 threads=7 spinningthreads=1 idlethreads=2 runqueue=0 [0 0 0 0]
SCHED 3ms: gomaxprocs=4 idleprocs=2 threads=7 spinningthreads=1 idlethreads=2 runqueue=0 [0 0 0 0]
SCHED 10ms: gomaxprocs=4 idleprocs=1 threads=7 spinningthreads=1 idlethreads=1 runqueue=0 [7 7 0 0]
SCHED 15ms: gomaxprocs=4 idleprocs=0 threads=7 spinningthreads=1 idlethreads=0 runqueue=0 [3 1 3 0]
SCHED 18ms: gomaxprocs=4 idleprocs=0 threads=7 spinningthreads=0 idlethreads=0 runqueue=0 [1 1 3 1]
SCHED 22ms: gomaxprocs=4 idleprocs=0 threads=8 spinningthreads=0 idlethreads=0 runqueue=1 [1 1 3 1]
SCHED 24ms: gomaxprocs=4 idleprocs=0 threads=8 spinningthreads=0 idlethreads=1 runqueue=0 [0 0 0 0]
SCHED 25ms: gomaxprocs=4 idleprocs=0 threads=8 spinningthreads=0 idlethreads=1 runqueue=0 [0 0 0 0]
SCHED 29ms: gomaxprocs=4 idleprocs=0 threads=8 spinningthreads=1 idlethreads=0 runqueue=0 [0 6 0 0]
SCHED 30ms: gomaxprocs=4 idleprocs=1 threads=8 spinningthreads=1 idlethreads=1 runqueue=0 [0 0 0 0]
```

我们这时设置的是 schedtrace=1000,即1秒输出一次，

- SCHED：每一行都表示一次调度器的调试信息，后面提示的毫秒数表示启动到现在的运行时间，输出的时间间隔受 schedtrace 的值影响。
- gomaxprocs：当前的 CPU 核心数（GOMAXPROCS 的当前值）。
- idleprocs：空闲的处理器数量，后面的数字表示当前的空闲数量。
- threads：OS 线程数量，后面的数字表示当前正在运行的线程数量。
- spinningthreads：自旋状态的 OS 线程数量。
- idlethreads：空闲的线程数量。
- runqueue：全局队列中 Goroutine 数量，而后面的 [3 1 3 0] 则分别代表这 4 个 P 的本地队列正在运行的 Goroutine 数量。
如果想打印更详情的信息，可能可以与 scheddetail 一起使用，如GODEBUG=schedtrace=1000,scheddetail=1 go run main.go ，会打印出每个G/P/M之间的切换状态详情。

推荐使用pprof查看，目前这种方式不是太友好。