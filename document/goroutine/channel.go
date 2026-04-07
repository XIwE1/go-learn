package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sync wg = waitgroup 用于等待一组 goroutine 完成
// wg.Add(n)：计数器 +n，表示“接下来有 n 个任务/协程要完成”
// wg.Done()：计数器 -1，表示“我这个任务/协程完成了”
// wg.Wait()：阻塞等待，直到计数器变成 0 才返回

func fanOut(input <-chan string, workerCount int) []<-chan string {
	outputs := make([]<-chan string, workerCount)

	// 创建用于接收结果的channel
	for i := 0; i < workerCount; i++ {
		outputs[i] = make(chan string)
	}

	// 启动工作线程
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, input, outputs[i], &wg)
	}

	// 所有工作线程完成后关闭所有输出channel
	go func() {
		wg.Wait()

		for _, output := range outputs {
			close(output)
		}
	}()

	result := make([]<-chan string, workerCount)
	for i, output := range outputs {
		result[i] = output
	}

	return result
}

func fanIn(channels ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	// 为每个channel启动一个goroutine来读取数据
	for _, channel := range channels {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				output <- msg
			}
		}(channel)
	}

	// 当所有输入 channels 关闭时关闭输出
	go func() {
		wg.Wait()
		close(output)	
	}()

	return output
}

func producer(id int) <-chan string {
	output := make(chan string)

	go func() {
		// defer关键字用于延迟执行，会在函数返回前执行(return)
		defer close(output)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			output <- fmt.Sprintf("Producer %d: Message %d", id, i)
		}
	}()

	return output
}

func worker(id int, input <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range input {
		// 模拟工作
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		output <- fmt.Sprintf("Worker %d processed: %s", id, task)
	}
}

func main() {
	// fan out 模式 ： 将工作从一个 channel 分发到多个工作线程，适合批量任务处理
	input := make(chan string)
	outputs := fanOut(input, 3)

	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- fmt.Sprintf("Task %d", i)
		}
	}

	var results []string
	for _, output := range outputs {
		for result := range output {
			results = append(results, result)
		}
	}

	fmt.Println("所有任务完成:", results)

	// fan in 模式 ： 将多个工作线程的结果合并到一个 channel 中，适合收集结果如并行查询 统一结果给下游
	producer1 := producer(1)
	producer2 := producer(2)
	producer3 := producer(3)

	combined := fanIn(producer1, producer2, producer3)

	var results2 []string
	for result := range combined {
		results2 = append(results2, result)
		fmt.Println("收到结果:", result)
	}
	fmt.Println("所有结果完成:", results2)
}