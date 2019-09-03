package awesomeProject

import (
	"fmt"
	"math/rand"
	"time"
)

func Destroy(url string, times int, concurrency int) {
	var (
		taskChan chan int = make(chan int, concurrency)	// 任务队列
		endChan chan byte = make(chan byte, concurrency) // 结束队列
	)

	// 拉起一些worker
	for i := 0; i < concurrency; i++ {
		go worker(url, taskChan, endChan)
	}

	// 开始填充任务
	for i := 0; i < times; i++ {
		taskChan <- i
	}

	// 通知任务耗尽
	close(taskChan)

	// 等待worker全部退出
	for i := 0; i < concurrency; i++ {
		<- endChan
	}
}

func worker(url string, taskChan chan int, endChan chan byte) {
	for {
		select {
		case times, valid := <- taskChan:
			if valid {
				// 模拟请求0~1秒
				time.Sleep(time.Duration(rand.Intn(1000))  * time.Millisecond)
				fmt.Printf("第%d次攻击\n", times)
			} else {
				// 告知退出
				endChan <- 1
				return
			}
		}
	}
}