package main

import "fmt"

func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	var a = adder1()
	var b = adder2()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b(10))
	fmt.Println(b(10))
	fmt.Println(b(10))
}
