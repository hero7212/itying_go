package main

import "fmt"

type Animaler1 interface {
	SetName(string)
}

type Animaler2 interface {
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

func main() {

	// dog实现Animal接口
	d := &Dog{
		Name: "狗",
	}
	var d1 Animaler1 = d // 让dog实现animer1接口
	var d2 Animaler2 = d // 让dog实现animer2接口
	fmt.Println(d2.GetName())
	d1.SetName("2狗")
	fmt.Println(d2.GetName())

}
