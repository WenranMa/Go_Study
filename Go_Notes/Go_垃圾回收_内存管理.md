# Go GC 与 内存管理

## Go 垃圾回收
Golang中的垃圾回收主要应用三色标记法，GC过程和其他用户goroutine可并发运行，但需要一定时间的STW(stop the world)，STW的过程中，CPU不执行用户代码，全部用于垃圾回收，这个过程的影响很大，Golang进行了多次的迭代优化来解决这个问题。

### Go V1.3之前的标记-清除(mark and sweep)算法
此算法主要有两个主要的步骤：
- 标记(Mark phase)
- 清除(Sweep phase)

第一步，暂停程序业务逻辑, 找出不可达的对象，然后做上标记。第二步，回收标记好的对象。

操作非常简单，但是有一点需要额外注意：mark and sweep算法在执行的时候，需要程序暂停！即 STW(stop the world)。也就是说，这段时间程序会卡在哪儿。

![step_1](/file/img/go_gc_mark_and_sweap_001.png)

第二步, 开始标记，程序找出它所有可达的对象，并做上标记。

![step_2](/file/img/go_gc_mark_and_sweap_002.png)

第三步, 标记完了之后，然后开始清除未标记的对象. 结果如下.

![step_3](/file/img/go_gc_mark_and_sweap_003.png)

第四步, 停止暂停，让程序继续跑。然后循环重复这个过程，直到process程序生命周期结束。

标记-清扫(mark and sweep)的缺点:
- STW，stop the world；让程序暂停，程序出现卡顿 (重要问题)。
- 标记需要扫描整个heap
- 清除数据会产生heap碎片

所以Go V1.3版本之前就是以上来实施的, 流程是

![stw_1](/file/img/go_gc_STW1.png)

Go V1.3 做了简单的优化,将STW提前, 减少STW暂停的时间范围.如下所示

![stw_2](/file/img/go_gc_STW2.png)

这里面最重要的问题就是：mark-and-sweep 算法会暂停整个程序 。

### Go V1.5的三色标记法

我们首先看一张图，大概就会对 三色标记法 有一个大致的了解：
![tricolor](/file/img/tricolor_01.webp)

白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收；
黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象；
灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

原理：

- 首先把所有的对象都放到白色的集合中
- 从根节点开始遍历对象，遍历到的白色对象从白色集合中放到灰色集合中
- 遍历灰色集合中的对象，把灰色对象引用的白色集合的对象放入到灰色集合中，同时把遍历过的灰色集合中的对象放到黑色的集合中
- 循环步骤3，直到灰色集合中没有对象
- 步骤4结束后，白色集合中的对象就是不可达对象，也就是垃圾，进行回收。

我们来看一下具体的过程.

第一步 , 就是只要是新创建的对象,默认的颜色都是标记为“白色”.

![tricolor_01](/file/img/go_gc_tricolor_001.png)

这里面需要注意的是, 所谓“程序”, 则是一些对象的跟节点集合.

![tricolor_02](/file/img/go_gc_tricolor_002.jpeg)

第二步, 每次GC回收开始, 然后从根节点开始遍历所有对象，把遍历到的对象从白色集合放入“灰色”集合。

![tricolor_03](/file/img/go_gc_tricolor_003.jpeg)

第三步, 遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色对象放入黑色集合

![tricolor_04](/file/img/go_gc_tricolor_004.jpeg)

第四步, 重复第三步, 直到灰色中无任何对象.

![tricolor_05](/file/img/go_gc_tricolor_005.jpeg)

![tricolor_06](/file/img/go_gc_tricolor_006.jpeg)

第五步: 回收所有的白色标记表的对象. 也就是回收垃圾.

![tricolor_07](/file/img/go_gc_tricolor_007.jpeg)

以上便是三色并发标记法。

但是如何实现并行的呢？标记-清除(mark and sweep)算法中的卡顿(stw，stop the world)问题如何解决？

### 如果没有STW
我们还是基于上述的三色并发标记法来说, 他是一定要依赖STW的. 因为如果不暂停程序, 程序的逻辑改变对象引用关系, 这种动作如果在标记阶段做了修改，会影响标记结果的正确性。我们举一个场景.

如果三色标记法, 标记过程不使用STW将会发生什么事情？

![tricolor_01](/file/img/go_gc_tricolor_without_STW_01.jpeg)

![tricolor_02](/file/img/go_gc_tricolor_without_STW_02.jpeg)

![tricolor_03](/file/img/go_gc_tricolor_without_STW_03.jpeg)

![tricolor_04](/file/img/go_gc_tricolor_without_STW_04.jpeg)

![tricolor_05](/file/img/go_gc_tricolor_without_STW_05.jpeg)

可以看出，有两个问题, 在三色标记法中,是不希望被发生的

