package main

import "fmt"

func main() {

	// a := [5]int{55, 56, 57, 58, 59}

	// b := a[:] // 获取数组里面的所有值

	// fmt.Printf("a是%v, 类型是%T \n", a, a)
	// fmt.Printf("b是%v, 类型是%T \n", b, b)

	// c := a[0:4]
	// fmt.Printf("c是%v, 类型是%T \n", c, c)

	// d := a[2:]
	// fmt.Printf("d是%v, 类型是%T \n", d, d)

	// e := a[:3]
	// fmt.Printf("e是%v, 类型是%T \n", e, e)

	// a := []string{"北京", "上海", "广州"}
	// b := a[1:]
	// fmt.Printf("b是%v, 类型是%T \n", b, b)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度是%d, 容量是%d \n", len(s), cap(s))

	a := s[2:]
	fmt.Printf("长度是%d, 容量是%d \n", len(a), cap(a))

	b := s[1:3]
	fmt.Printf("长度是%d, 容量是%d \n", len(b), cap(b))
}
