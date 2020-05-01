package main

import "fmt"

type MyInt int

func (m MyInt) printInfo() {
	fmt.Println("我是自定义类型里面的自定义方法")
}

func main() {

	var a MyInt = 20

	a.printInfo()
}
