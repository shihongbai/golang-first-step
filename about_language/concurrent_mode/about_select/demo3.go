package main

// 对于空的select语句，程序会被阻塞，准确的说是当前协程被阻塞，同时Golang自带死锁检测机制，当
// 发现当前协程再也没有机会被唤醒时，则会panic。所以上述程序会panic
func main() {
	select {}
}