1. 一个白色对象被黑色对象引用(白色被挂在黑色下)
2. 灰色对象与它之间的可达关系的白色对象遭到破坏(灰色同时丢了该白色)

当以上两个条件同时满足时, 就会出现对象丢失现象!

当然, 如果上述中的白色对象3, 如果他还有很多下游对象的话, 也会一并都清理掉.

​为了防止这种现象的发生，最简单的方式就是STW，直接禁止掉其他用户程序对对象引用关系的干扰，但是STW的过程有明显的资源浪费，对所有的用户程序都有很大影响，如何能在保证对象不丢失的情况下合理的尽可能的提高GC效率，减少STW时间呢？

答案就是, 那么我们只要使用一个机制,来破坏上面的两个条件就可以了.

### 屏障机制
我们让GC回收器,满足下面两种情况之一时,可保对象不丢失. 所以引出两种方式.

#### “强-弱” 三色不变式

**强三色不变式**

不存在黑色对象引用到白色对象的指针。

![强](/file/img/go_gc_tricolor_strong.jpeg)

**弱三色不变式**

所有被黑色对象引用的白色对象都处于灰色保护状态.

![弱](/file/img/go_gc_tricolor_weak.jpeg)

为了遵循上述的两个方式,最初有两种屏障方式“插入屏障”, “删除屏障”.

#### 插入屏障
具体操作: 在A对象引用B对象的时候，B对象被标记为灰色。(将B挂在A下游，B必须被标记为灰色)

满足: 强三色不变式. (不存在黑色对象引用白色对象的情况了， 因为白色会强制变成灰色)

伪码如下:
```
添加下游对象(当前下游对象slot, 新下游对象ptr) {   
    标记灰色(新下游对象ptr)   
    当前下游对象slot = 新下游对象ptr                    
}
```
场景：
```
A.添加下游对象(nil, B)   //A 之前没有下游， 新添加一个下游对象B，B被标记为灰色
A.添加下游对象(C, B)     //A 将下游对象C 更换为B，B被标记为灰色
```
​
这段伪码逻辑就是写屏障, 黑色对象的内存槽有两种位置, 栈和堆. 栈空间的特点是容量小,但是要求相应速度快,因为函数调用弹出频繁使用, 所以“插入屏障”机制,在栈空间的对象操作中不使用. 而仅仅使用在堆空间对象的操作中.

​接下来，我们用几张图，来模拟整个一个详细的过程， 希望您能够更可观的看清晰整体流程。

![insert_barrier](/file/img/go_gc_insert_barrier_01.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_02.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_03.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_04.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_05.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_06.jpeg)

但是如果栈不添加,当全部三色标记扫描之后,栈上有可能依然存在白色对象被引用的情况(如上图的对象9). 所以要对栈重新进行三色标记扫描, 但这次为了对象不丢失, 要对本次标记扫描启动STW暂停. 直到栈空间的三色标记结束.

![insert_barrier](/file/img/go_gc_insert_barrier_07.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_08.jpeg)
![insert_barrier](/file/img/go_gc_insert_barrier_09.jpeg)

最后将栈和堆空间 扫描剩余的全部 白色节点清除. 这次STW大约的时间在10~100ms间.

![insert_barrier](/file/img/go_gc_insert_barrier_10.jpeg)

#### 删除屏障
具体操作: 被删除的对象，如果自身为灰色或者白色，那么被标记为灰色。

满足: 弱三色不变式. (保护灰色对象到白色对象的路径不会断)

伪代码：
```
添加下游对象(当前下游对象slot， 新下游对象ptr) {
    if (当前下游对象slot是灰色 || 当前下游对象slot是白色) {
          标记灰色(当前下游对象slot)     //slot为被删除对象， 标记为灰色
    }
    当前下游对象slot = 新下游对象ptr
}
```
场景：
```
A.添加下游对象(B, nil)   //A对象，删除B对象的引用。B被A删除，被标记为灰(如果B之前为白)
A.添加下游对象(B, C)     //A对象，更换下游B变成C。B被A删除，被标记为灰(如果B之前为白)
```
接下来，我们用几张图，来模拟整个一个详细的过程， 希望您能够更可观的看清晰整体流程。

![delete_barrier](/file/img/go_gc_delete_barrier_01.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_02.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_03.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_04.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_05.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_06.jpeg)
![delete_barrier](/file/img/go_gc_delete_barrier_07.jpeg)

这种方式的回收精度低，一个对象即使被删除了最后一个指向它的指针也依旧可以活过这一轮，在下一轮GC中被清理掉。

### Go V1.8的混合写屏障(hybrid write barrier)机制
插入写屏障和删除写屏障的短板：
- 插入写屏障：结束时需要STW来重新扫描栈，标记栈上引用的白色对象的存活；
- 删除写屏障：回收精度低，GC开始时STW扫描堆栈来记录初始快照，这个过程会保护开始时刻的所有存活对象。

