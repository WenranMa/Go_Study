# Go GC 与 内存管理

## GC 设计原理
Go 语言使用自动的内存管理系统。

相信很多人对垃圾收集器的印象都是暂停程序（Stop the world，STW），随着用户程序申请越来越多的内存，系统中的垃圾也逐渐增多；当程序的内存占用达到一定阈值时，整个应用程序就会全部暂停，垃圾收集器会扫描已经分配的所有对象并回收不再使用的内存空间，当这个过程结束后，用户程序才可以继续执行，Go 语言在早期也使用这种策略实现垃圾收集，但是今天的实现已经复杂了很多。

### mutator-allocator-collector

用户程序（Mutator）会通过内存分配器（Allocator）在堆上申请内存，而垃圾收集器（Collector）负责回收堆上的内存空间，内存分配器和垃圾收集器共同管理着程序中的堆内存空间。

### 标记清除
标记清除（Mark-Sweep）算法是最常见的垃圾收集算法，标记清除收集器是跟踪式垃圾收集器，其执行过程可以分成标记（Mark）和清除（Sweep）两个阶段：

- 标记阶段 — 从根对象出发查找并标记堆中所有存活的对象；
- 清除阶段 — 遍历堆中的全部对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表；

如下图所示，内存空间中包含多个对象，我们从根对象出发依次遍历对象的子对象并将从根节点可达的对象都标记成存活状态，即 A、C 和 D 三个对象，剩余的 B、E 和 F 三个对象因为从根节点不可达，所以会被当做垃圾：

![mark-sweep-mark-phase](img/001_mark-sweep-mark-phase.png)
标记清除的标记阶段

标记阶段结束后会进入清除阶段，在该阶段中收集器会依次遍历堆中的所有对象，释放其中没有被标记的 B、E 和 F 三个对象并将新的空闲内存空间以链表的结构串联起来，方便内存分配器的使用。

![mark-sweep-mark-phase](img/002_mark-sweep-sweep-phase.png)
标记清除的清除阶段

这里介绍的是最传统的标记清除算法，垃圾收集器从垃圾收集的根对象出发，递归遍历这些对象指向的子对象并将所有可达的对象标记成存活；标记阶段结束后，垃圾收集器会依次遍历堆中的对象并清除其中的垃圾，整个过程需要标记对象的存活状态，用户程序在垃圾收集的过程中也不能执行，我们需要用到更复杂的机制来解决 STW 的问题。

### 三色抽象

我们首先看一张图，大概就会对 三色标记法 有一个大致的了解：
![tricolor](img/tricolor_01.webp)

白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收；
黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象；
灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

原理：

- 首先把所有的对象都放到白色的集合中
- 从根节点开始遍历对象，遍历到的白色对象从白色集合中放到灰色集合中
- 遍历灰色集合中的对象，把灰色对象引用的白色集合的对象放入到灰色集合中，同时把遍历过的灰色集合中的对象放到黑色的集合中
- 循环步骤3，直到灰色集合中没有对象
- 步骤4结束后，白色集合中的对象就是不可达对象，也就是垃圾，进行回收


### 屏障技术
内存屏障技术是一种屏障指令，它可以让 CPU 或者编译器在执行内存相关操作时遵循特定的约束，目前多数的现代处理器都会乱序执行指令以最大化性能，但是该技术能够保证内存操作的顺序性，在内存屏障前执行的操作一定会先于内存屏障后执行的操作。

想要在并发或者增量的标记算法中保证正确性，我们需要达成以下两种三色不变性（Tri-color invariant）中的一种：

强三色不变性 — 黑色对象不会指向白色对象，只会指向灰色对象或者黑色对象；
弱三色不变性 — 黑色对象指向的白色对象必须包含一条从灰色对象经由多个白色对象的可达路径；

我们在这里想要介绍的是 Go 语言中使用的两种写屏障技术，分别是 Dijkstra 提出的插入写屏障8和 Yuasa 提出的删除写屏障9，这里会分析它们如何保证三色不变性和垃圾收集器的正确性。

#### 插入写屏障
Dijkstra 在 1978 年提出了插入写屏障，通过如下所示的写屏障，用户程序和垃圾收集器可以在交替工作的情况下保证程序执行的正确性：

writePointer(slot, ptr):
    shade(ptr)
    *slot = ptr

上述插入写屏障的伪代码非常好理解，每当执行类似 *slot = ptr 的表达式时，我们会执行上述写屏障通过 shade 函数尝试改变指针的颜色。如果 ptr 指针是白色的，那么该函数会将该对象设置成灰色，其他情况则保持不变。


