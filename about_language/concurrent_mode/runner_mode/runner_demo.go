package runner_mode

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送终端信号时结束这些任务
type Runner struct {
	// interrupt 通道：报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete 通道：报告任务处理完成
	complete chan error

	// timeout 报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在仍无执行超时时返回.
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的时间时返回.
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新准备使用的runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将一个任务附加到Runner上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有仍无并监视通道事件
func (r *Runner) Start() error {
	// 接收所有的终端信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered task.
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals.
		signal.Stop(r.interrupt)
		return true

	// Continue running as normal.
	default:
		return false
	}
}
