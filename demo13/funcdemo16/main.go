package main

import "fmt"

func fn1(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	return a / b
}

func main() {
	fn1(10, 0)
	fmt.Println("结束")
	fn1(20, 0)
}