#### 删除写屏障
Yuasa 在 1990 年的论文 Real-time garbage collection on general-purpose machines 中提出了删除写屏障，因为一旦该写屏障开始工作，它会保证开启写屏障时堆上所有对象的可达，所以也被称作快照垃圾收集（Snapshot GC）10：

This guarantees that no objects will become unreachable to the garbage collector traversal all objects which are live at the beginning of garbage collection will be reached even if the pointers to them are overwritten.

该算法会使用如下所示的写屏障保证增量或者并发执行垃圾收集时程序的正确性：

writePointer(slot, ptr)
    shade(*slot)
    *slot = ptr
Go
上述代码会在老对象的引用被删除时，将白色的老对象涂成灰色，这样删除写屏障就可以保证弱三色不变性，老对象引用的下游对象一定可以被灰色对象引用。

#### 混合写屏障
在 Go 语言 v1.7 版本之前，运行时会使用 Dijkstra 插入写屏障保证强三色不变性，但是运行时并没有在所有的垃圾收集根对象上开启插入写屏障。因为应用程序可能包含成百上千的 Goroutine，而垃圾收集的根对象一般包括全局变量和栈对象，如果运行时需要在几百个 Goroutine 的栈上都开启写屏障，会带来巨大的额外开销，所以 Go 团队在实现上选择了在标记阶段完成时暂停程序、将所有栈对象标记为灰色并重新扫描，在活跃 Goroutine 非常多的程序中，重新扫描的过程需要占用 10 ~ 100ms 的时间。

Go 语言在 v1.8 组合 Dijkstra 插入写屏障和 Yuasa 删除写屏障构成了如下所示的混合写屏障，该写屏障会将被覆盖的对象标记成灰色并在当前栈没有扫描时将新对象也标记成灰色：

writePointer(slot, ptr):
    shade(*slot)
    if current stack is grey:
        shade(ptr)
    *slot = ptr
Go
为了移除栈的重扫描过程，除了引入混合写屏障之外，在垃圾收集的标记阶段，我们还需要将创建的所有新对象都标记成黑色，防止新分配的栈内存和堆内存中的对象被错误地回收，因为栈内存在标记阶段最终都会变为黑色，所以不再需要重新扫描栈空间。

### 回收流程
GO的GC是并行GC, 也就是GC的大部分处理和普通的go代码是同时运行的, 这让GO的GC流程比较复杂.
首先GC有四个阶段, 它们分别是:

- Sweep Termination: 对未清扫的span进行清扫, 只有上一轮的GC的清扫工作完成才可以开始新一轮的GC
- Mark: 扫描所有根对象, 和根对象可以到达的所有对象, 标记它们不被回收
- Mark Termination: 完成标记工作, 重新扫描部分根对象(要求STW)
- Sweep: 按标记结果清扫span

### 垃圾收集的触发

除了使用后台运行的系统监控器和强制垃圾收集助手触发垃圾收集之外，另外两个方法会从任意处理器上触发垃圾收集，这种不需要中心组件协调的方式是在 v1.6 版本中引入的，接下来我们将展开介绍这三种不同的触发时机。

- 后台触发

运行时会在应用程序启动时在后台开启一个用于强制触发垃圾收集的 Goroutine，该 Goroutine 的职责非常简单 — 调用 runtime.gcStart 尝试启动新一轮的垃圾收集：

- 手动触发

用户程序会通过 runtime.GC 函数在程序运行期间主动通知运行时执行，该方法在调用时会阻塞调用方直到当前垃圾收集循环完成，在垃圾收集期间也可能会通过 STW 暂停整个程序：

手动触发垃圾收集的过程不是特别常见，一般只会在运行时的测试代码中才会出现，不过如果我们认为触发主动垃圾收集是有必要的，我们也可以直接调用该方法，但是作者并不认为这是一种推荐的做法。

- 申请内存

最后一个可能会触发垃圾收集的就是 runtime.mallocgc 了，我们在上一节内存分配器中曾经介绍过运行时会将堆上的对象按大小分成微对象、小对象和大对象三类，这三类对象的创建都可能会触发新的垃圾收集循环：


### 练习

下面代码中的指针 p 为野指针，因为返回的栈内存在函数结束时会被释放？

```go
type TimesMatcher struct {
    base int
}

func NewTimesMatcher(base int) *TimesMatcher  {
    return &TimesMatcher{base:base}
}

func main() {
    p := NewTimesMatcher(3)
    fmt.Println(p)
}
// false
// Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），因此在 Go 语言中返回函数参数或临时变量是安全的。
```

GC演变：

Go 的 GC 回收有三次演进过程，
- GoV1.3 之前普通标记清除（mark and sweep）方法，整体过程需要启动 STW，效率极低。
- GoV1.5 三色标记法+写屏障，需要重新扫描一次栈(需要 STW)，效率普通。
- GoV1.8 三色标记法+混合写屏障机制：整个过程不要 STW，效率高。

