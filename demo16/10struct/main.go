package main

import "fmt"

type Person struct {
	name  string
	age   int
	hobby []string
	map1  map[string]string
}

func main() {

	var p Person
	p.name = "张三"
	p.age = 20
	p.hobby = make([]string, 6, 6)
	p.hobby[0] = "写代码"
	p.hobby[1] = "打篮球"
	p.hobby[2] = "睡觉"

	p.map1 = make(map[string]string)

	p.map1["address"] = "北京"
	p.map1["phone"] = "1312432423"

	fmt.Printf("%#v \n", p)

	fmt.Printf("%#v \n", p.hobby)
}
