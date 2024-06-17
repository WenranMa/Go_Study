# Go 设计题

## 设计消息队列

```go
package mq

import (
	"errors"
	"sync"
	"time"
)

type Broker interface {
	publish(topic string, msg interface{}) error
	subscribe(topic string) (<-chan interface{}, error)
	unsubscribe(topic string, sub <-chan interface{}) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{})
	setConditions(capacity int)
}

type BrokerImpl struct {
	exit         chan bool
	capacity     int
	topics       map[string][]chan interface{} // key： topic  value ： queue
	sync.RWMutex                               // 同步锁
}

func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}

func (b *BrokerImpl) publish(topic string, msg interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}
	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil
	}
	b.broadcast(msg, subscribers)
	return nil
}
func (b *BrokerImpl) broadcast(msg interface{}, subscribers []chan interface{}) {
	count := len(subscribers)
	concurrency := 1
	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}
	pub := func(start int) {
		for j := start; j < count; j += concurrency {
			select {
			case subscribers[j] <- msg:
			case <-time.After(time.Millisecond * 5):
			case <-b.exit:
				return
			}
		}
	}
	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}

func (b *BrokerImpl) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-b.exit:
		return nil, errors.New("broker closed")
	default:
	}
	ch := make(chan interface{}, b.capacity)
	b.Lock()
	b.topics[topic] = append(b.topics[topic], ch)
	b.Unlock()
	return ch, nil
}

func (b *BrokerImpl) unsubscribe(topic string, sub <-chan interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}
	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil
	}
	// delete subscriber
	var newSubs []chan interface{}
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}
	b.Lock()
	b.topics[topic] = newSubs
	b.Unlock()
	return nil
}

func (b *BrokerImpl) setCapacity(capacity int) {
	b.capacity = capacity
}

func (b *BrokerImpl) close() {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		b.topics = make(map[string][]chan interface{})
		b.Unlock()
	}
	return
}

type Client struct {
	bro *BrokerImpl
}

func NewClient() *Client {
	return &Client{
		bro: NewBroker(),
	}
}
func (c *Client) SetCapacity(capacity int) {
	c.bro.setCapacity(capacity)
}
func (c *Client) Publish(topic string, msg interface{}) error {
	return c.bro.publish(topic, msg)
}
func (c *Client) Subscribe(topic string) (<-chan interface{}, error) {
	return c.bro.subscribe(topic)
}
func (c *Client) Unsubscribe(topic string, sub <-chan interface{}) error {
	return c.bro.unsubscribe(topic, sub)
}
func (c *Client) Close() {
	c.bro.close()
}
func (c *Client) GetPayLoad(sub <-chan interface{}) interface{} {
	for val := range sub {
		if val != nil {
			return val
		}
	}
	return nil
}
```

