package main

import (
	"fmt"
	"sync"
	"time"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("i love test1 - ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器-1
}

var wg sync.WaitGroup

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("i love test2 - ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器-1
}

func main() {
	wg.Add(1)  // 协程计数器+1
	go test1() // 开启协程
	wg.Add(1)  // 协程计数器+1
	go test2() // 开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("i love main - ", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() // 等待协程执行完毕
	fmt.Println("主线程退出")
}
