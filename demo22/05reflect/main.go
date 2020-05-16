package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// v.Kind()	// 获取种类
	kind := v.Kind()

	switch kind {
	case reflect.Int64:
		fmt.Printf("int类型的原始值%v，计算后的值是%v \n", v.Int(), v.Int()+10)
	case reflect.Float32:
		fmt.Printf("float32类型的原始值%v \n", v.Float()+10.1)
	case reflect.Float64:
		fmt.Printf("float64类型的原始值%v \n", v.Float()+10.1)
	case reflect.String:
		fmt.Printf("string类型的原始值%v \n", v.String())
	default:
		fmt.Printf("没有判断")
	}
}

func main() {
	var a int64 = 100
	var b float32 = 3.14
	var c string = "你好golang"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}