```go
package main

import (
	"fmt"
	"time"

	"Z_Interview/mq"
)

func main() {
	OneTopic()
}

// one topic, two consumer, one publisher
func OneTopic() {
	m := mq.NewClient()
	defer m.Close()
	m.SetCapacity(10)
	ch, err := m.Subscribe("news")
	if err != nil {
		fmt.Println("subscribe failed")
		return
	}
	ch2, _ := m.Subscribe("news")
	go OnePub(m)
	go OneSub(ch, m)
	OneSub(ch2, m)
}

func OnePub(c *mq.Client) {
	// t := time.NewTicker(1 * time.Second)
	// defer t.Stop()
	// for {
	// 	select {
	// 	case <-t.C:
	// 		err := c.Publish("news", "good news")
	// 		if err != nil {
	// 			fmt.Println("pub message failed")
	// 		}
	// 	default:
	// 	}
	// }
	for i := 0; i < 10; i++ {
		err := c.Publish("news", fmt.Sprintf("good news_%d", i))
		if err != nil {
			fmt.Println("pub message failed")
		}
		time.Sleep(1 * time.Second)
	}
	c.Close()
}

func OneSub(m <-chan interface{}, c *mq.Client) {
	for {
		val := c.GetPayLoad(m)
		fmt.Printf("get message is %s\n", val)
	}
}

// 多个topic测试
func ManyTopic() {
	m := mq.NewClient()
	defer m.Close()
	m.SetCapacity(10)
	top := ""
	for i := 0; i < 10; i++ {
		top = fmt.Sprintf("Golang梦工厂_%02d", i)
		go Sub(m, top)
	}
	ManyPub(m)
}
func ManyPub(c *mq.Client) {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			for i := 0; i < 10; i++ {
				//多个topic 推送不同的消息
				top := fmt.Sprintf("Golang梦工厂_%02d", i)
				payload := fmt.Sprintf("asong真帅_%02d", i)
				err := c.Publish(top, payload)
				if err != nil {
					fmt.Println("pub message failed")
				}
			}
		default:
		}
	}
}
func Sub(c *mq.Client, top string) {
	ch, err := c.Subscribe(top)
	if err != nil {
		fmt.Printf("sub top:%s failed\n", top)
	}
	for {
		val := c.GetPayLoad(ch)
		if val != nil {
			fmt.Printf("%s get message is %s\n", top, val)
		}
	}
}
```

Ref:
https://segmentfault.com/a/1190000024518618
https://segmentfault.com/a/1190000043530116



## 设计并实现一个并发安全的Map （sync.Map 源码分析）
提供 Get(key), Set(key, value), 和 Delete(key) 方法，并确保在多 goroutine 环境下操作的线程安全。

### Sync.Map
go1.9之后加入了支持并发安全的Map

sync.Map 的主要思想就是读写分离，空间换时间。

看看 sync.map 优点：

- 空间换时间：通过冗余的两个数据结构(read、dirty)，实现加锁对性能的影响。
- 使用只读数据(read)，避免读写冲突。
- 动态调整，miss次数多了之后，将dirty数据迁移到read中。
- double-checking。
- 延迟删除。 删除一个键值只是打标记，只有在迁移dirty数据的时候才清理删除的数据。
- 优先从read读取、更新、删除，因为对read的读取不需要锁。

sync.Map 数据结构
在 src/sync/map.go 中
```go
type Map struct {
    // 当涉及到脏数据(dirty)操作时候，需要使用这个锁
    mu Mutex
    
    // read是一个只读数据结构，包含一个map结构，
    // 读不需要加锁，只需要通过 atomic 加载最新的指正即可
    read atomic.Value // readOnly
    
    // dirty 包含部分map的键值对，如果操作需要mutex获取锁
    // 最后dirty中的元素会被全部提升到read里的map去
    dirty map[interface{}]*entry
    
    // misses是一个计数器，用于记录read中没有的数据而在dirty中有的数据的数量。
    // 也就是说如果read不包含这个数据，会从dirty中读取，并misses+1
    // 当misses的数量等于dirty的长度，就会将dirty中的数据迁移到read中
    misses int
}
```
read的数据结构 readOnly：
```go
// readOnly is an immutable struct stored atomically in the Map.read field.
type readOnly struct {
    // m包含所有只读数据，不会进行任何的数据增加和删除操作 
    // 但是可以修改entry的指针因为这个不会导致map的元素移动
    m       map[interface{}]*entry
    
    // 标志位，如果为true则表明当前read只读map的数据不完整，dirty map中包含部分数据
    amended bool // true if the dirty map contains some key not in m.
}
```
只读map，对该map的访问不需要加锁，但是这个map也不会增加元素，元素会被先增加到dirty中，然后后续会迁移到read只读map中，通过原子操作所以不需要加锁操作。

readOnly.m和Map.dirty存储的值类型是*entry,它包含一个指针p, 指向用户存储的value值，结构如下：

```go
type entry struct {
    p unsafe.Pointer // *interface{}
}
```
p有三种值：

