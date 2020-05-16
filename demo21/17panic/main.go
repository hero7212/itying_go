package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello,world")
	}
}

func test() {
	// defer + recover
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("test()错误", err)
		}
	}()
	// 定义一个map
	var myMap map[int]string
	myMap[0] = "golang" // error
}

func main() {
	go sayHello()
	go test()

	time.Sleep(time.Second)
}
