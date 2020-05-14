package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("i love test - ", i)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	go test() // 开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("i love main - ", i)
		time.Sleep(time.Millisecond * 50)
	}
}