- nil: entry已被删除了，并且m.dirty为nil
- expunged: entry已被删除了，并且m.dirty不为nil，而且这个entry不存在于m.dirty中
- 其它： entry是一个正常的值

查找

根据key来查找 value， 函数为 Load()，源码如下：
```go
// src/sync/map.go

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
    // 首先从只读ready的map中查找，这时不需要加锁
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    
    // 如果没有找到，并且read.amended为true，说明dirty中有新数据，从dirty中查找，开始加锁了
    if !ok && read.amended {
        m.mu.Lock() // 加锁
        
       // 又在 readonly 中检查一遍，因为在加锁的时候 dirty 的数据可能已经迁移到了read中
        read, _ = m.read.Load().(readOnly)
        e, ok = read.m[key]
        
        // read 还没有找到，并且dirty中有数据
        if !ok && read.amended {
            e, ok = m.dirty[key] //从 dirty 中查找数据
            
            // 不管m.dirty中存不存在，都将misses + 1
            // missLocked() 中满足条件后就会把m.dirty中数据迁移到m.read中
            m.missLocked()
        }
        m.mu.Unlock()
    }
    if !ok {
        return nil, false
    }
    return e.load()
}
```
从函数可以看出，如果查询的键值正好在m.read中，不需要加锁，直接返回结果，优化了性能。即使不在read中，经过几次miss后，m.dirty中的数据也会迁移到m.read中，这时又可以从read中查找。所以对于`更新／增加`较少，加载存在的key很多的case，性能基本和无锁的map类似。

missLockerd() 迁移数据：
```go
// src/sync/map.go

func (m *Map) missLocked() {
    m.misses++
    if m.misses < len(m.dirty) {//misses次数小于 dirty的长度，就不迁移数据，直接返回
        return
    }
    m.read.Store(readOnly{m: m.dirty}) //开始迁移数据
    m.dirty = nil   //迁移完dirty就赋值为nil
    m.misses = 0  //迁移完 misses归0
}
```

新增和更新, 方法是 Store()， 更新或者新增一个 entry， 源码如下：
```go
// src/sync/map.go

// Store sets the value for a key.
func (m *Map) Store(key, value interface{}) {
   // 直接在read中查找值，找到了，就尝试 tryStore() 更新值
    read, _ := m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok && e.tryStore(&value) {
        return
    }
    
    // m.read 中不存在
    m.mu.Lock()
    read, _ = m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok {
        if e.unexpungeLocked() { // 未被标记成删除，前面讲到entry数据结构时，里面的p值有3种。1.nil 2.expunged，这个值含义有点复杂，可以看看前面entry数据结构 3.正常值
            
            m.dirty[key] = e // 加入到dirty里
        }
        e.storeLocked(&value) // 更新值
    } else if e, ok := m.dirty[key]; ok { // 存在于 dirty 中，直接更新
        e.storeLocked(&value)
    } else { // 新的值
        if !read.amended { // m.dirty 中没有新数据，增加到 m.dirty 中
            // We're adding the first new key to the dirty map.
            // Make sure it is allocated and mark the read-only map as incomplete.
            m.dirtyLocked() // 从 m.read中复制未删除的数据
            m.read.Store(readOnly{m: read.m, amended: true}) 
        }
        m.dirty[key] = newEntry(value) //将这个entry加入到m.dirty中
    }
    m.mu.Unlock()
}
```
操作都是先从m.read开始，不满足条件再加锁，然后操作m.dirty。

删除, 根据key删除一个值：
```go
// src/sync/map.go

// Delete deletes the value for a key.
func (m *Map) Delete(key interface{}) {
    // 从 m.read 中开始查找
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    
    if !ok && read.amended { // m.read中没有找到，并且可能存在于m.dirty中，加锁查找
        m.mu.Lock() // 加锁
        read, _ = m.read.Load().(readOnly) // 再在m.read中查找一次
        e, ok = read.m[key]
        if !ok && read.amended { //m.read中又没找到，amended标志位true，说明在m.dirty中
            delete(m.dirty, key) // 删除
        }
        m.mu.Unlock()
    }
    if ok { // 在 m.ready 中就直接删除
        e.delete()
    }
}
```
注意在使用sync.Map时切忌不要将其拷贝， go源码中有对sync.Map注释到” A Map must not be copied after first use.”因为当sync.Map被拷贝之后， Map类型的dirty还是那个map 但是read 和 锁却不是之前的read和锁(都不在一个世界你拿什么保护我)


