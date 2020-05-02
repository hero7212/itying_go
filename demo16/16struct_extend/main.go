package main

import "fmt"

// 父结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动", a.Name)
}

// 子结构体

type Dog struct {
	Age     int
	*Animal // 结构体嵌套  继承
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺", d.Name)
}

func main() {
	d := Dog{
		Age: 20,
		Animal: &Animal{
			Name: "阿奇",
		},
	}
	d.run()
	d.wang()
}
