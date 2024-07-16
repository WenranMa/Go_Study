# Go 设计题

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
	if is.addNewSeg(start, end, intensity) {
		return
	}
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

// Set
func (is *IntensitySegments) Set(start, end, intensity int) {
	if is.addNewSeg(start, end, intensity) {
		return
	}
	// insert new start
	if index, ok := is.m[start]; !ok {
		l := is.findLeft(start)
		seg := &Segment{start, intensity}
		is.insert(l, seg)
	} else {
		is.Segments[index].Intensity = intensity
	}
	is.updateIndex()
	// insert new end
	if _, ok := is.m[end]; !ok {
		r := is.findRight(end)
		seg := &Segment{end, 0}
		is.insert(r+1, seg)
	}
	is.updateIndex()
	// delete intensity between start and end
	l, r := is.m[start], is.m[end]
	if r-l > 1 {
		temp := is.Segments
		is.Segments = append([]*Segment{}, temp[0:l+1]...)
		is.Segments = append(is.Segments, temp[r:]...)
	}
	is.updateIndex()
	is.checkAndTrim()
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

func (is *IntensitySegments) addNewSeg(start, end, intensity int) bool {
	// Check for invalid arguments.
	if start >= end {
		panic("wrong arugments.")
	}
	// Handle the case where the collection is empty.
	if len(is.Segments) == 0 {
		is.Segments = append(is.Segments, &Segment{start, intensity}, &Segment{end, 0})
		is.checkAndTrim()
		return true
	}
	// Handle the case where the new segment is entirely before the first existing segment.
	if end < is.Segments[0].Start {
		is.Segments = append([]*Segment{{start, intensity}, {end, 0}}, is.Segments...)
		is.checkAndTrim()
		return true
	}
	// Handle the case where the new segment is entirely after the last existing segment.
	if start > is.Segments[len(is.Segments)-1].Start {
		is.Segments = append(is.Segments, &Segment{start, intensity}, &Segment{end, 0})
		is.checkAndTrim()
		return true
	}
	return false
}

// findLeft searches and returns the index of the first segment start position that is greater than the given start position.
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
	for i := len(is.Segments) - 1; i > 1 && is.Segments[i].Intensity == 0 && is.Segments[i-1].Intensity == 0; i -= 1 {
		delete(is.m, is.Segments[i].Start)
		is.Segments = is.Segments[:i]
	}
	var i int
	for i = 0; i < len(is.Segments); i += 1 {
		if is.Segments[i].Intensity == 0 {
			delete(is.m, is.Segments[i].Start)
		} else {
			break
		}
	}
	is.Segments = is.Segments[i:]
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
