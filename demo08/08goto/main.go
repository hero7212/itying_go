package main

import "fmt"

func main() {

	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto label3
	}

	fmt.Println("aaa")
	fmt.Println("bbb")
label3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}
