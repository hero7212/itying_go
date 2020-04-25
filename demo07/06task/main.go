package main

import "fmt"

func main() {
	// var a = 34
	// var b = 10
	// a, b = b, a
	// fmt.Println(a, b)

	var days = 100
	var week = days / 7
	var day = days % 7
	fmt.Printf("距离放假还有%v周零%v天", week, day)
}
