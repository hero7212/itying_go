package main

import "fmt"

func main() {
	flag := true // 全局变量
	if flag {
		fmt.Println("flag是true")
	}

	if age := 34; age > 20 { // age 是if的局部变量
		fmt.Println("成年人")
	}
}
