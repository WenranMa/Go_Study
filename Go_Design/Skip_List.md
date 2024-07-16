# # 设计并实现一个跳表（Skip List）

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
