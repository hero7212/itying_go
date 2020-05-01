package main

import "fmt"

func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("结束")
}

func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

func f3() (a int) {
	defer func() {
		a++
	}()
	return
}

func main() {
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")
	// f1()
	fmt.Println(f2())
	fmt.Println(f3())
}
