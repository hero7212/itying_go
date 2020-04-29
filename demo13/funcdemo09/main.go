package main

import "fmt"

func test() {
	fmt.Println("test...")
}

func main() {
	func() {
		fmt.Println("匿名, 自执行")
	}()

	fn := func(x, y int) int {
		return (x * y)
	}
	fmt.Println(fn(2, 3))

	func(x, y int) {
		fmt.Println(x, y)
	}(10, 20)
}
