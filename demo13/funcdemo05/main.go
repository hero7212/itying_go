package main

import "fmt"

var a = "全局变量"

func run(c string) {
	b := "局部变量"
	fmt.Println("run")
	fmt.Println(a)
	c = "变了"
	fmt.Println("b是", b)
	fmt.Println("c是：", c)
}

func main() {
	fmt.Println(a)
	run(a)
	fmt.Println(a)
}
