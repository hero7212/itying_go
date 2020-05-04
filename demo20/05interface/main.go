package main

import "fmt"

// 定义一个方法，可以传入任意数据类型，然后
func MyPrint1(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("int类型")
	}
}

// x.(type) 这个语法只能用在switch语句里面
func MyPrint2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	default:
		fmt.Println("传入错误")
	}
}

// 类型断言
func main() {
	// var a interface{}
	// a = "你好golang"
	// v, ok := a.(string)
	// if ok {
	// 	fmt.Println("a是string类型，值是：", v)
	// } else {
	// 	fmt.Println("断言失败")
	// }

	MyPrint1("朱莉")
	MyPrint2(1)
}
