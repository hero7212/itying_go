package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name1" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	fmt.Println(t.Kind())

	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是一个结构体指针")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体指针")
		return
	}

	name := v.Elem().FieldByName("Name")
	name.SetString("小莉")

	age := v.Elem().FieldByName("Age")
	age.SetInt(22)
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	reflectChangeStruct(&stu1)

	fmt.Printf("%#v \n", stu1)

}
