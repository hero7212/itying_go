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

func (p *Person) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

// golang中结构体是独立的，不会相互影响

func main() {

	var p1 = Person{
		name: "张三",
		age:  20,
		sex:  "男",
	}

	p1.PrintInfo() //姓名：张三 年龄：20

	var p2 = Person{
		name: "王五",
		age:  22,
		sex:  "男",
	}

	p1.SetInfo("朱莉", 34)
	p1.PrintInfo() //姓名：朱莉 年龄：34
	p2.PrintInfo() //姓名：王五 年龄：22
}
