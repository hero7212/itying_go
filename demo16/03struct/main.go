package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {

	var p2 = new(Person)
	p2.name = "李四"
	p2.sex = "男"
	p2.age = 30

	(*p2).name = "王五" // 这是底层的写法，也行

	fmt.Printf("值：%#v 类型：%T\n", p2, p2)

	var p3 = &Person{}
	p3.name = "朱莉"
	p3.age = 23
	p3.sex = "女"

	fmt.Printf("值：%#v 类型：%T\n", p3, p3)

	var p4 = Person{
		name: "哈哈",
		age:  20,
		sex:  "男",
	}

	fmt.Printf("值：%#v 类型：%T\n", p4, p4)

	var p5 = &Person{
		name: "王麻子",
		age:  20,
		sex:  "男",
	}

	fmt.Printf("值：%#v 类型：%T\n", p5, p5)

	var p6 = &Person{
		name: "程丽",
	}

	fmt.Printf("值：%#v 类型：%T\n", p6, p6)

	var p7 = &Person{
		"张敏",
		18,
		"女",
	}

	fmt.Printf("值：%#v 类型：%T\n", p7, p7)
}
