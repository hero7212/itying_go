package main

import "fmt"

type AInterface interface {
	SetName(string)
}

type BInterface interface {
	GetName() string
}

type Animaler interface {
	AInterface
	BInterface
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

func main() {

	// dog实现Animal接口
	d := &Dog{
		Name: "狗",
	}
	var d1 Animaler = d // 让dog实现Animaler接口
	fmt.Println(d1.GetName())
	d1.SetName("11狗")
	fmt.Println(d1.GetName())

}
