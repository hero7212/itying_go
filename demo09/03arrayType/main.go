package main

import "fmt"

func main() {

	// var a = 10
	// b := a
	// fmt.Println(a, b)
	// b++
	// fmt.Println(a, b)

	var sli1 = []int{1, 2, 3}
	sli2 := sli1
	sli2[1] = 10
	fmt.Println(sli1)
	fmt.Println(sli2)
}
