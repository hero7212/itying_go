package main

import "fmt"

func main() {
	// fmt.Println("你好 golang")
	// fmt.Print("你好 golang")
	// fmt.Printf("你好 golang")

	/*
		fmt.Println("你好 golang")
		fmt.Print("你好 golang")
	*/

	// var a int = 10
	// var b int = 3
	// var c int = 5

	// fmt.Printf("a=%v b=%v c=%v", a, b, c)

	a := 10
	b := 3
	c := 5

	fmt.Printf("a=%v b=%v c=%v", a, b, c)

	// 使用Printf打印一个变量的类型

	fmt.Printf("a=%v a的类型是%T", a, a)
}
