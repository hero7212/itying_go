package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {

	var p1 = Person{
		name: "哈哈",
		age:  20,
		sex:  "男",
	}

	p2 := p1
	p2.name = "李四"
	fmt.Printf("%#v\n", p1) //main.Person{name:"哈哈", age:20, sex:"男"}
	fmt.Printf("%#v\n", p2) //main.Person{name:"李四", age:20, sex:"男"}
}
