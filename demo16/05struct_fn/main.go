package main

import "fmt"

type Person struct {
	name   string
	age    int
	sex    string
	height int
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄：%v \n", p.name, p.age)
}

func main() {

	var p1 = Person{
		name: "张三",
		age:  20,
		sex:  "男",
	}

	var p2 = Person{
		name: "李四",
		age:  30,
		sex:  "男",
	}

	p1.PrintInfo()
	p2.PrintInfo()

}
