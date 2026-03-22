// goroutine 是 Go 语言中的轻量级线程，用于**并发编程**
// 
// 何时使用 Goroutine
// I/O 密集型操作: 网络请求、文件操作
// CPU 密集型操作: 数学计算（需要适当协调）
// 事件处理: 并发处理多个事件
// 后台任务: 清理、监控、日志记录。

// 并发：是通过交错执行来处理多个任务的能力
// 并行：是在不同 CPU 核心上同时执行多个任务的能力

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func simulateTask(name string, duration time.Duration) {
	fmt.Printf("%s 开始\n", name)
	time.Sleep(duration)
	fmt.Printf("%s 完成\n", name)
}

func task1() {
	fmt.Println("任务1开始")
}

func task2(param string) {
	fmt.Println("任务2开始 带参数", param)
}

func worker(id int, done chan<- int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Worker %d 完成工作\n", id)
	done <- id  // 发送结果到Channel
}

func main() {
	fmt.Println("主程序开始")
	
	// 启动 goroutine（并发执行）
	go simulateTask("任务 A", 2*time.Second)
	go simulateTask("任务 B", 1*time.Second)
	
	// 等待 goroutine 完成
	time.Sleep(3 * time.Second)
	
	fmt.Println("主程序结束")

	go func() {
		task1()
	}()

	go task2("hello")

	go func(name string) {
		fmt.Println("来自", name, "的问候")
	}("goroutine")

	// 等待 goroutine 完成
	time.Sleep(100 * time.Millisecond)

	// 运行channel例子
	channelExample()
}

// 并发编程的核心概念：
// 1. Goroutine: 轻量级线程，用于并发执行任务
// 2. Channel: chan 类型 用于 goroutine 间通信的管道
// 3. Select: 用于多路复用，处理多个 channel 的读写操作
// 4. Mutex: 用于保护共享资源的访问
// 5. WaitGroup: 用于等待一组 goroutine 完成

// Channel 是 Go 中 goroutine 间通信的主要机制 提供了一种安全的数据共享方式
// <- 的位置决定作用 done<-表写入 <-done表读取 在方法里<-限制管道的用法
// 如果 chan 不带箭头表示即可读也可写
// 通过 close(done) 关闭管道
func channelExample() {
	fmt.Println("启动工作线程...")

	// 创建用于同步的 channel
	done := make(chan int, 3)

	// 启动多个线程
	for i := 0; i < 3; i++ {
		go worker(i, done)
	}

	var results []int
	// 如果done里的值小于3会阻塞
	for i := 0; i < 3; i++ {
		result := <-done // 从 channel 取出值使用
		results = append(results, result)
	}
	fmt.Println("所有工作线程完成:", results)
}


