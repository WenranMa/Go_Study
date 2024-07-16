# 设计消息队列

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
