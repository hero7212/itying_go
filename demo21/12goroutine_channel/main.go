package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	// close(primeChan)	// 关闭了，其他循环中同名的协程就用不了了
	// 什么时候关闭prime管道呢？ 定义exitChan来存放prime关闭的标识
	// 给exitChan 放数据
	exitChan <- true
	wg.Done()

}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	l := 16
	start := time.Now().Unix()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, l) // 标识 prime管道  close

	wg.Add(1)
	go putNum(intChan)

	for i := 0; i < l; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < l; i++ {
			<-exitChan
		}
		// 关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().Unix()

	fmt.Println("执行完毕...", end-start, "毫秒")
}