参考

https://www.topgoer.cn/docs/gozhuanjia/chapter044.2-garbage_collection

https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee076rjgk7


## 内存

### 栈内存
内存管理中面临的最大两个问题：

- 内存的分配
- 内存的回收

有没有简单、高效、且通用的办法统一解决这个内存分配问题呢？

最简单、高效地分配和回收方式就是对一段连续内存的`线性分配`，`栈内存`的分配就采用了这种方式。

通过利用「栈内存」，CPU在执行指令过程中可以高效的存储临时变量。其次：

栈内存的分配过程：看起来像不像数据结构栈的入栈过程。
栈内存的释放过程：看起来像不像数据结构栈的出栈过程。

所以`栈内存`是计算机对连续内存的采取的`线性分配`管理方式，便于高效存储指令运行过程中的临时变量。

### 堆内存
为什么需要堆内存？

假如函数A内变量是个指针且被函数B外的代码依赖，如果对应变量内存被回收，这个指针就成了野指针不安全。怎么解决这个问题呢？

答：这就是`堆内存`存在的意义，Go语言会在代码编译期间通过`逃逸分析`把分配在`栈`上的变量分配到`堆`上去。

堆内存通过「垃圾回收器」回收。

### Go 内存分配

为了方便自主管理内存，做法便是先向系统申请一块内存，然后将内存切割成小块，通过一定的内存分配算法管理内存。

预申请的内存划分为`spans、bitmap、arena`三部分。其中`arena即为所谓的堆区`，应用中需要的内存从这里分配。其中spans和bitmap是为了管理arena区而存在的。

arena的大小为512G，为了方便管理把arena区域划分成一个个的page，每个page为8KB,一共有512GB/8KB个页；

spans区域存放span的指针，每个指针对应一个或多个page，所以span区域的大小为(512GB/8KB)*指针大小8byte = 512M

bitmap区域大小也是通过arena计算出来，不过主要用于GC。

#### span (mspan)
span是用于管理arena页的关键数据结构，每个span中包含1个或多个连续页，为了满足小对象分配，span中的一页会划分更小的粒度，而对于大对象比如超过页大小，则通过多页实现。

根据对象大小，划分了一系列class，每个class都代表一个固定大小的对象，以及每个span的大小。

#### cache (mcache) 
有了管理内存的基本单位span，还要有个数据结构来管理span，这个数据结构叫`mcentral`，各线程需要内存时从mcentral管理的span中申请内存，为了避免多线程申请内存时不断地加锁，Golang为每个线程分配了span的缓存，这个缓存即是`cache (mcache)`。

mcache在初始化时是没有任何span的，在使用过程中会动态地从central中获取并缓存下来，根据使用情况，每种class的span个数也不相同。上图所示，class 0的span数比class1的要多，说明本线程中分配的小对象要多一些。

#### central (mcentral)
cache作为线程的私有资源为单个线程服务，而central则是全局资源，为多个线程服务，当某个线程内存不足时会向central申请，当某个线程释放内存时又会回收进central。

#### heap (mheap)
从mcentral数据结构可见，每个mcentral对象只管理特定的class规格的span。事实上每种class都会对应一个mcentral,这个mcentral的集合存放于mheap数据结构中。

系统预分配的内存分为spans、bitmap、arean三个区域，通过mheap管理起来。

总结：
- arena区域按页划分成一个个小块，span管理一个或多个页
- mcentral管理多个span供线程申请使用
- mcache作为线程私有资源，资源来源于mcentral

#### 内存分配优先级
栈内存分配：

- 小于32KB的栈内存
    - 来源优先级1：线程缓存mcache
    - 来源优先级2：全局缓存stackpool
    - 来源优先级3：逻辑处理器结构p.pagecache
    - 来源优先级4：堆mheap
- 大于等于32KB的栈内存
    - 来源优先级1：全局缓存stackLarge
    - 来源优先级2：逻辑处理器结构p.pagecache
    - 来源优先级3：堆mheap

堆内存分配

- 微对象 0 < Micro Object < 16B
    - 来源优先级1：线程缓存mcache.tiny
    - 来源优先级2：线程缓存mcache.alloc
- 小对象 16B =< Small Object <= 32KB
    - 来源优先级1：线程缓存mcache.alloc
    - 来源优先级2：中央缓存mcentral
    - 来源优先级3：逻辑处理器结构p.pagecache
    - 来源优先级4：堆mheap
- 大对象 32KB < Large Object
    - 来源优先级1：逻辑处理器结构p.pagecache
    - 来源优先级2：堆mheap