## 设计并实现一个跳表（Skip List）
跳表本质上是一个链表, 它其实是由有序链表发展而来。跳表在链表之上做了一些优化，跳表在有序链表之上加入了若干层用于索引的有序链表。索引链表的结点来源于基础链表，不过只有部分结点会加入索引链表中，并且越高层的链表结点数越少。跳表查询从顶层链表开始查询，然后逐级展开，直到底层链表。这种查询方式与树结构非常类似，使得跳表的查询效率相近树结构。另外跳表使用概率均衡技术而不是使用强制性均衡，因此对于插入和删除结点比传统上的平衡树算法更为简洁高效。因此跳表适合增删操作比较频繁，并且对查询性能要求比较高的场景。

首先是基础链表，基础链表建立一层索引表，索引表只有基础链表结点的1/2。索引表相当于基础链表的目录，查询时首先从索引表开始查找，当遇到比待查数据大的结点时，再从基础链表中查找。由于索引表的结点只有基础链表的1/2，因此需要比较的结点大大减少，从而可以加快查询速度。

利用同样的方式，我们还为基础链表添加第二层索引表，第二层索引表结点的数量是第一层索引表的1/2，第二层索引表的结点的数量只有基础链表的1/4左右，查找时首先从第二层索引表开始查找，当遇到比待查数据大的结点时，再从第一层索引表开始查找，然后再从基础链表中查找数据。这种逐层展开的方式与树结构的查询非常类似，因此跳表的查询性能与树结构接近，时间的复杂度为O(logN)。

我们知道一个平衡的二叉树，在进行增加或删除结点后可能造成二叉树结构的不平衡，从而会降低二叉树的查询性能。按照上面的方式实现的跳表也会面临同样的问题，新增或删除结点后会破坏链表之间的比例关系，从而造成跳表查询性能的降低。平衡树通过旋转操作来保持平衡，而旋转一方面增加了代码实现的复杂性，同时也降低了增删操作的性能。

真正的跳表不会采用示例中的方式来建立上层链表，而是采用了一种概率均衡技术来创建上层链表， 并保证各层链表之间的比例关系。在为跳表增加一个结点时，会调用一个概率函数来计算一个结点的层次（level），例如若level=3，则结点除了出现在基础链表外，还会出现在第一层索引以及第二层索引链表中。结点层次的计算逻辑如下：

- 每个节点肯定都在基础链表中。
- 如果一个节点存在于第i层链表，那么它有第(i+1)层链表的概率为p。
- 节点最大的层数不允许超过一个最大值。

跳表每一个节点的层数是随机的，而且新插入一个节点不会影响其它节点的层数。因此插入操作只需要修改插入节点前后的指针，而不需要对很多节点都进行调整,这就降低了插入操作的复杂度。实际上这是跳表的一个很重要的特性，这让跳表在增删操作的性能上明显优于平衡树的方案。

