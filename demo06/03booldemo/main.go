package main

import "fmt"

func main() {

	var flag = true

	fmt.Printf("%v--%T\n", flag, flag)

	var i int
	fmt.Println(i)
	fmt.Println("------------")
	var f float64
	fmt.Println(f)

	var a = true
	if a {
		fmt.Println(a)
	}
}
