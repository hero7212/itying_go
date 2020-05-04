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

func (p *Phone) start() { // 指针接收者
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	// 错误写法
	/*
		var phone1 = Phone{
			Name: "小米",
		}

		var p2 Usber = phone1
		p2.start()
	*/

	var phone1 = &Phone{
		Name: "小米",
	}

	var p1 Usber = phone1
	p1.start()
}
