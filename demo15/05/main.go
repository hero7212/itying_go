package main

import "fmt"

func main() {
	var a = new(int) // int类型的指针

	fmt.Printf("值：%v 类型：%T 指针变量对应的值：%v \n", a, a, *a)

	var b *int
	b = new(int)
	*b = 100
	fmt.Println(*b)

	var f = new(bool)
	fmt.Println(*f)

}
