package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	// var num = 10 + x	// 10 + x (mismatched types int and interface {})
	b, _ := x.(int)
	var num = 10 + b
	fmt.Println(num)

	// 反射来实现
	// v := reflect.ValueOf(x)
	// fmt.Println(v)

	// var n = v + 12	// v + 12 (mismatched types reflect.Value and int)
	// fmt.Println(n)

	// 反射获取变量的原生值
	v := reflect.ValueOf(x)

	var m = v.Int() + 12
	fmt.Println(m)
}

func main() {

	var a = 13
	reflectValue(a)
}
