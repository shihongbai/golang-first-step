package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个goroutine间
// 共享的资源，被管理的子会员必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex // 资源池锁
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求(Acquire)了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New 用来创建一个管理资源的池
// 这个池需要一个可以分配新资源的函数， 并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Print("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil
	default:
		log.Print("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池中
func (p *Pool) Release(r io.Closer) {
	// 确保本操作和close的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池已经关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	// 试图将资源放入队列
	case p.resources <- r:
		log.Print("Release:", "In Queue")
	// 如果队列已满，关闭这个资源
	default:
		log.Print("Release", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭现有资源
func (p *Pool) Close() {
	// 保证操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果已经关闭，直接返回
	if p.closed {
		return
	}

	// 将池关闭
	p.closed = true

	// 在清空通道资源前，将通道关闭，如果不这么做会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
