// Go: 使用 goroutine 和 channel 的工作池
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkerPool struct {
	workerCount int
	tasks       chan string
	results     chan string
	wg          sync.WaitGroup
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		tasks:       make(chan string, workerCount),
		results:     make(chan string, workerCount),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for task := range wp.tasks {
		// 模拟工作
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		wp.results <- fmt.Sprintf("工作线程 %d 处理了: %s", id, task)
	}
}

func (wp *WorkerPool) Submit(tasks []string) {
	// 提交所有任务
	for _, task := range tasks {
		wp.tasks <- task
	}
	close(wp.tasks)
}

func (wp *WorkerPool) CollectResults() []string {
	var results []string
	
	// 在单独的 goroutine 中收集结果
	go func() {
		for result := range wp.results {
			results = append(results, result)
		}
	}()
	
	// 等待所有工作线程完成
	wp.wg.Wait()
	close(wp.results)
	
	return results
}

func main() {
	pool := NewWorkerPool(3)
	pool.Start()
	
	tasks := []string{"任务 0", "任务 1", "任务 2", "任务 3", "任务 4"}
	pool.Submit(tasks)
	
	results := pool.CollectResults()
	fmt.Println("所有任务完成:", results)
}