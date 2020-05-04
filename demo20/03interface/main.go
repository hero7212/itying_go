package main

import "fmt"

type A interface{} // 空接口 表示没有任何约束	任意类型都可以实现空接口

func main() {

	var a A
	var str = "朱莉"
	a = str
	fmt.Printf("值：%v, 类型：%T \n", a, a)

	num := 20
	a = num // 表示让int类型实现A接口
	fmt.Printf("值：%v, 类型：%T \n", a, a)

	flag := true
	a = flag
	fmt.Printf("值：%v, 类型：%T \n", a, a)
}
