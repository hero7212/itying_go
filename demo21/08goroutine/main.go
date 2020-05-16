package main

import "fmt"

func main() {

	// var ch1 = make(chan int, 10)
	// for i := 1; i <= 10; i++ {
	// 	ch1 <- i
	// }
	// close(ch1) // 关闭管道

	// // 管道没有key
	// for v := range ch1 {
	// 	fmt.Println(v)
	// }

	// 用for遍历管道到底要不要关闭？  用for循环，可以不关闭， 用for range就必须关闭
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	close(ch2) // 关闭管道

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
}