- 栈内存也来源于堆mheap

### 逃逸分析
本该分配到栈上的变量，跑到了堆上，这就导致了内存逃逸。栈是高地址到低地址，栈上的变量，函数结束后变量会跟着回收掉，不会有额外性能的开销。变量从栈逃逸到堆上，如果要回收掉，需要进行 gc，那么 gc 一定会带来额外的性能开销。

内存逃逸的情况如下：
- 方法内返回局部变量指针。
- 向 channel 发送指针数据。
- 在闭包中引用包外的值。
- 在 slice 或 map 中存储指针。
- 切片（扩容后）长度太大。
- interface 类型上调用方法。

```go
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}
//Fibonacci()函数中原本属于局部变量的a和b由于闭包的引用，不得不将二者放到堆上，以致产生逃逸
/*
$ go build -gcflags "-m -l" main.go
# command-line-arguments
./main.go:9:2: moved to heap: a
./main.go:9:5: moved to heap: b
./main.go:10:9: func literal escapes to heap 
./main.go:19:13: ... argument does not escape
./main.go:19:34: f() escapes to heap
*/
```

```go
func main() {
	s := "Escape"
	fmt.Println(s)	
    // 很多函数参数为interface类型，比如fmt.Println(a …interface{})，编译期间很难确定其参数的具体类型，也会产生逃逸。
}
/*
$ go build -gcflags "-m -l" main.go
# command-line-arguments
./main.go:7:13: ... argument does not escape
./main.go:7:14: s escapes to heap
*/
```

```go
type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆
	s.Name = name
	s.Age = age
	return s
}

func main() {
	fmt.Println(StudentRegister("Jim", 18))
}
/*
$ go build -gcflags "-m -l" main.go
# command-line-arguments
./main.go:10:22: leaking param: name
./main.go:11:10: new(Student) escapes to heap
./main.go:18:13: ... argument does not escape
*/
```

### 内存泄露
内存泄露有下面一些情况：
- 如果 goroutine 在执行时被阻塞而无法退出，就会导致 goroutine 的内存泄漏，一个 goroutine 的最低栈大小为 2KB，在高并发的场景下，对内存的消耗也是非常恐怖的。
- 互斥锁未释放或者造成死锁会造成内存泄漏
- time.Ticker 是每隔指定的时间就会向通道内写数据。作为循环触发器，必须调用 stop 方法才会停止，从而被 GC 掉，否则会一直占用内存空间。
- 字符串的截取引发临时性的内存泄漏
```go
func main() {
    var str0 = "12345678901234567890"
    str1 := str0[:10]
}
```

- 切片截取引起子切片内存泄漏
```go
func main() {
    var s0 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    s1 := s0[:3]
}
```

- 函数数组传参引发内存泄漏，如果我们在函数传参的时候用到了数组传参，且这个数组够大（我们假设数组大小为 100 万，64 位机上消耗的内存约为 800w 字节，即 8MB 内存），或者该函数短时间内被调用 N 次，那么可想而知，会消耗大量内存，对性能产生极大的影响，如果短时间内分配大量内存，而又来不及 GC，那么就会产生临时性的内存泄漏，对于高并发场景相当可怕。

排查方式：

一般通过 pprof 是 Go 的性能分析工具，在程序运行过程中，可以记录程序的运行信息，可以是 CPU 使用情况、内存使用情况、goroutine 运行情况等，当需要性能调优或者定位 Bug 时候，这些记录的信息是相当重要。

### 练习

Channel 分配在栈上还是堆上？那些在栈，哪些在堆？

Channel 被设计用来实现协程间通信的组件，其作用域和生命周期不可能仅限于某个函数内部，所以 golang 直接将其分配在堆上。

如果编译器不能确保变量在函数 return 之后不再被引用，编译器就会将变量分配到堆上。而且，如果一个局部变量非常大，那么它也应该被分配到堆上而不是栈上。如果一个变量被取地址，那么它就有可能被分配到堆上。如果函数 return 之后，变量不再被引用，则将其分配到栈上。

介绍一下大对象小对象，为什么小对象多了会造成 gc 压力？

小于等于 32k 的对象就是小对象，其它都是大对象。一般小对象通过 mspan 分配内存；大对象则直接由 mheap 分配内存。通常小对象过多会导致 GC 三色法消耗过多的 CPU。优化思路是，减少对象分配。

小对象：如果申请小对象时，发现当前内存空间不存在空闲跨度时，将会需要调用 nextFree 方法获取新的可用的对象，可能会触发 GC 行为。

大对象：如果申请大于 32k 以上的大对象时，可能会触发 GC 行为。