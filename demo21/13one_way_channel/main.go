package main

import "fmt"

func main() {
	// 声明管道, 双向
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)

	// 管道，只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 20
	// <-ch2	// receive from send-only type chan<- in

	// 管道，只读
	ch3 := make(<-chan int, 2)
	// ch3 <- 23	// send to receive-only type <-chan int

}
