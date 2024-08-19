# 362. 敲击计数器 Design Hit Counter

Design a hit counter which counts the number of hits received in the past 5 minutes.

Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.

It is possible that several hits arrive roughly at the same time.

Example:

```java
HitCounter counter = new HitCounter();

// hit at timestamp 1.
counter.hit(1);

// hit at timestamp 2.
counter.hit(2);

// hit at timestamp 3.
counter.hit(3);

// get hits at timestamp 4, should return 3.
counter.getHits(4);

// hit at timestamp 300.
counter.hit(300);

// get hits at timestamp 300, should return 4.
counter.getHits(300);

// get hits at timestamp 301, should return 3.
counter.getHits(301); 
```

Follow up:
What if the number of hits per second could be very large? Does your design scale?

### 题目大意
设计一个敲击计数器，使它可以统计在过去5分钟内被敲击次数。

每个函数会接收一个时间戳参数（以秒为单位），你可以假设最早的时间戳从1开始，且都是按照时间顺序对系统进行调用（即时间戳是单调递增）。

在同一时刻有可能会有多次敲击。

### 解：

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	fmt.Println(hc.GetHits(4))
	hc.Hit(300)
	fmt.Println(hc.GetHits(300))
	hc.GetHits(310)
	fmt.Println(hc.GetHits(310))
}

type HitCounter struct {
	hits *list.List
}

func Constructor() *HitCounter {
	return &HitCounter{
		hits: list.New(),
	}
}

// 记录一次点击
func (hc *HitCounter) Hit(timestamp int) {
	hc.hits.PushBack(timestamp)
	for e := hc.hits.Front(); e != nil; e = e.Next() {
		if e.Value.(int) < timestamp-300 {
			hc.hits.Remove(e)
		}
	}
}

// 获取过去5分钟内的点击数
func (hc *HitCounter) GetHits(timestamp int) int {
	count := 0
	for e := hc.hits.Front(); e != nil; e = e.Next() {
		if e.Value.(int) > timestamp-300 {
			count++
		}
	}
	return count
}

```