Go V1.8版本引入了混合写屏障机制（hybrid write barrier），避免了对栈re-scan的过程，极大的减少了STW的时间。结合了两者的优点。

#### 混合写屏障规则
具体操作:
1. GC开始将栈上的对象全部扫描并标记为黑色(之后不再进行第二次重复扫描，无需STW)，
2. GC期间，任何在栈上创建的新对象，均为黑色。
3. 被删除的对象标记为灰色。
4. 被添加的对象标记为灰色。

满足: 变形的弱三色不变式.

伪代码：
```
添加下游对象(当前下游对象slot, 新下游对象ptr) {
    标记灰色(当前下游对象slot)    //只要当前下游对象被移走，就标记灰色
    标记灰色(新下游对象ptr)
    当前下游对象slot = 新下游对象ptr
}
```
这里我们注意， 屏障技术是不在栈上应用的，因为要保证栈的运行效率。

#### 混合写屏障的具体场景分析
接下来，我们用几张图，来模拟整个一个详细的过程，注意混合写屏障是GC的一种屏障机制，所以只是当程序执行GC的时候，才会触发这种机制。

GC开始：扫描栈区，将可达对象全部标记为黑

![hybrid](/file/img/go_gc_hybrid_001.jpeg)
![hybrid](/file/img/go_gc_hybrid_002.jpeg)

**场景一**： 对象被一个堆对象删除引用，成为栈对象的下游
```
//前提：堆对象4->对象7 = 对象7；  //对象7 被 对象4引用
栈对象1->对象7 = 堆对象7；  //将堆对象7 挂在 栈对象1 下游
堆对象4->对象7 = null；    //对象4 删除引用 对象7
```
![hybrid](/file/img/go_gc_hybrid_011.jpeg)
![hybrid](/file/img/go_gc_hybrid_012.jpeg)

**场景二**： 对象被一个栈对象删除引用，成为另一个栈对象的下游
```
new 栈对象9；
对象9->对象3 = 对象3；      //将栈对象3 挂在 栈对象9 下游
对象2->对象3 = null；      //对象2 删除引用 对象3
```
![hybrid](/file/img/go_gc_hybrid_021.jpeg)
![hybrid](/file/img/go_gc_hybrid_022.jpeg)
![hybrid](/file/img/go_gc_hybrid_023.jpeg)

**场景三**：对象被一个堆对象删除引用，成为另一个堆对象的下游
```
堆对象10->对象7 = 堆对象7；       //将堆对象7 挂在 堆对象10 下游
堆对象4->对象7 = null；         //对象4 删除引用 对象7
```
![hybrid](/file/img/go_gc_hybrid_031.jpeg)
![hybrid](/file/img/go_gc_hybrid_032.jpeg)
![hybrid](/file/img/go_gc_hybrid_033.jpeg)

**场景四**：对象从一个栈对象删除引用，成为另一个堆对象的下游
```
栈对象1->对象2 = null；        //对象1 删除引用 对象2
堆对象4->对象7 = null；       //对象4 删除 对象7
堆对象4->对象2 = 对象2；       //将栈对象2 挂在 堆对象4 下游
```
![hybrid](/file/img/go_gc_hybrid_041.jpeg)
![hybrid](/file/img/go_gc_hybrid_042.jpeg)
![hybrid](/file/img/go_gc_hybrid_043.jpeg)

Golang中的混合写屏障满足弱三色不变式，结合了删除写屏障和插入写屏障的优点，只需要在开始时并发扫描各个goroutine的栈，使其变黑并一直保持，这个过程不需要STW，而标记结束后，因为栈在扫描后始终是黑色的，也无需再进行re-scan操作了，减少了STW的时间。

### 总结：
Go 的 GC 回收有三次演进过程：
- GoV1.3 之前普通标记清除（mark and sweep）方法，整体过程需要启动 STW，效率极低。
- GoV1.5 三色标记法+写屏障，需要重新扫描一次栈(需要 STW)，效率普通。
- GoV1.8 三色标记法+混合写屏障机制：整个过程不要 STW，效率高。

参考

https://www.topgoer.cn/docs/gozhuanjia/chapter044.2-garbage_collection

https://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmee076rjgk7


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

外部函数直接使用子函数的局部变量，再C/C++中不允许的，子函数的foo_val 的声明周期早就销毁了才对，如下面的C/C++代码：
```c
#include <stdio.h>

int *foo(int arg_val) {
    int foo_val = 11;
    return &foo_val;
}

int main()
{
    int *main_val = foo(666);
    printf("%d\n", *main_val);
}

// 编译
// $ gcc pro_1.c 
// pro_1.c: In function ‘foo’:
// pro_1.c:7:12: warning: function returns address of local variable [-Wreturn-local-addr]
//      return &foo_val;
//             ^~~~~~~~
```
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
