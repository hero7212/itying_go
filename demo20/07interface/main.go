package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 手机
type Phone struct {
	Name string
}

func (p Phone) start() { // 值接收者
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {

	var p1 = Phone{
		Name: "小米手机",
	}

	var p2 Usber = p1
	p2.start()

	var p3 = Phone{
		Name: "苹果手机",
	}

	var p4 Usber = p3
	p4.start()
}
