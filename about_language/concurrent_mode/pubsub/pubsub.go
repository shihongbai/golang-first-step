// Package pubsub 发布订阅模型 Package pubsub implements a simple multi-topic pub-sub library.
// 发布订阅（publish-and-subscribe）模型通常被简写为 pub/sub 模型。
// 在这个模型中，消息生产者成为发布者（publisher），而消息消费者则成为订阅者（subscriber），
// 生产者和消费者是 M:N 的关系。在传统生产者和消费者模型中，是将消息发送到一个队列中，
// 而发布订阅模型则是将消息发布给一个主题。
package pubsub

import (
	"sync"
	"time"
)

type (
	// 订阅者为一个通道
	subscriber chan interface{}
	// 主题为一个过滤器
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m           sync.RWMutex             // 读写锁
	buffer      int                      // 订阅队列的缓存大小
	timout      time.Duration            // 发布超时时间
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建一个发布者对象，可以设置发布超时间和缓存队列长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timout:      publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict 退出订阅
func (p *Publisher) Evict(ch chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, ch)
	close(ch)
}

// Publish 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var war sync.WaitGroup
	for sub, topic := range p.subscribers {
		war.Add(1)
		go p.sendTopic(sub, topic, v, &war)
	}
	war.Wait()
}

// Close 关闭发布者对象，并关闭所有订阅者
func (p *Publisher) Close() {
	p.m.RLock()
	defer p.m.RUnlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题, 并容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timout):
	}
}