代码：
```go
package skiplist

import (
	"fmt"
	"math/rand"
	"sync"
)

type SkipNodeInt struct {
	key   int64
	value interface{}
	next  []*SkipNodeInt
}

type SkipListInt struct {
	SkipNodeInt
	mutex  sync.RWMutex
	update []*SkipNodeInt
	maxl   int
	skip   int
	level  int
	length int32
}

func NewSkipListInt(skip ...int) *SkipListInt {
	list := &SkipListInt{}
	list.maxl = 32
	list.skip = 4
	list.level = 0
	list.length = 0
	list.SkipNodeInt.next = make([]*SkipNodeInt, list.maxl)
	list.update = make([]*SkipNodeInt, list.maxl)
	if len(skip) == 1 && skip[0] > 1 {
		list.skip = skip[0]
	}
	return list
}

func (list *SkipListInt) Get(key int64) interface{} {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	var prev = &list.SkipNodeInt
	var next *SkipNodeInt
	for i := list.level - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = prev.next[i]
		}
	}

	if next != nil && next.key == key {
		return next.value
	} else {
		return nil
	}
}

func (list *SkipListInt) Set(key int64, val interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	//获取每层的前驱节点=>list.update
	var prev = &list.SkipNodeInt
	var next *SkipNodeInt
	for i := list.level - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = prev.next[i]
		}
		list.update[i] = prev
	}
	fmt.Println("update 1: ", list.update)
	fmt.Println("list next 1: ", list.next)
	//如果key已经存在
	if next != nil && next.key == key {
		next.value = val
		return
	}

	//随机生成新结点的层数
	level := list.randomLevel()
	if level > list.level {
		level = list.level + 1
		list.level = level
		list.update[list.level-1] = &list.SkipNodeInt
	}
	fmt.Println("update 2: ", list.update)
	fmt.Println("list next 2: ", list.next)
	//申请新的结点
	node := &SkipNodeInt{}
	node.key = key
	node.value = val
	node.next = make([]*SkipNodeInt, level)

	//调整next指向
	for i := 0; i < level; i++ {
		node.next[i] = list.update[i].next[i]
		list.update[i].next[i] = node
	}
	fmt.Println("update 3: ", list.update)
	fmt.Println("list next 3: ", list.next)
	list.length++
}

func (list *SkipListInt) Remove(key int64) interface{} {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	//获取每层的前驱节点=>list.update
	var prev = &list.SkipNodeInt
	var next *SkipNodeInt
	for i := list.level - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && next.key < key {
			prev = next
			next = prev.next[i]
		}
		list.update[i] = prev
	}

	//结点不存在
	node := next
	if next == nil || next.key != key {
		return nil
	}

	//调整next指向
	for i, v := range node.next {
		if list.update[i].next[i] == node {
			list.update[i].next[i] = v
			if list.SkipNodeInt.next[i] == nil {
				list.level -= 1
			}
		}
		list.update[i] = nil
	}

	list.length--
	return node.value
}

func (list *SkipListInt) randomLevel() int {
	i := 1
	for ; i < list.maxl; i++ {
		r := rand.Int()
		fmt.Println("random n: ", r, " i: ", i)
		if r%list.skip != 0 {
			break
		}
	}
	fmt.Println("random Level", i)
	return i
}
```


## 设计一个分布式缓存系统
要求支持高并发访问和数据一致性保证。请讨论你的设计方案，包括但不限于缓存数据的分片、缓存失效策略、以及如何解决缓存雪崩问题。

## 设计一个分布式ID生成系统
要求生成的ID既要保证全局唯一，又要尽可能地有序。请讨论你的设计方案和实现细节。

雪花算法?


## Trie
LeetCode 0208, 0211

## LRU
LeetCode 0146

## InsertInverval
LeetCode 056, 057

## Randomset O(1)
LeetCode 0380


