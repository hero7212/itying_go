package main

import "fmt"

func main() {
	// var a = 10
	// fmt.Printf("a的值是：%v a的类型%T a的地址%p", a, a, &a)

	// var a = 10
	// var p = &a
	// fmt.Printf("a的值是：%v a的类型%T a的地址%p \n", a, a, &a)

	// fmt.Printf("p的值是：%v, p的类型%T \n", p, p)

	var a = 10
	var b = &a
	fmt.Printf("a的值是：%v a的类型%T a的地址%p \n", a, a, &a)

	fmt.Printf("b的值是：%v b的类型%T b的地址%p \n", b, b, &b)
}
