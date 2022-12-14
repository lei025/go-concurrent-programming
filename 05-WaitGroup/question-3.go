package main

import (
	"sync"
	"time"
)

// 常见问题三：前一个 Wait 还没结束就重用 WaitGroup
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done() // 计数器减1
		wg.Add(1) // 计数值加1
	}()
	wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}