## Intensity Segment
```go
package main

import (
	"fmt"
	"strings"
)

// Segment represents a start with an intensity value.
type Segment struct {
	Start     int
	Intensity int
}

// toString formats one segment like [1,2].
func (s *Segment) toString() string {
	return fmt.Sprintf("[%d,%d]", s.Start, s.Intensity)
}

// IntensitySegments manages a collection of segments
type IntensitySegments struct {
	m        map[int]int // A map to quickly locate the index of a segment by its start position. key is segment's Start, value is index in Segments
	Segments []*Segment  // [[start, intensity]...] and sorted by start
}

// constructor
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		m:        make(map[int]int, 0),
		Segments: make([]*Segment, 0),
	}
}

func (is *IntensitySegments) Add(start, end, intensity int) {
	// Check for invalid arguments.
	if start >= end {
		panic("wrong arugments.")
	}
	// Handle the case where the collection is empty.
	if len(is.Segments) == 0 {
		is.Segments = append(is.Segments, &Segment{start, intensity}, &Segment{end, 0})
		is.checkAndTrim()
		return
	}

	// Handle the case where the new segment is entirely before the first existing segment.
	if end < is.Segments[0].Start {
		is.Segments = append([]*Segment{{start, intensity}, {end, 0}}, is.Segments...)
		is.checkAndTrim()
		// Handle the case where the new segment is entirely after the last existing segment.
	} else if start > is.Segments[len(is.Segments)-1].Start {
		is.Segments = append(is.Segments, &Segment{start, intensity}, &Segment{end, 0})
		is.checkAndTrim()
	} else {
		// insert new start
		if index, ok := is.m[start]; !ok {
			l := is.findLeft(start)
			if l == 0 {
				is.Segments = append([]*Segment{{start, intensity}}, is.Segments...)
			} else {
				seg := &Segment{start, intensity + is.Segments[l-1].Intensity}
				is.insert(l, seg)
			}
		} else {
			is.Segments[index].Intensity += intensity
		}
		is.updateIndex()
		// insert new end
		if _, ok := is.m[end]; !ok {
			r := is.findRight(end)
			if r == len(is.Segments)-1 {
				is.Segments = append(is.Segments, &Segment{end, 0})
			} else if r < len(is.Segments)-1 {
				var seg *Segment
				if r > 0 && is.Segments[r].Start == start {
					seg = &Segment{end, is.Segments[r-1].Intensity}
				} else {
					seg = &Segment{end, is.Segments[r].Intensity}
				}
				is.insert(r+1, seg)
			}
		}
		is.updateIndex()
		// update intensity between start and end
		l, r := is.m[start], is.m[end]
		for i := l + 1; i < r && i < len(is.Segments); i++ {
			is.Segments[i].Intensity += intensity
		}
		is.checkAndTrim()
	}
}

// Set is an alias for Add.
func (is *IntensitySegments) Set(start, end, intensity int) {
	is.Add(start, end, intensity)
}

// ToString formats and prints the current segments like [[1,2],[3,0]]....
func (is *IntensitySegments) ToString() string {
	var segemntStrs []string
	for _, s := range is.Segments {
		segemntStrs = append(segemntStrs, s.toString())
	}
	segStr := strings.Join(segemntStrs, ",")
	fmt.Printf("[%s]\n", segStr)
	return segStr
}

//  findLeft searches and returns the index of the first segment start position that is greater than the given start position.

func (is *IntensitySegments) findLeft(start int) int {
	for i, s := range is.Segments {
		if s.Start > start {
			return i
		}
	}
	return 0
}

// findRight searches and returns the index of the last segment start position that is less than the given end position.

func (is *IntensitySegments) findRight(end int) int {
	for i := len(is.Segments) - 1; i >= 0; i-- {
		if is.Segments[i].Start < end {
			return i
		}
	}
	return len(is.Segments) - 1
}

// insert inserts a new segment at the specified index.
func (is *IntensitySegments) insert(index int, seg *Segment) {
	is.Segments = append(is.Segments, &Segment{0, 0})
	copy(is.Segments[index+1:], is.Segments[index:])
	is.Segments[index] = seg
}

// updateIndex updates the index mapping for all segments after potential changes.
func (is *IntensitySegments) updateIndex() {
	for i, n := range is.Segments {
		is.m[n.Start] = i
	}
}

// checkAndTrim removes any leading or trailing segments with zero intensity.

func (is *IntensitySegments) checkAndTrim() {
	if is.Segments[0].Intensity == 0 {
		delete(is.m, is.Segments[0].Start)
		is.Segments = is.Segments[1:]
	}
	l := len(is.Segments)
	if l >= 2 && is.Segments[l-1].Intensity == 0 && is.Segments[l-2].Intensity == 0 {
		delete(is.m, is.Segments[l-1].Start)
		is.Segments = is.Segments[:len(is.Segments)-1]
	}
	is.updateIndex()
}

func main() {
	s1 := NewIntensitySegments()
	s1.ToString() // Should be "[]"

	s1.Add(10, 30, 1)
	s1.ToString() // Should be: "[[10,1],[30,0]]"
	s1.Add(20, 40, 1)
	s1.ToString() // Should be: "[[10,1],[20,2],[30,1],[40,0]]"
	s1.Add(10, 40, -2)
	s1.ToString() // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"

	// Another example sequence:
	s2 := NewIntensitySegments()
	s2.ToString() // Should be "[]"
	s2.Add(10, 30, 1)
	s2.ToString() // Should be "[[10,1],[30,0]]"
	s2.Add(20, 40, 1)
	s2.ToString() // Should be "[[10,1],[20,2],[30,1],[40,0]]"
	s2.Add(10, 40, -1)
	s2.ToString() // Should be "[[20,1],[30,0]]"
	s2.Add(10, 40, -1)
	s2.ToString() // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"

	s := NewIntensitySegments()
	s.ToString() // []
	s.Add(1, 2, 3)
	s.ToString() // [[1,3],[2,0]]
	s.Add(2, 3, 1)
	s.ToString() //[[1,3],[2,1],[3,0]]
	s.Add(10, 40, 5)
	s.ToString() // [[1,3],[2,1],[3,0],[10,5],[40,0]]
	s.Add(13, 18, 4)
	s.ToString() // [[1,3],[2,1],[3,0],[10,5],[13,9],[18,5],[40,0]]
	s.Add(3, 30, 1)
	s.ToString() // [[1,3],[2,1],[3,1],[10,6],[13,10],[18,6],[30,5],[40,0]]
	s.Add(-9, 2, 10)
	s.ToString() // [[-9,10],[1,13],[2,1],[3,1],[10,6],[13,10],[18,6],[30,5],[40,0]]

}

/*
package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (sl intSlice) Len() int {
	return len(sl)
}
func (sl intSlice) Less(i, j int) bool {
	return sl[i] < sl[j]
}
func (sl intSlice) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

// IntensitySegments
//
//	Using a map (because of it dscreteness) to store: [segemnt] -> intensity
type IntensitySegments struct {
	it map[int]int
}

// NewIntensitySegments new a IntensitySegments Object
// return a pointer to the new object
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		it: map[int]int{},
	}
}

// set set a new range(from <-> to, to not included) of intensity based on existing one.
func (is *IntensitySegments) set(from, to int, amount int) {
	keys := is.orderedKeys()
	if len(keys) == 0 { // if no segments, in the intial case
		is.it[to] = 0
		is.it[from] = amount
		return
	}

	if to < keys[0] || to > keys[len(keys)-1] { // the end of set segment is outside of the existing one
		is.it[to] = 0
	} else if _, f := is.it[to]; !f { // 'to' is not found in existing segment numbers
		l := is.leftKey(to)
		is.it[to] = is.it[l]
	}

	is.it[from] = amount // directly set the segment beginning

	for _, k := range is.orderedKeys() { // all keys between from and to is invalid, delete them
		if k > from && k < to {
			delete(is.it, k)
		}
	}

	is.merge()
}

// Add Add a new range(from <-> to, to not included) of intensity to existing one.
func (is *IntensitySegments) Add(from, to int, amount int) {
	keys := is.orderedKeys()
	if len(keys) == 0 { // if no segments, in the intial case
		is.set(from, to, amount)
		return
	}

	if to < keys[0] || from > keys[len(keys)-1] { // no overlap case
		is.set(from, to, amount)
		return
	}

	// handle the 'to'
	if _, f := is.it[to]; !f { // if not found to, Add it
		l := is.leftKey(to)
		is.it[to] = is.it[l]
	}

	// handle the 'from'
	if from < keys[0] {
		is.it[from] = amount
	} else if from == keys[0] {
		is.it[from] += amount
	} else {
		if _, f := is.it[from]; !f {
			is.it[from] = amount
			if l := is.leftKey(from); l != -1 {
				is.it[from] += is.it[l]
			}
		} else {
			is.it[from] += amount
		}
	}

	// handle the keys between from and 'to'
	for _, k := range keys {
		if k > from && k < to {
			is.it[k] += amount
		}
	}

	// merge segments
	is.merge()
}

// merge merge continous segments to together if they have same intensity
func (is *IntensitySegments) merge() {
	keys := is.orderedKeys()
	if len(keys) == 0 {
		return
	}

	// clear prefix 0
	i := 0
	for ; i < len(keys) && is.it[keys[i]] == 0; i++ {
		delete(is.it, keys[i])
	}
	keys = keys[i:]

	for i = len(keys) - 1; i >= 0; i-- {
		if is.it[keys[i]] == 0 {
			if i-1 >= 0 && is.it[keys[i-1]] == 0 {
				delete(is.it, keys[i])
				keys = append(keys[0:i], keys[i+1:]...)
			}
		}

	}

	// clear the surffix segment with same intensity <> 0
	lastIntensity := is.it[keys[0]]
	for i := 1; i < len(keys); i++ {
		if is.it[keys[i]] == lastIntensity && lastIntensity != 0 {
			delete(is.it, keys[i])
		} else {
			lastIntensity = is.it[keys[i]]
		}
	}

	// clear the surffix segment with intensity == 0
	lastIntensity = 0
	for i := len(keys) - 2; i >= 0; i-- {
		if is.it[keys[i]] == 0 && lastIntensity == 0 {
			delete(is.it, keys[i])
		} else {
			lastIntensity = is.it[keys[i]]
		}
	}
}

// orderedKeys orderedKeys return a sorted keys of is.it
func (is *IntensitySegments) orderedKeys() []int {
	keys := intSlice{}
	for k := range is.it {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

// dumps dumps return a string of simple format, i.e. [[20 1] [30 2] [40 0]]
func (is *IntensitySegments) dumps() string {
	keys := is.orderedKeys()
	rlt := [][2]int{}
	for _, k := range keys {
		rlt = append(rlt, [2]int{k, is.it[k]})
	}
	return fmt.Sprintf("%v", rlt)
}

// leftKey return the left segment number of the given.
func (is *IntensitySegments) leftKey(x int) int {
	l := -1
	for _, k := range is.orderedKeys() {
		if k < x {
			l = k
		}
	}
	return l
}

// ToString print the dumped string simply
func (is *IntensitySegments) ToString() {
	fmt.Printf("%v\n", is.dumps())
}

func main() {
	segments1 := NewIntensitySegments()
	segments1.ToString() // Should be "[]"

	segments1.Add(10, 30, 1)
	segments1.ToString() // Should be: "[[10,1],[30,0]]"
	segments1.Add(1, 5, 1)
	segments1.ToString() // Should be: "[[10,1],[20,2],[30,1],[40,0]]"
	segments1.Add(10, 40, -2)
	segments1.ToString() // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"

	// Another example sequence:
	segments2 := NewIntensitySegments()
	segments2.ToString() // Should be "[]"
	segments2.Add(10, 30, 1)
	segments2.ToString() // Should be "[[10,1],[30,0]]"
	segments2.Add(20, 40, 1)
	segments2.ToString() // Should be "[[10,1],[20,2],[30,1],[40,0]]"
	segments2.Add(10, 40, -1)
	segments2.ToString() // Should be "[[20,1],[30,0]]"
	segments2.Add(10, 40, -1)
	segments2.ToString() // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
}
*/
```
