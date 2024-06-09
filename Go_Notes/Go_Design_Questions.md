# Go 设计题

## 设计消息队列

在Go语言中实现一个简单的消息队列，支持发布/订阅模式并能处理高并发场景

1. 设计消息队列结构

首先，定义消息队列的数据结构，包括消息通道、订阅者列表等。
```go
type Topic struct {
	name    string
	messages chan interface{}
	subscribers map[string]chan interface{}
	mutex sync.RWMutex
}

type SimpleMQ struct {
	topics map[string]*Topic
	mutex sync.RWMutex
}
```
2. 实现消息发布

为SimpleMQ添加发布消息的方法，确保消息能被所有订阅该主题的订阅者接收。
```go
func (s *SimpleMQ) Publish(topicName string, message interface{}) {
	s.mutex.RLock()
	topic, ok := s.topics[topicName]
	s.mutex.RUnlock()

	if !ok {
		return // 或者创建新的主题
	}

	topic.mutex.RLock()
	for _, sub := range topic.subscribers {
		sub <- message
	}
	topic.mutex.RUnlock()
}
```

3. 实现消息订阅

添加订阅和取消订阅的方法，以及确保订阅者能够接收到新发布的消息。

```go
func (s *SimpleMQ) Subscribe(topicName string) (<-chan interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	topic, ok := s.topics[topicName]
	if !ok {
		topic = &Topic{
			name:        topicName,
			messages:    make(chan interface{}),
			subscribers: make(map[string]chan interface{}),
		}
		s.topics[topicName] = topic
	}

	subChannel := make(chan interface{})
	topic.subscribers["uniqueSubscriberID"] = subChannel // 使用唯一ID替换"uniqueSubscriberID"

	go func() {
		for msg := range topic.messages {
			subChannel <- msg
		}
		close(subChannel)
	}()

	return subChannel, nil
}

func (s *SimpleMQ) Unsubscribe(topicName, subscriberID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	topic, ok := s.topics[topicName]
	if !ok {
		return
	}

	topic.mutex.Lock()
	delete(topic.subscribers, subscriberID)
	topic.mutex.Unlock()
}
```

4. 处理高并发
为了处理高并发情况，可以考虑以下几点优化：

使用带缓冲的channel：为messages channel设置合适的缓冲大小，以减少goroutine间的同步等待。
go
topic.messages = make(chan interface{}, bufferSize)
限制goroutine数量：在处理发布消息到多个订阅者时，如果订阅者数量非常大，可以考虑使用worker pool模式限制并发执行的goroutine数量。

锁的优化：合理使用读写锁(sync.RWMutex)，减少锁的竞争，提高并发性能。

5. 完整示例代码
以上代码片段展示了实现的基本框架，实际应用中还需要根据具体需求进行调整和完善。完整的实现可能还包括错误处理、日志记录、以及更复杂的订阅管理逻辑等。