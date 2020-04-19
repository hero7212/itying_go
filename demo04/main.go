package main

import "fmt"

func main() {

	// var username = "张三"
	// fmt.Println(username)

	// const (
	// 	A = "A"
	// 	B = "B"
	// 	C
	// )
	// fmt.Println(A, B, C)

	// const a = iota
	// fmt.Println(a)

	// const (
	// 	n1 = iota
	// 	_
	// 	n3
	// 	n4
	// )
	// fmt.Println(n1, n3, n4)

	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4
		n5, n6
	)
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
