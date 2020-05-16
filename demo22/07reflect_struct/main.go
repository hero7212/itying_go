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

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) setInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	field0 := t.Field(0)
	fmt.Printf("%#v \n", field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))
	fmt.Println("字段Tag：", field0.Tag.Get("form"))
	fmt.Println("---------------------------------")
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}

	var fieldCount = t.NumField()
	fmt.Printf("结构体有%v个属性 \n ", fieldCount)

	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println("--------------------------------")
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称：%v 属性值：%v 属性类型：%v 属性Tag:%v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}
}

func PrintStructFn(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	// v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	method0 := t.Method(0)
	fmt.Printf("%#v \n", method0)
	fmt.Println("字段名称：", method0.Name)
	fmt.Println("字段类型：", method0.Type)

	fmt.Println("------------------------------")

	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	fmt.Println("-------------------------------")

	// v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)

	info1 := v.MethodByName("GetInfo").Call(nil)

	fmt.Println(info1)

	var params = []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))

	v.MethodByName("SetInfo").Call(params)	// 执行方法传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	fmt.Println("方法数量：", t.NumMethod())

}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	// PrintStructField(stu1)
	PrintStructFn(&stu1)

	fmt.Printf("%#V\n", stu1)
}
