package main

import "fmt"

type Animaler interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c Cat) GetName() string {
	return c.Name
}

func main() {

	// dog实现Animal接口
	d := &Dog{
		Name: "狗",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("2狗")
	fmt.Println(d1.GetName())

	// cat实现Animal接口
	c := &Cat{
		Name: "猫",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())
	c1.SetName("2猫")
	fmt.Println(c1.GetName())
}
