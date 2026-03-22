// Go: 缓冲 vs 无缓冲 channel
package main

import (
	"fmt"
	"time"
)

func createChannel(name string, delay time.Duration) chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(delay)
		ch <- fmt.Sprintf("来自 %s 的数据", name)
	}()
	return ch
}

func main() {
	// 无缓冲 channel（同步）
	unbuffered := make(chan string)
	
	// 缓冲 channel（在缓冲区大小内异步）
	buffered := make(chan string, 2)
	
	// 无缓冲 channel 示例
	go func() {
		fmt.Println("发送到无缓冲 channel...")
		unbuffered <- "Hello" // 这会阻塞后续代码直到有人接收
		fmt.Println("已发送到无缓冲 channel")
	}()
	
	time.Sleep(100 * time.Millisecond)
	message := <-unbuffered
	fmt.Println("从无缓冲 channel 收到:", message)
	
	// 缓冲 channel 示例
	fmt.Println("发送到缓冲 channel...")
	buffered <- "第一条消息"  // 不会阻塞
	buffered <- "第二条消息" // 不会阻塞
	fmt.Println("已发送到缓冲 channel")
	
	// 从缓冲 channel 接收
	fmt.Println("收到:", <-buffered)
	fmt.Println("收到:", <-buffered)

	// select 语句允许 goroutine 等待多个 channel 操作：
	// 创建多个 channel
	chA := createChannel("A", 1*time.Second)
	chB := createChannel("B", 2*time.Second)
	chC := createChannel("C", 1500*time.Millisecond)

	// Select 等待第一个可用的 channel
	select {
	case msg := <-chA:
		fmt.Println("从 A 收到:", msg)
	case msg := <-chB:
		fmt.Println("从 B 收到:", msg)
	case msg := <-chC:
		fmt.Println("从 C 收到:", msg)
	}
	// 使用 select 和 time.After 提供超时功能
	case <-time.After(3 * time.Second):
		fmt.Println("操作超时")
	}
}