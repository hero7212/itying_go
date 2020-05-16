package main

import "fmt"

func main() {
	// 声明管道
	ch := make(chan int, 3)

	// 给管道存值
	ch <- 10
	ch <- 21
	ch <- 32

	// 去管道的值，给a
	a := <-ch

	fmt.Println(a)

	// 从管道里取值，只是不付给变量而已
	<-ch

	c := <-ch

	fmt.Println(c)

	ch <- 56
	fmt.Printf("%v 容量：%v 长度%v \n", ch, cap(ch), len(ch))

	ch1 := make(chan int, 4)

	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	// 管道是引用
	ch2 := ch1

	ch2 <- 25
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d)

	// 管道阻塞

	// ch6 := make(chan int, 1)

	// ch6 <- 34
	// ch6 <- 36 // 放不下了，阻塞

	// ch7 := make(chan string, 2)

	// ch7 <- "数据1"
	// ch7 <- "数据2"

	// m1 := <-ch7
	// m2 := <-ch7
	// m3 := <-ch7 // 取不到了，阻塞
	// fmt.Println(m1, m2, m3)

	ch8 := make(chan int, 1)
	ch8 <- 34
	<-ch8
	ch8 <- 67
	<-ch8
	m4 := <-ch8
	fmt.Println(m4)
}
