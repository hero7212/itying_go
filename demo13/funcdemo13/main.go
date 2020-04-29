package main

import "fmt"

var a = 12

func test() {
	var a = 3
	fmt.Println(a)
}

func main() {
	test()
}
