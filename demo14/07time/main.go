package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		ticker := time.NewTicker(time.Second)

		for t := range ticker.C {
			fmt.Println(t)
		}
	*/
	// ticker := time.NewTicker(time.Second)
	// n := 5
	// for t := range ticker.C {
	// 	n--
	// 	if n == 0 {
	// 		ticker.Stop() // 终止定时器
	// 		break
	// 	}
	// 	fmt.Println(t)
	// }

	fmt.Println("aaa")
	time.Sleep(time.Second)
	fmt.Println("aaa2")
	time.Sleep(time.Second)
	fmt.Println("aaa3")
	time.Sleep(time.Second * 5)
	fmt.Println("aaa4")

	for {
		time.Sleep(time.Second)
		fmt.Println("死循环")
	}
}
