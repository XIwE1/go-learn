package main

import (
	"context"
	"fmt"
	"time"
)

// what context 库（context 包）在 Go 里主要用来做 “取消/超时/请求范围传递”
// why 一个请求（比如一次 HTTP 请求）超时了，我希望相关的所有 goroutine 都尽快退出,下游不需要结果了
// how
//  取消 context.WithCancel(parent) 产生一个 ctx 和一个 cancel()
//  截止时间 context.WithTimeout(parent, d) 

func processWithContext(ctx context.Context, input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for {
			select {
				case value := <-input:
					time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				case <-ctx.Done():
					fmt.Println("Context cancelled")
					return
			}
		}
	}()

	return output
}

func main() {
	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建输入 channel
	input := make(chan int)

	// 开始处理
	output := processWithContext(ctx, input)

	// 发送数据
	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			select {
				case input <- i:
					fmt.Println("Sent:", i)
				case <-ctx.Done():
					fmt.Println("Context cancelled")
					return
			}
		}
	}()
	
	// 收集结果
	var results []int
	for {
		select {
		case result := <-output:
			results = append(results, result)
			fmt.Println("Received:", result)
		case <-ctx.Done():
			fmt.Println("Context cancelled")
			goto done
		}
	}

	done:
	fmt.Println("Results:", results)
}