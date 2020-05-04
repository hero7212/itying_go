package main

import "fmt"

// 空接口可以直接当做类型，表示任意类型

func show(a interface{}) {
	fmt.Printf("值：%v, 类型：%T \n", a, a)
}
func main() {

	// var a interface{}
	// a = 20
	// fmt.Printf("值：%v, 类型：%T \n", a, a)
	// a = "朱莉"
	// fmt.Printf("值：%v, 类型：%T \n", a, a)
	// a = true
	// fmt.Printf("值：%v, 类型：%T \n", a, a)

	// show(20)
	// show("你好")
	// slice := []interface{}{1, 2, 34, "朱莉"}
	// show(slice)

	var m1 = make(map[string]interface{})

	m1["username"] = "张三"
	m1["age"] = 20
	m1["married"] = true
	fmt.Println(m1)
}
